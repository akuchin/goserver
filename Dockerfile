FROM alpine:latest

COPY Server .

EXPOSE 9001

RUN ["chmod", "+x", "Server"]

ENTRYPOINT ["./Server"]