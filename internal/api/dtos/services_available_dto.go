package dtos

type ServiceAvailable struct {
	ServiceName      string `json:"serviceName"`
	ServiceNameId    string `json:"serviceNameId"`
	Specialization   string `json:"specialization"`
	SpecializationID string `json:"specializationId"`
}

type CreateServiceAvailableBody struct {
	Name             string `json:"name"`
	Specialization   string `json:"specialization"`
	SpecializationID string `json:"specializationId"`
}

type UpdateServiceAvailableBody struct {
	Name             string `json:"name"`
}
