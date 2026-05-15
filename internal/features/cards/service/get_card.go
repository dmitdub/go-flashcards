package cards_service

import (
	"context"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
)

func (s *CardsService) GetCard(
	ctx context.Context,
	id int,
) (domain.Card, error) {
	card, err := s.cardsRepository.GetCard(ctx, id)
	if err != nil {
		return domain.Card{}, fmt.Errorf("get card from repository: %w", err)
	}

	return card, nil
}
