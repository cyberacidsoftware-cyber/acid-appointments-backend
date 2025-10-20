package usecase

import (
	"fmt"

	"github.com/pkg/errors"
)

type UsecaseErrorKind int32

const (
	ErrorKindUnexpected UsecaseErrorKind = iota
	ErrorKindNotFound
	ErrorKindAlreadyExists
	ErrorKindInvalidArgument
	ErrorKindInvalidState
	ErrorKindAuth
)

type UsecaseError struct {
	Error error
	Kind  UsecaseErrorKind
}

func NewError(kind UsecaseErrorKind, err error) *UsecaseError {
	return &UsecaseError{Kind: kind, Error: err}
}

func NewUnexpectedError(err error) *UsecaseError {
	return NewError(ErrorKindUnexpected, err)
}

func NewNotFoundError(err error) *UsecaseError {
	return NewError(ErrorKindNotFound, err)
}

func NewInvalidArgumentError(err error) *UsecaseError {
	return NewError(ErrorKindInvalidArgument, err)
}

func NewAlreadyExistsError(err error) *UsecaseError {
	return NewError(ErrorKindAlreadyExists, err)
}

func NewResourceAlreadyExistsError(resource string) *UsecaseError {
	return NewAlreadyExistsError(ErrResourceAlreadyExists).InField(resource)
}

func NewUniqueInformationDuplicated(resource string) *UsecaseError {
	return NewAlreadyExistsError(ErrUniqueInformationDuplicated).InField(resource)
}

func NewInvalidStateError(err error) *UsecaseError {
	return NewError(ErrorKindInvalidState, err)
}

func NewAuthError(err error) *UsecaseError {
	return NewError(ErrorKindAuth, err)
}

func (e *UsecaseError) InField(field string) *UsecaseError {
	e.Error = fmt.Errorf("%s: %w", field, e.Error)
	return e
}

var (
	ErrUnreachable = errors.New("unreachable")

	ErrInvalidUuid     = errors.New("invalid uuid")
	ErrInvalidDate     = errors.New("invalid date")
	ErrInvalidTime     = errors.New("invalid time")
	ErrInvalidFormat   = errors.New("invalid format")
	ErrInvalidRange    = errors.New("invalid range")
	ErrInvalidDuration = errors.New("invalid duration")

	ErrExpectPositiveValue = errors.New("expect positive value")
)

var (
	ErrFieldIsRequired = errors.New("field is required")

	ErrInvalidAppointmentStatus = errors.New("invalid appointment status")
	ErrAppointmentsIntersection = errors.New("appointments intersection")

	ErrResourceNotFound = errors.New("resource not found")
)

var (
	ErrResourceAlreadyExists = errors.New("resource already exists")

	ErrUniqueInformationDuplicated = errors.New("unique information duplicated")
)
