FROM golang:1.15.3-alpine3.12 AS builder
WORKDIR /
COPY . .

RUN apk add --update git gcc musl-dev
RUN CGO_ENABLED=1 go build -tags musl -installsuffix cgo -o /gravity-exporter-kafka ./cmd/gravity-exporter-kafka/gravity-exporter-kafka.go

FROM alpine:3.12
COPY --from=builder /gravity-exporter-kafka /gravity-exporter-kafka
COPY ./configs .
EXPOSE 44144
CMD ["/gravity-exporter-kafka"]
