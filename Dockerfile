FROM golang:1.23.4-alpine

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o /app/gotest /app/cmd/bin/main.go

EXPOSE 8000

CMD ["/app/gotest"]