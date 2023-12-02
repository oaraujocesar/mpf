FROM golang:1.21.4

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

RUN apt-get update && apt-get install -y tar
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.darwin-arm64.tar.gz | tar xvz

COPY . .

RUN go mod tidy

EXPOSE 3333

ENTRYPOINT ["air"]
