package decks_service

import (
	"context"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
)

func (s *DecksService) PatchDeck(
	ctx context.Context,
	id int,
	patch domain.DeckPatch,
) (domain.Deck, error) {
	deck, err := s.decksRepository.GetDeck(ctx, id)
	if err != nil {
		return domain.Deck{}, fmt.Errorf("get deck: %w", err)
	}

	if err := deck.ApplyPatch(patch); err != nil {
		return domain.Deck{}, fmt.Errorf("apply deck patch: %w", err)
	}

	patchedDeck, err := s.decksRepository.PatchDeck(ctx, id, deck)
	if err != nil {
		return domain.Deck{}, fmt.Errorf("patch deck: %w", err)
	}

	return patchedDeck, nil
}
