package dtos

type Specialization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SpecializationInfoBody struct {
	Name string `json:"name"`
}