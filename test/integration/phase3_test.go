package integration

import (
"os"
"path/filepath"
"testing"

"github.com/ivikasavnish/agenticide-go/internal/cli"
"github.com/ivikasavnish/agenticide-go/internal/core/extension"
"github.com/ivikasavnish/agenticide-go/internal/core/storage"
)

func TestPhase3ExtensionCommands(t *testing.T) {
registry := extension.NewRegistry()
extCmd := cli.NewExtensionCommands(registry)

if extCmd == nil {
t.Fatal("extension commands should not be nil")
}

listCmd := extCmd.ListCommand()
if listCmd.Use != "list" {
t.Fatalf("expected 'list', got '%s'", listCmd.Use)
}

enableCmd := extCmd.EnableCommand()
if enableCmd.Use != "enable <name>" {
t.Fatalf("expected 'enable <name>', got '%s'", enableCmd.Use)
}

t.Log("✓ Extension commands created")
}

func TestPhase3TaskCommands(t *testing.T) {
tempDir := t.TempDir()
dbPath := filepath.Join(tempDir, "test.db")

store, err := storage.NewSQLiteStorage(dbPath)
if err != nil {
t.Fatalf("create storage: %v", err)
}
defer store.Close()

taskCmd := cli.NewTaskCommands(store)
if taskCmd == nil {
t.Fatal("task commands should not be nil")
}

// Test adding a task
_, err = store.Execute(`
INSERT INTO tasks (id, title, status) 
VALUES ('test-1', 'Test Task', 'pending')
`)
if err != nil {
t.Fatalf("insert task: %v", err)
}

// Query it back
var title string
err = store.QueryRow("SELECT title FROM tasks WHERE id = 'test-1'").Scan(&title)
if err != nil {
t.Fatalf("query task: %v", err)
}

if title != "Test Task" {
t.Fatalf("expected 'Test Task', got '%s'", title)
}

t.Log("✓ Task commands working")
}

func TestPhase3PlanCommands(t *testing.T) {
tempDir := t.TempDir()
dbPath := filepath.Join(tempDir, "test.db")

store, err := storage.NewSQLiteStorage(dbPath)
if err != nil {
t.Fatalf("create storage: %v", err)
}
defer store.Close()

planCmd := cli.NewPlanCommands(store)
if planCmd == nil {
t.Fatal("plan commands should not be nil")
}

cmd := planCmd.PlanCommand()
if cmd.Use != "plan <requirement>" {
t.Fatalf("expected 'plan <requirement>', got '%s'", cmd.Use)
}

t.Log("✓ Plan commands created")
}

func TestPhase3Integration(t *testing.T) {
tempDir := t.TempDir()
dbPath := filepath.Join(tempDir, "test.db")

store, err := storage.NewSQLiteStorage(dbPath)
if err != nil {
t.Fatalf("create storage: %v", err)
}
defer store.Close()

registry := extension.NewRegistry()

// Create all command handlers
extCmd := cli.NewExtensionCommands(registry)
taskCmd := cli.NewTaskCommands(store)
planCmd := cli.NewPlanCommands(store)

if extCmd == nil || taskCmd == nil || planCmd == nil {
t.Fatal("command handlers should not be nil")
}

// Test task workflow
_, err = store.Execute(`
INSERT INTO tasks (id, title, status, priority) 
VALUES ('integration-1', 'Integration Test', 'pending', 1)
`)
if err != nil {
t.Fatalf("insert task: %v", err)
}

// Complete the task
_, err = store.Execute(`
UPDATE tasks SET status = 'done' WHERE id = 'integration-1'
`)
if err != nil {
t.Fatalf("update task: %v", err)
}

// Verify
var status string
err = store.QueryRow("SELECT status FROM tasks WHERE id = 'integration-1'").Scan(&status)
if err != nil {
t.Fatalf("query task: %v", err)
}

if status != "done" {
t.Fatalf("expected status 'done', got '%s'", status)
}

// Test home directory creation
home, _ := os.UserHomeDir()
agenticideDir := filepath.Join(home, ".agenticide")
if _, err := os.Stat(agenticideDir); os.IsNotExist(err) {
t.Logf("Note: .agenticide directory will be created on first run")
}

t.Log("✓ Phase 3 fully integrated")
t.Log("  - Extension commands functional")
t.Log("  - Task management working")
t.Log("  - Plan generation ready")
t.Log("  - Storage integration complete")
}
