package usecase

import (
	"backend/internal/infra"
)

type CountCustomersArgs struct {
	Cpf   string
	Phone string
	Name  string
}

func (args *CountCustomersArgs) Validate() *UsecaseError {
	return nil
}

func CountCustomers(state State, args CountCustomersArgs) (int64, *UsecaseError) {
	count, err := state.Queries().CountCustomers(state.Context(), infra.CountCustomersParams{
		Name:  args.Name,
		Cpf:   args.Cpf,
		Phone: args.Phone,
	})
	if err == nil {
		return count, nil
	}
	return 0, NewUnexpectedError(err)
}
