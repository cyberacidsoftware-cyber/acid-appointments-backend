package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend/internal/api/dtos"
	"backend/internal/api/presenter"
	"backend/internal/usecase"

	"github.com/go-chi/render"
)

func (h *api) getSpecialist(w http.ResponseWriter, r *http.Request) {
	specialistID, ok := GetAndParseUuidParam(w, r, "specialist_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	specialist, err := usecase.GetSpecialist(rs, specialistID)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetSpecialist(specialist)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

func (h *api) listSpecialists(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	page := ParseIntOrDefault(query.Get("page"), 0)
	pageSize := ParseIntOrDefault(query.Get("pageSize"), 10)
	name := query.Get("name")
	cpf := query.Get("cpf")
	phone := query.Get("phone")

	// Validate e execute the usecase
	args := usecase.ListSpecialistsArgs{
		PaginationArgs: usecase.PaginationArgs{
			PageSize: pageSize,
			Page:     page,
		},
		CountArgs: usecase.CountSpecialistArgs{
			Cpf:   cpf,
			Phone: phone,
			Name:  name,
		},
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	specialists, err := usecase.ListSpecialists(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.ListSpecialists(specialists)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

func (h *api) countSpecialists(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	name := query.Get("name")
	cpf := query.Get("cpf")
	phone := query.Get("phone")

	// Validate e execute the usecase
	args := usecase.CountSpecialistArgs{
		Cpf:   cpf,
		Phone: phone,
		Name:  name,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	count, err := usecase.CountSpecialists(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.JSON(w, r, count)
	render.Status(r, http.StatusOK)
}

func (h *api) createSpecialist(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	var body dtos.SpecialistInfoBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	var args usecase.SpecialistWithServicesInfoArgs
	args.Specialist = usecase.SpecialistInfoArgs{
		Name:      body.Name,
		Email:     body.Email,
		Phone:     body.Phone,
		Birthdate: body.Birthdate,
		Cpf:       body.Cpf,
		Cnpj:      body.Cnpj,
	}
	args.Services = make([]usecase.SpecialistServiceInfoArgs, 0, len(body.Services))
	for _, s := range body.Services {
		sArgs := usecase.SpecialistServiceInfoArgs{
			ServiceNameIDRaw:  s.ServiceNameID,
			Price:             s.Price,
			DurationMin:       s.Duration,
			RequireSpecialist: false,
		}
		args.Services = append(args.Services, sArgs)
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	tx, txErr := h.pool.Begin(r.Context())
	if txErr != nil {
		SomethingWentWrong(w, fmt.Errorf("failed to start transaction: %w", txErr))
		return
	}
	defer tx.Rollback(r.Context())

	rs := NewRequestState(h.q.WithTx(tx), r)

	specialistID, err := usecase.CreateSpecialistWithServices(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}
	tx.Commit(r.Context())

	// Format the response
	render.JSON(w, r, specialistID.String())
	render.Status(r, http.StatusCreated)
}

func (h *api) updateSpecialist(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	specialistID, ok := GetAndParseUuidParam(w, r, "specialist_id")
	if !ok {
		return
	}

	var body dtos.SpecialistInfoBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	var args usecase.SpecialistWithServicesInfoArgs
	args.Specialist = usecase.SpecialistInfoArgs{
		Name:      body.Name,
		Email:     body.Email,
		Phone:     body.Phone,
		Birthdate: body.Birthdate,
		Cpf:       body.Cpf,
		Cnpj:      body.Cnpj,
	}
	args.Services = make([]usecase.SpecialistServiceInfoArgs, 0, len(body.Services))
	for _, s := range body.Services {
		sArgs := usecase.SpecialistServiceInfoArgs{
			SpecialistID:      specialistID,
			ServiceNameIDRaw:  s.ServiceNameID,
			Price:             s.Price,
			DurationMin:       s.Duration,
			RequireSpecialist: true,
		}
		args.Services = append(args.Services, sArgs)
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	tx, txErr := h.pool.Begin(r.Context())
	if txErr != nil {
		SomethingWentWrong(w, fmt.Errorf("failed to start transaction: %w", txErr))
		return
	}
	defer tx.Rollback(r.Context())

	rs := NewRequestState(h.q.WithTx(tx), r)

	specialist, err := usecase.UpdateSpecialistWithServices(rs, specialistID, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}
	tx.Commit(r.Context())

	// Format the response
	response := presenter.GetSpecialist(specialist)
	render.JSON(w, r, response)
	render.Status(r, http.StatusCreated)
}

func (h *api) deleteSpecialist(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	specialistID, ok := GetAndParseUuidParam(w, r, "specialist_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	err := usecase.DeleteSpecialist(rs, specialistID)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.NoContent(w, r)
}

func (h *api) getSpecialistAppointments(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	date := query.Get("date")

	specialistID, ok := GetAndParseUuidParam(w, r, "specialist_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	args := usecase.ListSpecialistAppointmentsArgs{
		SpecialistID: specialistID,
		DateRaw:      date,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	appointments, err := usecase.ListSpecialistAppointments(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetSpecialistAppointments(appointments)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

func (h *api) getSpecialistSpecializations(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	specialistID, ok := GetAndParseUuidParam(w, r, "specialist_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	specializations, err := rs.Queries().ListSpecializationsBySpecialistID(rs.Context(), specialistID)
	if err != nil {
		SomethingWentWrong(w, err)
		return
	}

	// Format the response
	var response []dtos.Specialization
	for _, s := range specializations {
		response = append(response, dtos.Specialization{
			ID:   s.ID.String(),
			Name: s.Name,
		})
	}

	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

func (h *api) getSpecialistServices(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	specialistID, ok := GetAndParseUuidParam(w, r, "specialist_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	services, err := usecase.GetServicesBySpecialistID(rs, specialistID)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetSpecialistServices(services)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

func (h *api) getSpecialistService(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	specialistID, ok := GetAndParseUuidParam(w, r, "specialist_id")
	if !ok {
		return
	}
	serviceNameId, ok := GetAndParseUuidParam(w, r, "service_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	service, err := usecase.GetSpecialistService(rs, specialistID, serviceNameId)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetService(service)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}
