# Build stage
FROM golang:1.20 as builder
WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN go build -o notification-app

# Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/notification-app .
EXPOSE 3003
ENTRYPOINT ["./notification-app"]
