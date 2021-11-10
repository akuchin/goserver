FROM alpine:latest

COPY Server .

RUN apk add --no-cache bash

EXPOSE 9001

RUN ["chmod", "+x", "Server"]

CMD ["./Server"]