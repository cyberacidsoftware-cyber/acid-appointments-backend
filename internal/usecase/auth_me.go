package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
)

func AuthMe(state State, userId uuid.UUID) (infra.GetIdentityByIDRow, *UsecaseError) {
	identity, err := state.Queries().GetIdentityByID(state.Context(), userId)
	if ErrorIsNoRows(err) {
		return identity, NewAuthError(err)
	}
	if err != nil {
		return identity, NewUnexpectedError(err)
	}
	return identity, nil
}
