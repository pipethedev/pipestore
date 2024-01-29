FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

ENV GO111MODULE=on

RUN go clean -modcache

RUN go mod download

COPY . .

ARG PORT

ENV PORT=${PORT}

RUN go build -o tmp/main .

# Clean up the Go module and ensure dependencies are correctly tracked
RUN go mod tidy

# Expose the port your application will listen on
EXPOSE ${PORT}

# Command to run the application
CMD ["./tmp/main"]