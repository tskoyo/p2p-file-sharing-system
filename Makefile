include .env
export $(shell sed 's/=.*//' .env)

.PHONY: build run lint test

build:
	docker build -t p2p-dev-env .

run-server:
	docker run -it --rm --name my-p2p-app --network p2p-network \
	 -v "$(shell pwd):/app" \
	 -w /app \
	 p2p-dev-env \
	 sh -c "\
	 	COMMAND=start-server \
		ID=${ID} \
		SERVER_ADDRESS=${SERVER_ADDRESS} \
		SERVER_PORT=${SERVER_PORT} \
		CLIENT_PORT=${CLIENT_PORT} \
		PEER_ADDRESS=${PEER_ADDRESS} \
		PEER_PORT=${PEER_PORT} air"

run-client:
	docker run -it --rm --name my-p2p-client --network p2p-network \
	 -v "$(shell pwd):/app" \
	 -w /app \
	 p2p-dev-env \
	 sh -c "\
	 	COMMAND=connect-to-peer \
		ID=${ID} \
		SERVER_ADDRESS=0 \
		SERVER_PORT=0 \
		CLIENT_PORT=${CLIENT_PORT} \
		PEER_ADDRESS=${PEER_ADDRESS} \
		PEER_PORT=${PEER_PORT} air"

lint:
	docker run -it --rm -v "$(shell pwd):/app" -w /app -e GOFLAGS="-buildvcs=false" p2p-dev-env golangci-lint run

test:
	docker run -it --rm -v "$(shell pwd):/app" -w /app p2p-dev-env go test -v ./...