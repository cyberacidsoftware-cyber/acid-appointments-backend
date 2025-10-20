package gen

//go:generate go run ./cmd/wenv/main.go tern migrate --migrations ./internal/infra/migrations/ --config ./internal/infra/migrations/tern.conf
//go:generate docker run --rm -v ".:/src" -w /src kjconroy/sqlc generate -f "./internal/infra/sqlc.yaml"
//ge:generate swag init --dir ./cmd/api/,./internal/api/ --output ./internal/api/docs
