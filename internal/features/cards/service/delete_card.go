package cards_service

import (
	"context"
	"fmt"
)

func (s *CardsService) DeleteCard(
	ctx context.Context,
	id int,
) error {
	if err := s.cardsRepository.DeleteCard(ctx, id); err != nil {
		return fmt.Errorf("delete card from repository: %w", err)
	}

	return nil
}
