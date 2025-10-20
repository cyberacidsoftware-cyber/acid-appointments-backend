package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
)

type UpdateServiceNameArgs struct {
	Name string
}

func (args *UpdateServiceNameArgs) Validate() *UsecaseError {
	if args.Name == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("name")
	}
	return nil
}

func UpdateServiceName(state State, serviceId uuid.UUID, args UpdateServiceNameArgs) (uuid.UUID, *UsecaseError) {
	exists, err := ServiceWithNameExists(state, args.Name, serviceId)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	if exists {
		return uuid.Nil, NewResourceAlreadyExistsError("service_name")
	}

	params := infra.UpdateServiceNameParams{
		ID:   serviceId,
		Name: args.Name,
	}
	id, err := state.Queries().UpdateServiceName(state.Context(), params)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	return id, nil
}
