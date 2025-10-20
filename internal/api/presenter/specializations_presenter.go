package presenter

import (
	"backend/internal/api/dtos"
	"backend/internal/infra"
)

func GetSpecializations(specializations []infra.Specialization) []dtos.Specialization {
	response := make([]dtos.Specialization, 0, len(specializations))
	
	for _, s := range specializations {
		response = append(response, GetSpecialization(s))
	}

	return response
}

func GetSpecialization(s infra.Specialization) dtos.Specialization {
	return dtos.Specialization{
		ID: s.ID.String(),
		Name: s.Name,
	}
}