package main

import (
"fmt"
"os"
"github.com/ivikasavnish/agenticide-go/internal/core/config"
"github.com/ivikasavnish/agenticide-go/internal/cli"
)

func main() {
cfg, err := config.NewConfig("")
if err != nil {
fmt.Fprintf(os.Stderr, "Config error: %v\n", err)
os.Exit(1)
}

fmt.Printf("setup_complete: %v\n", cfg.GetBool("setup_complete"))
fmt.Printf("CheckSetup: %v\n", cli.CheckSetup(cfg))

if !cli.CheckSetup(cfg) {
fmt.Println("Setup is NOT complete - should run wizard")
if err := cli.FirstRunSetup(cfg); err != nil {
fmt.Fprintf(os.Stderr, "Setup failed: %v\n", err)
}
} else {
fmt.Println("Setup IS complete - skip wizard")
}
}
