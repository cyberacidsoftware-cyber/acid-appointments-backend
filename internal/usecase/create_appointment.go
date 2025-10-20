package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateAppointmentArgs struct {
	CustomerIDRaw    string
	CustomerID       uuid.UUID
	ServiceIDRaw     string
	ServiceID        uuid.UUID
	ServiceNameIDRaw string
	ServiceNameID    uuid.UUID
	SpecialistIDRaw  string
	SpecialistID     uuid.UUID
	DateRaw          string
	Date             pgtype.Date
	TimeRaw          string
	Time             pgtype.Time
	Status           int32
}

func (args *CreateAppointmentArgs) Validate() *UsecaseError {
	if args.ServiceID == uuid.Nil {
		if err := args.ServiceID.Scan(args.ServiceIDRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidUuid).InField("serviceId")
		}
	}
	if args.SpecialistID == uuid.Nil {
		if err := args.SpecialistID.Scan(args.SpecialistIDRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidUuid).InField("specialistId")
		}
	}
	if args.ServiceNameID == uuid.Nil {
		if err := args.ServiceNameID.Scan(args.ServiceNameIDRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidUuid).InField("serviceNameId")
		}
	}
	if args.CustomerID == uuid.Nil {
		if err := args.CustomerID.Scan(args.CustomerIDRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidUuid).InField("customerId")
		}
	}
	if !args.Date.Valid {
		if err := args.Date.Scan(args.DateRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidDate).InField("date")
		}
	}
	if !args.Time.Valid {
		if err := args.Time.Scan(args.TimeRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidTime).InField("time")
		}
	}
	if args.Status < 0 || args.Status >= int32(AppointmentStatusCount) {
		return NewInvalidArgumentError(ErrInvalidAppointmentStatus).InField("status")
	}
	if args.ServiceNameID == uuid.Nil || args.SpecialistID == uuid.Nil {
		if args.ServiceID == uuid.Nil {
			return NewInvalidArgumentError(ErrFieldIsRequired).InField("serviceId")
		}
	}
	if args.CustomerID == uuid.Nil {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("customerId")
	}

	return nil
}

func CreateAppointment(state State, args CreateAppointmentArgs) (uuid.UUID, *UsecaseError) {
	var service infra.Service

	if args.ServiceID == uuid.Nil {
		maybeService, err := state.Queries().GetService(state.Context(), infra.GetServiceParams{
			ServiceNameID: args.ServiceNameID,
			SpecialistID:  args.SpecialistID,
		})
		if ErrorIsNoRows(err) {
			return uuid.Nil, NewNotFoundError(ErrResourceNotFound).InField("service")
		}
		if err != nil {
			return uuid.Nil, NewUnexpectedError(err)
		}
		service = maybeService
	} else {
		maybeService, err := state.Queries().GetServiceByID(state.Context(), args.ServiceID)
		if ErrorIsNoRows(err) {
			return uuid.Nil, NewNotFoundError(ErrResourceNotFound).InField("service")
		}
		if err != nil {
			return uuid.Nil, NewUnexpectedError(err)
		}
		service = maybeService
	}

	// filter correct intersect status
	appointmentsIntersects, err := state.Queries().AppointmentsIntersects(state.Context(), infra.AppointmentsIntersectsParams{
		Date:         args.Date,
		Time:         args.Time,
		Duration:     service.Duration,
		SpecialistID: service.SpecialistID,
	})
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	if appointmentsIntersects {
		return uuid.Nil, NewInvalidStateError(ErrAppointmentsIntersection)
	}

	params := infra.CreateAppointmentParams{
		CustomerID:    args.CustomerID,
		SpecialistID:  service.SpecialistID,
		ServiceNameID: service.ServiceNameID,
		Price:         service.Price,
		Duration:      service.Duration,
		Date:          args.Date,
		Time:          args.Time,
		Status:        args.Status,
	}
	id, err := state.Queries().CreateAppointment(state.Context(), params)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	return id, nil
}
