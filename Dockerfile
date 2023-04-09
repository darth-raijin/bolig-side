FROM golang:1.19-alpine AS builder

# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

#Install certs
RUN apk add --no-cache ca-certificates

# Working directory outside $GOPATH
WORKDIR /build

# Copy go module files and download dependencies
COPY go.* ./
RUN go mod download

# Copy source files
COPY . .

# Build source
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o loop .

# Minimal image for running the application
FROM scratch as final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/
# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Import the compiled executable from the first stage.
COPY --from=builder ["/build/loop", "/build/.env", "/"]

# Open port
EXPOSE 8080

# Will run as unprivileged user/group
USER nobody:nobody

ENTRYPOINT ["/loop"]