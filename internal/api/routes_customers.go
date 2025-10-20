package api

import (
	"encoding/json"
	"net/http"

	"backend/internal/api/dtos"
	"backend/internal/api/presenter"
	"backend/internal/usecase"

	"github.com/go-chi/render"
)

// @Summary      Get customer
// @Security     ApiKeyAuth
// @Description  Get a customer by id
// @Tags         Customers
// @Produce      json
// @Success      200 {object}  dtos.Customer
// @Router       /customers/{customer_id} [get]
func (h *api) getCustomer(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	customerId, ok := GetAndParseUuidParam(w, r, "customer_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	customer, err := usecase.GetCustomer(rs, customerId)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetCustomer(customer)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      List customers
// @Security     ApiKeyAuth
// @Description  Get a page of customers with filters
// @Tags         Customers
// @Produce      json
// @Param        page      query int     false "Page"
// @Param        pageSize  query int     false "Page size"
// @Param        name      query string  false "Name"
// @Param        cpf       query string  false "CPF"
// @Param        phone     query string  false "Phone"
// @Success      200 {object}  []dtos.Customer
// @Router       /customers [get]
func (h *api) getCustomers(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	page := ParseIntOrDefault(query.Get("page"), 0)
	pageSize := ParseIntOrDefault(query.Get("pageSize"), 10)
	name := query.Get("name")
	cpf := query.Get("cpf")
	phone := query.Get("phone")

	// Validate e execute the usecase
	args := usecase.ListCustomersArgs{
		PaginationArgs: usecase.PaginationArgs{
			PageSize: pageSize,
			Page:     page,
		},
		CountArgs: usecase.CountCustomersArgs{
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

	customers, err := usecase.ListCustomers(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetCustomers(customers)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      Count customers
// @Security     ApiKeyAuth
// @Description  Count customers with filters
// @Tags         Customers
// @Produce      json
// @Param        name      query string  false "Name"
// @Param        cpf       query string  false "CPF"
// @Param        phone     query string  false "Phone"
// @Success      200 {object}  int
// @Router       /customers/count [get]
func (h *api) countCustomers(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	name := query.Get("name")
	cpf := query.Get("cpf")
	phone := query.Get("phone")

	// Validate e execute the usecase
	args := usecase.CountCustomersArgs{
		Cpf:   cpf,
		Phone: phone,
		Name:  name,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	count, err := usecase.CountCustomers(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.JSON(w, r, count)
	render.Status(r, http.StatusOK)
}

// @Summary      Create customer
// @Security     ApiKeyAuth
// @Description  Create an new customer
// @Tags         Customers
// @Produce      json
// @Param 		  data body dtos.CustomerInfoBody true "Customer information"
// @Success      200 {object}  string
// @Router       /customers [post]
func (h *api) createCustomer(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	var body dtos.CustomerInfoBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.CustomerInfoArgs{
		Name:      body.Name,
		Email:     body.Email,
		Phone:     body.Phone,
		Birthdate: body.Birthdate,
		Cpf:       body.Cpf,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	customer, err := usecase.CreateCustomer(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := customer.ID.String()
	render.JSON(w, r, response)
	render.Status(r, http.StatusCreated)
}

// @Summary      Update customer
// @Security     ApiKeyAuth
// @Description  Update a customer by id
// @Tags         Customers
// @Produce      json
// @Success      200 {object}  dtos.Customer
// @Router       /customers/{customer_id} [put]
func (h *api) updateCustomer(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	customerId, ok := GetAndParseUuidParam(w, r, "customer_id")
	if !ok {
		return
	}
	var body dtos.CustomerInfoBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.CustomerInfoArgs{
		Name:      body.Name,
		Email:     body.Email,
		Phone:     body.Phone,
		Birthdate: body.Birthdate,
		Cpf:       body.Cpf,
	}
	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	customer, err := usecase.UpdateCustomer(rs, customerId, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetCustomer(customer)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      Delete customer
// @Security     ApiKeyAuth
// @Description  Delete a customer by id
// @Tags         Customers
// @Produce      json
// @Success      204
// @Router       /customers/{customer_id} [delete]
func (h *api) deleteCustomer(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	customerId, ok := GetAndParseUuidParam(w, r, "customer_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	err := usecase.DeleteCustomer(rs, customerId)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.NoContent(w, r)
}
