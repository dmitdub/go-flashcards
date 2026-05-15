package decks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
	core_postgres_pool "github.com/dmitdub/go-flashcards/internal/core/repository/postgres/pool"
)

func (r *DecksRepository) PatchDeck(
	ctx context.Context,
	id int,
	deck domain.Deck,
) (domain.Deck, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE flashcards.decks
	SET
		title=$1,
		description=$2,
		version=version + 1

	WHERE id=$3 AND version=$4

	RETURNING
		id,
		version,
		title,
		description,
		author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		deck.Title,
		deck.Description,
		id,
		deck.Version,
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
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Deck{}, fmt.Errorf(
				"deck with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.Deck{}, fmt.Errorf("scan error: %w", err)
	}

	deckDomain := deckDomainFromModel(deckModel)

	return deckDomain, nil
}
