# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/wombat/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /wombat

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /wombat /wombat

EXPOSE 2137

USER nonroot:nonroot

ENTRYPOINT ["/wombat"]

# source: https://docs.docker.com/language/golang/build-images/#multi-stage-builds