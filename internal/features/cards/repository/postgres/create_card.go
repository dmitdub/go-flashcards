package cards_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
	core_postgres_pool "github.com/dmitdub/go-flashcards/internal/core/repository/postgres/pool"
)

func (r *CardsRepository) CreateCard(
	ctx context.Context,
	card domain.Card,
) (domain.Card, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO flashcards.cards (front, back, learned, parent_deck_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id, version, front, back, learned, parent_deck_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		card.Front,
		card.Back,
		card.Learned,
		card.ParentDeckID,
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
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Card{}, fmt.Errorf(
				"%v: deck with id='%d': %w",
				err,
				card.ParentDeckID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Card{}, fmt.Errorf("scan error: %w", err)
	}

	cardDomain := domain.NewCard(
		cardModel.ID,
		cardModel.Version,
		cardModel.Front,
		cardModel.Back,
		cardModel.Learned,
		cardModel.ParentDeckID,
	)

	return cardDomain, nil
}
