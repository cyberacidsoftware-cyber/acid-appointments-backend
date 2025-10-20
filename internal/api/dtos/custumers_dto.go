package dtos

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Cpf       string `json:"cpf"`
}

type CustomerInfoBody struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Cpf       string `json:"cpf"`
}
