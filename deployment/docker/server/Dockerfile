FROM golang:latest
WORKDIR /usr/src/app

COPY . .

RUN cd ./cmd/server && go build .

CMD cmd/server/server