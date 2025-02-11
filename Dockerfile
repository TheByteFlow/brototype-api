# Use official Golang image
FROM golang:1.23

# Install air for live code reloading
RUN go install github.com/cosmtrek/air@latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files first to leverage Docker caching
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the project files
COPY . .

# Expose application port
EXPOSE 8080

# Set the default command for development using air
CMD ["air"]
