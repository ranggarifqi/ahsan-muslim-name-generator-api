# builder
FROM golang:1.16.5-alpine3.13

RUN apk update && \
    apk upgrade && \
    apk add make

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["make", "run"]