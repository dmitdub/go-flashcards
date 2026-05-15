package decks_transport_http

import (
	"net/http"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_logger "github.com/dmitdub/go-flashcards/internal/core/logger"
	core_http_request "github.com/dmitdub/go-flashcards/internal/core/transport/http/request"
	core_http_response "github.com/dmitdub/go-flashcards/internal/core/transport/http/response"
)

type CreateDeckRequest struct {
	Title        string  `json:"title" validate:"required,min=1,max=100"`
	Description  *string `json:"description" validate:"omitempty,min=1,max=500"`
	AuthorUserID int     `json:"author_user_id" validate:"required"`
}

type CreateDeckResponse DeckDTOResponse

func (h *DecksHTTPHandler) CreateDeck(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateDeckRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	deckDomain := domain.NewTaskUninitialized(
		request.Title,
		request.Description,
		request.AuthorUserID,
	)

	deckDomain, err := h.decksService.CreateDeck(ctx, deckDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create deck",
		)

		return
	}

	response := CreateDeckResponse(deckDTOFromDomain(deckDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}
