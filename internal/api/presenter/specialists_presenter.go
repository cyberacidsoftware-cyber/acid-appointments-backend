package presenter

import (
	"backend/internal/api/dtos"
	"backend/internal/infra"
)

func GetSpecialistServices(services []infra.ListServicesBySpecialistIDRow) []dtos.SpecialistService {
	response := make([]dtos.SpecialistService, 0, len(services))
	for _, s := range services {
		response = append(response, dtos.SpecialistService{
			ID:               s.ID.String(),
			ServiceName:      s.ServiceName,
			ServiceNameID:    s.ServiceNameID.String(),
			SpecializationID: s.SpecializationID.String(),
			Price:            s.Price,
			Duration:         MicrosToMinutes(s.Duration.Microseconds),
		})
	}
	return response
}

func GetSpecialistAppointments(appointments []infra.ListAppointmentsBySpecialistIDRow) []dtos.SpecialistAppointment {
	response := make([]dtos.SpecialistAppointment, 0, len(appointments))
	for _, a := range appointments {
		response = append(response, dtos.SpecialistAppointment{
			ID:            a.ID.String(),
			CustomerName:  a.CustomerName,
			CustomerID:    a.CustomerID.String(),
			ServiceName:   a.ServiceName,
			ServiceNameID: a.ServiceNameID.String(),
			Price:         a.Price,
			Duration:      MicrosToMinutes(a.Duration.Microseconds),
			Date:          DateToString(a.Date),
			Time:          TimeToString(a.Time),
		})
	}
	return response
}

func GetSpecialist(specialist infra.Specialist) dtos.Specialist {
	return dtos.Specialist{
		ID:        specialist.ID.String(),
		Name:      specialist.Name,
		Email:     specialist.Email,
		Phone:     specialist.Phone,
		Birthdate: DateToString(specialist.Birthdate),
		Cpf:       specialist.Cpf,
		Cnpj:      specialist.Cnpj.String,
	}
}

func ListSpecialists(specialists []infra.Specialist) []dtos.Specialist {
	response := make([]dtos.Specialist, 0, len(specialists))
	for _, s := range specialists {
		response = append(response, GetSpecialist(s))
	}
	return response
}
