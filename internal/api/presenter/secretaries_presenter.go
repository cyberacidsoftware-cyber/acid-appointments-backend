package presenter

import (
	"backend/internal/api/dtos"
	"backend/internal/infra"
)

func GetSecretaries(secretaries []infra.Secretary) []dtos.Secretary {
	response := make([]dtos.Secretary, 0, len(secretaries))
	for _, c := range secretaries {
		response = append(response, GetSecretary(c))
	}
	return response
}

func GetSecretary(c infra.Secretary) dtos.Secretary {
	return dtos.Secretary{
		ID:        c.ID.String(),
		Name:      c.Name,
		Email:     c.Email,
		Phone:     c.Phone,
		Birthdate: DateToString(c.Birthdate),
		Cpf:       c.Cpf,
		Cnpj:      c.Cnpj.String,
	}
}
