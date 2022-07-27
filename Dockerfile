# syntax=docker/dockerfile:1
FROM docker.io/golang:1.18 AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /dtek-api

FROM gcr.io/distroless/base-debian11
COPY --from=build /dtek-api /
CMD ["/dtek-api"]