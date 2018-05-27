# Dependencies of the microservices

All of the services need a database to persist information:

author-service: Postgres
publisher-service: MongoDB
book-service: MongoDB

author-service needs NATS messaging for sending emails

