## Build
FROM golang:1.20.0-alpine3.17

WORKDIR /wombat-src

COPY . .
RUN go mod download
RUN go build -o wombat

RUN addgroup --gid 2137 certgroup
RUN adduser --disabled-password --gecos "" --ingroup certgroup wombat
USER wombat

EXPOSE 8080
EXPOSE 8081

ENTRYPOINT ["./wombat"]
