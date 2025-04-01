FROM golang:1.24.1-alpine AS builder

COPY --link . .

RUN go mod download

RUN go build -o /go/bin/app .

CMD ["/go/bin/app"]
