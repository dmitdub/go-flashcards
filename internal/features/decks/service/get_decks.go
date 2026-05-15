package decks_service

import (
	"context"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
)

func (s *DecksService) GetDecks(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Deck, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	decks, err := s.decksRepository.GetDecks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get decks from repository: %w", err)
	}

	return decks, nil
}
