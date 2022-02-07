FROM golang:1.15.2-alpine as builder
COPY . /home
WORKDIR /home
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go build
FROM alpine:latest
WORKDIR /home
COPY --from=builder /home/news .
EXPOSE 3007
CMD ["./news"]