package usecase

import (
	"backend/internal/infra"
	"strings"

	"github.com/google/uuid"
)

type SpecializationInfoArgs struct {
	Name string
}

func (args *SpecializationInfoArgs) Validate() *UsecaseError {
	args.Name = strings.TrimSpace(args.Name)
	if args.Name == "" {
		return NewInvalidArgumentError(ErrFieldIsRequired).InField("name")
	}
	return nil
}

func SpecializationWithNameExists(state State, name string, exceptId uuid.UUID) (bool, error) {
	specialization, err := state.Queries().GetSpecializationByName(state.Context(), name)
	if ErrorIsNoRows(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if specialization.ID == exceptId {
		return false, nil
	}
	return true, nil
}

func CreateSpecialization(state State, args SpecializationInfoArgs) (uuid.UUID, *UsecaseError) {
	exists, err := SpecializationWithNameExists(state, args.Name, uuid.Nil)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	if exists {
		return uuid.Nil, NewResourceAlreadyExistsError("specialization.name")
	}

	id, err := state.Queries().CreateSpecialization(state.Context(), args.Name)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	return id, nil
}

func UpdateSpecialization(state State, specializationId uuid.UUID, args SpecializationInfoArgs) (uuid.UUID, *UsecaseError) {
	exists, err := SpecializationWithNameExists(state, args.Name, specializationId)
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	if exists {
		return uuid.Nil, NewResourceAlreadyExistsError("specialization.name")
	}

	id, err := state.Queries().UpdateSpecialization(state.Context(), infra.UpdateSpecializationParams{
		ID: specializationId,
		Name: args.Name,
	})
	if err != nil {
		return uuid.Nil, NewUnexpectedError(err)
	}
	return id, nil
}
