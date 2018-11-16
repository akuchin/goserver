FROM golang:1.8

COPY Server.go .

EXPOSE 9001

RUN go build Server.go
RUN ["chmod", "+x", "Server"]

CMD ["./Server"]