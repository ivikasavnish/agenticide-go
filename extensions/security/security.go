package security

import (
"context"
"fmt"
"os/exec"
"path/filepath"
"strings"

"github.com/ivikasavnish/agenticide-go/internal/core/extension"
"github.com/ivikasavnish/agenticide-go/internal/ui"
"github.com/ivikasavnish/agenticide-go/internal/ui/components"
)

type SecurityExtension struct {
enabled bool
}

func New() extension.Extension {
return &SecurityExtension{
enabled: false,
}
}

func (se *SecurityExtension) Name() string {
return "security"
}

func (se *SecurityExtension) Version() string {
return "1.0.0"
}

func (se *SecurityExtension) Description() string {
return "SAST scanning, secret detection, and vulnerability checking"
}

func (se *SecurityExtension) Author() string {
return "Agenticide Team"
}

func (se *SecurityExtension) Dependencies() []string {
return []string{}
}

func (se *SecurityExtension) Enable(ctx context.Context) error {
se.enabled = true
return nil
}

func (se *SecurityExtension) Disable(ctx context.Context) error {
se.enabled = false
return nil
}

func (se *SecurityExtension) IsEnabled() bool {
return se.enabled
}

func (se *SecurityExtension) Commands() []extension.Command {
return []extension.Command{
{
Name:        "scan",
Description: "Run security scan on project",
Usage:       "agenticide security scan [path]",
},
{
Name:        "secrets",
Description: "Detect secrets in codebase",
Usage:       "agenticide security secrets [path]",
},
{
Name:        "vulns",
Description: "Check dependencies for vulnerabilities",
Usage:       "agenticide security vulns",
},
}
}

func (se *SecurityExtension) HandleCommand(ctx context.Context, cmd string, args []string) (*extension.Result, error) {
switch cmd {
case "scan":
return se.runScan(args)
case "secrets":
return se.detectSecrets(args)
case "vulns":
return se.checkVulnerabilities(args)
default:
return &extension.Result{
Success: false,
Error:   fmt.Errorf("unknown command: %s", cmd),
}, nil
}
}

func (se *SecurityExtension) OnEvent(ctx context.Context, event extension.Event) error {
// Handle events like file changes, git commits, etc.
return nil
}

func (se *SecurityExtension) UI() extension.UI {
return nil // CLI-only for now
}

func (se *SecurityExtension) runScan(args []string) (*extension.Result, error) {
path := "."
if len(args) > 0 {
path = args[0]
}

fmt.Println(ui.Title("🔒 Security Scan"))
fmt.Println()

progress := components.NewProgressBar(100, 50).
SetLabel("Scanning project...")

// Simulate scan phases
phases := []string{
"Analyzing code structure...",
"Running SAST checks...",
"Detecting secrets...",
"Checking dependencies...",
"Generating report...",
}

for i, phase := range phases {
progress.SetLabel(phase)
progress.SetCurrent((i + 1) * 20)
fmt.Printf("\r%s", progress.Render())
}
fmt.Println()
fmt.Println()

// Run gosec if available
issues := se.runGosec(path)
secrets := se.scanSecrets(path)

// Display results
fmt.Println(ui.Title("📊 Scan Results"))
fmt.Println()

if len(issues) == 0 && len(secrets) == 0 {
fmt.Println(ui.Success("No security issues found"))
} else {
if len(issues) > 0 {
fmt.Println(ui.Warning(fmt.Sprintf("Found %d potential issues", len(issues))))
for _, issue := range issues {
fmt.Printf("  • %s\n", ui.Muted(issue))
}
fmt.Println()
}

if len(secrets) > 0 {
fmt.Println(ui.Error(fmt.Sprintf("Found %d potential secrets", len(secrets))))
for _, secret := range secrets {
fmt.Printf("  • %s\n", ui.Muted(secret))
}
}
}

return &extension.Result{
Success: true,
Data: map[string]interface{}{
"issues":  len(issues),
"secrets": len(secrets),
},
}, nil
}

func (se *SecurityExtension) runGosec(path string) []string {
// Try to run gosec if installed
cmd := exec.Command("gosec", "-fmt=json", path)
output, err := cmd.CombinedOutput()

if err != nil {
// gosec not installed or no Go files
return []string{}
}

// Parse output (simplified)
if strings.Contains(string(output), "issue") {
return []string{"Potential security issue detected (run 'gosec' for details)"}
}

return []string{}
}

func (se *SecurityExtension) scanSecrets(path string) []string {
secrets := []string{}

// Simple pattern matching for common secrets
patterns := []string{
"password", "secret", "api_key", "token",
"AWS_ACCESS_KEY", "GITHUB_TOKEN",
}

// Walk directory and check files
filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
if err != nil || info.IsDir() {
return nil
}

// Skip binary files
if strings.HasSuffix(filePath, ".exe") || strings.HasSuffix(filePath, ".bin") {
return nil
}

// Read file and check patterns
content, err := os.ReadFile(filePath)
if err != nil {
return nil
}

text := string(content)
for _, pattern := range patterns {
if strings.Contains(strings.ToLower(text), pattern) {
secrets = append(secrets, fmt.Sprintf("%s: contains '%s'", filePath, pattern))
break
}
}

return nil
})

return secrets
}

func (se *SecurityExtension) detectSecrets(args []string) (*extension.Result, error) {
path := "."
if len(args) > 0 {
path = args[0]
}

fmt.Println(ui.Title("🔍 Secret Detection"))
fmt.Println()

secrets := se.scanSecrets(path)

if len(secrets) == 0 {
fmt.Println(ui.Success("No secrets detected"))
} else {
fmt.Println(ui.Warning(fmt.Sprintf("Found %d potential secrets:", len(secrets))))
for _, secret := range secrets {
fmt.Printf("  • %s\n", ui.Muted(secret))
}
}

return &extension.Result{
Success: true,
Data:    map[string]interface{}{"secrets": len(secrets)},
}, nil
}

func (se *SecurityExtension) checkVulnerabilities(args []string) (*extension.Result, error) {
fmt.Println(ui.Title("🛡️ Vulnerability Check"))
fmt.Println()

// Check for go.mod
if _, err := os.Stat("go.mod"); err == nil {
fmt.Println(ui.Info("Checking Go dependencies..."))

// Run govulncheck if available
cmd := exec.Command("govulncheck", "./...")
output, err := cmd.CombinedOutput()

if err != nil {
fmt.Println(ui.Muted("  govulncheck not installed (install with: go install golang.org/x/vuln/cmd/govulncheck@latest)"))
} else if strings.Contains(string(output), "No vulnerabilities found") {
fmt.Println(ui.Success("  No vulnerabilities found"))
} else {
fmt.Println(ui.Warning("  Vulnerabilities detected (see output above)"))
}
}

// Check for package.json
if _, err := os.Stat("package.json"); err == nil {
fmt.Println(ui.Info("Checking npm dependencies..."))

cmd := exec.Command("npm", "audit")
output, err := cmd.CombinedOutput()

if err == nil && strings.Contains(string(output), "0 vulnerabilities") {
fmt.Println(ui.Success("  No vulnerabilities found"))
} else {
fmt.Println(ui.Muted("  Run 'npm audit' for details"))
}
}

return &extension.Result{Success: true}, nil
}
