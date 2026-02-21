# Agenticide

**The Best Agentic Layer for Developers** - Simple, powerful, extensible.

AI-powered development assistant with a marketplace of extensions. Think of it as `npx` for AI agents.

## Quick Start

```bash
# Just run it - chat is default
agenticide

# Or explicitly
agenticide chat
```

## Installation

### Download Binary

See [agenticide-releases](https://github.com/ivikasavnish/agenticide-releases) for pre-built binaries.

```bash
# macOS (Apple Silicon)
curl -L https://github.com/ivikasavnish/agenticide-releases/releases/latest/download/agenticide-darwin-arm64 -o agenticide
chmod +x agenticide
sudo mv agenticide /usr/local/bin/
```

### From Source

```bash
git clone https://github.com/ivikasavnish/agenticide-go
cd agenticide-go
go build -o agenticide ./cmd/agenticide
```

## Extension Marketplace

Discover and install extensions like `npx`:

```bash
# Search for extensions
agenticide search security
agenticide search

# Install extension
agenticide install security
agenticide install code-analyzer

# List installed
agenticide list

# Get extension info
agenticide info security

# Uninstall
agenticide uninstall security
```

### Available Extensions

**Security & Quality**
- `security` - SAST scanning, secret detection, vulnerability checks
- `code-analyzer` - Complexity metrics, dead code detection
- `test-generator` - Auto-generate unit tests

**Development**
- `project-runner` - Auto-detect and run any project
- `git-ops` - Git operations, PR reviews
- `ai-recipes` - Pre-built AI workflows

**Operations**
- `deployment` - Deploy to AWS/GCP/Azure/Vercel
- `monitoring` - Health checks, alerts
- `cost-controller` - Track and optimize costs

**Data & Analytics**
- `db-analytics` - Database queries, terminal charts
- `web-search` - Multi-engine search

**Design**
- `ui-design` - Lovable design system integration

## Usage

### Chat (Default)

```bash
# Start chat
agenticide

# With flags
agenticide chat --ultraloop --ultrathink
```

### Task Management

```bash
agenticide task list
agenticide task add "Implement feature X"
agenticide task complete task-123
agenticide task graph
```

### Plan Mode

```bash
agenticide plan "Add authentication to API"
```

### Launch Modes

```bash
agenticide window    # Full-screen TUI
agenticide micro     # Floating overlay
agenticide web       # Browser interface
agenticide server start  # Background daemon
```

## Registration

Agenticide requires a license for use:

```bash
# Sign up (provides: email, mobile, LinkedIn, use case)
agenticide signup

# Activate license
agenticide activate AGNT-XXXX-XXXX-XXXX-XXXX

# Check status
agenticide status
```

## Architecture

### Core
- **Extension Registry**: Thread-safe plugin system
- **Event Bus**: Pub/sub messaging
- **Task System**: Dependency-aware tracking
- **Storage**: SQLite persistence
- **Config**: Viper-based configuration
- **Logging**: Zap structured logs

### Beautiful UI
- Lipgloss-styled terminal UI
- Tables, progress bars, lists, panels, charts
- Status icons (✓ ◐ ○ ⚠)
- Color themes

## Extension Development

Create custom extensions:

```go
package myext

import "github.com/ivikasavnish/agenticide-go/internal/core/extension"

type MyExtension struct {
    enabled bool
}

func (e *MyExtension) Name() string { return "my-extension" }
func (e *MyExtension) Version() string { return "1.0.0" }
func (e *MyExtension) Description() string { return "My custom extension" }
// ... implement extension.Extension interface
```

Publish to marketplace:
```bash
# Create manifest.json
# Push to GitHub
# Submit to registry
```

## Performance

- **Startup**: < 15ms
- **Memory**: ~50MB
- **Binary**: 5-6MB (single file)
- **Response**: < 50ms

## Philosophy

### Simple
- One command to start: `agenticide`
- No complex configuration
- Beautiful, intuitive UI

### Powerful
- Full agentic capabilities
- Extensible architecture
- Task orchestration

### Orchestrator
- Coordinate multiple agents
- Dependency management
- Event-driven workflows

## Launch Modes

1. **CLI** (default) - Command-line interface
2. **Window** - Full-screen terminal UI
3. **Micro** - Floating overlay window
4. **Server** - Background daemon + API
5. **Web** - Browser-based interface

## Repositories

- **Source**: [agenticide-go](https://github.com/ivikasavnish/agenticide-go)
- **Releases**: [agenticide-releases](https://github.com/ivikasavnish/agenticide-releases)

## License

Proprietary - Commercial use requires license.

## Support

- **Email**: support@agenticide.dev
- **Issues**: [GitHub Issues](https://github.com/ivikasavnish/agenticide-go/issues)
- **Docs**: [Documentation](https://github.com/ivikasavnish/agenticide-go/docs)

---

**Built with ❤️ for developers who want the best agentic experience.**
