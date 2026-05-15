package cards_postgres_repository

import (
	"context"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
)

func (r *CardsRepository) GetCards(
	ctx context.Context,
	deckID *int,
	limit *int,
	offset *int,
) ([]domain.Card, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, front, back, learned, parent_deck_id
	FROM flashcards.cards
	%s
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	args := []any{limit, offset}

	if deckID != nil {
		query = fmt.Sprintf(query, "WHERE parent_deck_id=$3")
		args = append(args, deckID)
	} else {
		query = fmt.Sprintf(query, "")
	}

	rows, err := r.pool.Query(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("select cards: %w", err)
	}
	defer rows.Close()

	var cardModels []CardModel

	for rows.Next() {
		var cardModel CardModel

		err := rows.Scan(
			&cardModel.ID,
			&cardModel.Version,
			&cardModel.Front,
			&cardModel.Back,
			&cardModel.Learned,
			&cardModel.ParentDeckID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan cards: %w", err)
		}

		cardModels = append(cardModels, cardModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	cardDomains := cardDomainsFromModels(cardModels)

	return cardDomains, nil
}
