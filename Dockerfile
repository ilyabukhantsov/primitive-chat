# Stage 1: Build the binary
FROM golang:1.26.3-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy dependency files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application 
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd

# Stage 2: Create the final production image
FROM alpine:latest  

# Set work directory
WORKDIR /root/

# Copy only the compiled binary from the builder stage
COPY --from=builder /main .

# Optional: Expose the port your app runs on
EXPOSE 8080

# Run the binary
CMD ["./main"]

