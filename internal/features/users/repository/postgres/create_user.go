package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO flashcards.users (nickname, phone)
	VALUES ($1, $2)
	RETURNING id, version, nickname, phone;
	`

	row := r.pool.QueryRow(ctx, query, user.Nickname, user.Phone)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Nickname,
		&userModel.Phone,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Nickname,
		userModel.Phone,
	)

	return userDomain, nil
}
