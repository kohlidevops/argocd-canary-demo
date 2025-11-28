# builder
FROM golang:1.24 AS builder
WORKDIR /src
COPY app/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# runtime
FROM gcr.io/distroless/base-debian11
COPY --from=builder /src/app /usr/local/bin/app
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/app"]
