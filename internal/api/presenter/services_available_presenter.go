package presenter

import (
	"backend/internal/api/dtos"
	"backend/internal/infra"
)

func GetServiceAvailable(s infra.GetServiceAvailableByIDRow) dtos.ServiceAvailable {
	return dtos.ServiceAvailable{
		ServiceName: s.ServiceName,
		ServiceNameId: s.ServiceNameID.String(),
		Specialization: s.Specialization,
		SpecializationID: s.SpecializationID.String(),
	}
}
