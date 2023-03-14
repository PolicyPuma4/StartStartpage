FROM golang:1.20.2-alpine3.17 AS build

WORKDIR /usr/src/startstartpage

COPY go.mod ./
COPY main.go ./

RUN go build -o /usr/local/bin/startstartpage cmd/startstartpage/main.go

FROM alpine:3.17

COPY --from=build /usr/local/bin/startstartpage /startstartpage

CMD ["/startstartpage"]
