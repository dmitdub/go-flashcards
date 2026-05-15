package decks_service

import (
	"context"
	"fmt"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
)

func (s *DecksService) GetDeck(
	ctx context.Context,
	id int,
) (domain.Deck, error) {
	deck, err := s.decksRepository.GetDeck(ctx, id)
	if err != nil {
		return domain.Deck{}, fmt.Errorf("get deck from repository: %w", err)
	}

	return deck, nil
}
