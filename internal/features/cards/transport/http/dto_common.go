package cards_transport_http

import "github.com/dmitdub/go-flashcards/internal/core/domain"

type CardDTOResponse struct {
	ID           int    `json:"id"`
	Version      int    `json:"version"`
	Front        string `json:"front"`
	Back         string `json:"back"`
	Learned      bool   `json:"learned"`
	ParentDeckID int    `json:"parent_deck_id"`
}

func cardDTOFromDomain(card domain.Card) CardDTOResponse {
	return CardDTOResponse{
		ID:           card.ID,
		Version:      card.Version,
		Front:        card.Front,
		Back:         card.Back,
		Learned:      card.Learned,
		ParentDeckID: card.ParentDeckID,
	}
}

func cardDTOsFromDomains(cards []domain.Card) []CardDTOResponse {
	dtos := make([]CardDTOResponse, len(cards))

	for i, card := range cards {
		dtos[i] = cardDTOFromDomain(card)
	}

	return dtos
}
