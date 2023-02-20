FROM golang:1.19.6-alpine3.17 as builder

WORKDIR /build

RUN apk add --no-cache \
    build-base \
    libpcap-dev

COPY . .

RUN make dist


FROM alpine:3.16

WORKDIR /app

RUN apk add --no-cache \
    libpcap-dev

COPY --from=builder /build/dist .

ENTRYPOINT ["./balena-go"]
