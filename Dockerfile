FROM golang:1.19.2-alpine3.16 AS builder

RUN apk update \
 && apk upgrade --no-cache

COPY . /app
WORKDIR /app
RUN go build -o app \
 && echo "${SOURCE_VERSION}" > revision.txt

FROM alpine:3.16 AS runtime

RUN apk update \
 && apk upgrade --no-cache

WORKDIR /app
COPY --from=builder /app/app /app/app
COPY --from=builder /app/public /app/public
COPY --from=builder /app/revision.txt /app/revision.txt
ENTRYPOINT ["/app/app"]
