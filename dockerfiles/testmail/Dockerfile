FROM golang:1.21-alpine3.18 as build

WORKDIR /usr/src
COPY go.mod /usr/src/go.mod
COPY go.sum /usr/src/go.sum
RUN go mod download

COPY cmd /usr/src/cmd
COPY internal /usr/src/internal

WORKDIR /usr/src/cmd/testmail

RUN go build -o /usr/src/testmail


FROM alpine:3.18
USER 1000
WORKDIR /usr/app
COPY --from=build --chown=1000:1000 /usr/src/testmail /usr/bin/testmail
ENTRYPOINT [ "/usr/bin/testmail", "serve" ]