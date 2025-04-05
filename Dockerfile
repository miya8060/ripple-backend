FROM golang:1.24.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

# air
RUN go install github.com/air-verse/air@latest

# sqlc
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# golang-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

EXPOSE 8080

CMD ["air"]