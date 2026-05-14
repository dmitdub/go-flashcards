package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
	core_postgres_pool "github.com/dmitdub/go-flashcards/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, nickname, phone
	FROM flashcards.users
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Nickname,
		&userModel.Phone,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Nickname,
		userModel.Phone,
	)

	return userDomain, nil
}
