package decks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
)

func (r *DecksRepository) GetDecks(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Deck, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, author_user_id
	FROM flashcards.decks
	%s
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	args := []any{limit, offset}

	if userID != nil {
		query = fmt.Sprintf(query, "WHERE author_user_id=$3")
		args = append(args, userID)
	} else {
		query = fmt.Sprintf(query, "")
	}

	rows, err := r.pool.Query(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("select decks: %w", err)
	}
	defer rows.Close()

	var deckModels []DeckModel

	for rows.Next() {
		var deckModel DeckModel

		err := rows.Scan(
			&deckModel.ID,
			&deckModel.Version,
			&deckModel.Title,
			&deckModel.Description,
			&deckModel.AuthorUserID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan decks: %w", err)
		}

		deckModels = append(deckModels, deckModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	deckDomains := deckDomainsFromModels(deckModels)

	return deckDomains, nil
}
