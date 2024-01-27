FROM golang:1.21-alpine3.18 as build

WORKDIR /usr/src
COPY go.mod /usr/src/go.mod
COPY go.sum /usr/src/go.sum
RUN go mod download

COPY cmd /usr/src/cmd
COPY internal /usr/src/internal

WORKDIR /usr/src/cmd/s3-tester

RUN go build -o /usr/src/s3-tester


FROM alpine:3.18
USER 1000
WORKDIR /usr/app
COPY --from=build --chown=1000:1000 /usr/src/s3-tester /usr/bin/s3-tester
ENTRYPOINT [ "/usr/bin/s3-tester", "serve" ]