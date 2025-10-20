package usecase

import (
	"backend/internal/infra"
	"backend/internal/validate"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type SpecialistInfoArgs struct {
	Name          string
	Email         string
	Phone         string
	Birthdate     string
	BirthdateDate pgtype.Date
	Cpf           string
	Cnpj          string
	CnpjText      pgtype.Text
}

func (args *SpecialistInfoArgs) Validate() *UsecaseError {
	if !args.BirthdateDate.Valid {
		if err := args.BirthdateDate.Scan(args.Birthdate); err != nil {
			return NewInvalidArgumentError(ErrInvalidDate).InField("birthdate")
		}
	}
	if err := args.CnpjText.Scan(args.Cnpj); err != nil {
		return NewUnexpectedError(err).InField("cnpj")
	}
	if args.Name == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("name")
	}
	if args.Email == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("email")
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
	if args.CnpjText.String != "" {
		if !validate.IsCnpjValid(args.CnpjText.String) {
			return NewInvalidArgumentError(ErrInvalidFormat).InField("cnpj")
		}
	}
	return nil
}

func ServiceWithEmailExists(state State, email string, exceptId uuid.UUID) (bool, error) {
	specialist, err := state.Queries().GetSpecialistByEmail(state.Context(), email)
	if ErrorIsNoRows(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if specialist.ID == exceptId {
		return false, nil
	}
	return true, nil
}

func CreateSpecialist(state State, args SpecialistInfoArgs) (uuid.UUID, *UsecaseError) {
	exists, err := ServiceWithEmailExists(state, args.Email, uuid.Nil)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	if exists {
		return uuid.Nil, NewResourceAlreadyExistsError("specialist.email")
	}

	params := infra.CreateSpecialistParams{
		Name:      args.Name,
		Email:     args.Email,
		Phone:     args.Phone,
		Birthdate: args.BirthdateDate,
		Cpf:       args.Cpf,
		Cnpj:      args.CnpjText,
	}
	id, err := state.Queries().CreateSpecialist(state.Context(), params)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	return id, nil
}

func UpdateSpecialist(state State, specialistId uuid.UUID, args SpecialistInfoArgs) (infra.Specialist, *UsecaseError) {
	var none infra.Specialist
	exists, err := ServiceWithEmailExists(state, args.Email, specialistId)
	if err != nil {
		return none, NewUnexpectedError(err)
	}
	if exists {
		return none, NewResourceAlreadyExistsError("specialist.email")
	}
	params := infra.UpdateSpecialistParams{
		ID:        specialistId,
		Name:      args.Name,
		Email:     args.Email,
		Phone:     args.Phone,
		Birthdate: args.BirthdateDate,
		Cpf:       args.Cpf,
		Cnpj: 	   args.CnpjText,
	}
	specialist, err := state.Queries().UpdateSpecialist(state.Context(), params)
	if err != nil {
		return specialist, NewUnexpectedError(err)
	}
	return specialist, nil
}
