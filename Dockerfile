### Step 1: Baixar dependenciar e compilar o binario
FROM golang:1.23-alpine as stage1

WORKDIR /app

# Copia o go.mod e faz o download das dependencias.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o api

FROM scratch

WORKDIR /app

COPY --from=stage1 /app/cmd/api/api .

ENTRYPOINT ["/app/api"]
