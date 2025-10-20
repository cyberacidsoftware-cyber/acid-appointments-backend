package api

import (
	"net/http"

	"backend/internal/api/presenter"
	"backend/internal/usecase"

	"github.com/go-chi/render"
)

func (h *api) getServiceGroups(w http.ResponseWriter, r *http.Request) {
	// Collect query parameters, path parameters, and request body

	// Validate e execute the usecase
	rs := NewRequestState(h.q, r)

	specializations, serviceNames, err := usecase.ListServiceGroups(rs)
	if err != nil {
		presenter.UsecaseError(w, err)
		return
	}

	// Format the response
	response := presenter.GetServiceGroups(specializations, serviceNames)
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}
