FROM golang:1.13-alpine

LABEL maintainer="Edwin Vautier"

# Set the working directory inside the container
WORKDIR /go/src
# Copy the full project to current directory
COPY . .
# Run command to install dependencies
RUN go mod download

EXPOSE 8080