package users_postgres_repository

import "github.com/dmitdub/go-flashcards/internal/core/domain"

type UserModel struct {
	ID       int
	Version  int
	Nickname string
	Phone    *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.Nickname,
			user.Phone,
		)
	}

	return userDomains
}
