FROM golang:1.25.1-alpine

RUN apk add --no-cache git curl && mkdir /app

WORKDIR /app

COPY go.* ./
COPY internal ./internal
COPY cmd/fiber_app ./cmd/fiber_app

RUN go mod download

RUN go build -o ./fiber ./cmd/fiber_app/main.go

ENTRYPOINT [ "./fiber" ]