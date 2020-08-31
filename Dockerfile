FROM golang:1.14.3 AS build-env
WORKDIR /go/src/github.com/lucasreed/datetimeapi/

ARG VERSION=dev

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o dtapi -ldflags "-X main.version=${VERSION} -s -w"

FROM alpine:3.12
RUN apk --no-cache --update add ca-certificates tzdata && update-ca-certificates

USER nobody
COPY --from=build-env /go/src/github.com/lucasreed/datetimeapi /

WORKDIR /opt/app
ENTRYPOINT ["/dtapi"]
