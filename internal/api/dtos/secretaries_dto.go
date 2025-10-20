package dtos

type Secretary struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Cpf       string `json:"cpf"`
	Cnpj      string `json:"cnpj"`
}

type SecretaryInfoBody struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Password  string `json:"password"`
	Cpf       string `json:"cpf"`
	Cnpj      string `json:"cnpj"`
}
