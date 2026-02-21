package extension

import (
"context"
"fmt"
"sync"
)

type registry struct {
extensions map[string]Extension
enabled    map[string]bool
mu         sync.RWMutex
}

func NewRegistry() Registry {
return &registry{
extensions: make(map[string]Extension),
enabled:    make(map[string]bool),
}
}

func (r *registry) Register(ext Extension) error {
r.mu.Lock()
defer r.mu.Unlock()

name := ext.Name()
if _, exists := r.extensions[name]; exists {
return fmt.Errorf("extension already registered: %s", name)
}

r.extensions[name] = ext
r.enabled[name] = false
return nil
}

func (r *registry) Get(name string) (Extension, error) {
r.mu.RLock()
defer r.mu.RUnlock()

ext, exists := r.extensions[name]
if !exists {
return nil, fmt.Errorf("extension not found: %s", name)
}
return ext, nil
}

func (r *registry) List() []Extension {
r.mu.RLock()
defer r.mu.RUnlock()

extensions := make([]Extension, 0, len(r.extensions))
for _, ext := range r.extensions {
extensions = append(extensions, ext)
}
return extensions
}

func (r *registry) Enable(name string) error {
r.mu.Lock()
defer r.mu.Unlock()

ext, exists := r.extensions[name]
if !exists {
return fmt.Errorf("extension not found: %s", name)
}

if r.enabled[name] {
return nil
}

if err := ext.Enable(context.Background()); err != nil {
return err
}

r.enabled[name] = true
return nil
}

func (r *registry) Disable(name string) error {
r.mu.Lock()
defer r.mu.Unlock()

ext, exists := r.extensions[name]
if !exists {
return fmt.Errorf("extension not found: %s", name)
}

if !r.enabled[name] {
return nil
}

if err := ext.Disable(context.Background()); err != nil {
return err
}

r.enabled[name] = false
return nil
}

func (r *registry) IsEnabled(name string) bool {
r.mu.RLock()
defer r.mu.RUnlock()
return r.enabled[name]
}
