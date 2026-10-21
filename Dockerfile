FROM golang:1.25-alpine AS builder
WORKDIR /app
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o cache_consumer .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/cache_consumer .
COPY --from=builder /app/etc ./etc
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
CMD ["./cache_consumer", "-f", "./etc/config.yaml"]
