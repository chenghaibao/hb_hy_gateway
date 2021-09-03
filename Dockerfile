FROM golang:1.16-alpine as builder

WORKDIR  /usr/src/app

COPY . .

RUN go build -ldflags "-s -w" -o hb main.go

CMD ["/usr/src/app/hb", "http"]