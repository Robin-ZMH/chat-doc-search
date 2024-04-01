FROM golang:1.22.1 AS build-stage

WORKDIR /app

COPY ./ ./

RUN GOPROXY=https://goproxy.cn go mod download

RUN CGO_ENABLED=0 go build -o engine

FROM alpine:latest AS release-stage

WORKDIR /app

COPY --from=build-stage /app/engine /app/engine

EXPOSE 9999

STOPSIGNAL SIGINT

CMD ["/app/engine"]