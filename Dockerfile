FROM golang:1.9.2 as BUILD

WORKDIR /go/src/github.com/reynn/merge-kubeconfig

COPY . .

RUN go get -u -v github.com/Masterminds/glide \
    && glide install \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /tmp/merge-config .

FROM alpine:3.7

COPY --from=BUILD /tmp/merge-config /usr/bin/merge-config

ENTRYPOINT ["/usr/bin/merge-config"]