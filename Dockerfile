# Builder
FROM golang:1.18.4-alpine3.16 as builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /foodstore-server

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN go build -o app *.go

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && apk --update --no-cache add tzdata && mkdir /foodstore-server

WORKDIR /foodstore-server 

# Expose port 9090 to the outside world
EXPOSE 3001

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /foodstore-server/app /foodstore-server

# Command to run the executable
CMD ["/foodstore-server/app", "start"]