package decks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
	core_postgres_pool "github.com/dmitdub/go-flashcards/internal/core/repository/postgres/pool"
)

func (r *DecksRepository) CreateDeck(
	ctx context.Context,
	deck domain.Deck,
) (domain.Deck, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO flashcards.decks (title, description, author_user_id)
	VALUES ($1, $2, $3)
	RETURNING id, version, title, description, author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		deck.Title,
		deck.Description,
		deck.AuthorUserID,
	)

	var deckModel DeckModel

	err := row.Scan(
		&deckModel.ID,
		&deckModel.Version,
		&deckModel.Title,
		&deckModel.Description,
		&deckModel.AuthorUserID,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Deck{}, fmt.Errorf(
				"%v: user with id='%d': %w",
				err,
				deck.AuthorUserID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Deck{}, fmt.Errorf("scan error: %w", err)
	}

	deckDomain := domain.NewDeck(
		deckModel.ID,
		deckModel.Version,
		deckModel.Title,
		deckModel.Description,
		deckModel.AuthorUserID,
	)

	return deckDomain, nil
}
