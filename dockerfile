FROM golang:1.21 AS dep
RUN apt update
RUN apt -y install build-essential libsqlite3-dev

WORKDIR /go/src
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .

RUN sh ./build.sh

ENTRYPOINT [ "sh", "./output/boot.sh" ]