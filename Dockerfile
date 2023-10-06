FROM golang:1.19 as buildbase

WORKDIR /go/src/gitlab.com/rarimo/identity/verify-proxy
COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /usr/local/bin/verify-proxy /go/src/gitlab.com/rarimo/identity/verify-proxy

FROM alpine:3.18.2

RUN apk --update add --no-cache musl libstdc++ gcompat libgomp ca-certificates

WORKDIR /

COPY --from=buildbase "/go/pkg/mod/github.com/iden3/wasmer-go@v0.0.1" "/go/pkg/mod/github.com/iden3/wasmer-go@v0.0.1"

COPY --from=buildbase /usr/local/bin/verify-proxy /usr/local/bin/verify-proxy

ENTRYPOINT ["verify-proxy"]
