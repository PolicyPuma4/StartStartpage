FROM alpine:3.17 AS build

RUN apk update
RUN apk upgrade
RUN apk add --update go

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o /ssp

FROM alpine:3.17

WORKDIR /

COPY --from=build /ssp /ssp

CMD ["/ssp"]
