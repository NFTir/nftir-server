# Build stage
## Pulling golang image from Dockerhub
FROM golang:alpine AS builder

## Choose work directory
WORKDIR /NFTir

## Copy local project to docker container
COPY . .

## Run build command
RUN go build -o agent .

### Technically, the dockerfile could be done here but the image will be heavy since golang image is already 840MB itself. 
### Golang image contains all the dependencies even though we only need the executable binary file to run the app.

# Run stage
## alpine:latest image is a light linux image
FROM alpine:latest

## Choose work directory
WORKDIR /NFTir

## Copy only the executable binary file from the last stage to this stage
COPY --from=builder /NFTir/agent .


# Execute the build
CMD ["./agent"]