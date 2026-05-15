package cards_transport_http

import (
	"net/http"

	core_logger "github.com/dmitdub/go-flashcards/internal/core/logger"
	core_http_request "github.com/dmitdub/go-flashcards/internal/core/transport/http/request"
	core_http_response "github.com/dmitdub/go-flashcards/internal/core/transport/http/response"
)

type GetCardResponse CardDTOResponse

func (h *CardsHTTPHandler) GetCard(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	cardID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get cardID path value",
		)

		return
	}

	cardDomain, err := h.cardsService.GetCard(ctx, cardID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get card",
		)

		return
	}

	response := GetCardResponse(cardDTOFromDomain(cardDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}
