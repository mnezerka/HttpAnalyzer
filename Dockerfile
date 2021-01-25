FROM golang:alpine AS builder

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"
CMD ["./http-analyze"]

FROM alpine:latest AS alpine
COPY --from=builder /app/http-analyzer /app/http-analyzer
WORKDIR /app/
EXPOSE 8088
CMD ["./http-analyzer"]
