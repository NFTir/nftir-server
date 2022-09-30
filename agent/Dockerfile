# Pulling golang image from Dockerhub
FROM golang:latest

# Choose work directory
WORKDIR /NFTir

# Copy local project to docker container
COPY . .

# Run build command
RUN go build .

# Execute the build
CMD ["./agent"]
