package cards_transport_http

import (
	"fmt"
	"net/http"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_logger "github.com/dmitdub/go-flashcards/internal/core/logger"
	core_http_request "github.com/dmitdub/go-flashcards/internal/core/transport/http/request"
	core_http_response "github.com/dmitdub/go-flashcards/internal/core/transport/http/response"
	core_http_types "github.com/dmitdub/go-flashcards/internal/core/transport/http/types"
)

type PatchCardRequest struct {
	Front   core_http_types.Nullable[string] `json:"front"`
	Back    core_http_types.Nullable[string] `json:"back"`
	Learned core_http_types.Nullable[bool]   `json:"learned"`
}

func (r *PatchCardRequest) Validate() error {
	if r.Front.Set {
		if r.Front.Value == nil {
			return fmt.Errorf("`Front` can't be NULL")
		}

		frontLen := len([]rune(*r.Front.Value))
		if frontLen < 1 || frontLen > 200 {
			return fmt.Errorf("`Front` must be between 1 and 200 symbols")
		}
	}

	if r.Back.Set {
		if r.Back.Value == nil {
			return fmt.Errorf("`Back` can't be NULL")
		}

		backLen := len([]rune(*r.Back.Value))
		if backLen < 1 || backLen > 500 {
			return fmt.Errorf("`Back` must be between 1 and 500 symbols")
		}
	}

	if r.Learned.Set {
		if r.Learned.Value == nil {
			return fmt.Errorf("`Learned` can't be NULL")
		}
	}

	return nil
}

type PatchDeckResponse CardDTOResponse

func (h *CardsHTTPHandler) PatchCard(rw http.ResponseWriter, r *http.Request) {
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

	var request PatchCardRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	cardPatch := cardPatchFromRequest(request)

	cardDomain, err := h.cardsService.PatchCard(ctx, cardID, cardPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch card",
		)

		return
	}

	response := PatchDeckResponse(cardDTOFromDomain(cardDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func cardPatchFromRequest(request PatchCardRequest) domain.CardPatch {
	return domain.NewCardPatch(
		request.Front.ToDomain(),
		request.Back.ToDomain(),
		request.Learned.ToDomain(),
	)
}
