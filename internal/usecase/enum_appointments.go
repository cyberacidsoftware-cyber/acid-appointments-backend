package usecase

type AppointmentStatus int32

const (
	AppointmentStatusNone AppointmentStatus = iota
	AppointmentStatusPending
	AppointmentStatusRealized
	AppointmentStatusCanceled
	AppointmentStatusCount
)
