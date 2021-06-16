FROM golang:1.14.9-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build

# Multistage build - output of build to alpine
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/crud-gin /app/
WORKDIR /app
CMD ["./crud-gin"]
