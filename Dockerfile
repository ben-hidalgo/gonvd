FROM golang:1.12.1-stretch

ENV GO111MODULE=on

#
WORKDIR /usr/local/gocache

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download


WORKDIR /usr/local/gonvd

COPY entrypoint.sh entrypoint.sh

#
#RUN go build ./...
#RUN go test ./...

EXPOSE 8000

CMD [ "bash", "entrypoint.sh" ]
