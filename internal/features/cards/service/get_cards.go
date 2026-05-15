package cards_service

import (
	"context"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
)

func (s *CardsService) GetCards(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Card, error) {
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

	cards, err := s.cardsRepository.GetCards(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get cards from repository: %w", err)
	}

	return cards, nil
}
