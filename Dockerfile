# syntax=docker/dockerfile:1
FROM docker.io/golang:1.18-alpine AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /dtek-api

FROM gcr.io/distroless/static-debian11
COPY --from=build /dtek-api /
CMD ["/dtek-api"]