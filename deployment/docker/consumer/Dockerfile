FROM golang:latest
WORKDIR /usr/src/app

COPY . .

RUN cd ./cmd/consumer && go build .

CMD cmd/consumer/consumer