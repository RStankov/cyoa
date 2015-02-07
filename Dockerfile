FROM ubuntu:trusty

RUN apt-get update
RUN apt-get install -y build-essential git wget curl

ENV GOPATH /go
ENV PATH /go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games
ENV APPPATH /go/src/github.com/rstankov/choose_your_own_adventure

RUN wget -qO- http://golang.org/dl/go1.4.1.linux-amd64.tar.gz | tar -C /usr/local -xzf -

RUN mkdir -p /go/bin
RUN mkdir -p /go/pkg
RUN mkdir -p /go/src

RUN go get github.com/pilu/fresh

ADD . $APPPATH

WORKDIR $APPPATH
