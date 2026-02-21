# Agenticide Extension Marketplace

Discover, install, and manage extensions like `npx` for Node.js.

## Quick Start

```bash
# Search for extensions
agenticide search

# Install an extension
agenticide install security

# List installed extensions
agenticide list

# Get extension details
agenticide info security

# Uninstall
agenticide uninstall security
```

## How It Works

### Like npx

```bash
# NPM/npx way
npx create-react-app my-app
npm install -g typescript

# Agenticide way
agenticide install security
agenticide install code-analyzer
```

### Remote Installation

Extensions are:
1. **Discovered** via marketplace registry
2. **Downloaded** from GitHub releases or registry
3. **Installed** to `~/.agenticide/extensions/`
4. **Activated** automatically
5. **Updated** via registry

## Extension Categories

### Security & Quality
- `security` - SAST, secrets, vulnerabilities
- `code-analyzer` - Metrics, complexity, dead code
- `test-generator` - Auto-generate tests

### Development
- `project-runner` - Run any project type
- `git-ops` - Git workflows, PR reviews
- `ai-recipes` - Pre-built AI workflows

### Operations
- `deployment` - Multi-cloud deployment
- `monitoring` - Health checks, alerts
- `cost-controller` - Cost tracking

### Data
- `db-analytics` - Database tools
- `web-search` - Web research

### Design
- `ui-design` - UI component tools

## Extension Structure

Each extension includes:

```
~/.agenticide/extensions/security/
├── security              # Executable (plugin)
├── manifest.json         # Metadata
└── README.md            # Documentation
```

### Manifest Format

```json
{
  "name": "security",
  "version": "1.0.0",
  "description": "SAST scanning and security checks",
  "author": "Agenticide Team",
  "repository": "https://github.com/ivikasavnish/agenticide-go",
  "tags": ["security", "sast", "scanner"],
  "category": "security",
  "install_url": "https://github.com/.../security"
}
```

## Creating Extensions

### 1. Implement Interface

```go
package myext

import "github.com/ivikasavnish/agenticide-go/internal/core/extension"

type MyExtension struct {
    enabled bool
}

func (e *MyExtension) Name() string {
    return "my-extension"
}

func (e *MyExtension) Version() string {
    return "1.0.0"
}

func (e *MyExtension) Description() string {
    return "My custom extension"
}

func (e *MyExtension) Commands() []extension.Command {
    return []extension.Command{
        {
            Name: "hello",
            Description: "Say hello",
            Handler: e.handleHello,
        },
    }
}

func (e *MyExtension) HandleCommand(ctx context.Context, cmd string, args []string) (*extension.Result, error) {
    // Handle command
    return &extension.Result{Success: true}, nil
}
```

### 2. Create Manifest

```json
{
  "name": "my-extension",
  "version": "1.0.0",
  "description": "My custom extension",
  "author": "Your Name",
  "repository": "https://github.com/you/my-extension",
  "tags": ["custom"],
  "category": "development"
}
```

### 3. Build & Release

```bash
# Build
go build -o my-extension

# Create release
gh release create v1.0.0 my-extension

# Submit to registry
# (Coming soon: automated submission)
```

## Registry API

Extensions are discovered via:

- **Official Registry**: `https://registry.agenticide.dev/v1`
- **GitHub Releases**: Auto-discovered
- **Custom Registry**: Self-hosted option

### Search API

```bash
GET https://registry.agenticide.dev/v1/search?q=security

Response:
{
  "extensions": [
    {
      "name": "security",
      "version": "1.0.0",
      "description": "...",
      "install_url": "..."
    }
  ]
}
```

## Community Extensions

Anyone can create and publish extensions:

1. Build extension following interface
2. Host on GitHub with releases
3. Submit to community registry
4. Users discover via `agenticide search`

### Examples

- `security` - Official security scanner
- `git-ops` - Community git tools
- `test-generator` - Community test automation

## Best Practices

### Extension Design
- ✅ Single responsibility
- ✅ Clear command names
- ✅ Beautiful terminal output
- ✅ Error handling
- ✅ Documentation

### Performance
- ✅ Fast startup (< 50ms)
- ✅ Minimal memory
- ✅ Async operations
- ✅ Progress indicators

### Security
- ✅ Validate inputs
- ✅ Sandbox execution
- ✅ Secure dependencies
- ✅ Signed releases

## Coming Soon

- [ ] Extension ratings & reviews
- [ ] Automatic updates
- [ ] Dependency management
- [ ] Extension marketplace web UI
- [ ] Premium extensions
- [ ] Team extensions

## Support

- **Registry Issues**: [GitHub Issues](https://github.com/ivikasavnish/agenticide-go/issues)
- **Extension Development**: [Developer Guide](docs/EXTENSION_DEV.md)
- **Submission**: marketplace@agenticide.dev

---

**Marketplace makes Agenticide infinitely extensible.**
