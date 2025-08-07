# Use the official Go image as the base image
FROM golang:1.24.6 AS builder

# Set the working directory inside the container
WORKDIR /app

COPY . .
# Download all dependencies
RUN go mod download
# Build the Go binary
# RUN go build -o /app/shorty
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/shorty

# Use a minimal base image for the final stage
FROM scratch

COPY --from=builder /app/shorty /app/shorty
WORKDIR /app

# Expose the application port
EXPOSE 8080

# Command to run the binary
CMD ["/app/shorty"]
