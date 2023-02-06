## Build
FROM golang:1.20.0-alpine3.17 AS build

WORKDIR /wombat-src

COPY . .
RUN go mod download

RUN go build -o wombat

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /wombat-src/wombat /wombat
COPY --from=build /wombat-src/static /static

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/wombat"]
