# [STAGE] Build
FROM golang:1.19.0-alpine3.16 AS builder

# Set working directory to home
WORKDIR /app

# Copy all the golang source files to /app dir
COPY . .
# Clean go mod
RUN go mod tidy

# Build the app
RUN go build -o zbyte_user main.go


# [STAGE] Run
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/template_service .

# Expose docker port
EXPOSE 7011

# Run the app
CMD ["/app/template_service"]