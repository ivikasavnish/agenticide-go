package extension

import (
"context"
"time"
)

type Extension interface {
Name() string
Version() string
Description() string
Author() string
Dependencies() []string
Enable(ctx context.Context) error
Disable(ctx context.Context) error
IsEnabled() bool
Commands() []Command
HandleCommand(ctx context.Context, cmd string, args []string) (*Result, error)
OnEvent(ctx context.Context, event Event) error
UI() UI
}

type Command struct {
Name        string
Description string
Usage       string
Handler     CommandHandler
}

type CommandHandler func(ctx context.Context, args []string) (*Result, error)

type Result struct {
Success bool
Data    interface{}
Error   error
UI      string
}

type Event struct {
Type      string
Source    string
Timestamp time.Time
Data      interface{}
}

type UI interface {
Render() string
Update(msg interface{}) UI
}

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
