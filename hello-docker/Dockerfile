# === Build stage ===
FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app .

# === Runtime stage ===
FROM gcr.io/distroless/static-debian12:nonroot AS runtime

WORKDIR /src

COPY --from=builder /build/app /src/app

EXPOSE 8000

USER nonroot:nonroot
ENTRYPOINT [ "./app" ]
