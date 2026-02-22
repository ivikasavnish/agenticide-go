#!/bin/bash
set -e

VERSION="0.4.0"
DIST_DIR="dist"

echo "Building Agenticide v${VERSION} for all platforms..."
mkdir -p ${DIST_DIR}

# darwin-arm64 (Apple Silicon)
echo "Building darwin-arm64..."
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build \
  -ldflags="-s -w" \
  -o ${DIST_DIR}/agenticide-darwin-arm64 \
  ./cmd/agenticide

# darwin-amd64 (Intel Mac)
echo "Building darwin-amd64..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build \
  -ldflags="-s -w" \
  -o ${DIST_DIR}/agenticide-darwin-amd64 \
  ./cmd/agenticide

# linux-amd64
echo "Building linux-amd64..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags="-s -w" \
  -o ${DIST_DIR}/agenticide-linux-amd64 \
  ./cmd/agenticide

# linux-arm64
echo "Building linux-arm64..."
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build \
  -ldflags="-s -w" \
  -o ${DIST_DIR}/agenticide-linux-arm64 \
  ./cmd/agenticide

echo ""
echo "Build complete! Binaries in ${DIST_DIR}:"
ls -lh ${DIST_DIR}/
