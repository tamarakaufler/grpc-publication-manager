# Networking

## Requirements

    - Individual services need to connect to their respective databases
    - Some services need to talk to other service(s)

## Local development

Create a custom bridge and connect running containers to it.
Run databases in named containers and use their container name as their hostname.

a) custom network

    docker network create --driver bridge pm-net-bridge

b) postgres database

    docker run --network=pm-net-bridge --rm -d -e POSTGRES_PASSWORD=mypass -p 5432:5432 -d --name author-postgres author-postgres

c)
    ca) start a container connected to the sustom bridge

    docker run --network=pm-net-bridge --rm -d -e DB_HOST=author-postgres:5432 -e DB_USER=author -e DB_PASSWORD=yyyy -p 50051:50051 --name author-service quay.io/tamarakaufler/grpc-author-service:v1alpha1

    OR

    ca) connect a running container to the sustom bridge

    docker network connect pm-net-bridge author-postgres        (to connect a running container)
    docker network connect pm-net-bridge author-service
