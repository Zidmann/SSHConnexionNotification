############################################################################
# Builder Image
FROM golang:1.16.7 as builder

# Create and set working directory
RUN mkdir -p /app
WORKDIR /app

# Install dependencies
ADD go.mod go.mod
ADD go.sum go.sum
RUN go mod download

# Add code
ADD . .

# Build the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

############################################################################
# Final Image
FROM alpine:3.14

# Add ca-certificates
RUN apk update && apk --no-cache  add ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main ./main

############################################################################
# Run the binary
CMD ["./main"]

############################################################################
