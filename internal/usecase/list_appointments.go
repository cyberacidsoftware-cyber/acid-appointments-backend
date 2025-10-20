package usecase

import (
	"backend/internal/infra"
)

type ListAppointmentsArgs struct {
	CountArgs      CountAppointmentsArgs
	PaginationArgs PaginationArgs
}

func (args *ListAppointmentsArgs) Validate() *UsecaseError {
	if err := args.CountArgs.Validate(); err != nil {
		return err
	}

	if err := args.PaginationArgs.Validate(); err != nil {
		return err
	}

	return nil
}

func ListAppointments(state State, args ListAppointmentsArgs) ([]infra.ListAppointmentsRow, *UsecaseError) {
	appointments, err := state.Queries().ListAppointments(state.Context(), infra.ListAppointmentsParams{
		Column1: args.CountArgs.StartDate,
		Column2: args.CountArgs.EndDate,
		Column3: args.CountArgs.CustomerName,
		Column4: args.CountArgs.SpecialistName,
		Column5: args.CountArgs.ServiceName,
		Column6: args.CountArgs.Status,
		Limit:   args.PaginationArgs.PageSize,
		Offset:  args.PaginationArgs.PageSize * args.PaginationArgs.Page,
	})
	if err == nil {
		return appointments, nil
	}
	return nil, NewUnexpectedError(err)
}
