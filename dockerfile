FROM golang:1.23.0

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code.
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Optional: Document the port the application listens on.
EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]
