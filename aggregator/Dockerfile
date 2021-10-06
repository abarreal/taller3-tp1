FROM golang:1.16-alpine

# Create a directory for GOPATH
WORKDIR /opt/gopath
ENV GOPATH=/opt/gopath

# Create a directory for the app.
WORKDIR /opt/app

# Copy go.mod and go.sum, and install deps.
# This avoids having to download them
# multiple times, each time code changes.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy source code and build.
COPY . .

# Do build the program.
RUN go build -o main

# Define the command that will be executed
# when the image is launched with docker run. 
CMD ["/opt/app/main"]