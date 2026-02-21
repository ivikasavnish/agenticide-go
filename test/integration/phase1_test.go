package integration

import (
"context"
"os"
"path/filepath"
"testing"
"time"

"github.com/ivikasavnish/agenticide-go/internal/core/config"
"github.com/ivikasavnish/agenticide-go/internal/core/extension"
"github.com/ivikasavnish/agenticide-go/internal/core/logger"
"github.com/ivikasavnish/agenticide-go/internal/core/storage"
)

func TestPhase1Context(t *testing.T) {
ctx := extension.NewContext(context.Background(), 5*time.Second)

ctx.Set("test_key", "test_value")
val, ok := ctx.Get("test_key")
if !ok {
t.Fatal("expected key to exist")
}
if val != "test_value" {
t.Fatalf("expected 'test_value', got %v", val)
}

ctx.SetMetadata("user", "testuser")
user, ok := ctx.GetMetadata("user")
if !ok {
t.Fatal("expected metadata to exist")
}
if user != "testuser" {
t.Fatalf("expected 'testuser', got %s", user)
}

select {
case <-ctx.Done():
t.Fatal("context should not be done yet")
default:
}

ctx.Cancel()
<-ctx.Done()

if ctx.Err() == nil {
t.Fatal("expected error after cancel")
}

t.Log("✓ Context system working")
}

func TestPhase1Storage(t *testing.T) {
tempDir := t.TempDir()
dbPath := filepath.Join(tempDir, "test.db")

store, err := storage.NewSQLiteStorage(dbPath)
if err != nil {
t.Fatalf("create storage: %v", err)
}
defer store.Close()

_, err = store.Execute(
"INSERT INTO tasks (id, title, status) VALUES (?, ?, ?)",
"test-task", "Test Task", "pending",
)
if err != nil {
t.Fatalf("insert task: %v", err)
}

var title, status string
err = store.QueryRow(
"SELECT title, status FROM tasks WHERE id = ?",
"test-task",
).Scan(&title, &status)
if err != nil {
t.Fatalf("query task: %v", err)
}

if title != "Test Task" || status != "pending" {
t.Fatalf("expected 'Test Task' and 'pending', got '%s' and '%s'", title, status)
}

t.Log("✓ SQLite storage working")
}

func TestPhase1Config(t *testing.T) {
tempDir := t.TempDir()
configPath := filepath.Join(tempDir, "config.yaml")

os.WriteFile(configPath, []byte(`
log_level: debug
storage_path: /tmp/test.db
ultraloop_max_retries: 20
`), 0644)

cfg, err := config.NewConfig(configPath)
if err != nil {
t.Fatalf("create config: %v", err)
}

if cfg.GetString("log_level") != "debug" {
t.Fatalf("expected log_level=debug, got %s", cfg.GetString("log_level"))
}

if cfg.GetInt("ultraloop_max_retries") != 20 {
t.Fatalf("expected ultraloop_max_retries=20, got %d", cfg.GetInt("ultraloop_max_retries"))
}

cfg.Set("test_key", "test_value")
if cfg.GetString("test_key") != "test_value" {
t.Fatal("set/get failed")
}

t.Log("✓ Config system working")
}

func TestPhase1Logger(t *testing.T) {
log, err := logger.NewLogger("info")
if err != nil {
t.Fatalf("create logger: %v", err)
}
defer log.Sync()

log.Info("test info message")
log.Warn("test warning message")

namedLog := log.Named("test-extension")
namedLog.Info("extension-specific log")

sugar := log.Sugar()
sugar.Infow("structured log",
"key1", "value1",
"key2", 123,
)

t.Log("✓ Logger system working")
}

func TestPhase1Integration(t *testing.T) {
tempDir := t.TempDir()

cfg, err := config.NewConfig("")
if err != nil {
t.Fatalf("create config: %v", err)
}
cfg.Set("storage_path", filepath.Join(tempDir, "test.db"))

log, err := logger.NewLogger(cfg.GetString("log_level"))
if err != nil {
t.Fatalf("create logger: %v", err)
}
defer log.Sync()

store, err := storage.NewSQLiteStorage(cfg.GetString("storage_path"))
if err != nil {
t.Fatalf("create storage: %v", err)
}
defer store.Close()

ctx := extension.NewContext(context.Background(), 10*time.Second)
ctx.Set("config", cfg)
ctx.Set("logger", log)
ctx.Set("storage", store)

if _, ok := ctx.Get("config"); !ok {
t.Fatal("config not in context")
}
if _, ok := ctx.Get("logger"); !ok {
t.Fatal("logger not in context")
}
if _, ok := ctx.Get("storage"); !ok {
t.Fatal("storage not in context")
}

log.Info("Phase 1 integration test complete",
zap.String("storage", cfg.GetString("storage_path")),
)

t.Log("✓ Phase 1 fully integrated")
}
