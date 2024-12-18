# Gunakan image Go resmi
FROM golang:1.23.3-alpine3.20 AS builder

# Set working directory di dalam container
WORKDIR /app

# Install dependencies sistem yang diperlukan
RUN apk add --no-cache git curl

# Copy go.mod and go.sum
COPY go.* ./

# Download dependency
RUN go mod download

# Copy semua file ke dalam container
COPY . .

# Build go app
RUN go build -o main ./cmd/api/main.go

# Stage 2
FROM alpine:3.20

# Set working directory di dalam container
WORKDIR /app

# Copy binary dari stage 1
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run binary
CMD ["./main"]