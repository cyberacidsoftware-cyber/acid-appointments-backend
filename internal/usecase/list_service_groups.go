package usecase

import (
	"backend/internal/infra"
)

func ListServiceGroups(state State) ([]infra.Specialization, []infra.ServiceName, *UsecaseError) {
	specializations, err := state.Queries().ListSpecializations(state.Context())
	if err != nil {
		return nil, nil, NewUnexpectedError(err)
	}

	serviceNames, err := state.Queries().ListServiceNames(state.Context())
	if err != nil {
		return nil, nil, NewUnexpectedError(err)
	}

	return specializations, serviceNames, nil
}
