FROM golang:1.18-alpine AS builder

COPY . /github.com/ksusonic/relax-tg-bot/
WORKDIR /github.com/ksusonic/relax-tg-bot/

RUN go mod download
RUN go build -o ./bin/bot cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/ksusonic/relax-tg-bot/bin/bot .
#COPY --from=0 /github.com/ksusonic/relax-tg-bot/configs configs/

#EXPOSE 80

CMD ["./bot"]
