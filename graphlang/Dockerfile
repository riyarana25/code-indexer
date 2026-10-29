# Explicitly set builder platform
FROM --platform=linux/amd64 golang:1.27-trixie AS builder

WORKDIR /analyzer
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the cmd entrypoint (graphlang itself is a library package).
# Tree-sitter grammars need cgo, so build against glibc and run on Debian.
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o graphlang-analyzer ./cmd/graphlang-analyzer

# Final stage uses the same Debian release as the builder so the cgo binary's libc matches
FROM --platform=linux/amd64 debian:trixie-slim

WORKDIR /analyzer
COPY --from=builder /analyzer/graphlang-analyzer /analyzer/
CMD ["./graphlang-analyzer"]
