package usecase

import (
	"backend/internal/infra"

	"golang.org/x/crypto/bcrypt"
)

type AuthLoginArgs struct {
	Email    string
	Password string
}

func (args *AuthLoginArgs) Validate() *UsecaseError {
	return nil
}

func AuthLogin(state State, args AuthLoginArgs) (infra.GetIdentityByEmailRow, *UsecaseError) {
	identity, err := state.Queries().GetIdentityByEmail(state.Context(), args.Email)
	if ErrorIsNoRows(err) {
		return identity, NewAuthError(err)
	}
	if err != nil {
		return identity, NewUnexpectedError(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(identity.Password), []byte(args.Password))
	if err != nil {
		return identity, NewAuthError(err)
	}
	return identity, nil
}
