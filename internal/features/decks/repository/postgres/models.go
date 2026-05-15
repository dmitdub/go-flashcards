package decks_postgres_repository

import "github.com/dmitdub/go-flashcards/internal/core/domain"

type DeckModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	AuthorUserID int
}

func deckDomainFromModel(deckModel DeckModel) domain.Deck {
	return domain.NewDeck(
		deckModel.ID,
		deckModel.Version,
		deckModel.Title,
		deckModel.Description,
		deckModel.AuthorUserID,
	)
}

func deckDomainsFromModels(deckModels []DeckModel) []domain.Deck {
	domains := make([]domain.Deck, len(deckModels))

	for i, model := range deckModels {
		domains[i] = deckDomainFromModel(model)
	}

	return domains
}
