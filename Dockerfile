# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go mod init urlshortener
RUN go get github.com/gorilla/mux
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Command to run the application
CMD ["./main"]
