FROM golang

RUN mkdir /arcadia-typerace-server

ADD . /arcadia-typerace-server

WORKDIR /arcadia-typerace-server

RUN go build -o main .

CMD ["/typerace/socket-server"]