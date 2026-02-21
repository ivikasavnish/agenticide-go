package event

import (
"context"
"sync"

"github.com/ivikasavnish/agenticide-go/internal/core/extension"
)

type bus struct {
handlers map[string][]extension.EventHandler
mu       sync.RWMutex
}

func NewBus() extension.EventBus {
return &bus{
handlers: make(map[string][]extension.EventHandler),
}
}

func (b *bus) Publish(ctx context.Context, event extension.Event) error {
b.mu.RLock()
handlers, exists := b.handlers[event.Type]
b.mu.RUnlock()

if !exists {
return nil
}

for _, handler := range handlers {
if err := handler(ctx, event); err != nil {
return err
}
}

return nil
}

func (b *bus) Subscribe(eventType string, handler extension.EventHandler) error {
b.mu.Lock()
defer b.mu.Unlock()

if _, exists := b.handlers[eventType]; !exists {
b.handlers[eventType] = make([]extension.EventHandler, 0)
}

b.handlers[eventType] = append(b.handlers[eventType], handler)
return nil
}
