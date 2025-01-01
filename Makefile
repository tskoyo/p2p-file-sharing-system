.PHONY: build run lint test

build:
	docker build -t p2p-dev-env .

run:
	docker run -it --rm --name my-p2p-app -v "$(shell pwd):/app" -w /app p2p-dev-env air

lint:
	docker run -it --rm -v "$(shell pwd):/app" -w /app -e GOFLAGS="-buildvcs=false" p2p-dev-env golangci-lint run

test:
	docker run -it --rm -v "$(shell pwd):/app" -w /app p2p-dev-env go test -v ./...