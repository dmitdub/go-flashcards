package cards_transport_http

import (
	"net/http"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_logger "github.com/dmitdub/go-flashcards/internal/core/logger"
	core_http_request "github.com/dmitdub/go-flashcards/internal/core/transport/http/request"
	core_http_response "github.com/dmitdub/go-flashcards/internal/core/transport/http/response"
)

type CreateCardRequest struct {
	Front        string `json:"front" validate:"required,min=1,max=200"`
	Back         string `json:"back" validate:"required,min=1,max=500"`
	ParentDeckID int    `json:"parent_deck_id" validate:"required"`
}

type CreateCardResponse CardDTOResponse

func (h *CardsHTTPHandler) CreateCard(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateCardRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	cardDomain := domain.NewCardUninitialized(
		request.Front,
		request.Back,
		request.ParentDeckID,
	)

	cardDomain, err := h.cardsService.CreateCard(ctx, cardDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create card",
		)

		return
	}

	response := CreateCardResponse(cardDTOFromDomain(cardDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}
