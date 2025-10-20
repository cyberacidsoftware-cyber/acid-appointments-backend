package usecase

import (
	"backend/internal/infra"
	"backend/internal/validate"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type SecretaryInfoArgs struct {
	Name          string
	Email         string
	Phone         string
	Password      string
	Birthdate     string
	BirthdateDate pgtype.Date
	Cpf           string
	Cnpj          string
	CnpjText      pgtype.Text
	Update        bool
}

func (args *SecretaryInfoArgs) Validate() *UsecaseError {
	if !args.CnpjText.Valid {
		if err := args.CnpjText.Scan(args.Cnpj); err != nil {
			return NewUnexpectedError(err)
		}
	}
	if args.Birthdate != "" {
		if err := DateFromString(&args.BirthdateDate, args.Birthdate); err != nil {
			return NewInvalidArgumentError(ErrInvalidDate).InField("birthdate")
		}
	}
	if args.Name == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("name")
	}
	if args.Phone == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("phone")
	}
	if args.Cpf == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("cpf")
	}
	if !validate.IsCpfValid(args.Cpf) {
		return NewInvalidArgumentError(ErrInvalidFormat).InField("cpf")
	}
	if args.Cnpj == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("cnpj")
	}
	if !validate.IsCnpjValid(args.Cnpj) {
		return NewInvalidArgumentError(ErrInvalidFormat).InField("cnpj")
	}
	if args.Update {
		if args.Password != "" && !validate.IsPasswordValid(args.Password) {
			return NewInvalidArgumentError(ErrInvalidFormat).InField("password")
		}
	} else {
		if args.Password == "" {
			return NewInvalidArgumentError(ErrFieldIsRequired).InField("password")
		}
		if !validate.IsPasswordValid(args.Password) {
			return NewInvalidArgumentError(ErrInvalidFormat).InField("password")
		}
	}
	return nil
}

func SecretaryWithEmailExists(state State, email string, exceptId uuid.UUID) (infra.Secretary, bool, error) {
	secretary, err := state.Queries().GetSecretaryByEmail(state.Context(), email)
	if ErrorIsNoRows(err) {
		return secretary, false, nil
	}
	if err != nil {
		return secretary, false, err
	}
	if secretary.ID == exceptId {
		return secretary, false, nil
	}
	return secretary, true, nil
}

func CreateSecretary(state State, args SecretaryInfoArgs) (infra.Secretary, *UsecaseError) {
	var none infra.Secretary
	_, exists, err := SecretaryWithEmailExists(state, args.Email, uuid.Nil)
	if err != nil {
		return none, NewUnexpectedError(err)
	}
	if exists {
		return none, NewResourceAlreadyExistsError("secretary.email")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcrypt.DefaultCost)
	if err != nil {
		return none, NewUnexpectedError(err)
	}
	args.Password = string(hashedPassword)

	params := infra.CreateSecretaryParams{
		Name:      args.Name,
		Email:     args.Email,
		Phone:     args.Phone,
		Birthdate: args.BirthdateDate,
		Cpf:       args.Cpf,
		Cnpj:      args.CnpjText,
		Password:  args.Password,
	}
	secretary, err := state.Queries().CreateSecretary(state.Context(), params)
	if err != nil {
		return none, NewUnexpectedError(err)
	}
	return secretary, nil
}

func UpdateSecretary(state State, secretaryId uuid.UUID, args SecretaryInfoArgs) (infra.Secretary, *UsecaseError) {
	var none infra.Secretary
	secretary0, exists, err := SecretaryWithEmailExists(state, args.Email, secretaryId)
	if err != nil {
		return none, NewUnexpectedError(err)
	}
	if exists {
		return none, NewResourceAlreadyExistsError("secretary.email")
	}

	if args.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcrypt.DefaultCost)
		if err != nil {
			return none, NewUnexpectedError(err)
		}
		args.Password = string(hashedPassword)
	} else {
		args.Password = secretary0.Password
	}

	params := infra.UpdateSecretaryParams{
		ID:        secretaryId,
		Name:      args.Name,
		Email:     args.Email,
		Phone:     args.Phone,
		Birthdate: args.BirthdateDate,
		Cpf:       args.Cpf,
		Cnpj:      args.CnpjText,
		Password:  args.Password,
	}
	secretary, err := state.Queries().UpdateSecretary(state.Context(), params)
	if err != nil {
		return secretary, NewUnexpectedError(err)
	}
	return secretary, nil
}
