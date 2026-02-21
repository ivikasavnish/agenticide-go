# Agenticide Extension Contract

Complete interface specification for building extensions.

## Core Interface

Every extension must implement the `Extension` interface:

```go
package extension

import (
    "context"
    "time"
)

type Extension interface {
    // Identity
    Name() string                           // Unique extension name (e.g., "security")
    Version() string                        // Semantic version (e.g., "1.0.0")
    Description() string                    // One-line description
    Author() string                         // Author/team name
    
    // Dependencies
    Dependencies() []string                 // List of required extensions
    
    // Lifecycle
    Enable(ctx context.Context) error       // Called when extension is enabled
    Disable(ctx context.Context) error      // Called when extension is disabled
    IsEnabled() bool                        // Current enabled state
    
    // Commands
    Commands() []Command                    // Available commands
    HandleCommand(ctx context.Context, cmd string, args []string) (*Result, error)
    
    // Events
    OnEvent(ctx context.Context, event Event) error  // Handle system events
    
    // UI
    UI() UI                                 // Optional UI component
}
```

## Supporting Types

### Command
```go
type Command struct {
    Name        string           // Command name (e.g., "scan")
    Description string           // Human-readable description
    Usage       string           // Usage example
    Handler     CommandHandler   // Optional direct handler
}

type CommandHandler func(ctx context.Context, args []string) (*Result, error)
```

### Result
```go
type Result struct {
    Success bool         // Command succeeded
    Data    interface{}  // Result data (can be any type)
    Error   error        // Error if failed
    UI      string       // Rendered UI output
}
```

### Event
```go
type Event struct {
    Type      string      // Event type (e.g., "file.changed", "git.commit")
    Source    string      // Extension that emitted event
    Timestamp time.Time   // When event occurred
    Data      interface{} // Event payload
}
```

### UI (Optional)
```go
type UI interface {
    Render() string                    // Render UI to string
    Update(msg interface{}) UI         // Update UI state
}
```

## Registry & Event Bus

Extensions interact with the system through:

```go
type Registry interface {
    Register(ext Extension) error
    Get(name string) (Extension, error)
    List() []Extension
    Enable(name string) error
    Disable(name string) error
    IsEnabled(name string) bool
}

type EventBus interface {
    Publish(ctx context.Context, event Event) error
    Subscribe(eventType string, handler EventHandler) error
}

type EventHandler func(ctx context.Context, event Event) error
```

## Complete Example

Here's a full extension implementation:

```go
package myext

import (
    "context"
    "fmt"
    
    "github.com/ivikasavnish/agenticide-go/internal/core/extension"
    "github.com/ivikasavnish/agenticide-go/internal/ui"
)

type MyExtension struct {
    enabled bool
}

// Constructor (exported as New)
func New() extension.Extension {
    return &MyExtension{
        enabled: false,
    }
}

// Identity methods
func (m *MyExtension) Name() string {
    return "my-extension"
}

func (m *MyExtension) Version() string {
    return "1.0.0"
}

func (m *MyExtension) Description() string {
    return "My custom extension for doing X"
}

func (m *MyExtension) Author() string {
    return "Your Name"
}

// Dependencies
func (m *MyExtension) Dependencies() []string {
    return []string{} // Or ["security", "git-ops"] if depends on others
}

// Lifecycle
func (m *MyExtension) Enable(ctx context.Context) error {
    m.enabled = true
    
    // Initialize resources
    // Connect to databases
    // Start background workers
    
    return nil
}

func (m *MyExtension) Disable(ctx context.Context) error {
    m.enabled = false
    
    // Cleanup resources
    // Close connections
    // Stop background workers
    
    return nil
}

func (m *MyExtension) IsEnabled() bool {
    return m.enabled
}

// Commands
func (m *MyExtension) Commands() []extension.Command {
    return []extension.Command{
        {
            Name:        "hello",
            Description: "Say hello",
            Usage:       "agenticide my hello [name]",
        },
        {
            Name:        "status",
            Description: "Show extension status",
            Usage:       "agenticide my status",
        },
    }
}

func (m *MyExtension) HandleCommand(ctx context.Context, cmd string, args []string) (*extension.Result, error) {
    switch cmd {
    case "hello":
        return m.handleHello(args)
    case "status":
        return m.handleStatus(args)
    default:
        return &extension.Result{
            Success: false,
            Error:   fmt.Errorf("unknown command: %s", cmd),
        }, nil
    }
}

func (m *MyExtension) handleHello(args []string) (*extension.Result, error) {
    name := "World"
    if len(args) > 0 {
        name = args[0]
    }
    
    message := fmt.Sprintf("Hello, %s!", name)
    
    return &extension.Result{
        Success: true,
        Data:    message,
        UI:      ui.Success(message),
    }, nil
}

func (m *MyExtension) handleStatus(args []string) (*extension.Result, error) {
    status := map[string]interface{}{
        "enabled": m.enabled,
        "version": m.Version(),
    }
    
    return &extension.Result{
        Success: true,
        Data:    status,
        UI:      ui.Info("Extension is running"),
    }, nil
}

// Events
func (m *MyExtension) OnEvent(ctx context.Context, event extension.Event) error {
    switch event.Type {
    case "file.changed":
        // React to file changes
        return m.handleFileChange(event)
    case "git.commit":
        // React to git commits
        return m.handleGitCommit(event)
    default:
        // Ignore unknown events
        return nil
    }
}

func (m *MyExtension) handleFileChange(event extension.Event) error {
    // Process file change
    return nil
}

func (m *MyExtension) handleGitCommit(event extension.Event) error {
    // Process git commit
    return nil
}

// UI (optional - can return nil)
func (m *MyExtension) UI() extension.UI {
    return nil // Or implement custom UI
}
```

