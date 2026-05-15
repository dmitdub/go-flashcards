package domain

import (
	"fmt"

	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
)

type Deck struct {
	ID      int
	Version int

	Title       string
	Description *string

	AuthorUserID int
}

func NewDeck(
	id int,
	version int,
	title string,
	description *string,
	authorUserID int,
) Deck {
	return Deck{
		ID:           id,
		Version:      version,
		Title:        title,
		Description:  description,
		AuthorUserID: authorUserID,
	}
}

func NewDeckUninitialized(
	title string,
	description *string,
	authorUserID int,
) Deck {
	return NewDeck(
		UninitializedID,
		UninitializedVersion,
		title,
		description,
		authorUserID,
	)
}

func (d *Deck) Validate() error {
	titleLen := len([]rune(d.Title))
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf(
			"invalid `Title` len: %d: %w",
			titleLen,
			core_errors.ErrInvalidArgument,
		)
	}

	if d.Description != nil {
		descriptionLen := len([]rune(*d.Description))
		if descriptionLen < 1 || descriptionLen > 500 {
			return fmt.Errorf(
				"invalid `Description` len: %d: %w",
				descriptionLen,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

type DeckPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
}

func NewDeckPatch(
	title Nullable[string],
	description Nullable[string],
) DeckPatch {
	return DeckPatch{
		Title:       title,
		Description: description,
	}
}

func (p *DeckPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf(
			"`Title` can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (d *Deck) ApplyPatch(patch DeckPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate deck patch: %w", err)
	}

	tmp := *d

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched deck: %w", err)
	}

	*d = tmp

	return nil
}
