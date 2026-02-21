package config

import (
"fmt"
"os"
"path/filepath"
"sync"

"github.com/spf13/viper"
)

type Config struct {
v  *viper.Viper
mu sync.RWMutex
}

func NewConfig(configPath string) (*Config, error) {
v := viper.New()

if configPath != "" {
v.SetConfigFile(configPath)
} else {
home, err := os.UserHomeDir()
if err != nil {
return nil, fmt.Errorf("get home dir: %w", err)
}

configDir := filepath.Join(home, ".agenticide")
if err := os.MkdirAll(configDir, 0755); err != nil {
return nil, fmt.Errorf("create config dir: %w", err)
}

v.AddConfigPath(configDir)
v.SetConfigName("config")
v.SetConfigType("yaml")
}

v.SetEnvPrefix("AGENTICIDE")
v.AutomaticEnv()

v.SetDefault("log_level", "info")
v.SetDefault("storage_path", filepath.Join(os.TempDir(), "agenticide.db"))
v.SetDefault("extensions_enabled", []string{})
v.SetDefault("ultraloop_max_retries", 10)
v.SetDefault("ultrathink_timeout", "5m")

if err := v.ReadInConfig(); err != nil {
if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
return nil, fmt.Errorf("read config: %w", err)
}
}

return &Config{v: v}, nil
}

func (c *Config) Get(key string) interface{} {
c.mu.RLock()
defer c.mu.RUnlock()
return c.v.Get(key)
}

func (c *Config) GetString(key string) string {
c.mu.RLock()
defer c.mu.RUnlock()
return c.v.GetString(key)
}

func (c *Config) GetInt(key string) int {
c.mu.RLock()
defer c.mu.RUnlock()
return c.v.GetInt(key)
}

func (c *Config) GetBool(key string) bool {
c.mu.RLock()
defer c.mu.RUnlock()
return c.v.GetBool(key)
}

func (c *Config) GetStringSlice(key string) []string {
c.mu.RLock()
defer c.mu.RUnlock()
return c.v.GetStringSlice(key)
}

func (c *Config) Set(key string, value interface{}) {
c.mu.Lock()
defer c.mu.Unlock()
c.v.Set(key, value)
}

func (c *Config) WriteConfig() error {
c.mu.Lock()
defer c.mu.Unlock()
return c.v.WriteConfig()
}
