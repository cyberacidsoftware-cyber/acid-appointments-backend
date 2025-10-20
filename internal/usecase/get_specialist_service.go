package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
)

func GetSpecialistService(state State, specialistId, serviceNameId uuid.UUID) (infra.Service, *UsecaseError) {
	service, err := state.Queries().GetService(state.Context(), infra.GetServiceParams{
		SpecialistID:  specialistId,
		ServiceNameID: serviceNameId,
	})
	if err == nil {
		return service, nil
	}
	if ErrorIsNoRows(err) {
		return service, NewNotFoundError(ErrResourceNotFound).InField("service")
	}
	return service, NewUnexpectedError(err)
}
