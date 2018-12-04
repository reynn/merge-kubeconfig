FROM golang:1.11.2-alpine3.8 as BUILD

WORKDIR /go/merge

COPY . .

RUN apk add --update --no-cache git \
    && GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /tmp/merge-config .

FROM alpine:3.8

COPY --from=BUILD /tmp/merge-config /usr/bin/merge-config

ENTRYPOINT ["/usr/bin/merge-config"]
