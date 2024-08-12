FROM golang:1.22-alpine3.20 AS build

WORKDIR /opt/isitrunning

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:3.20

RUN addgroup -S isitrunning && adduser -S isitrunning -G isitrunning

USER isitrunning
WORKDIR /opt/isitrunning

COPY --from=build --chown=isitrunning:isitrunning /opt/isitrunning/app ./

ENTRYPOINT ["./app"]