build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o author-cli .
	docker build -t author-cli .

run:
	docker run --net=pm-net-bridge \
	--name=author-cli \
	--rm \
	-e SERVICE_HOST=author-service:50051 \
	author-cli
