# [STAGE] Build
FROM golang:1.21.7-alpine3.19 AS builder

# Set working directory to home
WORKDIR /app

# Copy all the golang source files to /app dir
COPY . .
# Clean go mod
RUN go mod tidy

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -o server src/main.go


# [STAGE] Run
FROM golang:1.21.7-alpine3.19
WORKDIR /app
COPY --from=builder /app/server .

# Expose docker port
EXPOSE 7006

# Run the app
CMD ["/app/server"]