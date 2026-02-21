package cli

import (
"database/sql"
"fmt"
"time"

"github.com/spf13/cobra"

"github.com/ivikasavnish/agenticide-go/internal/core/storage"
"github.com/ivikasavnish/agenticide-go/internal/ui"
"github.com/ivikasavnish/agenticide-go/internal/ui/components"
)

type TaskCommands struct {
storage *storage.SQLiteStorage
}

func NewTaskCommands(storage *storage.SQLiteStorage) *TaskCommands {
return &TaskCommands{
storage: storage,
}
}

func (tc *TaskCommands) ListCommand() *cobra.Command {
return &cobra.Command{
Use:   "list",
Short: "List all tasks",
Run: func(cmd *cobra.Command, args []string) {
fmt.Println(ui.Title("📋 Tasks"))
fmt.Println()

rows, err := tc.storage.Query(`
SELECT id, title, status, priority, created_at 
FROM tasks 
ORDER BY priority DESC, created_at ASC
`)
if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Query failed: %v", err)))
return
}
defer rows.Close()

table := components.NewTable("ID", "Title", "Status", "Priority", "Created")
hasRows := false

for rows.Next() {
hasRows = true
var id, title, status string
var priority int
var createdAt string

if err := rows.Scan(&id, &title, &status, &priority, &createdAt); err != nil {
continue
}

statusText := status
switch status {
case "done":
statusText = ui.Success("done")
case "in_progress":
statusText = ui.Info("in progress")
case "pending":
statusText = ui.Muted("pending")
case "blocked":
statusText = ui.Warning("blocked")
}

table.AddRow(id, title, statusText, fmt.Sprintf("%d", priority), createdAt[:10])
}

if !hasRows {
fmt.Println(ui.Muted("  No tasks found"))
return
}

fmt.Println(table.Render())
},
}
}

func (tc *TaskCommands) AddCommand() *cobra.Command {
return &cobra.Command{
Use:   "add <description>",
Short: "Add a new task",
Args:  cobra.MinimumNArgs(1),
Run: func(cmd *cobra.Command, args []string) {
title := args[0]
description := ""
if len(args) > 1 {
description = args[1]
}

id := fmt.Sprintf("task-%d", time.Now().Unix())

_, err := tc.storage.Execute(`
INSERT INTO tasks (id, title, description, status, priority)
VALUES (?, ?, ?, 'pending', 0)
`, id, title, description)

if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Failed to add task: %v", err)))
return
}

fmt.Println(ui.Success(fmt.Sprintf("Task added: %s", id)))
},
}
}

func (tc *TaskCommands) CompleteCommand() *cobra.Command {
return &cobra.Command{
Use:   "complete <id>",
Short: "Mark task as complete",
Args:  cobra.ExactArgs(1),
Run: func(cmd *cobra.Command, args []string) {
id := args[0]

result, err := tc.storage.Execute(`
UPDATE tasks 
SET status = 'done', completed_at = CURRENT_TIMESTAMP
WHERE id = ?
`, id)

if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Failed: %v", err)))
return
}

rows, _ := result.RowsAffected()
if rows == 0 {
fmt.Println(ui.Warning(fmt.Sprintf("Task not found: %s", id)))
return
}

fmt.Println(ui.Success(fmt.Sprintf("Task completed: %s", id)))
},
}
}

func (tc *TaskCommands) GraphCommand() *cobra.Command {
return &cobra.Command{
Use:   "graph",
Short: "Show task dependency graph",
Run: func(cmd *cobra.Command, args []string) {
fmt.Println(ui.Title("🌳 Task Dependency Graph"))
fmt.Println()

rows, err := tc.storage.Query(`
SELECT t.id, t.title, t.status, 
       GROUP_CONCAT(td.depends_on) as deps
FROM tasks t
LEFT JOIN task_dependencies td ON t.id = td.task_id
GROUP BY t.id
ORDER BY t.priority DESC
`)
if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Query failed: %v", err)))
return
}
defer rows.Close()

for rows.Next() {
var id, title, status string
var deps sql.NullString

if err := rows.Scan(&id, &title, &status, &deps); err != nil {
continue
}

statusIcon := "○"
switch status {
case "done":
statusIcon = ui.Success("")
case "in_progress":
statusIcon = ui.Info("◐")
case "blocked":
statusIcon = ui.Warning("⚠")
}

fmt.Printf("  %s %s\n", statusIcon, title)
if deps.Valid && deps.String != "" {
fmt.Printf("    %s depends on: %s\n", ui.Muted("↳"), ui.Muted(deps.String))
}
}
},
}
}
