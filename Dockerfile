# syntax=docker/dockerfile:1

FROM golang:1.17.6-alpine3.15
WORKDIR /app
ARG MYSQL__HOST=172.27.42.11
ENV MYSQL__HOST $MYSQL__HOST
COPY . /app
RUN go mod download
RUN go mod vendor
RUN cd cmd/server && go build -o bim .
RUN mkdir -p statik && cp -p cmd/server/bim . && cp -p cmd/server/statik/statik.go statik
CMD ["./bim"]