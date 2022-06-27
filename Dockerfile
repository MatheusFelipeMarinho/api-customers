# Dockerfile Reference: https://docs.docker.com/engine/reference/builder/
FROM golang:1.18-alpine

# For alpine image
RUN apk update && apk upgrade && \
    apk add --no-cache make

WORKDIR /go/src/app
COPY . .
# RUN go get -d -v ./...

# Fix following error on alpine image
# cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in $PATH
ENV CGO_ENABLED=0

EXPOSE 8080

# Start the server
CMD [ "build" ]