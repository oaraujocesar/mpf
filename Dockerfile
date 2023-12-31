FROM golang:1.21

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY . .

RUN go mod tidy

EXPOSE 3333

ENV GOFLAGS=-buildvcs=false

ENTRYPOINT ["air"]
