FROM golang:alpine as builder
LABEL authors="artisan"

WORKDIR /

COPY .. .

RUN go build -o ./bin/app ./cmd/price/price.go

FROM alpine

COPY --from=builder /bin/app .
COPY --from=builder .env .

CMD ./app