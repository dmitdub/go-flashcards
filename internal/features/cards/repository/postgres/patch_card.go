package cards_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
	core_postgres_pool "github.com/dmitdub/go-flashcards/internal/core/repository/postgres/pool"
)

func (r *CardsRepository) PatchCard(
	ctx context.Context,
	id int,
	card domain.Card,
) (domain.Card, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE flashcards.cards
	SET
		front=$1,
		back=$2,
		learned=$3,
		version=version + 1

	WHERE id=$4 AND version=$5

	RETURNING
		id,
		version,
		front,
		back,
		learned
		parent_deck_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		card.Front,
		card.Back,
		card.Learned,
		id,
		card.Version,
	)

	var cardModel CardModel

	err := row.Scan(
		&cardModel.ID,
		&cardModel.Version,
		&cardModel.Front,
		&cardModel.Back,
		&cardModel.Learned,
		&cardModel.ParentDeckID,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Card{}, fmt.Errorf(
				"card with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.Card{}, fmt.Errorf("scan error: %w", err)
	}

	cardDomain := cardDomainFromModel(cardModel)

	return cardDomain, nil
}
