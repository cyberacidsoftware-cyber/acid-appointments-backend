package usecase

import (
	"backend/internal/infra"

	"github.com/jackc/pgx/v5/pgtype"
)

type CountAppointmentsArgs struct {
	CustomerName   string
	ServiceName    string
	SpecialistName string
	StartDateRaw   string
	StartDate      pgtype.Date
	EndDateRaw     string
	EndDate        pgtype.Date
	Status         int32
}

func (args *CountAppointmentsArgs) Validate() *UsecaseError {
	if args.StartDateRaw != "" {
		if err := args.StartDate.Scan(args.StartDateRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidDate).InField("startDate")
		}
	}
	if args.EndDateRaw != "" {
		if err := args.EndDate.Scan(args.EndDateRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidDate).InField("endDate")
		}
	}
	return nil
}

func CountAppointments(state State, args CountAppointmentsArgs) (int64, *UsecaseError) {
	count, err := state.Queries().CountAppointments(state.Context(), infra.CountAppointmentsParams{
		Column1: args.StartDate,
		Column2: args.EndDate,
		Column3: args.CustomerName,
		Column4: args.SpecialistName,
		Column5: args.ServiceName,
		Column6: args.Status,
	})
	if err == nil {
		return count, nil
	}
	return 0, NewUnexpectedError(err)
}
