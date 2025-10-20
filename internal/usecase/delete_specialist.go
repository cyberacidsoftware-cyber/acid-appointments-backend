package usecase

import (
	"github.com/google/uuid"
)

func DeleteSpecialist(state State, id uuid.UUID) *UsecaseError {
	count, err := state.Queries().DeleteSpecialistByID(state.Context(), id)
	if err != nil {
		return NewError(ErrorKindUnexpected, err)
	}
	if count == 0 {
		return NewNotFoundError(ErrResourceNotFound).InField("specialist")
	}
	return nil
}
