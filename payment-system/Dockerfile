FROM golang:1.23.4
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build cmd/paymentsystem/main.go

CMD ["./main"]
