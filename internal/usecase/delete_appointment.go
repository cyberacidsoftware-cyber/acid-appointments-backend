package usecase

import (
	"github.com/google/uuid"
)

func DeleteAppointment(state State, id uuid.UUID) *UsecaseError {
	count, err := state.Queries().DeleteAppointment(state.Context(), id)
	if err != nil {
		return NewError(ErrorKindUnexpected, err)
	}
	if count == 0 {
		return NewNotFoundError(ErrResourceNotFound).InField("appointment")
	}
	return nil
}
