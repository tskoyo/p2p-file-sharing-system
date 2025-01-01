FROM golang:1.23.3

WORKDIR /app

# Pre-copy these files to leverage Docker cache
COPY go.mod go.sum ./

RUN go mod download && go mod verify

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.62.2

COPY . .

RUN go build -o /usr/local/bin/app ./cmd

EXPOSE 8080

CMD ["app"]
