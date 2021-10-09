FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/echenim/distributed-web/src
RUN cd /build && git clone https://github.com/echenim/distributed-web.git

RUN cd /build/distributed-web/src && go build

EXPOSE 30000

ENTRYPOINT [ "/build/distributed-web/src/main" ]