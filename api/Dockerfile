FROM golang:1.16.6-alpine

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN go build

EXPOSE 8080

CMD ["/app/product-pack-test-go"]