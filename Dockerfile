FROM golang:1.17

RUN mkdir /tmp/app
ADD . /tmp/app
WORKDIR /tmp/app

RUN go build -o main .
CMD ["/tmp/app/main"]
EXPOSE 3000
