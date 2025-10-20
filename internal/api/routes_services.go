package api

import (
	"encoding/json"
	"net/http"

	"backend/internal/api/dtos"
	"backend/internal/api/presenter"
	"backend/internal/usecase"

	"github.com/go-chi/render"
)

// @Summary      Get service
// @Security     ApiKeyAuth
// @Description  Get a service by id
// @Tags         Services
// @Produce      json
// @Success      200 {object}  dtos.Service
// @Router       /services/{service_id} [get]
func (h *api) getService(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	serviceId, ok := GetAndParseUuidParam(w, r, "service_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	service, err := usecase.GetService(rs, serviceId)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetService(service)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      List services
// @Security     ApiKeyAuth
// @Description  Get a page of services with filters
// @Tags         Services
// @Produce      json
// @Param        page				query int     false "Page"
// @Param        pageSize  		query int     false "Page size"
// @Param        service			query string  false "Service"
// @Param        specialist		query string  false "Specialist"
// @Param        specialization	query string  false "Specialization"
// @Success      200 {object}  []dtos.Service
// @Router       /services [get]
func (h *api) getServices(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	page := ParseIntOrDefault(query.Get("page"), 0)
	pageSize := ParseIntOrDefault(query.Get("pageSize"), 10)
	service := query.Get("service")
	specialist := query.Get("specialist")
	specialization := query.Get("specialization")

	// Validate e execute the usecase
	args := usecase.ListServicesEnrichedArgs{
		PageSize:           pageSize,
		Page:               page,
		SpecialistName:     specialist,
		SpecializationName: specialization,
		ServiceName:        service,
	}

	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	services, err := usecase.ListServicesEnriched(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetServices(services)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      Count services
// @Security     ApiKeyAuth
// @Description  Count services with filters
// @Tags         Services
// @Produce      json
// @Param        service        query string  false "Service"
// @Param        specialist		query string  false "Specialist"
// @Param        specialization	query string  false "Specialization"
// @Success      200 {object}  int
// @Router       /services/count [get]
func (h *api) countServices(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	service := query.Get("service")
	specialist := query.Get("specialist")
	specialization := query.Get("specialization")

	// Validate e execute the usecase
	args := usecase.CountServicesEnrichedArgs{
		SpecialistName:     specialist,
		SpecializationName: specialization,
		ServiceName:        service,
	}

	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	count, err := usecase.CountServicesEnriched(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.JSON(w, r, count)
	render.Status(r, http.StatusOK)
}

// @Summary      Create service
// @Security     ApiKeyAuth
// @Description  Create an new service
// @Tags         Services
// @Produce      json
// @Param 		  data body dtos.ServiceInfoBody true "Service information"
// @Success      200 {object}  dtos.Service
// @Router       /services [post]
func (h *api) createService(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	var body dtos.ServiceInfoBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.SpecialistServiceInfoArgs{
		ServiceNameIDRaw:  body.ServiceNameID,
		SpecialistIDRaw:   body.SpecialistID,
		Price:             body.Price,
		DurationMin:       body.Duration,
		RequireSpecialist: true,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	id, err := usecase.CreateSpecialistService(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.JSON(w, r, id.String())
	render.Status(r, http.StatusCreated)
}

// @Summary      Create service
// @Security     ApiKeyAuth
// @Description  Create an new service
// @Tags         Services
// @Produce      json
// @Param 		  data body dtos.ServiceInfoBody true "Service information"
// @Success      200 {object}  dtos.Service
// @Router       /services/{service_id} [put]
func (h *api) updateService(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	serviceId, ok := GetAndParseUuidParam(w, r, "service_id")
	if !ok {
		return
	}

	var body dtos.ServiceInfoBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.SpecialistServiceInfoArgs{
		ServiceNameIDRaw:  body.ServiceNameID,
		Price:             body.Price,
		DurationMin:       body.Duration,
		RequireSpecialist: false,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	id, err := usecase.UpdateSpecialistService(rs, serviceId, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.JSON(w, r, id.String())
	render.Status(r, http.StatusCreated)
}

// @Summary      Delete service
// @Security     ApiKeyAuth
// @Description  Delete an service by id
// @Tags         Services
// @Produce      json
// @Success      204
// @Router       /services/{service_id} [delete]
func (h *api) deleteService(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	serviceId, ok := GetAndParseUuidParam(w, r, "service_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	err := usecase.DeleteSpecialistService(rs, serviceId)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.NoContent(w, r)
}
