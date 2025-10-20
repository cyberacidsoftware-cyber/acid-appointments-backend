package api

import (
	"net/http"

	_ "backend/internal/api/docs"
	"backend/internal/env"
	"backend/internal/infra"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type api struct {
	q    *infra.Queries
	pool *pgxpool.Pool
	auth *jwtauth.JWTAuth
}

// @title			Appointsments API
// @version		1.0
// @description	This is an API for managing appointments
// @termsOfService	http://swagger.io/terms/
// @contact.name	Daniel A. R. Farina
// @contact.url	http://www.github.com/danielbom
// @contact.email	daniel.rodrigues.45g@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host		localhost:8080
// @BasePath	/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						"Authorization"
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func NewApi(pool *pgxpool.Pool, auth *jwtauth.JWTAuth) http.Handler {
	q := infra.New(pool)
	h := &api{q, pool, auth}

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/health", h.health)

	r.Post("/auth/login", h.authLogin)

	r.With(h.JWT).Post("/auth/refresh", h.authRefresh)
	r.With(h.JWT).Get("/auth/me", h.authMe)

	r.With(h.JWT).Get("/appointments", h.getAppointments)
	r.With(h.JWT).Get("/appointments/count", h.countAppointments)
	r.With(h.JWT).Get("/appointments/calendar", h.getAppointmentsCalendar)
	r.With(h.JWT).Get("/appointments/calendar/count", h.getAppointmentsCalendarCount)
	r.With(h.JWT).Get("/appointments/{appointment_id}", h.getAppointment)
	r.With(h.JWT).Post("/appointments", h.createAppointment)
	r.With(h.JWT).Put("/appointments/{appointment_id}", h.updateAppointment)
	r.With(h.JWT).Delete("/appointments/{appointment_id}", h.deleteAppointment)

	r.With(h.JWT).Get("/customers", h.getCustomers)
	r.With(h.JWT).Get("/customers/count", h.countCustomers)
	r.With(h.JWT).Get("/customers/{customer_id}", h.getCustomer)
	r.With(h.JWT).Post("/customers", h.createCustomer)
	r.With(h.JWT).Put("/customers/{customer_id}", h.updateCustomer)
	r.With(h.JWT).Delete("/customers/{customer_id}", h.deleteCustomer)

	r.With(h.JWT).Get("/secretaries", h.getSecretaries)
	r.With(h.JWT).Get("/secretaries/count", h.countSecretaries)
	r.With(h.JWT).Get("/secretaries/{secretary_id}", h.getSecretary)
	r.With(h.JWT).Post("/secretaries", h.createSecretary)
	r.With(h.JWT).Put("/secretaries/{secretary_id}", h.updateSecretary)
	r.With(h.JWT).Delete("/secretaries/{secretary_id}", h.deleteSecretary)

	r.With(h.JWT).Get("/specialists", h.listSpecialists)
	r.With(h.JWT).Get("/specialists/count", h.countSpecialists)
	r.With(h.JWT).Get("/specialists/{specialist_id}", h.getSpecialist)
	r.With(h.JWT).Get("/specialists/{specialist_id}/appointments", h.getSpecialistAppointments)
	r.With(h.JWT).Get("/specialists/{specialist_id}/specializations", h.getSpecialistSpecializations)
	r.With(h.JWT).Get("/specialists/{specialist_id}/services", h.getSpecialistServices)
	r.With(h.JWT).Get("/specialists/{specialist_id}/services/{service_id}", h.getSpecialistService)
	r.With(h.JWT).Post("/specialists", h.createSpecialist)
	r.With(h.JWT).Put("/specialists/{specialist_id}", h.updateSpecialist)
	r.With(h.JWT).Delete("/specialists/{specialist_id}", h.deleteSpecialist)

	r.With(h.JWT).Get("/specializations", h.getSpecializations)
	r.With(h.JWT).Post("/specializations", h.createSpecialization)
	r.With(h.JWT).Put("/specializations/{specialization_id}", h.updateSpecialization)
	r.With(h.JWT).Delete("/specializations/{specialization_id}", h.deleteSpecialization)

	r.With(h.JWT).Get("/service-groups", h.getServiceGroups)

	r.With(h.JWT).Get("/services", h.getServices)
	r.With(h.JWT).Get("/services/count", h.countServices)
	r.With(h.JWT).Get("/services/{service_id}", h.getService)
	r.With(h.JWT).Post("/services", h.createService)
	r.With(h.JWT).Put("/services/{service_id}", h.updateService)
	r.With(h.JWT).Delete("/services/{service_id}", h.deleteService)

	r.With(h.JWT).Get("/services-available", h.getServicesAvailable)
	// r.With(h.JWT).Get("/services-available/count", h.countAvailableServices)
	r.With(h.JWT).Get("/services-available/{service_id}", h.getServiceAvailable)
	r.With(h.JWT).Post("/services-available", h.createServiceAvailable)
	r.With(h.JWT).Put("/services-available/{service_id}", h.updateServiceAvailable)
	r.With(h.JWT).Delete("/services-available/{service_id}", h.deleteServiceAvailable)

	if env.Get(env.APP_ENVIRONMENT) == "TEST" {
		r.Get("/test/stats", h.getTestStats)
		r.Get("/test/init", h.testInit)
	}

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	return r
}
