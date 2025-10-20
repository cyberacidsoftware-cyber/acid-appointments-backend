package usecase

import "github.com/google/uuid"

func DeleteSpecialistService(state State, serviceID uuid.UUID) *UsecaseError {
	count, err := state.Queries().DeleteService(state.Context(), serviceID)
	if err != nil {
		return NewError(ErrorKindUnexpected, err)
	}
	if count == 0 {
		return NewNotFoundError(ErrResourceNotFound).InField("service")
	}
	return nil
}
