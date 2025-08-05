FROM golang:1.23 AS builder
WORKDIR /app

# Copy go.mod and go.sum files first (for caching dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application (replace 'myapp' with your desired binary name)
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

# Run stage (minimal, portable container)
FROM gcr.io/distroless/base-debian11
WORKDIR /

COPY --from=builder /app/myapp /myapp
# Copy templates to /app/templates to match Go code expectation
COPY --from=builder /app/templates /app/templates


# Expose application port if needed (replace with your app's port)
EXPOSE 8000

CMD ["/myapp"]