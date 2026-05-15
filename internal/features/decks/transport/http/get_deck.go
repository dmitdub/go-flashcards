package decks_transport_http

import (
	"net/http"

	core_logger "github.com/dmitdub/go-flashcards/internal/core/logger"
	core_http_request "github.com/dmitdub/go-flashcards/internal/core/transport/http/request"
	core_http_response "github.com/dmitdub/go-flashcards/internal/core/transport/http/response"
)

type GetDeckResponse DeckDTOResponse

func (h *DecksHTTPHandler) GetDeck(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	deckID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get deckID path value",
		)

		return
	}

	deckDomain, err := h.decksService.GetDeck(ctx, deckID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get deck",
		)

		return
	}

	response := GetDeckResponse(deckDTOFromDomain(deckDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}
