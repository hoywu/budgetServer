# 构建层
FROM golang:1.22.1-alpine AS builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server .

# 运行层
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/server .
CMD ["/app/server"]

# Use `--network=host` for docker build if you have a proxy.
# Example: `docker build --network=host --tag budget .`

# Run Example: `docker run -d --name budget --network host -v /path/to/config.yaml:/app/config.yaml budget`
