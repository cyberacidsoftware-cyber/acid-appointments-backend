package dtos

type Specialist struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Cpf       string `json:"cpf"`
	Cnpj      string `json:"cnpj"`
}

type SpecialistService struct {
	ID               string `json:"id"`
	SpecializationID string `json:"specializationId"`
	ServiceName      string `json:"serviceName"`
	ServiceNameID    string `json:"serviceNameId"`
	Price            int32  `json:"price"`
	Duration         int32  `json:"duration"`
}

type SpecialistAppointment struct {
	ID            string `json:"id"`
	CustomerName  string `json:"customerName"`
	CustomerID    string `json:"customerId"`
	ServiceName   string `json:"serviceName"`
	ServiceNameID string `json:"serviceId"`
	Price         int32  `json:"price"`
	Duration      int32  `json:"duration"`
	Date          string `json:"date"`
	Time          string `json:"time"`
}

type SpecialistServiceInfoBody struct {
	ServiceNameID string `json:"serviceNameId"`
	Price         int32  `json:"price"`
	Duration      int32  `json:"duration"` // in minutes
}

type SpecialistInfoBody struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Cpf       string `json:"cpf"`
	Cnpj      string `json:"cnpj"`
	Services  []SpecialistServiceInfoBody
}
