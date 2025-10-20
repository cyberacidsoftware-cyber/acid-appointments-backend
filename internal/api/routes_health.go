package api

import (
	"net/http"

	"github.com/go-chi/render"
)

// @Summary     Health check
// @Description Check if the service is healthy
// @Tags        API
// @Produce     json
// @Success     200 {object}  string
// @Router      /health [get]
func (h *api) health(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body

	// Validate e execute the usecase

	// Format the response
	response := "Healthy!"
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}
