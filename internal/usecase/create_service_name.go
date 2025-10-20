package usecase

import (
	"backend/internal/infra"
	"strings"

	"github.com/google/uuid"
)

type CreateServiceNameArgs struct {
	Name                string
	Specialization      string
	SpecializationIDRaw string
	SpecializationID    uuid.UUID
}

func (args *CreateServiceNameArgs) Validate() *UsecaseError {
	args.Name = strings.TrimSpace(args.Name)
	if args.Specialization == "" && args.SpecializationID == uuid.Nil {
		if err := args.SpecializationID.Scan(args.SpecializationIDRaw); err != nil {
			return NewInvalidArgumentError(ErrInvalidUuid).InField("specializationId")
		}
	}
	if args.Name == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("name")
	}
	return nil
}

func ServiceWithNameExists(state State, name string, exceptId uuid.UUID) (bool, error) {
	service, err := state.Queries().GetServiceNameByName(state.Context(), name)
	if ErrorIsNoRows(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if service.ID == exceptId {
		return false, nil
	}
	return true, nil
}

func CreateServiceName(state State, args CreateServiceNameArgs) (uuid.UUID, *UsecaseError) {
	exists, err := ServiceWithNameExists(state, args.Name, uuid.Nil)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	if exists {
		return uuid.Nil, NewResourceAlreadyExistsError("service_name")
	}

	if args.Specialization != "" {
		specializationId, err := state.Queries().CreateSpecialization(state.Context(), args.Specialization)
		if err != nil {
			return uuid.Nil, NewUnexpectedError(err)
		}
		args.SpecializationID = specializationId
	}

	params := infra.CreateServiceNameParams{
		Name:             args.Name,
		SpecializationID: args.SpecializationID,
	}
	id, err := state.Queries().CreateServiceName(state.Context(), params)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	return id, nil
}
