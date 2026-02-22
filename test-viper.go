package main

import (
"fmt"
"os"
"path/filepath"
"github.com/spf13/viper"
)

func main() {
v := viper.New()
home, _ := os.UserHomeDir()
configDir := filepath.Join(home, ".agenticide")

v.AddConfigPath(configDir)
v.SetConfigName("config")
v.SetConfigType("yaml")

v.SetDefault("setup_complete", false)

fmt.Printf("Before ReadInConfig - setup_complete: %v\n", v.GetBool("setup_complete"))

if err := v.ReadInConfig(); err != nil {
if _, ok := err.(viper.ConfigFileNotFoundError); ok {
fmt.Println("Config file not found - using defaults")
} else {
fmt.Printf("Error reading config: %v\n", err)
}
} else {
fmt.Printf("Loaded config from: %s\n", v.ConfigFileUsed())
}

fmt.Printf("After ReadInConfig - setup_complete: %v\n", v.GetBool("setup_complete"))
}
