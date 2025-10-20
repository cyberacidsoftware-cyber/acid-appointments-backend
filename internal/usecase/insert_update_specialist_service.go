package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type SpecialistServiceInfoArgs struct {
	ServiceNameIDRaw  string
	ServiceNameID     uuid.UUID
	SpecialistIDRaw   string
	SpecialistID      uuid.UUID
	Price             int32
	DurationMin       int32
	Duration          pgtype.Interval
	RequireSpecialist bool
}

func (args *SpecialistServiceInfoArgs) Validate() *UsecaseError {
	if args.ServiceNameID == uuid.Nil {
		if args.ServiceNameIDRaw == "" {
			return NewInvalidArgumentError(ErrInvalidUuid).InField("serviceNameId")
		}
		if err := args.ServiceNameID.Scan(args.ServiceNameIDRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidUuid).InField("serviceNameId")
		}
	}
	if args.RequireSpecialist && args.SpecialistID == uuid.Nil {
		if args.SpecialistIDRaw == "" {
			return NewInvalidArgumentError(ErrInvalidUuid).InField("specialistId")
		}
		if err := args.SpecialistID.Scan(args.SpecialistIDRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidUuid).InField("specialistId")
		}
	}
	if !args.Duration.Valid {
		if args.DurationMin <= 0 {
			return NewInvalidArgumentError(ErrExpectPositiveValue).InField("duration")
		}
		args.Duration = DurationFromMinutes(args.DurationMin)
	}
	if args.Price < 0 {
		return NewInvalidArgumentError(ErrExpectPositiveValue).InField("price")
	}
	return nil
}

func CreateSpecialistService(state State, args SpecialistServiceInfoArgs) (uuid.UUID, *UsecaseError) {
	_, err := state.Queries().GetService(state.Context(), infra.GetServiceParams{
		SpecialistID:  args.SpecialistID,
		ServiceNameID: args.ServiceNameID,
	})
	if err == nil {
		return uuid.Nil, NewResourceAlreadyExistsError("service")
	} else if !ErrorIsNoRows(err) {
		return uuid.Nil, NewUnexpectedError(err)
	}

	id, err := state.Queries().CreateService(state.Context(), infra.CreateServiceParams{
		ServiceNameID: args.ServiceNameID,
		SpecialistID:  args.SpecialistID,
		Price:         args.Price,
		Duration:      args.Duration,
	})
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	return id, nil
}

func UpdateSpecialistService(state State, serviceId uuid.UUID, args SpecialistServiceInfoArgs) (uuid.UUID, *UsecaseError) {
	params := infra.UpdateServiceParams{
		ID:       serviceId,
		Price:    args.Price,
		Duration: args.Duration,
	}
	id, err := state.Queries().UpdateService(state.Context(), params)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	return id, nil
}
