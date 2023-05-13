# Use an official Golang runtime as a parent image
FROM golang:latest
# Set the working directory to /app
WORKDIR /app

# Create the MongoDB data directory
RUN mkdir -p /data/db

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main .

# Expose port 8080 for the Golang app and port 27017 for MongoDB
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
