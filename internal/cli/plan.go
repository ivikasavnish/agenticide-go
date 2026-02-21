package cli

import (
"fmt"
"strings"
"time"

"github.com/spf13/cobra"

"github.com/ivikasavnish/agenticide-go/internal/core/storage"
"github.com/ivikasavnish/agenticide-go/internal/ui"
"github.com/ivikasavnish/agenticide-go/internal/ui/components"
)

type PlanCommands struct {
storage *storage.SQLiteStorage
}

func NewPlanCommands(storage *storage.SQLiteStorage) *PlanCommands {
return &PlanCommands{
storage: storage,
}
}

func (pc *PlanCommands) PlanCommand() *cobra.Command {
return &cobra.Command{
Use:   "plan <requirement>",
Short: "Generate implementation plan with task decomposition",
Args:  cobra.MinimumNArgs(1),
Run: func(cmd *cobra.Command, args []string) {
requirement := strings.Join(args, " ")

fmt.Println(ui.Title("📝 Generating Plan"))
fmt.Println()
fmt.Printf("Requirement: %s\n", ui.Highlight(requirement))
fmt.Println()

progress := components.NewProgressBar(100, 50).
SetLabel("Analyzing requirement...")

for i := 0; i <= 100; i += 20 {
progress.SetCurrent(i)
fmt.Printf("\r%s", progress.Render())
time.Sleep(200 * time.Millisecond)
}
fmt.Println()
fmt.Println()

// Generate decomposed tasks (placeholder - would use AI)
tasks := pc.decomposeTasks(requirement)

fmt.Println(ui.Title("🌳 Task Breakdown"))
fmt.Println()

list := components.NewList()
for _, task := range tasks {
list.AddItem(task.Title)
}

statuses := make([]string, len(tasks))
for i := range statuses {
statuses[i] = "pending"
}
fmt.Println(list.RenderWithStatus(statuses))
fmt.Println()

// Save tasks to database
fmt.Print("Save tasks to database? (y/n): ")
var response string
fmt.Scanln(&response)

if strings.ToLower(response) == "y" {
for _, task := range tasks {
pc.storage.Execute(`
INSERT INTO tasks (id, title, description, status, priority)
VALUES (?, ?, ?, 'pending', ?)
`, task.ID, task.Title, task.Description, task.Priority)

for _, dep := range task.Dependencies {
pc.storage.Execute(`
INSERT INTO task_dependencies (task_id, depends_on)
VALUES (?, ?)
`, task.ID, dep)
}
}

fmt.Println()
fmt.Println(ui.Success(fmt.Sprintf("Saved %d tasks", len(tasks))))
fmt.Println(ui.Muted("Use 'agenticide task list' to see all tasks"))
}
},
}
}

type Task struct {
ID           string
Title        string
Description  string
Priority     int
Dependencies []string
}

func (pc *PlanCommands) decomposeTasks(requirement string) []Task {
// Placeholder task decomposition
// In production, this would use AI to intelligently break down requirements

tasks := []Task{
{
ID:          fmt.Sprintf("task-%d-1", time.Now().Unix()),
Title:       "Research and design",
Description: fmt.Sprintf("Research best practices for: %s", requirement),
Priority:    3,
},
{
ID:          fmt.Sprintf("task-%d-2", time.Now().Unix()),
Title:       "Implement core functionality",
Description: "Build the main implementation",
Priority:    2,
Dependencies: []string{fmt.Sprintf("task-%d-1", time.Now().Unix())},
},
{
ID:          fmt.Sprintf("task-%d-3", time.Now().Unix()),
Title:       "Write tests",
Description: "Add comprehensive test coverage",
Priority:    2,
Dependencies: []string{fmt.Sprintf("task-%d-2", time.Now().Unix())},
},
{
ID:          fmt.Sprintf("task-%d-4", time.Now().Unix()),
Title:       "Documentation",
Description: "Write user and developer documentation",
Priority:    1,
Dependencies: []string{fmt.Sprintf("task-%d-2", time.Now().Unix())},
},
}

return tasks
}
