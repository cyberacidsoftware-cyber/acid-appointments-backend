package api

import (
	"context"
	"net/http"

	"backend/internal/infra"
)

type RequestState struct {
	q *infra.Queries
	r *http.Request
}

func NewRequestState(q *infra.Queries, r *http.Request) *RequestState {
	return &RequestState{q: q, r: r}
}

func (rs *RequestState) Queries() *infra.Queries {
	return rs.q
}

func (rs *RequestState) Context() context.Context {
	return rs.r.Context()
}
