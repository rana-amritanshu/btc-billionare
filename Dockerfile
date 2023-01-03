
FROM golang:1.19 AS builder
WORKDIR /go/src/github.com/rana-amritanshu/btc-billionare/
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/rana-amritanshu/btc-billionare/main .
COPY --from=builder /go/src/github.com/rana-amritanshu/btc-billionare/docs ./docs
COPY --from=builder /go/src/github.com/rana-amritanshu/btc-billionare/swagger-ui ./swagger-ui
CMD ["./main"]