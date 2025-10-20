package usecase

import (
	"github.com/google/uuid"
)

func DeleteSpecialization(state State, id uuid.UUID) *UsecaseError {
	count, err := state.Queries().DeleteSpecializationByID(state.Context(), id)
	if err != nil {
		return NewError(ErrorKindUnexpected, err)
	}
	if count == 0 {
		return NewNotFoundError(ErrResourceNotFound).InField("specialization")
	}
	return nil
}
