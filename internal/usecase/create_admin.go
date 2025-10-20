package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const CreateAdminArgsName = "CreateAdminArgs"
const CreateAdminUsecaseName = "CreateAdminUsecase"

type CreateAdminArgs struct {
	Name     string
	Email    string
	Password string
}

func (args *CreateAdminArgs) Validate() *UsecaseError {
	if args.Name == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("name")
	}
	if args.Email == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("email")
	}
	if args.Password == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("password")
	}
	return nil
}

func CreateAdmin(state State, args CreateAdminArgs) (uuid.UUID, *UsecaseError) {
	_, err := state.Queries().GetIdentityByEmail(state.Context(), args.Email)
	if err == nil {
		return uuid.Nil, NewResourceAlreadyExistsError("identity")
	} else if !ErrorIsNoRows(err) {
		return uuid.Nil, NewUnexpectedError(err)
	}

	params := infra.CreateAdminParams{
		Name:     args.Name,
		Email:    args.Email,
		Password: args.Password,
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	params.Password = string(hashedPassword)

	id, err := state.Queries().CreateAdmin(state.Context(), params)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	return id, nil
}
