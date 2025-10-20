package presenter

import (
	"backend/internal/api/dtos"
	"backend/internal/infra"
)

func GetServices(services []infra.ListServicesEnrichedRow) []dtos.Service {
	response := make([]dtos.Service, 0, len(services))
	for _, s := range services {
		response = append(response, dtos.Service{
			ID:               s.ID.String(),
			ServiceName:      s.ServiceName,
			ServiceNameID:    s.ServiceNameID.String(),
			Specialist:       s.SpecialistName,
			SpecialistID:     s.SpecialistID.String(),
			Specialization:   s.SpecializationName,
			SpecializationID: s.SpecializationID.String(),
			Price:            s.Price,
			Duration:         MicrosToMinutes(s.Duration.Microseconds),
		})
	}
	return response
}

func GetService(s infra.Service) dtos.ServiceBase {
	return dtos.ServiceBase{
		ID:               s.ID.String(),
		ServiceNameID:    s.ServiceNameID.String(),
		SpecialistID:     s.SpecialistID.String(),
		Price:            s.Price,
		Duration:         MicrosToMinutes(s.Duration.Microseconds),
	}
}
