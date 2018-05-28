IMAGE_TAG=v1alpha1
QUAY_PASS?=biggestsecret

net:
	docker network create --driver bridge pm-net-bridge

protoc:
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/tamarakaufler/grpc-publication-manager/author-service \
    proto/author.proto
	
build: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o author-service -a -installsuffix cgo main.go \
	datastore.go store.go service.go token-service.go
	docker build -t quay.io/tamarakaufler/grpc-author-service:$(IMAGE_TAG) .
	docker login quay.io -u tamarakaufler -p $(QUAY_PASS)
	docker push quay.io/tamarakaufler/grpc-author-service:$(IMAGE_TAG)

dev: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o author-service -a -installsuffix cgo main.go \
	datastore.go store.go service.go token-service.go
	docker build -t quay.io/tamarakaufler/grpc-author-service:$(IMAGE_TAG) .

devrun:
	docker run \
	--name=author-service \
	--network=pm-net-bridge \
	-e DB_HOST=author-postgres:5432 \
	--rm \
	-p 50051:50051 \
	quay.io/tamarakaufler/grpc-author-service:$(IMAGE_TAG)

run:
	docker run \
	--name=author-service \
	--network=pm-net-bridge \
	-e DB_HOST=author-postgres:5432 \
	-e DB_USER=author_user \
	-e DB_PASS=authorpass \
	--rm \
	-p 50051:50051 \
	quay.io/tamarakaufler/grpc-author-service:$(IMAGE_TAG)
