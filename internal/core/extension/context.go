package extension

import (
"context"
"sync"
"time"
)

type ExtensionContext struct {
ctx      context.Context
cancel   context.CancelFunc
values   map[string]interface{}
mu       sync.RWMutex
metadata map[string]string
}

func NewContext(parent context.Context, timeout time.Duration) *ExtensionContext {
if parent == nil {
parent = context.Background()
}

ctx, cancel := context.WithTimeout(parent, timeout)

return &ExtensionContext{
ctx:      ctx,
cancel:   cancel,
values:   make(map[string]interface{}),
metadata: make(map[string]string),
}
}

func NewContextWithCancel(parent context.Context) *ExtensionContext {
if parent == nil {
parent = context.Background()
}

ctx, cancel := context.WithCancel(parent)

return &ExtensionContext{
ctx:      ctx,
cancel:   cancel,
values:   make(map[string]interface{}),
metadata: make(map[string]string),
}
}

func (ec *ExtensionContext) Set(key string, value interface{}) {
ec.mu.Lock()
defer ec.mu.Unlock()
ec.values[key] = value
}

func (ec *ExtensionContext) Get(key string) (interface{}, bool) {
ec.mu.RLock()
defer ec.mu.RUnlock()
val, ok := ec.values[key]
return val, ok
}

func (ec *ExtensionContext) SetMetadata(key, value string) {
ec.mu.Lock()
defer ec.mu.Unlock()
ec.metadata[key] = value
}

func (ec *ExtensionContext) GetMetadata(key string) (string, bool) {
ec.mu.RLock()
defer ec.mu.RUnlock()
val, ok := ec.metadata[key]
return val, ok
}

func (ec *ExtensionContext) Context() context.Context {
return ec.ctx
}

func (ec *ExtensionContext) Cancel() {
if ec.cancel != nil {
ec.cancel()
}
}

func (ec *ExtensionContext) Done() <-chan struct{} {
return ec.ctx.Done()
}

func (ec *ExtensionContext) Err() error {
return ec.ctx.Err()
}
