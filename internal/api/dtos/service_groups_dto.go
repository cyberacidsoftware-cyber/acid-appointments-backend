package dtos

type ServiceGroupItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ServiceGroup struct {
	ID    string             `json:"id"`
	Name  string             `json:"name"`
	Items []ServiceGroupItem `json:"items"`
}
