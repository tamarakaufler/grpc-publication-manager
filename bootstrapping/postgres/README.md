# To set up the PostgresSQL for use with the author-service microservice

## build the image (includes setting up microservice db user, the database and its table)
  - docker build -t author-postgres .

## run the Postgres application in a container
  - docker run --rm -e POSTGRES_PASSWORD=mypass -p 5432:5432 --name author-postgres author-postgres       (does not detach and provides insight into what is happening during the setup)
  - docker run --rm -e POSTGRES_PASSWORD=mypass -p 5432:5432 -d --name author-postgres author-postgres

  - docker build -t author-postgres . && docker run --rm -d -p 5432:5432 --name author-postgres author-postgres

## testing the setup works
  - docker run --rm -e POSTGRES_PASSWORD=mypass -p 5432:5432 --name author-postgres author-postgres
  - psql -h localhost -p 5432 -U postgres                                 (connecting locally from a terminal)
  - docker exec -it author-postgres /bin/bash                             (entering the container to bash shell)
  - docker exec -it author-postgres psql publication_manager postgres     (entering the container to psql shell)
  - docker exec -it author-postgres psql publication_manager author_user  (entering the container to psql shell)
