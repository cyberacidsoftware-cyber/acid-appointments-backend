package usecase

import (
	"backend/internal/infra"

	"github.com/jackc/pgx/v5/pgtype"
)

type ListAppointmentsCalendarArgs struct {
	Year          int
	StartDate     string
	StartDateDate pgtype.Date
	EndDate       string
	EndDateDate   pgtype.Date
}

func (args *ListAppointmentsCalendarArgs) Validate() *UsecaseError {
	if args.Year > 0 {
		if err := DateFromYear(&args.StartDateDate, args.Year); err != nil {
			return NewInvalidArgumentError(ErrInvalidDate).InField("startDate")
		}
		if err := args.EndDateDate.Scan(args.StartDateDate.Time.AddDate(1, 0, -1)); err != nil {
			return NewInvalidArgumentError(ErrInvalidDate).InField("endDate")
		}
	} else {
		if args.StartDate != "" {
			if err := args.StartDateDate.Scan(args.StartDate); err != nil {
				return NewInvalidArgumentError(ErrInvalidDate).InField("startDate")
			}
		}
		if args.EndDate != "" {
			if err := args.EndDateDate.Scan(args.EndDate); err != nil {
				return NewInvalidArgumentError(ErrInvalidDate).InField("endDate")
			}
		}
	}
	return nil
}

func ListAppointmentsCalendar(state State, args ListAppointmentsCalendarArgs) ([]infra.ListAppointmentsCalendarRow, *UsecaseError) {
	appointments, err := state.Queries().ListAppointmentsCalendar(state.Context(), infra.ListAppointmentsCalendarParams{
		Date:   args.StartDateDate,
		Date_2: args.EndDateDate,
	})
	if err == nil {
		return appointments, nil
	}
	return nil, NewUnexpectedError(err)
}
