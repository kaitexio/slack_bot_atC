FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN go build -o /main

FROM alpine:3.9

COPY --from=builder /main .

ENTRYPOINT ["/main"]