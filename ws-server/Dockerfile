FROM golang

RUN mkdir /arcadia-typerace-server

ADD . /arcadia-typerace-server

WORKDIR /arcadia-typerace-server

RUN go mod vendor

RUN go build -o socket-server .

CMD ["/arcadia-typerace-server/socket-server"]