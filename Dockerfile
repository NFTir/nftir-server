# Build stage
## Pull golang image from the hub
FROM golang:alpine AS builder

## Set up ENV vars
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 


## Set up args for build
ARG AWS_ACCESS_KEY_ID
ARG AWS_SECRET_ACCESS_KEY
ARG REGION

## Set up aws-conf
RUN apk add --no-cache aws-cli

RUN aws configure set aws_access_key_id $AWS_ACCESS_KEY_ID \
    && aws configure set aws_secret_access_key $AWS_SECRET_ACCESS_KEY \
    && aws configure set region $REGION \
    && aws configure set output json


## Choose work directory
WORKDIR /NFTir

## Copy local project to docker container
COPY . .

## Run build command
RUN go build -o server .


# Run stage
## alpine:latest image is a light linux image
FROM alpine:latest AS runner

## Choose work directory
WORKDIR /NFTir

## Copy the executable binary file and .env file from the last stage to the new stage
COPY --from=builder /NFTir/server .
COPY --from=builder /root/.aws /root/.aws


ARG PORT
EXPOSE $PORT 

# Execute the build
CMD ["./server"]