## We specify the base image we need for our
## go application
FROM golang:1.18.3

RUN apt-get update

## We create an /build directory within our image
## that will hold our application sourcefiles
RUN mkdir /build

## We copy everything in the root directory
## into our /build directory
ADD . /build

## We specify that we now wish to execute any further
## commands inside our /build directory
WORKDIR /build

ENV GO111MODULE=on

RUN go mod tidy

## we run go build to compile the binary
## executable of our Go program
RUN env GOOS=linux GOARCH=amd64 \
    go build -o bin/link-easy-go -v app/main.go

# So in case when other container tries to connect 
# the 8080 port of the linnk-easy-go container, 
# the EXPOSE instruction is what 
# makes this possible.
EXPOSE 8080

## Our start command which kicks off our
# newly created binary executable
CMD ["/build/bin/link-easy-go"]


# --- Note ---

# Build docker image
# docker build -t link-easy-go:1.0 .

# -t for tagging, default is 'latest'

# Run docker container
# docker run -p 8080:8080 -it --rm link-easy-go:1.0

# -p - for <OUR_HOST_PORT>:<CONTAINER:PORT> 
# -it - This flag specifies that we want to run this image in
#       interactive mode with a tty for this container process.
# --rm - Automatically remove the container when it exits