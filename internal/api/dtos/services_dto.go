package dtos

type Service struct {
	ID               string `json:"id"`
	ServiceName      string `json:"serviceName"`
	ServiceNameID    string `json:"serviceNameId"`
	Specialist       string `json:"specialistName"`
	SpecialistID     string `json:"specialistId"`
	Specialization   string `json:"specialization"`
	SpecializationID string `json:"specializationId"`
	Price            int32  `json:"price"`
	Duration         int32  `json:"duration"`
}

type ServiceInfoBody struct {
	SpecialistID  string `json:"specialistId"`
	ServiceNameID string `json:"serviceNameId"`
	Price         int32  `json:"price"`
	Duration      int32  `json:"duration"` // in minutes
}

type ServiceBase struct {
	ID            string `json:"id"`
	SpecialistID  string `json:"specialistId"`
	ServiceNameID string `json:"serviceNameId"`
	Price         int32  `json:"price"`
	Duration      int32  `json:"duration"` // in minutes
}
