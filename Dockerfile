# Start from a base Go image
FROM golang:1.22-alpine

# Set the Current Working Directory inside the container
WORKDIR /agent-service

# Copy go mod and sum files
COPY ./src/go.mod ./src/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./src .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]