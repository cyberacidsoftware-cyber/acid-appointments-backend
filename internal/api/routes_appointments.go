package api

import (
	"encoding/json"
	"net/http"

	"backend/internal/api/dtos"
	"backend/internal/api/presenter"
	"backend/internal/usecase"

	"github.com/go-chi/render"
)

// @Summary      Get appointment
// @Security     ApiKeyAuth
// @Description  Get an appointment by id
// @Tags         Appointments
// @Produce      json
// @Success      200 {object}  dtos.Appointment
// @Router       /appointments/{appointment_id} [get]
func (h *api) getAppointment(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	appointmentID, ok := GetAndParseUuidParam(w, r, "appointment_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	appointment, err := usecase.GetAppointment(rs, appointmentID)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetAppointment(appointment)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      List appointments
// @Security     ApiKeyAuth
// @Description  Get a page of appointments with filters sorted by date and time
// @Tags         Appointments
// @Produce      json
// @Param        page      query int     false "Page"
// @Param        pageSize  query int     false "Page size"
// @Param        startDate query string  false "Start date"
// @Param        endDate   query string  false "End date"
// @Success      200 {object}  []dtos.Appointment
// @Router       /appointments [get]
func (h *api) getAppointments(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	page := ParseIntOrDefault(query.Get("page"), 0)
	pageSize := ParseIntOrDefault(query.Get("pageSize"), 10)
	startDate := query.Get("startDate")
	endDate := query.Get("endDate")
	serviceName := query.Get("serviceName")
	specialist := query.Get("specialist")
	customer := query.Get("customer")

	// Validate e execute the usecase
	args := usecase.ListAppointmentsArgs{
		PaginationArgs: usecase.PaginationArgs{
			Page:     page,
			PageSize: pageSize,
		},
		CountArgs: usecase.CountAppointmentsArgs{
			ServiceName:    serviceName,
			SpecialistName: specialist,
			CustomerName:   customer,
			StartDateRaw:   startDate,
			EndDateRaw:     endDate,
		},
	}

	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	appointments, err := usecase.ListAppointments(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetAppointments(appointments)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      Count appointments
// @Security     ApiKeyAuth
// @Description  Count appointments with filters
// @Tags         Appointments
// @Produce      json
// @Param        startDate 	query string  false "Start date"
// @Param        endDate   	query string  false "End date"
// @Param        serviceNameId	query string  false "Service Name ID"
// @Param        specialistId	query string  false "Specialist ID"
// @Param        customerId	query string  false "Customer ID"
// @Success      200 {object}  int
// @Router       /appointments/count [get]
func (h *api) countAppointments(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	startDate := query.Get("startDate")
	endDate := query.Get("endDate")
	serviceNameID := query.Get("serviceNameId")
	specialistID := query.Get("specialistId")
	customerID := query.Get("customerId")

	// Validate e execute the usecase
	args := usecase.CountAppointmentsArgs{
		ServiceName:    serviceNameID,
		SpecialistName: specialistID,
		CustomerName:   customerID,
		StartDateRaw:   startDate,
		EndDateRaw:     endDate,
	}

	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	count, err := usecase.CountAppointments(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.JSON(w, r, count)
	render.Status(r, http.StatusOK)
}

// @Summary      List appointments for calendar
// @Security     ApiKeyAuth
// @Description  Get appointments for calendar in a interval of date
// @Tags         Appointments
// @Produce      json
// @Param        year      query int     false "Year"
// @Param        startDate query string  false "Start date"
// @Param        endDate   query string  false "End date"
// @Success      200 {object}  []dtos.AppointmentCalendar
// @Router       /appointments/calendar [get]
func (h *api) getAppointmentsCalendar(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	year := ParseIntOrDefault(query.Get("year"), 0)
	startDate := query.Get("startDate")
	endDate := query.Get("endDate")

	// Validate e execute the usecase
	args := usecase.ListAppointmentsCalendarArgs{
		Year:         int(year),
		StartDate: startDate,
		EndDate:   endDate,
	}

	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	appointmentsCalendar, err := usecase.ListAppointmentsCalendar(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetAppointmentsCalendar(appointmentsCalendar)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      Count appointments for calendar
// @Security     ApiKeyAuth
// @Description  Get appointments count for calendar
// @Tags         Appointments
// @Produce      json
// @Param        year      query int     false "Year"
// @Param        startDate query string  false "Start date"
// @Param        endDate   query string  false "End date"
// @Success      200 {object}  []dtos.AppointmentCalendarCount
// @Router       /appointments/calendar/count [get]
func (h *api) getAppointmentsCalendarCount(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	query := r.URL.Query()
	year := ParseIntOrDefault(query.Get("year"), 0)
	startDate := query.Get("startDate")
	endDate := query.Get("endDate")

	// Validate e execute the usecase
	args := usecase.ListAppointmentsCalendarCountArgs{
		Year:         int(year),
		StartDate: startDate,
		EndDate:   endDate,
	}

	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	appointmentsCalendarCount, err := usecase.ListAppointmentsCalendarCount(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetAppointmentsCalendarCount(appointmentsCalendarCount)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      Create appointment
// @Security     ApiKeyAuth
// @Description  Create a new appointment for a customer and service
// @Tags         Appointments
// @Accept       json
// @Produce      json
// @Param        request body dtos.CreateAppointmentBody true "Appointments info"
// @Success      204 {object}  string
// @Router       /appointments [post]
func (h *api) createAppointment(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	var body dtos.CreateAppointmentBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.CreateAppointmentArgs{
		CustomerIDRaw: body.CustomerID,
		ServiceIDRaw:  body.ServiceID,
		DateRaw:       body.Date,
		TimeRaw:       body.Time,
		Status:        int32(usecase.AppointmentStatusPending),
	}

	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	id, err := usecase.CreateAppointment(rs, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.JSON(w, r, id.String())
	render.Status(r, http.StatusCreated)
}

// @Summary      Create appointment
// @Security     ApiKeyAuth
// @Description  Create a new appointment for a customer and service
// @Tags         Appointments
// @Accept       json
// @Produce      json
// @Param        request body dtos.UpdateAppointmentBody true "Appointments info"
// @Success      200 {object}  string
// @Router       /appointments/{appointment_id} [put]
func (h *api) updateAppointment(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	appointmentID, ok := GetAndParseUuidParam(w, r, "appointment_id")
	if !ok {
		return
	}

	var body dtos.UpdateAppointmentBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		InvalidJson(w)
		return
	}

	// Validate e execute the usecase
	args := usecase.UpdateAppointmentArgs{
		Date: body.Date,
		Time: body.Time,
		Status:  body.Status,
	}

	if err := args.Validate(); err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	rs := NewRequestState(h.q, r)

	appointment, err := usecase.UpdateAppointment(rs, appointmentID, args)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := appointment.ID.String()
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

// @Summary      Delete appointment
// @Security     ApiKeyAuth
// @Description  Delete an appointment by id
// @Tags         Appointments
// @Produce      json
// @Success      204
// @Router       /appointments/{appointment_id} [delete]
func (h *api) deleteAppointment(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body
	appointmentID, ok := GetAndParseUuidParam(w, r, "appointment_id")
	if !ok {
		return
	}

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	err := usecase.DeleteAppointment(rs, appointmentID)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	render.NoContent(w, r)
}
