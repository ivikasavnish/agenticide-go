package main

import (
"fmt"

"github.com/ivikasavnish/agenticide-go/internal/ui"
"github.com/ivikasavnish/agenticide-go/internal/ui/components"
)

func main() {
fmt.Println(ui.Title("🚀 Agenticide Go - UI Demo"))
fmt.Println()

fmt.Println(ui.Title("Extension Registry"))
table := components.NewTable("Extension", "Status", "Version").
AddRow("Security Agent", "enabled", "1.0.0").
AddRow("Code Analyzer", "enabled", "1.0.0").
AddRow("Web Search", "disabled", "0.9.0").
AddRow("Project Runner", "enabled", "1.0.0").
AddRow("Cost Controller", "disabled", "0.5.0")
fmt.Println(table.Render())

fmt.Println(ui.Title("Analysis Progress"))
progress := components.NewProgressBar(100, 60).
SetCurrent(73).
SetLabel("Scanning project files...")
fmt.Println(progress.Render())
fmt.Println()

fmt.Println(ui.Title("Task Pipeline"))
list := components.NewList().
AddItem("Initialize core system").
AddItem("Load extensions").
AddItem("Configure storage").
AddItem("Start event bus").
AddItem("Run health checks")
fmt.Println(list.RenderWithStatus([]string{"done", "done", "in_progress", "pending", "pending"}))
fmt.Println()

fmt.Println(ui.Title("Security Scan"))
panel := components.NewPanel("").
SetContent(
ui.Success("No vulnerabilities found") + "\n" +
ui.Success("No secrets detected") + "\n" +
ui.Success("All dependencies up to date") + "\n" +
ui.Info("Last scan: 2 minutes ago"),
).
SetWidth(60)
fmt.Println(panel.Render())

fmt.Println(ui.Title("Code Coverage"))
chart := components.NewBarChart(40).
AddBar("Frontend", 87).
AddBar("Backend", 92).
AddBar("Tests", 45).
AddBar("Docs", 23)
fmt.Println(chart.Render())

fmt.Println()
fmt.Println(ui.Success("Phase 1 & 2 Complete!"))
fmt.Println(ui.Info("Ready for Phase 3: CLI Commands"))
}
