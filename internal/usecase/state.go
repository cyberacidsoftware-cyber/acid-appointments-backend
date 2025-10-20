package usecase

import (
	"context"

	"backend/internal/infra"
)

type State interface {
	Queries() *infra.Queries
	Context() context.Context
}
