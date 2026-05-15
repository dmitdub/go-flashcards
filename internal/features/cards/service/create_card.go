package cards_service

import (
	"context"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
)

func (s *CardsService) CreateCard(
	ctx context.Context,
	card domain.Card,
) (domain.Card, error) {
	if err := card.Validate(); err != nil {
		return domain.Card{}, fmt.Errorf("validate card domain: %w", err)
	}

	card, err := s.cardsRepository.CreateCard(ctx, card)
	if err != nil {
		return domain.Card{}, fmt.Errorf("create card: %w", err)
	}

	return card, nil
}
