FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./web ./web

# Set the working directory
WORKDIR /app/cmd

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o task_list

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
EXPOSE 8080

# Run
CMD ["/app/cmd/task_list"]
