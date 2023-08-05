# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


ENV GO_CRUD_MONGO_URI=uri

# Expose the port that your Go application listens on
EXPOSE $PORT

# Run the Go application
CMD ["./main"]
