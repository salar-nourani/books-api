FROM golang:1.25-alpine AS builder

RUN apk add --no-cache  git make

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin/books-api .

COPY --from=builder /app/.env.example .env

EXPOSE 8080

CMD ["./books-api"]


