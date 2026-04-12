FROM golang:1.26.1-alpine AS builder

RUN apk add --no-cache git

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg/mod \
  go build -v -installsuffix cgo -o /weatherie ./cmd/weatherie

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /weatherie .

CMD ["./weatherie"]
