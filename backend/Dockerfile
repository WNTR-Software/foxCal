# syntax=docker/dockerfile:1
FROM debian:latest AS base
WORKDIR /app
EXPOSE 8080

FROM docker.io/library/golang:1.25 AS dependencies
WORKDIR /dist
COPY go.mod go.sum ./
RUN go mod download

FROM dependencies AS build
WORKDIR /dist
COPY . .
RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build -o ./foxCal

FROM base AS final
WORKDIR /app
COPY --from=build /dist .
ENTRYPOINT [ "/app/foxCal" ]
