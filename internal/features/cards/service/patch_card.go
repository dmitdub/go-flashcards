package cards_service

import (
	"context"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
)

func (s *CardsService) PatchCard(
	ctx context.Context,
	id int,
	patch domain.CardPatch,
) (domain.Card, error) {
	card, err := s.cardsRepository.GetCard(ctx, id)
	if err != nil {
		return domain.Card{}, fmt.Errorf("get card: %w", err)
	}

	if err := card.ApplyPatch(patch); err != nil {
		return domain.Card{}, fmt.Errorf("apply card patch: %w", err)
	}

	patchedCard, err := s.cardsRepository.PatchCard(ctx, id, card)
	if err != nil {
		return domain.Card{}, fmt.Errorf("patch card: %w", err)
	}

	return patchedCard, nil
}
