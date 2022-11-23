FROM golang:1.19 AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /dtek-api

FROM gcr.io/distroless/base-debian11:latest
COPY --from=build /dtek-api /
CMD ["/dtek-api"]