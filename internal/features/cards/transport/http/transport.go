package cards_transport_http

import (
	"context"
	"net/http"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_http_server "github.com/dmitdub/go-flashcards/internal/core/transport/http/server"
)

type CardsHTTPHandler struct {
	cardsService CardsService
}

type CardsService interface {
	CreateCard(
		ctx context.Context,
		card domain.Card,
	) (domain.Card, error)

	GetCards(
		ctx context.Context,
		deckID *int,
		limit *int,
		offset *int,
	) ([]domain.Card, error)

	GetCard(
		ctx context.Context,
		id int,
	) (domain.Card, error)

	DeleteCard(
		ctx context.Context,
		id int,
	) error

	PatchCard(
		ctx context.Context,
		id int,
		patch domain.CardPatch,
	) (domain.Card, error)
}

func NewCardsHTTPHandler(
	cardsService CardsService,
) *CardsHTTPHandler {
	return &CardsHTTPHandler{
		cardsService: cardsService,
	}
}

func (h *CardsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/cards",
			Handler: h.CreateCard,
		},
		{
			Method:  http.MethodGet,
			Path:    "/cards",
			Handler: h.GetCards,
		},
		{
			Method:  http.MethodGet,
			Path:    "/cards/{id}",
			Handler: h.GetCard,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/cards/{id}",
			Handler: h.DeleteCard,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/cards/{id}",
			Handler: h.PatchCard,
		},
	}
}
