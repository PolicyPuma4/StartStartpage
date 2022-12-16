FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod main.go ./

RUN go build -o /usr/local/bin/app ./...

EXPOSE 3000

CMD ["app"]
