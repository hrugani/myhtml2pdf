# Use the official Golang image to create a build artifact.
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY ../../go.mod ../../go.sum ./
RUN go mod download

# Copy the entire project source code
COPY ../../ .

# Build the Go app for the webapi
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myhtml2pdf ./cmd/webapi

# Use a slim Debian image for the final container
FROM debian:11-slim

# Install necessary dependencies
RUN apt-get update && apt-get install -y pdftk && rm -rf /var/lib/apt/lists/*

# Expose the port the app runs on
EXPOSE 8080

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myhtml2pdf /

# Copy the wkhtmltopdf binary from the source
COPY --from=builder /app/cmd/webapi/wkhtmltopdf /usr/local/bin/

# Command to run the executable
CMD ["/myhtml2pdf"]
