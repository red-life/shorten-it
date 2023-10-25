FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY ./cmd/ ./cmd
COPY ./internal/ ./internal/
COPY ./pkg/ ./pkg/
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o shorten-it ./cmd

FROM alpine:3.14
WORKDIR /app
COPY ./.env ./.env
COPY --from=builder /src/shorten-it ./shorten-it
CMD ["./shorten-it"]