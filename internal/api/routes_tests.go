package api

import (
	"net/http"

	"backend/internal/api/presenter"
	"backend/internal/usecase"

	"github.com/go-chi/render"
)

func (h *api) getTestStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message":  "Environment: TEST",
		"database": h.pool.Config().ConnString(),
	}
	if response["database"] == "" {
		response["database"] = "<empty>"
	}
	render.JSON(w, r, response)
	render.Status(r, http.StatusOK)
}

func (h *api) testInit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := h.pool.Begin(ctx)
	if err != nil {
		http.Error(w, "failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback(ctx)

	qTx := h.q.WithTx(tx)
	rs := NewRequestState(qTx, r)

	ucErr := usecase.ResetDb(rs)
	if ucErr != nil {
		presenter.UsecaseError(w, ucErr)
		return
	}

	_, ucErr = usecase.CreateAdmin(rs, usecase.CreateAdminArgs{
		Name:     "Admin Test",
		Email:    "admin@test.com",
		Password: "123mudar",
	})

	if ucErr != nil {
		presenter.UsecaseError(w, ucErr)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		http.Error(w, "failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("System initialized to be tested"))
}
