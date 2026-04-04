FROM golang:1.26.1-alpine AS builder

RUN apk add --no-cache git

ENV CGO_ENABLED=0

WORKDIR /app

COPY . .

RUN go mod verify

RUN go mod download

RUN go build -o /weatherie ./cmd/weatherie 

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /weatherie .
COPY --from=builder /app/.env .

CMD ["./weatherie"]
