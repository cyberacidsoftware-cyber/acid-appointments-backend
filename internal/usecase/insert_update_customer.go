package usecase

import (
	"backend/internal/infra"
	"backend/internal/validate"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type CustomerInfoArgs struct {
	Name          string
	Email         string
	EmailText     pgtype.Text
	Phone         string
	Birthdate     string
	BirthdateDate pgtype.Date
	Cpf           string
}

func (args *CustomerInfoArgs) Validate() *UsecaseError {
	if !args.EmailText.Valid {
		if err := args.EmailText.Scan(args.Email); err != nil {
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
	return nil
}

func CustomerWithPhoneExists(state State, phone string, exceptId uuid.UUID) (bool, error) {
	customer, err := state.Queries().GetCustomerByPhone(state.Context(), phone)
	if ErrorIsNoRows(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if customer.ID == exceptId {
		return false, nil
	}
	return true, nil
}

func CreateCustomer(state State, args CustomerInfoArgs) (infra.Customer, *UsecaseError) {
	var none infra.Customer
	exists, err := CustomerWithPhoneExists(state, args.Phone, uuid.Nil)
	if err != nil {
		return none, NewUnexpectedError(err)
	}
	if exists {
		return none, NewResourceAlreadyExistsError("customer.phone")
	}
	params := infra.CreateCustomerParams{
		Name:      args.Name,
		Email:     args.EmailText,
		Phone:     args.Phone,
		Birthdate: args.BirthdateDate,
		Cpf:       args.Cpf,
	}
	customer, err := state.Queries().CreateCustomer(state.Context(), params)
	if err != nil {
		return none, NewUnexpectedError(err)
	}
	return customer, nil
}

func UpdateCustomer(state State, customerId uuid.UUID, args CustomerInfoArgs) (infra.Customer, *UsecaseError) {
	var none infra.Customer
	exists, err := CustomerWithPhoneExists(state, args.Phone, customerId)
	if err != nil {
		return none, NewUnexpectedError(err)
	}
	if exists {
		return none, NewResourceAlreadyExistsError("customer.phone")
	}
	params := infra.UpdateCustomerParams{
		ID:        customerId,
		Name:      args.Name,
		Email:     args.EmailText,
		Phone:     args.Phone,
		Birthdate: args.BirthdateDate,
		Cpf:       args.Cpf,
	}
	customer, err := state.Queries().UpdateCustomer(state.Context(), params)
	if err != nil {
		return customer, NewUnexpectedError(err)
	}
	return customer, nil
}
