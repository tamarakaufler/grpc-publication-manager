# To set up the PostgresSQL for use with the author-service microservice

## build the image (includes setting up microservice db user, the database and its table)
  - docker build -t author_postgres .

## run the Postgres application in a container
  - docker run --rm -e POSTGRES_PASSWORD=mypass -p 5432:5432 --name author_postgres author_postgres       (does not detach and provides insight into what is happening during the setup)
  - docker run --rm -e POSTGRES_PASSWORD=mypass -p 5432:5432 -d --name author_postgres author_postgres

  - docker build -t author_postgres . && docker run --rm -d -p 5432:5432 --name author_postgres author_postgres

## testing it works
  - docker run --rm -e POSTGRES_PASSWORD=mypass -p 5432:5432 --name author_postgres author_postgres
  - psql -h localhost -p 5432 -U postgres               (connecting locally from a terminal)
  - docker exec -it author_postgres /bin/bash             (entering the container for inspection to bash shell)
  - docker exec -it author_postgres psql -U postgres      (entering the container for inspection to psql shell)
