package usecase

import (
	"backend/internal/infra"
)

type CountServicesEnrichedArgs struct {
	SpecialistName     string
	SpecializationName string
	ServiceName        string
}

func (args *CountServicesEnrichedArgs) Validate() *UsecaseError {
	return nil
}

func CountServicesEnriched(state State, args CountServicesEnrichedArgs) (int64, *UsecaseError) {
	count, err := state.Queries().CountServicesEnriched(state.Context(), infra.CountServicesEnrichedParams{
		Specialist:     args.SpecialistName,
		Specialization: args.SpecializationName,
		ServiceName:    args.ServiceName,
	})
	if err == nil {
		return count, nil
	}
	return 0, NewUnexpectedError(err)
}
