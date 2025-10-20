package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
)

func GetCustomer(state State, id uuid.UUID) (infra.Customer, *UsecaseError) {
	customer, err := state.Queries().GetCustomerByID(state.Context(), id)
	if err == nil {
		return customer, nil
	}
	if ErrorIsNoRows(err) {
		return customer, NewNotFoundError(ErrResourceNotFound).InField("customer")
	}
	return customer, NewUnexpectedError(err)
}
