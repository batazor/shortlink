FROM golang:1.16-alpine as builder

ARG CI_COMMIT_TAG

WORKDIR /go/src/github/batazor/shortlink
COPY . .

# Load dependencies
RUN go mod vendor

# Build project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build \
  -a \
  -mod vendor \
  -ldflags "-s -w -X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo \
  -trimpath \
  -o app ./cmd/metadata

FROM alpine:3.13

# 50051: gRPC
# 9090: metrics
EXPOSE 50051 9090

# Install dependencies
RUN \
    apk update && \
    apk add --no-cache curl

RUN addgroup -S metadata && adduser -S -g metadata metadata
USER metadata

# TODO: fix it
#HEALTHCHECK \
#  --interval=5s \
#  --timeout=5s \
#  --retries=3 \
#  CMD curl -f localhost:9090/ready || exit 1

WORKDIR /app/
COPY --from=builder /go/src/github/batazor/shortlink/app .
CMD ["./app"]
