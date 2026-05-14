package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/dmitdub/go-flashcards/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	Nickname string
	Phone    *string
}

func NewUser(
	id int,
	version int,
	nickname string,
	phone *string,
) User {
	return User{
		ID:       id,
		Version:  version,
		Nickname: nickname,
		Phone:    phone,
	}
}

func NewUserUninitialized(
	nickname string,
	phone *string,
) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		nickname,
		phone,
	)
}

func (u *User) Validate() error {
	nicknameLen := len([]rune(u.Nickname))
	if nicknameLen < 3 || nicknameLen > 30 {
		return fmt.Errorf(
			"invalid `Nickname` len: %d: %w",
			nicknameLen,
			core_errors.ErrInvalidArgument,
		)
	}

	if u.Phone != nil {
		phoneLen := len([]rune(*u.Phone))
		if phoneLen < 10 || phoneLen > 15 {
			return fmt.Errorf(
				"invalid `Phone` len: %d: %w",
				phoneLen,
				core_errors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.Phone) {
			return fmt.Errorf(
				"invalid `Phone` format: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

type UserPatch struct {
	Nickname Nullable[string]
	Phone    Nullable[string]
}

func NewUserPatch(
	nickname Nullable[string],
	phone Nullable[string],
) UserPatch {
	return UserPatch{
		Nickname: nickname,
		Phone:    phone,
	}
}

func (p *UserPatch) Validate() error {
	if p.Nickname.Set && p.Nickname.Value == nil {
		return fmt.Errorf(
			"`Nickname` can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u

	if patch.Nickname.Set {
		tmp.Nickname = *patch.Nickname.Value
	}

	if patch.Phone.Set {
		tmp.Phone = patch.Phone.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil
}
