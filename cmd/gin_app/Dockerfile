FROM golang:alpine

RUN apk add --no-cache git curl && mkdir /app

WORKDIR /app

COPY go.* ./
COPY internal ./internal
COPY cmd/gin_app ./cmd/gin_app

RUN go mod download

RUN go build -o ./gin ./cmd/gin_app/main.go

ENTRYPOINT [ "./gin" ]