# Using the offical Golang image as the base
FROM golang:1.26-alpine AS builder

# Setting env 
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build

# getting dependency 
COPY go.mod ./

#Download dependencies
RUN go mod download

# Copy the entire application source
COPY . .
RUN go build -ldflags="-w -s" -o herald ./herald.go

FROM gcr.io/distroless/static-debian12

WORKDIR /

# CRITICAL: Copy the compiled binary from the builder stage
COPY --from=builder /build/herald /herald

EXPOSE 8080

# built-in safe nonroot user included in distroless
USER nonroot:nonroot

# Run the application
ENTRYPOINT ["/herald"]