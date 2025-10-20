package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
)

type ListSpecialistAppointmentsArgs struct {
	SpecialistID    uuid.UUID
	SpecialistIDRaw string
	DateRaw         string
}

func (args *ListSpecialistAppointmentsArgs) Validate() *UsecaseError {
	if args.SpecialistID == uuid.Nil {
		if err := args.SpecialistID.Scan(args.SpecialistIDRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidUuid).InField("specialistId")
		}
	}

	return nil
}

func ListSpecialistAppointments(rs State, args ListSpecialistAppointmentsArgs) ([]infra.ListAppointmentsBySpecialistIDRow, *UsecaseError) {
	appointments, err := rs.Queries().ListAppointmentsBySpecialistID(rs.Context(), infra.ListAppointmentsBySpecialistIDParams{
		SpecialistID: args.SpecialistID,
		Column2:      args.DateRaw,
	})
	if err != nil {
		return nil, NewUnexpectedError(err)
	}
	return appointments, nil
}
