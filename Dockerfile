FROM golang:1.20.2-alpine3.17 AS build

WORKDIR /usr/src/app

COPY go.mod ./
COPY cmd ./cmd

RUN go build -ldflags="-s -w" -o /usr/local/bin/app cmd/startstartpage/main.go

FROM alpine:3.17

COPY --from=build /usr/local/bin/app /app

CMD ["/app"]
