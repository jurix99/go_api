# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app
ENV GIN_MODE=release
# Copy the Go API source code to the container
COPY . .

RUN go build -o main .

# Expose port 8080 for the Go API
EXPOSE 8080

# Run the Go API executable
CMD ["./main"]
