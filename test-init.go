package main

import (
"fmt"
"os"
"github.com/ivikasavnish/agenticide-go/internal/core/config"
"github.com/ivikasavnish/agenticide-go/internal/cli"
)

func main() {
fmt.Println("Starting config init...")
cfg, err := config.NewConfig("")
if err != nil {
fmt.Fprintf(os.Stderr, "Config error: %v\n", err)
os.Exit(1)
}
fmt.Println("Config loaded")

setupComplete := cfg.GetBool("setup_complete")
fmt.Printf("setup_complete value: %v (type: %T)\n", setupComplete, setupComplete)

fmt.Println("Calling RequireSetup...")
if err := cli.RequireSetup(cfg); err != nil {
fmt.Fprintf(os.Stderr, "Setup error: %v\n", err)
os.Exit(1)
}
fmt.Println("RequireSetup completed")

fmt.Println("Final setup_complete:", cfg.GetBool("setup_complete"))
}
