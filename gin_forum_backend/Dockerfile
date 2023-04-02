FROM golang:alpine AS builder


ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /build


COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o gin_forum .

FROM debian:stretch-slim


COPY ./conf /conf
COPY ./wait-for-it.sh /


COPY --from=builder /build/gin_forum /
RUN chmod 755 wait-for-it.sh
