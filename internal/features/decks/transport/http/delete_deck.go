package decks_transport_http

import (
	"net/http"

	core_logger "github.com/dmitdub/go-flashcards/internal/core/logger"
	core_http_request "github.com/dmitdub/go-flashcards/internal/core/transport/http/request"
	core_http_response "github.com/dmitdub/go-flashcards/internal/core/transport/http/response"
)

func (h *DecksHTTPHandler) DeleteDeck(rw http.ResponseWriter, r *http.Request) {
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

	if err := h.decksService.DeleteDeck(ctx, deckID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete deck",
		)

		return
	}

	responseHandler.NoContentResponse()
}
