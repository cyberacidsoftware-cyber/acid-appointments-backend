package usecase

import (
	"backend/internal/infra"
)

type ListServicesEnrichedArgs struct {
	PageSize           int32
	Page               int32
	SpecialistName     string
	SpecializationName string
	ServiceName        string
}

func (args *ListServicesEnrichedArgs) Validate() *UsecaseError {
	if args.PageSize == 0 {
		args.PageSize = 10
	}

	if args.PageSize < 0 {
		return NewInvalidArgumentError(ErrExpectPositiveValue).InField("pageSize")
	}
	if args.Page < 0 {
		return NewInvalidArgumentError(ErrExpectPositiveValue).InField("page")
	}
	return nil
}

func ListServicesEnriched(state State, args ListServicesEnrichedArgs) ([]infra.ListServicesEnrichedRow, *UsecaseError) {
	services, err := state.Queries().ListServicesEnriched(state.Context(), infra.ListServicesEnrichedParams{
		Limit:          args.PageSize,
		Offset:         args.Page * args.PageSize,
		Specialist:     args.SpecialistName,
		Specialization: args.SpecializationName,
		ServiceName:    args.ServiceName,
	})
	if err != nil {
		return nil, NewUnexpectedError(err)
	}
	return services, nil
}
