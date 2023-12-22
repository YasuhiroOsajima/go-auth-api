FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o go-auth-api /app/cmd/main.go


FROM golang:latest
WORKDIR /app
COPY --from=builder /app/go-auth-api .

EXPOSE 9999
CMD ["/app/go-auth-api"]
