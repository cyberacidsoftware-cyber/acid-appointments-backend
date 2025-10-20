package usecase

import (
	"backend/internal/infra"
)

type ListCustomersArgs struct {
	CountArgs      CountCustomersArgs
	PaginationArgs PaginationArgs
}

func (args *ListCustomersArgs) Validate() *UsecaseError {
	if err := args.CountArgs.Validate(); err != nil {
		return err
	}

	if err := args.PaginationArgs.Validate(); err != nil {
		return err
	}

	return nil
}

func ListCustomers(state State, args ListCustomersArgs) ([]infra.Customer, *UsecaseError) {
	customers, err := state.Queries().ListCustomers(state.Context(), infra.ListCustomersParams{
		Name:   args.CountArgs.Name,
		Cpf:    args.CountArgs.Cpf,
		Phone:  args.CountArgs.Phone,
		Limit:  args.PaginationArgs.PageSize,
		Offset: args.PaginationArgs.Page * args.PaginationArgs.PageSize,
	})
	if err == nil {
		return customers, nil
	}
	return nil, NewUnexpectedError(err)
}
