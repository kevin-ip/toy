FROM golang:1.24.1-alpine AS downloader
COPY --link go.mod go.sum ./
RUN go mod download

FROM downloader AS builder
COPY --link . .
RUN go build -o /go/bin/app .

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app

CMD ["/go/bin/app"]
