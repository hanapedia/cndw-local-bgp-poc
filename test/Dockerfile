FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o server main.go

# Stage 2: runtime
FROM alpine:3.20
WORKDIR /app

# Install curl for debugging
RUN apk add --no-cache curl

# Copy built binary
COPY --from=builder /app/server .

EXPOSE 8080
ENTRYPOINT ["./server"]
