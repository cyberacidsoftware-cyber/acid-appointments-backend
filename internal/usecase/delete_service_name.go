package usecase

import (
	"github.com/google/uuid"
)

func DeleteServiceName(state State, id uuid.UUID) *UsecaseError {
	count, err := state.Queries().DeleteServiceNameByID(state.Context(), id)
	if err != nil {
		return NewError(ErrorKindUnexpected, err)
	}
	if count == 0 {
		return NewNotFoundError(ErrResourceNotFound).InField("service_name")
	}
	return nil
}
