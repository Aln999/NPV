# ---- Build Stage ----
    FROM golang:1.23 as builder
    WORKDIR /app
    
    # Copy Go modules and download dependencies
    COPY backend/go.mod backend/go.sum ./
    RUN go mod download
    
    # Copy the application source code
    COPY backend/ . 
    
    # Build the Go binary
    RUN CGO_ENABLED=0 GOOS=linux go build -o /app/network-policy-visualizer
    
    # ---- Final Image ----
    FROM alpine:3.18
    WORKDIR /app
    
    # Create a non-root user
    RUN addgroup -S appgroup && adduser -S appuser -G appgroup
    
    # Copy the built Go binary from the builder stage
    COPY --from=builder /app/network-policy-visualizer .
    
    # Copy frontend files
    COPY frontend/ ./frontend/
    
    # Ensure the binary has execution permissions
    RUN chmod +x /app/network-policy-visualizer
    
    # Set user to non-root for security
    USER appuser
    
    # Expose the application port
    EXPOSE 8080
    
    # Run the application
    CMD ["/app/network-policy-visualizer"]
    
