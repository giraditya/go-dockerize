FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && \
	go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/test-comstek ./cmd

ENTRYPOINT ["test-comstek"] 

EXPOSE 3000
