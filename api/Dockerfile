
FROM golang:1.12-stretch

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY . /usr/src/app

RUN go build

CMD ["/usr/src/app/api", "prod"]
