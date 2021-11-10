FROM ubuntu:latest

COPY Server .

EXPOSE 9001

RUN ["chmod", "+x", "Server"]

CMD ["./Server"]