package integration

import (
"testing"

"github.com/ivikasavnish/agenticide-go/internal/ui"
"github.com/ivikasavnish/agenticide-go/internal/ui/components"
)

func TestPhase2Styles(t *testing.T) {
title := ui.RenderTitle("Test Title")
if title == "" {
t.Fatal("title should not be empty")
}

success := ui.RenderSuccess("Test success")
if success == "" {
t.Fatal("success should not be empty")
}

warning := ui.RenderWarning("Test warning")
if warning == "" {
t.Fatal("warning should not be empty")
}

t.Log("✓ Styles working")
t.Logf("\n%s", title)
t.Logf("%s", success)
t.Logf("%s", warning)
}

func TestPhase2Table(t *testing.T) {
table := components.NewTable("Name", "Status", "Version").
AddRow("Extension1", "enabled", "1.0.0").
AddRow("Extension2", "disabled", "2.1.0").
AddRow("Extension3", "enabled", "0.5.0")

output := table.Render()
if output == "" {
t.Fatal("table output should not be empty")
}

t.Log("✓ Table component working")
t.Logf("\n%s", output)
}

func TestPhase2Progress(t *testing.T) {
progress := components.NewProgressBar(100, 40).
SetCurrent(65).
SetLabel("Analyzing code...")

output := progress.Render()
if output == "" {
t.Fatal("progress output should not be empty")
}

t.Log("✓ Progress component working")
t.Logf("\n%s", output)
}

func TestPhase2List(t *testing.T) {
list := components.NewList().
AddItem("Task 1: Setup infrastructure").
AddItem("Task 2: Implement features").
AddItem("Task 3: Write tests").
SetSelected(1)

output := list.Render()
if output == "" {
t.Fatal("list output should not be empty")
}

statusList := list.RenderWithStatus([]string{"done", "in_progress", "pending"})
if statusList == "" {
t.Fatal("status list output should not be empty")
}

t.Log("✓ List component working")
t.Logf("\n%s", output)
t.Logf("\n%s", statusList)
}

func TestPhase2Panel(t *testing.T) {
panel := components.NewPanel("Security Scan Results").
SetContent("✓ No vulnerabilities found\n✓ No secrets detected\n✓ All dependencies up to date").
SetWidth(50)

output := panel.Render()
if output == "" {
t.Fatal("panel output should not be empty")
}

t.Log("✓ Panel component working")
t.Logf("\n%s", output)
}

func TestPhase2Chart(t *testing.T) {
chart := components.NewBarChart(30).
AddBar("Frontend", 45).
AddBar("Backend", 82).
AddBar("Tests", 23)

output := chart.Render()
if output == "" {
t.Fatal("chart output should not be empty")
}

sparkline := components.NewSparkLine(20, 5).
AddValue(10).AddValue(15).AddValue(25).
AddValue(20).AddValue(30).AddValue(45).
AddValue(40).AddValue(35).AddValue(50)

sparkOutput := sparkline.Render()
if sparkOutput == "" {
t.Fatal("sparkline output should not be empty")
}

t.Log("✓ Chart component working")
t.Logf("\n%s", output)
t.Logf("Sparkline: %s", sparkOutput)
}

func TestPhase2Integration(t *testing.T) {
t.Log(ui.RenderTitle("🎨 Phase 2 UI Demo"))
t.Log("")

table := components.NewTable("Extension", "Status", "Version").
AddRow("Security Agent", "enabled", "1.0.0").
AddRow("Code Analyzer", "enabled", "1.0.0").
AddRow("Web Search", "disabled", "0.9.0")
t.Logf("\n%s", table.Render())

progress := components.NewProgressBar(100, 50).
SetCurrent(75).
SetLabel("Scanning project...")
t.Logf("\n%s\n", progress.Render())

list := components.NewList().
AddItem("Initialize core system").
AddItem("Load extensions").
AddItem("Start event bus").
AddItem("Configure storage")
t.Logf("\n%s\n", list.RenderWithStatus([]string{"done", "done", "in_progress", "pending"}))

panel := components.NewPanel("System Status").
SetContent(ui.RenderSuccess("All systems operational") + "\n" +
ui.RenderInfo("Memory: 45MB / 50MB") + "\n" +
ui.RenderMuted("Uptime: 2h 15m")).
SetWidth(60)
t.Logf("\n%s", panel.Render())

t.Log("\n" + ui.RenderSuccess("Phase 2 UI Framework Complete"))
}