## Lifecycle Flow

```
1. Extension registered:     registry.Register(ext)
2. Extension enabled:        ext.Enable(ctx)
3. Commands available:       ext.Commands()
4. Command invoked:          ext.HandleCommand(ctx, cmd, args)
5. Events received:          ext.OnEvent(ctx, event)
6. Extension disabled:       ext.Disable(ctx)
```

## Context Usage

Extensions receive `context.Context` for:

```go
func (m *MyExtension) Enable(ctx context.Context) error {
    // 1. Check cancellation
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }
    
    // 2. Pass to child operations
    result, err := someOperation(ctx)
    
    // 3. Get values from context
    logger := ctx.Value("logger").(*logger.Logger)
    config := ctx.Value("config").(*config.Config)
    
    return nil
}
```

## Event Publishing

Extensions can publish events:

```go
func (m *MyExtension) HandleCommand(ctx context.Context, cmd string, args []string) (*extension.Result, error) {
    // Get event bus from context
    bus := ctx.Value("eventBus").(extension.EventBus)
    
    // Publish event
    bus.Publish(ctx, extension.Event{
        Type:      "my-extension.action",
        Source:    m.Name(),
        Timestamp: time.Now(),
        Data:      map[string]string{"status": "complete"},
    })
    
    return &extension.Result{Success: true}, nil
}
```

## Event Subscription

Extensions can subscribe to events:

```go
func (m *MyExtension) Enable(ctx context.Context) error {
    bus := ctx.Value("eventBus").(extension.EventBus)
    
    // Subscribe to file changes
    bus.Subscribe("file.changed", func(ctx context.Context, event extension.Event) error {
        return m.OnEvent(ctx, event)
    })
    
    // Subscribe to multiple event types
    bus.Subscribe("git.*", func(ctx context.Context, event extension.Event) error {
        return m.OnEvent(ctx, event)
    })
    
    m.enabled = true
    return nil
}
```

## Beautiful UI

Use Agenticide's UI components:

```go
import (
    "github.com/ivikasavnish/agenticide-go/internal/ui"
    "github.com/ivikasavnish/agenticide-go/internal/ui/components"
)

func (m *MyExtension) handleStatus(args []string) (*extension.Result, error) {
    // Title
    output := ui.Title("🔍 Extension Status")
    output += "\n\n"
    
    // Table
    table := components.NewTable("Property", "Value")
    table.AddRow("Name", m.Name())
    table.AddRow("Version", m.Version())
    table.AddRow("Enabled", fmt.Sprintf("%v", m.enabled))
    output += table.Render()
    
    // Progress bar
    progress := components.NewProgressBar(100, 50)
    progress.SetCurrent(75).SetLabel("Loading...")
    output += "\n" + progress.Render()
    
    // List
    list := components.NewList()
    list.AddItem("Feature 1")
    list.AddItem("Feature 2")
    output += "\n" + list.Render()
    
    // Status messages
    output += "\n" + ui.Success("Everything working!")
    output += "\n" + ui.Info("Tip: Use --verbose for details")
    output += "\n" + ui.Warning("Update available")
    output += "\n" + ui.Error("Something failed")
    
    return &extension.Result{
        Success: true,
        UI:      output,
    }, nil
}
```

