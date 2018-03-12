FROM golang:1.10.0-alpine3.7 as BUILD

WORKDIR /go/src/github.com/reynn/merge-kubeconfig

COPY . .

RUN apk add --updaate --no-cache git \
    && go get -u -v github.com/golang/dep/cmd/dep \
    && dep ensure \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /tmp/merge-config .

FROM alpine:3.7

COPY --from=BUILD /tmp/merge-config /usr/bin/merge-config

ENTRYPOINT ["/usr/bin/merge-config"]