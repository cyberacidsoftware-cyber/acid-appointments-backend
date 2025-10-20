package usecase

import (
	"backend/internal/infra"
)

type ListSecretariesArgs struct {
	CountArgs      CountSecretariesArgs
	PaginationArgs PaginationArgs
}

func (args *ListSecretariesArgs) Validate() *UsecaseError {
	if err := args.CountArgs.Validate(); err != nil {
		return err
	}

	if err := args.PaginationArgs.Validate(); err != nil {
		return err
	}

	return nil
}

func ListSecretaries(state State, args ListSecretariesArgs) ([]infra.Secretary, *UsecaseError) {
	secretaries, err := state.Queries().ListSecretaries(state.Context(), infra.ListSecretariesParams{
		Name:   args.CountArgs.Name,
		Cpf:    args.CountArgs.Cpf,
		Phone:  args.CountArgs.Phone,
		Limit:  args.PaginationArgs.PageSize,
		Offset: args.PaginationArgs.Page * args.PaginationArgs.PageSize,
	})
	if err == nil {
		return secretaries, nil
	}
	return nil, NewUnexpectedError(err)
}
