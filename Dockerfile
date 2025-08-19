# Use the official Golang image to create a build artifact.
# This is known as a multi-stage build.
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myhtml2pdf ./cmd/googleCloudRun

# Use a small Debian image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:11-slim

# Install pdftk
RUN apt-get update && apt-get install -y pdftk && rm -rf /var/lib/apt/lists/*

# The container listens on this port
EXPOSE 8080

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/myhtml2pdf /

# Copy the wkhtmltopdf binary to the production image from the builder stage.
COPY --from=builder /app/cmd/webapi/wkhtmltopdf /usr/local/bin/

# Run the web service on container startup.
CMD ["/myhtml2pdf"]
