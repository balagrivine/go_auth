# syntax=docker/dockerfile:1

# The official Go image with all necessary tools required
# to build and run the Go application
FROM golang:1.22

WORKDIR /app

# Copy go files into the root directory and download modules
COPY go.mod go.sum ./
COPY config ./config
COPY handler ./handler
RUN go mod download

# Copy source code into the image
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-auth

EXPOSE 8080

# Run
CMD ["/go-auth"]
