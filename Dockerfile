FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go build -o binary

ENTRYPOINT ["/app/binary"]
