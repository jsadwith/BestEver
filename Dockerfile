FROM golang:1.7.4

RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /go/src/github.com/jsadwith/BestEver
COPY . /go/src/github.com/jsadwith/BestEver

RUN go build .

EXPOSE 80

CMD ["/go/src/github.com/KargoGlobal/jsadwith/BestEver"]