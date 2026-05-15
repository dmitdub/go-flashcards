package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_logger "github.com/dmitdub/go-flashcards/internal/core/logger"
	core_pgx_pool "github.com/dmitdub/go-flashcards/internal/core/repository/postgres/pool/pgx"
	core_http_middlware "github.com/dmitdub/go-flashcards/internal/core/transport/http/middleware"
	core_http_server "github.com/dmitdub/go-flashcards/internal/core/transport/http/server"
	decks_postgres_repository "github.com/dmitdub/go-flashcards/internal/features/decks/repository/postgres"
	decks_service "github.com/dmitdub/go-flashcards/internal/features/decks/service"
	decks_transport_http "github.com/dmitdub/go-flashcards/internal/features/decks/transport/http"
	users_postgres_repository "github.com/dmitdub/go-flashcards/internal/features/users/repository/postgres"
	users_service "github.com/dmitdub/go-flashcards/internal/features/users/service"
	users_transport_http "github.com/dmitdub/go-flashcards/internal/features/users/transport/http"
	"go.uber.org/zap"
)

var timeZone = time.UTC

func main() {
	time.Local = timeZone

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("application time zone", zap.Any("zone", timeZone))

	logger.Debug("initializing postgres connection pool")
	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing feature", zap.String("feature", "decks"))
	decksRepository := decks_postgres_repository.NewDecksRepository(pool)
	decksService := decks_service.NewDecksService(decksRepository)
	decksTransportHTTP := decks_transport_http.NewDecksHTTPHandler(decksService)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middlware.RequestID(),
		core_http_middlware.Logger(logger),
		core_http_middlware.Trace(),
		core_http_middlware.Panic(),
	)
	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRoutes(decksTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouterV1)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
