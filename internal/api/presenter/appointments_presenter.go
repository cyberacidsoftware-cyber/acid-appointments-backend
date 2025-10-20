package presenter

import (
	"backend/internal/api/dtos"
	"backend/internal/infra"
	"backend/internal/usecase"
)

func GetAppointments(appointments []infra.ListAppointmentsRow) []dtos.Appointment {
	response := make([]dtos.Appointment, 0, len(appointments))
	for _, appointment := range appointments {
		response = append(response, GetAppointment(appointment))
	}
	return response
}

func GetAppointment(a infra.ListAppointmentsRow) dtos.Appointment {
	return dtos.Appointment{
		ID:             a.ID.String(),
		CustomerName:   a.CustomerName,
		CustomerID:     a.CustomerID.String(),
		ServiceName:    a.ServiceName,
		ServiceNameID:  a.ServiceNameID.String(),
		SpecialistName: a.SpecialistName,
		SpecialistID:   a.SpecialistID.String(),
		Price:          a.Price,
		Duration:       MicrosToMinutes(a.Duration.Microseconds),
		Date: 					DateToString(a.Date),
		Time:           TimeToString(a.Time),
		Status:         int32(a.Status),
	}
}

func GetAppointmentsCalendar(appointments []infra.ListAppointmentsCalendarRow) []dtos.AppointmentCalendar {
	response := make([]dtos.AppointmentCalendar, 0, len(appointments))
	for _, appointment := range appointments {
		response = append(response, dtos.AppointmentCalendar{
			ID:             appointment.ID.String(),
			Date: 					DateToString(appointment.Date),
			Time:           TimeToString(appointment.Time),
			SpecialistName: appointment.SpecialistName,
			Status:         int32(appointment.Status),
		})
	}
	return response
}

func GetAppointmentsCalendarCount(counts []infra.ListAppointmentsCalendarCountRow) []dtos.AppointmentCalendarCount {
	response := make([]dtos.AppointmentCalendarCount, 12)
	for index := range response {
		response[index].Month = int32(index)
	}
	for _, count := range counts {
		switch count.Status {
		case int32(usecase.AppointmentStatusPending):
			response[count.Month-1].PendingCount += count.Count
		case int32(usecase.AppointmentStatusRealized):
			response[count.Month-1].RealizedCount += count.Count
		case int32(usecase.AppointmentStatusCanceled):
			response[count.Month-1].CanceledCount += count.Count
		}
	}
	return response
}
