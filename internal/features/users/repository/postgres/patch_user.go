package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
	core_postgres_pool "github.com/dmitdub/go-flashcards/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	id int,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE flashcards.users
	SET
		nickname=$1
		phone=$2
		version=version+1
	WHERE id=$3 AND version=$4
	RETURNING
		id,
		version,
		nickname,
		phone;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.Nickname,
		user.Phone,
		id,
		user.Version,
	)

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
				"user with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
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
