package usecase

type UsecaseOperation int32

const (
	OperationNoop UsecaseOperation = iota
	OperationCreate
	OperationUpdate
	OperationDelete
	OperationError
)
