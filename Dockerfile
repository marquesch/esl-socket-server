FROM golang:1.26-trixie

WORKDIR /app

COPY server/ /app

CMD go run .
