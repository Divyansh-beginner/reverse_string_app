# Use official Go base image for your version
FROM golang:1.22.2

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (to leverage Docker cache)
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Now copy the rest of the project files, including Go code and scripts
COPY . ./
# Build your Go application
RUN go build -o dockerised_server

# Expose the port your Go server listens on (change if needed)
EXPOSE 8080

# Start your server (container command)
CMD ["./dockerised_server"]
