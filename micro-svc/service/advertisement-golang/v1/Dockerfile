FROM golang:1.9.6-alpine as builder
COPY . /app
WORKDIR /app
RUN apk update && apk add git \
    && go get -v -u -insecure github.com/gin-gonic/gin \
    && go build
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 3003
CMD ["./app"]