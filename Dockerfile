FROM golang:1.13.7 AS builder

WORKDIR /build/waitforit

COPY main.go .
RUN GOOS=linux GOARCH=amd64 go build .


FROM debian:stretch-slim
# don't place it into $GOPATH/bin because Drone mounts $GOPATH as volume
COPY --from=builder /build/waitforit/waitforit /usr/bin/

CMD ["waitforit"]