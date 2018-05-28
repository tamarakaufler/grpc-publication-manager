FROM alpine:3.7

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY ./author-cli /app

ENTRYPOINT ["./author-cli"]

CMD ["./author-cli"]
