package domain

import (
	"fmt"

	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
)

type Card struct {
	ID      int
	Version int

	Front   string
	Back    string
	Learned bool

	ParentDeckID int
}

func NewCard(
	id int,
	version int,
	front string,
	back string,
	learned bool,
	parentDeckID int,
) Card {
	return Card{
		ID:           id,
		Version:      version,
		Front:        front,
		Back:         back,
		Learned:      learned,
		ParentDeckID: parentDeckID,
	}
}

func NewCardUninitialized(
	front string,
	back string,
	parentDeckID int,
) Card {
	return NewCard(
		UninitializedID,
		UninitializedVersion,
		front,
		back,
		false,
		parentDeckID,
	)
}

func (c *Card) Validate() error {
	frontLen := len([]rune(c.Front))
	if frontLen < 1 || frontLen > 200 {
		return fmt.Errorf(
			"invalid `Front` len: %d: %w",
			frontLen,
			core_errors.ErrInvalidArgument,
		)
	}

	backLen := len([]rune(c.Back))
	if backLen < 1 || backLen > 500 {
		return fmt.Errorf(
			"invalid `Back` len: %d: %w",
			backLen,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

type CardPatch struct {
	Front   Nullable[string]
	Back    Nullable[string]
	Learned Nullable[bool]
}

func NewCardPatch(
	front Nullable[string],
	back Nullable[string],
	learned Nullable[bool],
) CardPatch {
	return CardPatch{
		Front:   front,
		Back:    back,
		Learned: learned,
	}
}

func (p *CardPatch) Validate() error {
	if p.Front.Set && p.Front.Value == nil {
		return fmt.Errorf(
			"`Front` can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Back.Set && p.Back.Value == nil {
		return fmt.Errorf(
			"`Back` can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Learned.Set && p.Learned.Value == nil {
		return fmt.Errorf(
			"`Learned` can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (c *Card) ApplyPatch(patch CardPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate card patch: %w", err)
	}

	tmp := *c

	if patch.Front.Set {
		tmp.Front = *patch.Front.Value
	}

	if patch.Back.Set {
		tmp.Back = *patch.Back.Value
	}

	if patch.Learned.Set {
		tmp.Learned = *patch.Learned.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched card: %w", err)
	}

	*c = tmp

	return nil
}
