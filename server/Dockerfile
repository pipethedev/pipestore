FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

ENV GO111MODULE=on

RUN go clean -modcache

RUN go mod download

COPY . .

ARG PORT

ENV PORT=57710

RUN go build -o tmp/main .

RUN go mod tidy

EXPOSE ${PORT}

CMD ["./tmp/main"]