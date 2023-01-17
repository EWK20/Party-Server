FROM golang:1.19.3-alpine3.16

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/app cmd/app/main.go

EXPOSE 4000

CMD ["./bin/app"]
