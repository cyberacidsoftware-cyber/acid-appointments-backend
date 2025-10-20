package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
)

func GetServicesBySpecialistID(state State, specialistID uuid.UUID) ([]infra.ListServicesBySpecialistIDRow, *UsecaseError) {
	services, err := state.Queries().ListServicesBySpecialistID(state.Context(), specialistID)
	if err != nil {
		return nil, NewUnexpectedError(err)
	}
	return services, nil
}
