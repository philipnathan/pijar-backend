# Gunakan image Go resmi
FROM golang:1.23.3-alpine3.20

# Set working directory di dalam container
WORKDIR /app

# Install dependencies sistem yang diperlukan
RUN apk add --no-cache git curl

# Install command for air
RUN go install github.com/air-verse/air@latest

# # Copy go.mod dan go.sum
# COPY go.* ./

# # Download dependency
# RUN go mod download

# # Copy semua file ke dalam container
# COPY . .

# Ekspor port aplikasi (misalnya 8080)
EXPOSE 8080

# Jalankan Air saat container dimulai
CMD ["air"]