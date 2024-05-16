# Build stage
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]