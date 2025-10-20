package dtos

type Appointment struct {
	ID             string `json:"id"`
	CustomerName   string `json:"customerName"`
	CustomerID     string `json:"customerId"`
	ServiceName    string `json:"serviceName"`
	ServiceNameID  string `json:"serviceNameId"`
	SpecialistName string `json:"specialistName"`
	SpecialistID   string `json:"specialistId"`
	Price          int32  `json:"price"`
	Duration       int32  `json:"duration"`
	Date           string `json:"date"`
	Time           string `json:"time"`
	Status         int32  `json:"status"` // TODO: remove
}

type CreateAppointmentBody struct {
	CustomerID string `json:"customerId"`
	ServiceID  string `json:"serviceId"`
	Date       string `json:"date"`
	Time       string `json:"time"`
}

type UpdateAppointmentBody struct {
	Date       string `json:"date"`
	Time       string `json:"time"`
	Status     int32  `json:"status"`
}

type AppointmentCalendar struct {
	ID             string `json:"id"`
	Date           string `json:"date"`
	Time           string `json:"time"`
	SpecialistName string `json:"specialistName"`
	Status         int32  `json:"status"`
}

type AppointmentCalendarCount struct {
	Month         int32 `json:"month"`
	PendingCount  int32 `json:"pendingCount"`
	RealizedCount int32 `json:"realizedCount"`
	CanceledCount int32 `json:"canceledCount"`
}
