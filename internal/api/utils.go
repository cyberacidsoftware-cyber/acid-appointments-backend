package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func InvalidJson(w http.ResponseWriter) {
	http.Error(w, "invalid json", http.StatusBadRequest)
}

func SomethingWentWrong(w http.ResponseWriter, err error) {
	slog.Error("something went wrong", "error", err)
	http.Error(w, "something went wrong", http.StatusInternalServerError)
}

func GetAndParseUuidParam(w http.ResponseWriter, r *http.Request, paramName string) (uuid.UUID, bool) {
	rawId := chi.URLParam(r, paramName)
	return ParseUuidParam(w, rawId, paramName)
}

func ParseUuidParam(w http.ResponseWriter, rawID string, paramName string) (uuid.UUID, bool) {
	id, err := uuid.Parse(rawID)
	if err != nil {
		slog.Warn("invalid param", paramName, rawID)
		http.Error(w, "invalid uuid param: "+paramName, http.StatusBadRequest)
		return uuid.Nil, false
	}
	return id, true
}

func ParseIntOrDefault(text string, defaultValue int32) int32 {
	if text == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(text)
	if err != nil {
		return defaultValue
	}
	return int32(value)
}
