package main

import (
"fmt"
"os"
"github.com/ivikasavnish/agenticide-go/internal/core/config"
)

func main() {
cfg, _ := config.NewConfig("")
cfg.Set("user.email", "test@example.com")
cfg.Set("user.phone", "555-1234")
cfg.Set("ai.provider", "openai")
cfg.Set("ai.api_key", "sk-test")
cfg.Set("setup_complete", true)

fmt.Println("Attempting to save...")
if err := cfg.Save(); err != nil {
fmt.Printf("Save error: %v\n", err)
} else {
fmt.Println("Save successful!")
}
}
