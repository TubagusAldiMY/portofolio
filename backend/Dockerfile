# Stage 1: Build
FROM golang:alpine AS builder

WORKDIR /app

# Copy go.mod + go.sum dulu untuk cache dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source
COPY knowledge_base ./knowledge_base
COPY . .

# Build dari folder cmd (bukan .)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backend ./cmd

# Stage 2: final image
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/backend /app/backend
COPY --from=builder /app/knowledge_base /app/knowledge_base

ENV PORT=8000
EXPOSE 8000

ENTRYPOINT ["/app/backend"]
