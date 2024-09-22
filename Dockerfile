FROM golang:1.23.1-alpine3.20

WORKDIR /app

RUN go install github.com/air-verse/air@latest

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY . .

RUN ln .env cmd/server/.env

RUN go mod tidy

EXPOSE 3333

ENV GOFLAGS=-buildvcs=false

ENTRYPOINT ["air"]
