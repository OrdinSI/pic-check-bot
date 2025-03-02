FROM golang:1.23.4 AS builder

ENV CGO_ENABLED=0

WORKDIR /bot
ADD . /bot

RUN go build -o bot ./cmd/bot


FROM alpine:3.20

COPY --from=builder /bot/bot  /srv/bot

WORKDIR /srv

CMD ["/srv/bot"]
