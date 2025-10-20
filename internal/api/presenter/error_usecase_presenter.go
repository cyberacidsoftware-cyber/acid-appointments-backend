package presenter

import (
	"backend/internal/usecase"
	"log/slog"
	"net/http"
)

func UsecaseError(w http.ResponseWriter, err *usecase.UsecaseError) {
	if err.Kind == usecase.ErrorKindUnexpected {
		slog.Error("something went wrong", "error", err.Error)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	if err.Kind == usecase.ErrorKindAuth {
		slog.Info("auth error", "error", err.Error)
		http.Error(w, "invalid credentials", http.StatusBadRequest)
		return
	}
	if err.Kind == usecase.ErrorKindInvalidArgument {
		slog.Info("invalid argument", "error", err.Error)
		http.Error(w, "invalid argument: "+err.Error.Error(), http.StatusBadRequest)
		return
	}
	http.Error(w, err.Error.Error(), http.StatusBadRequest)
}
