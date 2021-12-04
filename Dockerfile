# syntax=docker/dockerfile:1
FROM golang:1.17
COPY . /usr/src/app
WORKDIR /usr/src/app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY .env .
COPY --from=0 /usr/src/app/main ./
CMD ["./main"]
