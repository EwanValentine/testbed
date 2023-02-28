FROM golang:1.19-alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git openssh-client openssh-server

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy && go get ./...

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service .


# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/service .

ENTRYPOINT ["./service"]
