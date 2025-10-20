package cli

import (
	"context"
	"log"

	"backend/internal/env"
	"backend/internal/infra"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type State struct {
	pool *pgxpool.Pool
	tx   pgx.Tx
	ctx  context.Context
	q    *infra.Queries
}

func NewState() *State {
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, env.GetDatabaseConnection())
	if err != nil {
		log.Fatal(err)
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	q := infra.New(pool)

	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}

	q = q.WithTx(tx)
	return &State{pool, tx, ctx, q}
}

func (s *State) Queries() *infra.Queries {
	return s.q
}

func (s *State) Context() context.Context {
	return s.ctx
}

func (c *State) Close() {
	c.pool.Close()
}

func (s *State) CreateAdmin(params infra.CreateAdminParams) (uuid.UUID, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}
	params.Password = string(hashedPassword)
	return s.q.CreateAdmin(s.ctx, params)
}

func (s *State) GetIdentity(email, password string) (infra.GetIdentityByEmailRow, error) {
	identity, err := s.q.GetIdentityByEmail(s.ctx, email)
	if err == nil {
		return identity, nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(identity.Password), []byte(password))
	return identity, err
}
