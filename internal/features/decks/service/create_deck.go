package decks_service

import (
	"context"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
)

func (s *DecksService) CreateDeck(
	ctx context.Context,
	deck domain.Deck,
) (domain.Deck, error) {
	if err := deck.Validate(); err != nil {
		return domain.Deck{}, fmt.Errorf("validate deck domain: %w", err)
	}

	deck, err := s.decksRepository.CreateDeck(ctx, deck)
	if err != nil {
		return domain.Deck{}, fmt.Errorf("create deck: %w", err)
	}

	return deck, nil
}
