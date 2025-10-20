package presenter

import (
	"backend/internal/api/dtos"
	"backend/internal/infra"
)

func GetCustomers(customers []infra.Customer) []dtos.Customer {
	response := make([]dtos.Customer, 0, len(customers))
	for _, c := range customers {
		response = append(response, GetCustomer(c))
	}
	return response
}

func GetCustomer(c infra.Customer) dtos.Customer {
	return dtos.Customer{
		ID:        c.ID.String(),
		Name:      c.Name,
		Email:     c.Email.String,
		Phone:     c.Phone,
		Birthdate: DateToString(c.Birthdate),
		Cpf:       c.Cpf,
	}
}
