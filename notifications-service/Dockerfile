# Stage 1: Build Go binary
FROM golang:1.20 as builder
WORKDIR /app
COPY go.mod .
COPY main.go .

# Important: Build a static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Stage 2: Run in minimal container
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 3003
ENTRYPOINT ["./app"]
