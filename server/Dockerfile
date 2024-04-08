# Use the official Go image as the base
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy your Go files into the container
COPY . .

# Build the Go app
RUN go build -o myserver

# Expose the port your app will listen on
EXPOSE 8080

# Command to run your app
CMD ["./server"]

