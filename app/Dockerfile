FROM golang:1.19

WORKDIR /usr/local/src

COPY ./main.go ./

RUN go build -o /opt/app main.go

EXPOSE 18080

CMD ["/opt/app"]
