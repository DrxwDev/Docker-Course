# --- Build stage ---
FROM golang:tip-alpine3.22 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the requirements file first (for better caching)
COPY go.mod go.sum ./

# Copy the rest of the application code into the container
COPY . .

# Download the dependencies
RUN go mod tidy && go build -v -o app .

# --- Runtime stage ---
FROM alpine:3.22.1

WORKDIR /app
COPY --from=builder /app/app .

CMD [ "./app" ]
