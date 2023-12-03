FROM golang:1.21.4

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

RUN apt-get update && apt-get install -y tar
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.darwin-arm64.tar.gz | tar xvz
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

RUN go mod tidy

EXPOSE 3333

ENV GOFLAGS=-buildvcs=false

ENTRYPOINT ["air"]
