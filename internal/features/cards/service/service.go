package cards_service

import (
	"context"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
)

type CardsService struct {
	cardsRepository CardsRepository
}

type CardsRepository interface {
	CreateCard(
		ctx context.Context,
		card domain.Card,
	) (domain.Card, error)

	GetCards(
		ctx context.Context,
		deckID *int,
		limit *int,
		offset *int,
	) ([]domain.Card, error)

	GetCard(
		ctx context.Context,
		id int,
	) (domain.Card, error)

	DeleteCard(
		ctx context.Context,
		id int,
	) error

	PatchCard(
		ctx context.Context,
		id int,
		card domain.Card,
	) (domain.Card, error)
}

func NewCardsService(
	cardsRepository CardsRepository,
) *CardsService {
	return &CardsService{
		cardsRepository: cardsRepository,
	}
}
