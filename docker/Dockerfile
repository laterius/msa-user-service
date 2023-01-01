FROM golang:1.18-alpine3.16 as builder
WORKDIR /app
COPY ./app/go.mod .
COPY ./app/go.sum .
RUN go mod download
COPY ./app .
RUN cd cmd/service && go build -o service

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/cmd/service/service .
COPY ./app .
COPY ./docker/files/entrypoint.sh /
COPY ./docker/files/wait-for-it.sh /
