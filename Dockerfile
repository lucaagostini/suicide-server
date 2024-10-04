# Build stage
FROM alpine AS builder

# Final stage
FROM scratch

# Copy the binary from the builder stage
COPY target/main /main

# Copy SSL certificates from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["/main"]
