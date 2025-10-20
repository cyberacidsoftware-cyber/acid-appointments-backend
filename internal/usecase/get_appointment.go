package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
)

func GetAppointment(state State, id uuid.UUID) (infra.ListAppointmentsRow, *UsecaseError) {
	appointmentEnriched, err := state.Queries().GetAppointmentEnrichedByID(state.Context(), id)
	appointment := infra.ListAppointmentsRow(appointmentEnriched) /* synced with `ListAppointments` */
	if err == nil {
		return appointment, nil
	}
	if ErrorIsNoRows(err) {
		return appointment, NewNotFoundError(ErrResourceNotFound).InField("appointment")
	}
	return appointment, NewUnexpectedError(err)
}
