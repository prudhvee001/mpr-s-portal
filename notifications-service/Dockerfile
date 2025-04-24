# Stage 1: Build the Go binary
FROM golang:1.20 as builder
WORKDIR /app
COPY go.mod .
COPY main.go .
RUN go build -o app

# Stage 2: Run the Go binary
FROM alpine
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 3003
ENTRYPOINT ["./app"]
