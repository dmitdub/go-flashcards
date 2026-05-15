package decks_transport_http

import "github.com/dmitdub/go-flashcards/internal/core/domain"

type DeckDTOResponse struct {
	ID           int     `json:"id"`
	Version      int     `json:"version"`
	Title        string  `json:"title"`
	Description  *string `json:"description"`
	AuthorUserID int     `json:"author_user_id"`
}

func deckDTOFromDomain(deck domain.Deck) DeckDTOResponse {
	return DeckDTOResponse{
		ID:           deck.ID,
		Version:      deck.Version,
		Title:        deck.Title,
		Description:  deck.Description,
		AuthorUserID: deck.AuthorUserID,
	}
}

func deckDTOsFromDomains(decks []domain.Deck) []DeckDTOResponse {
	dtos := make([]DeckDTOResponse, len(decks))

	for i, deck := range decks {
		dtos[i] = deckDTOFromDomain(deck)
	}

	return dtos
}
