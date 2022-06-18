FROM golang:1.18-alpine3.16 AS builder

RUN apk update \
 && apk upgrade --no-cache

COPY . /app
WORKDIR /app
RUN go build -o app

FROM alpine:3.16 AS runtime

RUN apk update \
 && apk upgrade --no-cache

WORKDIR /app
COPY --from=builder /app/app /app/app
COPY --from=builder /app/public /app/public
ENTRYPOINT ["/app/app"]
