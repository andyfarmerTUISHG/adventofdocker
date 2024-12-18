# Build stage
# 1 The first stage (named builder) uses the full Go image to compile our application
FROM golang:latest AS builder

WORKDIR /app
COPY . /app/src

RUN go build -o main /app/src/main.go

# Final stage
# 2 The second stage starts fresh from a minimal Alpine Linux image
# 3 We copy only the compiled binary from the builder stage
# 4 Everything else from the build stage is discarded
FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]