## Error Handling

Best practices:

```go
func (m *MyExtension) HandleCommand(ctx context.Context, cmd string, args []string) (*extension.Result, error) {
    // Validate input
    if len(args) == 0 {
        return &extension.Result{
            Success: false,
            Error:   fmt.Errorf("missing required argument"),
            UI:      ui.Error("Usage: agenticide my hello <name>"),
        }, nil
    }
    
    // Handle operation errors
    result, err := someOperation(args[0])
    if err != nil {
        return &extension.Result{
            Success: false,
            Error:   err,
            UI:      ui.Error(fmt.Sprintf("Operation failed: %v", err)),
        }, nil
    }
    
    // Return success
    return &extension.Result{
        Success: true,
        Data:    result,
        UI:      ui.Success("Operation completed"),
    }, nil
}
```

## Storage & Configuration

Access system storage and config:

```go
func (m *MyExtension) Enable(ctx context.Context) error {
    // Get storage
    store := ctx.Value("storage").(*storage.SQLiteStorage)
    
    // Create extension table
    _, err := store.Execute(`
        CREATE TABLE IF NOT EXISTS my_extension_data (
            id INTEGER PRIMARY KEY,
            key TEXT UNIQUE,
            value TEXT
        )
    `)
    if err != nil {
        return err
    }
    
    // Get config
    cfg := ctx.Value("config").(*config.Config)
    apiKey := cfg.GetString("my_extension.api_key")
    
    m.enabled = true
    return nil
}
```

## Testing Extensions

```go
package myext

import (
    "context"
    "testing"
    
    "github.com/ivikasavnish/agenticide-go/internal/core/extension"
)

func TestMyExtension(t *testing.T) {
    ext := New()
    
    // Test identity
    if ext.Name() != "my-extension" {
        t.Errorf("Expected name 'my-extension', got '%s'", ext.Name())
    }
    
    // Test lifecycle
    ctx := context.Background()
    if err := ext.Enable(ctx); err != nil {
        t.Fatalf("Enable failed: %v", err)
    }
    
    if !ext.IsEnabled() {
        t.Error("Extension should be enabled")
    }
    
    // Test command
    result, err := ext.HandleCommand(ctx, "hello", []string{"Test"})
    if err != nil {
        t.Fatalf("HandleCommand failed: %v", err)
    }
    
    if !result.Success {
        t.Error("Command should succeed")
    }
    
    // Test cleanup
    if err := ext.Disable(ctx); err != nil {
        t.Fatalf("Disable failed: %v", err)
    }
}
```

## Packaging for Marketplace

### 1. Create Manifest

`manifest.json`:
```json
{
  "name": "my-extension",
  "version": "1.0.0",
  "description": "My custom extension",
  "author": "Your Name",
  "repository": "https://github.com/you/my-extension",
  "tags": ["utility", "automation"],
  "category": "development",
  "install_url": "https://github.com/you/my-extension/releases/latest/download/my-extension"
}
```

### 2. Build Binary

```bash
go build -o my-extension ./cmd/my-extension
```

### 3. Create Release

```bash
gh release create v1.0.0 my-extension manifest.json
```

### 4. Users Install

```bash
agenticide install my-extension
```

## Best Practices

### ✅ DO
- Single responsibility per extension
- Fast Enable/Disable (< 100ms)
- Graceful error handling
- Beautiful UI with Lipgloss
- Comprehensive error messages
- Context cancellation support
- Clean up resources in Disable()

### ❌ DON'T
- Block in Enable/Disable
- Ignore context cancellation
- Leave resources open
- Hard-code paths/URLs
- Assume extensions are available
- Panic on errors

## Standard Event Types

System publishes these events:

- `file.created` - New file created
- `file.changed` - File modified
- `file.deleted` - File removed
- `git.commit` - Git commit made
- `git.push` - Git push completed
- `task.created` - Task added
- `task.completed` - Task finished
- `extension.enabled` - Extension activated
- `extension.disabled` - Extension deactivated

## System Dependencies

Available from context:

```go
ctx.Value("config")    // *config.Config
ctx.Value("logger")    // *logger.Logger
ctx.Value("storage")   // *storage.SQLiteStorage
ctx.Value("registry")  // extension.Registry
ctx.Value("eventBus")  // extension.EventBus
```

---

**Ready to build your extension?**

1. Implement the `Extension` interface
2. Create `manifest.json`
3. Build binary
4. Submit to marketplace
5. Users discover with `agenticide search`
6. Users install with `agenticide install <name>`
