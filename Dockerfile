FROM golang:1.18-alpine AS build
WORKDIR /app

COPY ["go.mod", "go.sum"] ./

RUN go mod download

COPY **/*.go ./

RUN go build -o /dtek-api

FROM gcr.io/distroless/staic
COPY --from=build /dtek-api /
ENTRYPOINT ["/dtek-api"]