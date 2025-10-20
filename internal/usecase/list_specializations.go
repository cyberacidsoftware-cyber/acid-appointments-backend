package usecase

import (
	"backend/internal/infra"
)

func ListSpecializations(state State) ([]infra.Specialization, *UsecaseError) {
	specializations, err := state.Queries().ListSpecializations(state.Context())
	if err != nil {
		return nil, NewUnexpectedError(err)
	}
	return specializations, nil
}
