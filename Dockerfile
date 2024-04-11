FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /lhshop

EXPOSE 8000

CMD ["/lhshop"]
