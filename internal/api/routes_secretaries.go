package api

import (
	"encoding/json"
	"net/http"

	"backend/internal/api/dtos"
	"backend/internal/api/presenter"
	"backend/internal/usecase"

	"github.com/go-chi/render"
)

// @Summary      Get secretary
// @Security     ApiKeyAuth
// @Description  Get a secretary by id
// @Tags         Secretaries
// @Produce      json
// @Success      200 {object}  dtos.Secretary
// @Router       /secretaries/{secretary_id} [get]
func (h *api) getSecretary(w http.ResponseWriter, r *http.Request) {
	// Authorize access
	jwtData := GetJwtData(r)
	if !jwtData.HasAccess("secretary") {
		http.Error(w, "invalid access", http.StatusForbidden)
		return
	}

	// Collect query parameters, path parameters, and request body
	secretaryId, ok := GetAndParseUuidParam(w, r, "secretary_id")
	if !ok {
		return
	}

	if jwtData.Role == "secretary" && jwtData.UserID != secretaryId.String() {
		http.Error(w, "invalid access", http.StatusForbidden)
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	secretary, err := usecase.GetSecretary(rs, secretaryId)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetSecretary(secretary)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      List secretaries
// @Security     ApiKeyAuth
// @Description  Get a page of secretaries with filters
// @Tags         Secretaries
// @Produce      json
// @Param        page      query int     false "Page"
// @Param        pageSize  query int     false "Page size"
// @Param        name      query string  false "Name"
// @Param        cpf       query string  false "CPF"
// @Param        phone     query string  false "Phone"
// @Success      200 {object}  []dtos.Secretary
// @Router       /secretaries [get]
func (h *api) getSecretaries(w http.ResponseWriter, r *http.Request) {
	// Authorize access
	jwtData := GetJwtData(r)
	if !jwtData.HasAccess() {
		http.Error(w, "invalid access", http.StatusForbidden)
		return
	}

	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	page := ParseIntOrDefault(query.Get("page"), 0)
	pageSize := ParseIntOrDefault(query.Get("pageSize"), 10)
	name := query.Get("name")
	cpf := query.Get("cpf")
	phone := query.Get("phone")

	// Validate e execute the usecase
	args := usecase.ListSecretariesArgs{
		PaginationArgs: usecase.PaginationArgs{
			PageSize: pageSize,
			Page:     page,
		},
		CountArgs: usecase.CountSecretariesArgs{
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

	secretaries, err := usecase.ListSecretaries(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetSecretaries(secretaries)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      Count secretaries
// @Security     ApiKeyAuth
// @Description  Count secretaries with filters
// @Tags         Secretaries
// @Produce      json
// @Param        name      query string  false "Name"
// @Param        cpf       query string  false "CPF"
// @Param        phone     query string  false "Phone"
// @Success      200 {object}  int
// @Router       /secretaries/count [get]
func (h *api) countSecretaries(w http.ResponseWriter, r *http.Request) {
	// Authorize access
	jwtData := GetJwtData(r)
	if !jwtData.HasAccess() {
		http.Error(w, "invalid access", http.StatusForbidden)
		return
	}

	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	name := query.Get("name")
	cpf := query.Get("cpf")
	phone := query.Get("phone")

	// Validate e execute the usecase
	args := usecase.CountSecretariesArgs{
		Cpf:   cpf,
		Phone: phone,
		Name:  name,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	count, err := usecase.CountSecretaries(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.JSON(w, r, count)
	render.Status(r, http.StatusOK)
}

// @Summary      Create secretary
// @Security     ApiKeyAuth
// @Description  Create an new secretary
// @Tags         Secretaries
// @Produce      json
// @Param 		  data body dtos.SecretaryInfoBody true "Secretary information"
// @Success      200 {object}  dtos.Secretary
// @Router       /secretaries [post]
func (h *api) createSecretary(w http.ResponseWriter, r *http.Request) {
	// Authorize access
	jwtData := GetJwtData(r)
	if !jwtData.HasAccess() {
		http.Error(w, "invalid access", http.StatusForbidden)
		return
	}

	// Collect query parameters, path parameters, and request body
	var body dtos.SecretaryInfoBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.SecretaryInfoArgs{
		Name:      body.Name,
		Email:     body.Email,
		Phone:     body.Phone,
		Birthdate: body.Birthdate,
		Password:  body.Password,
		Cpf:       body.Cpf,
		Cnpj:      body.Cnpj,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	secretary, err := usecase.CreateSecretary(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetSecretary(secretary)
	render.JSON(w, r, response)
	render.Status(r, http.StatusCreated)
}

// @Summary      Update secretary
// @Security     ApiKeyAuth
// @Description  Update a secretary by id
// @Tags         Secretaries
// @Produce      json
// @Success      200 {object}  dtos.Secretary
// @Router       /secretaries/{secretary_id} [put]
func (h *api) updateSecretary(w http.ResponseWriter, r *http.Request) {
	// Authorize access
	jwtData := GetJwtData(r)
	if !jwtData.HasAccess("secretary") {
		http.Error(w, "invalid access", http.StatusForbidden)
		return
	}

	// Collect query parameters, path parameters, and request body
	secretaryId, ok := GetAndParseUuidParam(w, r, "secretary_id")
	if !ok {
		return
	}

	if jwtData.Role == "secretary" && jwtData.UserID != secretaryId.String() {
		http.Error(w, "invalid access", http.StatusForbidden)
		return
	}

	var body dtos.SecretaryInfoBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.SecretaryInfoArgs{
		Name:      body.Name,
		Email:     body.Email,
		Phone:     body.Phone,
		Birthdate: body.Birthdate,
		Password:  body.Password,
		Cpf:       body.Cpf,
		Cnpj:      body.Cnpj,
		Update:    true,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	secretary, err := usecase.UpdateSecretary(rs, secretaryId, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetSecretary(secretary)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      Delete secretary
// @Security     ApiKeyAuth
// @Description  Delete a secretary by id
// @Tags         Secretaries
// @Produce      json
// @Success      204
// @Router       /secretaries/{secretary_id} [delete]
func (h *api) deleteSecretary(w http.ResponseWriter, r *http.Request) {
	// Authorize access
	jwtData := GetJwtData(r)
	if !jwtData.HasAccess() {
		http.Error(w, "invalid access", http.StatusForbidden)
		return
	}

	// Collect query parameters, path parameters, and request body
	secretaryId, ok := GetAndParseUuidParam(w, r, "secretary_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	err := usecase.DeleteSecretary(rs, secretaryId)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.NoContent(w, r)
}
