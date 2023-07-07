FROM golang:1.20-alpine3.16 as golang-build

WORKDIR /usr/src
COPY go.mod /usr/src/go.mod
COPY go.sum /usr/src/go.sum
RUN go mod download

COPY cmd /usr/src/cmd
COPY internal /usr/src/internal

WORKDIR /usr/src/cmd/testmail

RUN go build -o /usr/src/testmail


FROM alpine:3.16 as application
USER 1000
WORKDIR /usr/app
COPY --from=golang-build --chown=1000:1000 /usr/src/testmail /usr/bin/testmail
ENTRYPOINT [ "/usr/bin/testmail serve" ]