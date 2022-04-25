FROM golang:1.17.8-alpine3.15 AS builder

ENV GO111MODULE=on
WORKDIR $GOPATH/src

RUN mkdir -p app
WORKDIR $GOPATH/src/app

RUN apk update && apk add libc-dev && apk add gcc && apk add make

COPY . $GOPATH/src/app/

RUN go get

RUN go build -a -o app .

FROM alpine:3.15.0 AS production

COPY --from=builder /go/src/app/app /usr/local/bin/app

CMD ["app"]