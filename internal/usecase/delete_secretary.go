package usecase

import (
	"github.com/google/uuid"
)

func DeleteSecretary(state State, id uuid.UUID) *UsecaseError {
	count, err := state.Queries().DeleteSecretaryByID(state.Context(), id)
	if err != nil {
		return NewError(ErrorKindUnexpected, err)
	}
	if count == 0 {
		return NewNotFoundError(ErrResourceNotFound).InField("secretary")
	}
	return nil
}
