package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_logger "github.com/dmitdub/go-flashcards/internal/core/logger"
	core_http_request "github.com/dmitdub/go-flashcards/internal/core/transport/http/request"
	core_http_response "github.com/dmitdub/go-flashcards/internal/core/transport/http/response"
	core_http_types "github.com/dmitdub/go-flashcards/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	Nickname core_http_types.Nullable[string] `json:"nickname"`
	Phone    core_http_types.Nullable[string] `json:"phone"`
}

func (r *PatchUserRequest) Validate() error {
	if r.Nickname.Set {
		if r.Nickname.Value == nil {
			return fmt.Errorf("`Nickname` can't be NULL")
		}

		nicknameLen := len([]rune(*r.Nickname.Value))
		if nicknameLen < 3 || nicknameLen > 30 {
			return fmt.Errorf("`Nickname` must be between 3 and 30 symbols")
		}
	}

	if r.Phone.Set {
		if r.Phone.Value != nil {
			phoneLen := len([]rune(*r.Phone.Value))
			if phoneLen < 10 || phoneLen > 15 {
				return fmt.Errorf("`Phone` must be between 10 and 15 symbols")
			}

			if !strings.HasPrefix(*r.Phone.Value, "+") {
				return fmt.Errorf("`Phone` must startswith '+' symbol")
			}
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)

		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.Nickname.ToDomain(),
		request.Phone.ToDomain(),
	)
}
