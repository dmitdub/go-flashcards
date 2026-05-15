package decks_service

import (
	"context"

	"github.com/dmitdub/go-flashcards/internal/core/domain"
)

type DecksService struct {
	decksRepository DecksRepository
}

type DecksRepository interface {
	CreateDeck(
		ctx context.Context,
		deck domain.Deck,
	) (domain.Deck, error)

	GetDecks(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Deck, error)

	GetDeck(
		ctx context.Context,
		id int,
	) (domain.Deck, error)

	DeleteDeck(
		ctx context.Context,
		id int,
	) error

	PatchDeck(
		ctx context.Context,
		id int,
		deck domain.Deck,
	) (domain.Deck, error)
}

func NewDecksService(
	decksRepository DecksRepository,
) *DecksService {
	return &DecksService{
		decksRepository: decksRepository,
	}
}
