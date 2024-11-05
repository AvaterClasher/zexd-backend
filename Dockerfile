FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o zexd ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/zexd .
EXPOSE 8080
CMD ["./zexd"]