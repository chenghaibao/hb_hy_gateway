FROM golang:1.16-alpine as builder

WORKDIR /usr/src/app

COPY . .

RUN go build -ldflags "-s -w" -o gateway main.go

FROM alpine:3.11 as runtime

RUN apk add --no-cache tzdata ca-certificates \
 && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
 && apk del tzdata \
 && rm -Rf /var/cache/apk/*

COPY --from=builder /usr/src/app/gateway /usr/local/bin/

EXPOSE 9002

WORKDIR /root

CMD ["gateway http"]