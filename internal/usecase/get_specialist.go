package usecase

import (
	"github.com/google/uuid"

	"backend/internal/infra"
)

func GetSpecialist(state State, specialistID uuid.UUID) (infra.Specialist, *UsecaseError) {
	specialist, err := state.Queries().GetSpecialistByID(state.Context(), specialistID)
	if err == nil {
		return specialist, nil
	}
	if ErrorIsNoRows(err) {
		return specialist, NewNotFoundError(ErrResourceNotFound).InField("specialist")
	}
	return specialist, NewUnexpectedError(err)
}
