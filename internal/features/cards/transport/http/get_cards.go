package cards_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/dmitdub/go-flashcards/internal/core/logger"
	core_http_request "github.com/dmitdub/go-flashcards/internal/core/transport/http/request"
	core_http_response "github.com/dmitdub/go-flashcards/internal/core/transport/http/response"
)

type GetCardsResponse []CardDTOResponse

func (h *CardsHTTPHandler) GetCards(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	deckID, limit, offset, err := getDeckIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get deckID/limit/offset query params",
		)

		return
	}

	cardsDomains, err := h.cardsService.GetCards(ctx, deckID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get cards",
		)

		return
	}

	response := GetCardsResponse(cardDTOsFromDomains(cardsDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getDeckIDLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		deckIDQueryParamKey = "deck_id"
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	deckID, err := core_http_request.GetIntQueryParam(r, deckIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'deck_id' query param: %w", err)
	}

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return deckID, limit, offset, nil
}
