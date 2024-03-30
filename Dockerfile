FROM golang:1.22.1 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /lhshop

EXPOSE 8080

CMD ["/lhshop"]
