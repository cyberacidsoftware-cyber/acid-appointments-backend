package usecase

import (
	"backend/internal/infra"
)

type CountSecretariesArgs struct {
	Cpf   string
	Phone string
	Name  string
}

func (args *CountSecretariesArgs) Validate() *UsecaseError {
	return nil
}

func CountSecretaries(state State, args CountSecretariesArgs) (int64, *UsecaseError) {
	count, err := state.Queries().CountSecretaries(state.Context(), infra.CountSecretariesParams{
		Name:  args.Name,
		Cpf:   args.Cpf,
		Phone: args.Phone,
	})
	if err == nil {
		return count, nil
	}
	return 0, NewUnexpectedError(err)
}
