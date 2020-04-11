FROM golang:1.13-alpine

ARG ARG_BOT_TOKEN

ENV BOT_TOKEN=$ARG_BOT_TOKEN

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]