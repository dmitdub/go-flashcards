package cards_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
	core_postgres_pool "github.com/dmitdub/go-flashcards/internal/core/repository/postgres/pool"
)

func (r *CardsRepository) GetCard(
	ctx context.Context,
	id int,
) (domain.Card, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, front, back, learned, parent_deck_id
	FROM flashcards.cards
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)

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
				"card with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}

		return domain.Card{}, fmt.Errorf("scan error: %w", err)
	}

	cardDomain := cardDomainFromModel(cardModel)

	return cardDomain, nil
}
