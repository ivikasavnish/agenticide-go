# Go CLI Testing Status

**Date:** 2026-02-22  
**Version:** v0.5.0  
**Status:** Binaries Available

## Build Status

Pre-built binaries are available in the releases repository:
- macOS ARM64 (8.7 MB)
- macOS AMD64 (8.5 MB)
- Linux AMD64 (8.3 MB)
- Linux ARM64 (8.0 MB)
- Windows AMD64 (8.8 MB)

## Installation

Download from GitHub releases:
```bash
# Linux AMD64
curl -L https://github.com/ivikasavnish/agenticide-releases/releases/download/v0.5.0/agenticide-linux-amd64 -o agenticide
chmod +x agenticide
sudo mv agenticide /usr/local/bin/
```

## Features (Phase 4)

The Go rewrite includes:
- ✅ Extension marketplace (12 extensions)
- ✅ NPX-style search/install
- ✅ Beautiful Lipgloss UI
- ✅ Multi-launch modes
- ✅ Signup & licensing
- ✅ Security extension

## Build Environment Note

The project requires CGO for SQLite support. Building locally may encounter issues with mise-installed Go. Use CI/CD or Docker with standard Go toolchain.

## Current Status

- Code: Complete (3,900+ lines)
- Docs: Complete (2,500+ lines)
- Binaries: Available in releases
- Testing: Ready for CI/CD testing

## Node.js Version

For immediate use, the Node.js version (v3.1.0) is production-ready and fully tested. See main repository for installation.

---
Updated: 2026-02-22
