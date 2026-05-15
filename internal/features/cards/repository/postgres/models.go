package cards_postgres_repository

import "github.com/dmitdub/go-flashcards/internal/core/domain"

type CardModel struct {
	ID           int
	Version      int
	Front        string
	Back         string
	Learned      bool
	ParentDeckID int
}

func cardDomainFromModel(cardModel CardModel) domain.Card {
	return domain.NewCard(
		cardModel.ID,
		cardModel.Version,
		cardModel.Front,
		cardModel.Back,
		cardModel.Learned,
		cardModel.ParentDeckID,
	)
}

func cardDomainsFromModels(cardModels []CardModel) []domain.Card {
	domains := make([]domain.Card, len(cardModels))

	for i, model := range cardModels {
		domains[i] = cardDomainFromModel(model)
	}

	return domains
}
