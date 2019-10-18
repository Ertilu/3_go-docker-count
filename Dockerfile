FROM golang:latest

LABEL maintainer="Reza Raka Nugraha"

COPY go.mod go.sum ./

RUN mkdir /app

ADD . /app

ADD config /app/config

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD ["/app/main"]

ENTRYPOINT /app/main