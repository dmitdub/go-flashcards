package users_transport_http

import "github.com/dmitdub/go-flashcards/internal/core/domain"

type UserDTOResponse struct {
	ID       int     `json:"id"`
	Version  int     `json:"version"`
	Nickname string  `json:"nickname"`
	Phone    *string `json:"phone"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:       user.ID,
		Version:  user.Version,
		Nickname: user.Nickname,
		Phone:    user.Phone,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
