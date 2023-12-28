FROM golang:1.21.4 

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

RUN apt-get update && apt-get install -y tar
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY . .

RUN go mod tidy

EXPOSE 3333

ENV GOFLAGS=-buildvcs=false

ENTRYPOINT ["air"]
