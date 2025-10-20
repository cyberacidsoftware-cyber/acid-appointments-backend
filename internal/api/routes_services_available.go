package api

import (
	"encoding/json"
	"net/http"

	"backend/internal/api/dtos"
	"backend/internal/api/presenter"
	"backend/internal/usecase"

	"github.com/go-chi/render"
)

func (h *api) getServicesAvailable(w http.ResponseWriter, r *http.Request) {
	h.getServiceGroups(w, r)
}

func (h *api) getServiceAvailable(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	serviceId, ok := GetAndParseUuidParam(w, r, "service_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	service, err := usecase.GetServiceAvailable(rs, serviceId)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetServiceAvailable(service)
	render.JSON(w, r, response)
	render.Status(r, http.StatusCreated)
}

func (h *api) createServiceAvailable(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	var body dtos.CreateServiceAvailableBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.CreateServiceNameArgs{
		Name:                body.Name,
		Specialization:      body.Specialization,
		SpecializationIDRaw: body.SpecializationID,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	id, err := usecase.CreateServiceName(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.JSON(w, r, id.String())
	render.Status(r, http.StatusCreated)
}

func (h *api) updateServiceAvailable(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	serviceId, ok := GetAndParseUuidParam(w, r, "service_id")
	if !ok {
		return
	}

	var body dtos.UpdateServiceAvailableBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.UpdateServiceNameArgs{
		Name: body.Name,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	id, err := usecase.UpdateServiceName(rs, serviceId, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.JSON(w, r, id.String())
	render.Status(r, http.StatusCreated)
}

func (h *api) deleteServiceAvailable(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	serviceId, ok := GetAndParseUuidParam(w, r, "service_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	err := usecase.DeleteServiceName(rs, serviceId)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.NoContent(w, r)
}
