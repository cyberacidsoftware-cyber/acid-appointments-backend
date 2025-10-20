package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
)

func GetSecretary(state State, id uuid.UUID) (infra.Secretary, *UsecaseError) {
	secretary, err := state.Queries().GetSecretaryByID(state.Context(), id)
	if err == nil {
		return secretary, nil
	}
	if ErrorIsNoRows(err) {
		return secretary, NewNotFoundError(ErrResourceNotFound).InField("secretary")
	}
	return secretary, NewUnexpectedError(err)
}
