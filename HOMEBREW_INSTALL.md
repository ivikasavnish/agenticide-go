# Homebrew Installation Guide

Install Agenticide via Homebrew on macOS and Linux.

## Quick Install

```bash
brew tap ivikasavnish/agenticide
brew install agenticide
```

## What Gets Installed

- **Binary**: `/usr/local/bin/agenticide` (or `/opt/homebrew/bin/agenticide` on Apple Silicon)
- **Version**: 0.4.0
- **Size**: ~5-6 MB

## Verify Installation

```bash
agenticide --version
# Output: agenticide version 0.4.0
```

## First Run

```bash
# Start chat (default)
agenticide

# Registration required
agenticide signup
agenticide activate <license-key>
```

## Usage Examples

### Extension Marketplace
```bash
# Search extensions
agenticide search security

# Install extension
agenticide install security

# List installed
agenticide list
```

### Task Management
```bash
agenticide task add "Implement feature X"
agenticide task list
agenticide task complete <id>
```

### Launch Modes
```bash
agenticide              # CLI (default)
agenticide window       # Full-screen TUI
agenticide micro        # Floating overlay
agenticide server start # Background daemon
```

## Update

```bash
brew update
brew upgrade agenticide
```

## Uninstall

```bash
brew uninstall agenticide
brew untap ivikasavnish/agenticide
```

## Troubleshooting

### Permission Issues
```bash
# Fix permissions
sudo chown -R $(whoami) /usr/local/bin/agenticide
chmod +x /usr/local/bin/agenticide
```

### Command Not Found
```bash
# Add to PATH (if needed)
echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

### Apple Silicon Issues
```bash
# Homebrew on Apple Silicon uses /opt/homebrew
echo 'export PATH="/opt/homebrew/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

## Manual Installation (Alternative)

If Homebrew doesn't work:

```bash
# macOS (Apple Silicon)
curl -L https://github.com/ivikasavnish/agenticide-releases/releases/latest/download/agenticide-darwin-arm64 -o agenticide
chmod +x agenticide
sudo mv agenticide /usr/local/bin/

# macOS (Intel)
curl -L https://github.com/ivikasavnish/agenticide-releases/releases/latest/download/agenticide-darwin-amd64 -o agenticide
chmod +x agenticide
sudo mv agenticide /usr/local/bin/
```

## Links

- **Homebrew Tap**: https://github.com/ivikasavnish/homebrew-agenticide
- **Releases**: https://github.com/ivikasavnish/agenticide-releases
- **Source**: https://github.com/ivikasavnish/agenticide-go
- **Marketplace**: https://github.com/ivikasavnish/agenticide-go/blob/main/MARKETPLACE.md

## Support

Issues with Homebrew installation? Report at:
https://github.com/ivikasavnish/homebrew-agenticide/issues
