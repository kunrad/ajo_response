FROM golang:1.14-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/go-app

COPY . .

# Build the Go app
RUN go build -o ./out/go-app .

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/go-app/out/go-app /app/go-app

# This container exposes port 8080 to the outside world
EXPOSE 3001

# Run the binary program produced by `go install`
CMD ["/app/go-app"]