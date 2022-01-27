FROM golang:1.17 AS build

WORKDIR $GOPATH/src/go-sample

# prime the modules into their own layer for caching
ENV GO111MODULE on
RUN go get github.com/go-delve/delve/cmd/dlv

COPY go.mod go.sum ./
RUN go mod download

COPY ../.. .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-sample

FROM ubuntu:latest

COPY --from=build /go/bin/dlv /usr/local/bin/
COPY --from=build /go/src/go-sample/go-sample /usr/local/bin/

ENV HOST_PORT "80"

EXPOSE 80 40003
CMD ["/usr/local/bin/go-sample"]
# CMD ["/usr/local/bin/dlv","--listen=:40003","--headless=true","--api-version=2","exec","/usr/local/bin/go-sample"]
