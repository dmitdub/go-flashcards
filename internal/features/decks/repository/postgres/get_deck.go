package decks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
	core_postgres_pool "github.com/dmitdub/go-flashcards/internal/core/repository/postgres/pool"
)

func (r *DecksRepository) GetDeck(
	ctx context.Context,
	id int,
) (domain.Deck, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, author_user_id
	FROM flashcards.decks
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)

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
				"deck with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}

		return domain.Deck{}, fmt.Errorf("scan error: %w", err)
	}

	deckDomain := deckDomainFromModel(deckModel)

	return deckDomain, nil
}
