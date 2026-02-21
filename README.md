# Agenticide Go

High-performance AI coding assistant built in Go. Production-ready commercial service with pluggable extensions and beautiful terminal UI.

## Features

- 🚀 **10-50x faster** than JavaScript version
- 🔒 **Source code protection** via compiled binary
- 🎨 **Beautiful terminal UI** with Lipgloss
- 🧩 **11 pluggable extensions**
- 📊 **Real-time task management** with dependency trees
- 🔄 **Ultraloop & Ultrathink modes**
- 💾 **SQLite persistence**
- 🎯 **Zero dependencies** - single binary

## Installation

### Binary (Recommended)
```bash
# Download latest release
curl -L https://github.com/ivikasavnish/agenticide-go/releases/latest/download/agenticide-$(uname -s)-$(uname -m) -o agenticide
chmod +x agenticide
sudo mv agenticide /usr/local/bin/
```

### From Source
```bash
git clone https://github.com/ivikasavnish/agenticide-go
cd agenticide-go
go build -o agenticide ./cmd/agenticide
```

## Quick Start

```bash
# Check version
agenticide --version

# Start interactive chat
agenticide chat

# With ultraloop (retry until complete)
agenticide chat --ultraloop

# With ultrathink (deep reasoning)
agenticide chat --ultrathink

# List extensions
agenticide ext list

# Generate task plan
agenticide plan "Add authentication to API"

# Execute tasks
agenticide task execute --ultraloop
```

## Extensions

1. **Web Search** - Multi-engine search with content extraction
2. **UI Design** - Lovable design system integration
3. **Security Agent** - SAST, secrets, vulnerabilities
4. **Code Analyzer** - AST, complexity, dead code
5. **Project Runner** - Run any language
6. **DB & Analytics** - Query builder, charts
7. **LLM Recipes** - Pre-built AI workflows
8. **Deployment** - AWS, GCP, Azure, Vercel
9. **Cost Controller** - Track and optimize costs
10. **Monitoring** - Health checks and alerts
11. **Task System** - Dependency-aware task management

## Architecture

```
Extension System
├── Registry (thread-safe management)
├── EventBus (pub/sub messaging)
├── Context (request-scoped data)
├── Storage (SQLite)
└── Config (hot reload)
```

## Performance

- Startup: < 10ms
- Memory: < 50MB
- Binary: < 20MB
- Response: < 100ms

## Development

```bash
# Install dependencies
go mod download

# Run tests
go test ./...

# Build
go build -o agenticide ./cmd/agenticide

# Run
./agenticide --help
```

## License

Proprietary - All rights reserved. See LICENSE for details.

## Commercial Use

This is a commercial product. For licensing inquiries: contact@agenticide.dev
