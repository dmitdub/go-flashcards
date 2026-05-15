package decks_transport_http

import (
	"context"
	"net/http"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_http_server "github.com/dmitdub/go-flashcards/internal/core/transport/http/server"
)

type DecksHTTPHandler struct {
	decksService DecksService
}

type DecksService interface {
	CreateDeck(
		ctx context.Context,
		deck domain.Deck,
	) (domain.Deck, error)

	GetDecks(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Deck, error)

	GetDeck(
		ctx context.Context,
		id int,
	) (domain.Deck, error)

	DeleteDeck(
		ctx context.Context,
		id int,
	) error

	PatchDeck(
		ctx context.Context,
		id int,
		patch domain.DeckPatch,
	) (domain.Deck, error)
}

func NewDecksHTTPHandler(
	decksService DecksService,
) *DecksHTTPHandler {
	return &DecksHTTPHandler{
		decksService: decksService,
	}
}

func (h *DecksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/decks",
			Handler: h.CreateDeck,
		},
		{
			Method:  http.MethodGet,
			Path:    "/decks",
			Handler: h.GetDecks,
		},
		{
			Method:  http.MethodGet,
			Path:    "/decks/{id}",
			Handler: h.GetDeck,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/decks/{id}",
			Handler: h.DeleteDeck,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/decks/{id}",
			Handler: h.PatchDeck,
		},
	}
}
