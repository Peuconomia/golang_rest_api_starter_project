FROM golang:1.19.1-alpine3.16 AS builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./.

#CMD ["app"]

FROM alpine:3.16

WORKDIR /usr/local/bin

COPY --from=builder /usr/local/bin/app ./

CMD ["./app"]