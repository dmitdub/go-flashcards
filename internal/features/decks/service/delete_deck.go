package decks_service

import (
	"context"
	"fmt"
)

func (s *DecksService) DeleteDeck(
	ctx context.Context,
	id int,
) error {
	if err := s.decksRepository.DeleteDeck(ctx, id); err != nil {
		return fmt.Errorf("delete deck from repository: %w", err)
	}

	return nil
}
