# grpc-publication-manager

An implementation of a system, publication-manager, through which authors can
have their books published. The microservice architecture was chosen for the 
project. There will be three microservices. Each one persisting data in a database.
Development and deployment takes advantage of containerization (Docker).

The system is a bit contrived, with the main purpose of examining and showing
how certain features (using JWT etc) and tasks (bootstapping databases, running
several microservices locally etc) can be implemented.

WIP

## Microservices

A set of microservices written in Go, using grpc communication.

Each microservice lives in its own directory, which includes:
  - codebase
  - Dockerfile
  - Makefile

### author-service

Stores user data in a Postgres database. Provides user authentication through
JWT (Javascript web token).

Dockerfile uses a Docker multi stage build feature (requires Docker 17.05 or higher
on the daemon and client). This reduces the size and build time of the final image
by building/compiling an application in one image and running it in a final one.  

## Requirements

 - grpc binary and libraries:
    - go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    - go get -u google.golang.org/grpc

 - Docker (https://docs.docker.com/install/#desktop)

## Bootstrapping

### Databases

Each microservice uses a database for persistent storage. Bootstrapping of the
database for testing purposes is described in the bootstapping section (directory).

### Networking

Setting up container networking is described in bootstrapping/network doc.

## Clients

Clients are used to test a microservice implementation. Client codebases are stored
in the cmd section.

## Credits
- https://ewanvalentine.io/microservices-in-golang-part-1/
- https://medium.com/@beld_pro/quick-tip-creating-a-postgresql-container-with-default-user-and-password-8bb2adb82342
