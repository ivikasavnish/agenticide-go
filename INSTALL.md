# Agenticide Installation Guide

## Quick Install (Linux/macOS)

### Direct Download (Recommended)

**Linux:**
```bash
curl -L https://github.com/ivikasavnish/agenticide-go/releases/download/v0.5.0/agenticide-linux-amd64 -o agenticide
chmod +x agenticide
sudo mv agenticide /usr/local/bin/
agenticide
```

**macOS (Apple Silicon):**
```bash
curl -L https://github.com/ivikasavnish/agenticide-go/releases/download/v0.5.0/agenticide-darwin-arm64 -o agenticide
chmod +x agenticide
sudo mv agenticide /usr/local/bin/
agenticide
```

**macOS (Intel):**
```bash
curl -L https://github.com/ivikasavnish/agenticide-go/releases/download/v0.5.0/agenticide-darwin-amd64 -o agenticide
chmod +x agenticide
sudo mv agenticide /usr/local/bin/
agenticide
```

### Homebrew (macOS/Linux)

```bash
brew tap ivikasavnish/agenticide
brew install agenticide
```

## First Run Setup

On first run, you'll see the setup wizard:

```
Welcome to Agenticide!

Let's get you set up. This will only take a moment.

Your email: [enter your email]
Your phone (optional): [enter phone or press enter]

Select your AI provider:
  1) OpenAI (GPT-4)
  2) Anthropic (Claude)
  3) GitHub Copilot
  4) Local (Ollama)
Choice [1-4]: [choose 1-4]

OpenAI API Key: [paste your API key]

✓ Setup complete!
```

Configuration is saved to `~/.agenticide/config.yaml`

## Commands

```bash
agenticide              # Start chat (runs setup on first run)
agenticide search       # Browse extension marketplace
agenticide install <ext># Install extension
agenticide list         # List installed extensions
agenticide task add     # Add task
agenticide --help       # See all commands
```

## Troubleshooting

### Clear Homebrew Cache (if updating)
```bash
brew uninstall agenticide
brew untap ivikasavnish/agenticide
rm -rf $(brew --cache)
brew tap ivikasavnish/agenticide
brew install agenticide
```

### Reset Configuration
```bash
rm ~/.agenticide/config.yaml
agenticide  # Will show setup wizard again
```
