FROM golang:1.16.6-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build

EXPOSE 8000

CMD ["/app/product-pack-test-go"]