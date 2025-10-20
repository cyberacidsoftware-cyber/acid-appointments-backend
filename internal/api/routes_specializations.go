package api

import (
	"encoding/json"
	"net/http"

	"backend/internal/api/dtos"
	"backend/internal/api/presenter"
	"backend/internal/usecase"

	"github.com/go-chi/render"
)

// @Summary      List specializations
// @Security     ApiKeyAuth
// @Description  Get a page of specializations with filters
// @Tags         Specializations
// @Produce      json
// @Param        page      query int     false "Page"
// @Param        pageSize  query int     false "Page size"
// @Success      200 {object}  []dtos.Specialization
// @Router       /specializations [get]
func (h *api) getSpecializations(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	specializations, err := usecase.ListSpecializations(rs)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetSpecializations(specializations)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      Create specialization
// @Security     ApiKeyAuth
// @Description  Create an new specialization
// @Tags         Specializations
// @Produce      json
// @Param 		  data body dtos.SpecializationInfoBody true "Specialization information"
// @Success      200 {object}  dtos.Specialization
// @Router       /specializations [post]
func (h *api) createSpecialization(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	var body dtos.SpecializationInfoBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.SpecializationInfoArgs{
		Name: body.Name,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	id, err := usecase.CreateSpecialization(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	// response := presenter.GetSpecialization(specialization)
	render.JSON(w, r, id.String())
	render.Status(r, http.StatusCreated)
}

// @Summary      Update specialization
// @Security     ApiKeyAuth
// @Description  Update a specialization by id
// @Tags         Specializations
// @Produce      json
// @Success      200 {object}  dtos.Specialization
// @Router       /specializations/{specialization_id} [put]
func (h *api) updateSpecialization(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	specializationId, ok := GetAndParseUuidParam(w, r, "specialization_id")
	if !ok {
		return
	}
	var body dtos.SpecializationInfoBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.SpecializationInfoArgs{
		Name:      body.Name,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	id, err := usecase.UpdateSpecialization(rs, specializationId, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	// response := presenter.GetSpecialization(specialization)
	render.JSON(w, r, id.String())
	render.Status(r, http.StatusCreated)
}

// @Summary      Delete specialization
// @Security     ApiKeyAuth
// @Description  Delete a specialization by id
// @Tags         Specializations
// @Produce      json
// @Success      204
// @Router       /specializations/{specialization_id} [delete]
func (h *api) deleteSpecialization(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	specializationId, ok := GetAndParseUuidParam(w, r, "specialization_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	err := usecase.DeleteSpecialization(rs, specializationId)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.NoContent(w, r)
}
