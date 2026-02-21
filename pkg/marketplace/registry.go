package marketplace

import (
"encoding/json"
"fmt"
"io"
"net/http"
"os"
"path/filepath"
"strings"
)

type ExtensionInfo struct {
Name        string   `json:"name"`
Version     string   `json:"version"`
Description string   `json:"description"`
Author      string   `json:"author"`
Repository  string   `json:"repository"`
Tags        []string `json:"tags"`
InstallURL  string   `json:"install_url"`
Category    string   `json:"category"`
}

type Registry struct {
baseURL string
}

func NewRegistry() *Registry {
return &Registry{
baseURL: "https://registry.agenticide.dev/v1",
}
}

// Search for extensions in the marketplace
func (r *Registry) Search(query string) ([]ExtensionInfo, error) {
// In production, this would query the registry API
// For now, return curated list

allExtensions := r.getBuiltInExtensions()

if query == "" {
return allExtensions, nil
}

// Filter by query
var results []ExtensionInfo
query = strings.ToLower(query)

for _, ext := range allExtensions {
if strings.Contains(strings.ToLower(ext.Name), query) ||
strings.Contains(strings.ToLower(ext.Description), query) ||
containsTag(ext.Tags, query) {
results = append(results, ext)
}
}

return results, nil
}

func containsTag(tags []string, query string) bool {
for _, tag := range tags {
if strings.Contains(strings.ToLower(tag), query) {
return true
}
}
return false
}

// Get extension details
func (r *Registry) Get(name string) (*ExtensionInfo, error) {
extensions := r.getBuiltInExtensions()

for _, ext := range extensions {
if ext.Name == name {
return &ext, nil
}
}

return nil, fmt.Errorf("extension not found: %s", name)
}

// Install extension from registry
func (r *Registry) Install(name string) error {
ext, err := r.Get(name)
if err != nil {
return err
}

// Download extension
if ext.InstallURL == "" {
return fmt.Errorf("no install URL for extension: %s", name)
}

home, _ := os.UserHomeDir()
extDir := filepath.Join(home, ".agenticide", "extensions", name)

if err := os.MkdirAll(extDir, 0755); err != nil {
return fmt.Errorf("failed to create extension directory: %w", err)
}

// Download extension binary/script
resp, err := http.Get(ext.InstallURL)
if err != nil {
return fmt.Errorf("failed to download extension: %w", err)
}
defer resp.Body.Close()

extFile := filepath.Join(extDir, name)
out, err := os.Create(extFile)
if err != nil {
return fmt.Errorf("failed to create extension file: %w", err)
}
defer out.Close()

_, err = io.Copy(out, resp.Body)
if err != nil {
return fmt.Errorf("failed to save extension: %w", err)
}

// Make executable
os.Chmod(extFile, 0755)

// Save metadata
metaFile := filepath.Join(extDir, "manifest.json")
metaData, _ := json.MarshalIndent(ext, "", "  ")
os.WriteFile(metaFile, metaData, 0644)

return nil
}

// List installed extensions
func (r *Registry) ListInstalled() ([]ExtensionInfo, error) {
home, _ := os.UserHomeDir()
extDir := filepath.Join(home, ".agenticide", "extensions")

if _, err := os.Stat(extDir); os.IsNotExist(err) {
return []ExtensionInfo{}, nil
}

entries, err := os.ReadDir(extDir)
if err != nil {
return nil, err
}

var installed []ExtensionInfo

for _, entry := range entries {
if !entry.IsDir() {
continue
}

metaFile := filepath.Join(extDir, entry.Name(), "manifest.json")
if data, err := os.ReadFile(metaFile); err == nil {
var info ExtensionInfo
if err := json.Unmarshal(data, &info); err == nil {
installed = append(installed, info)
}
}
}

return installed, nil
}

// Uninstall extension
func (r *Registry) Uninstall(name string) error {
home, _ := os.UserHomeDir()
extDir := filepath.Join(home, ".agenticide", "extensions", name)

return os.RemoveAll(extDir)
}

// Get built-in extension catalog
func (r *Registry) getBuiltInExtensions() []ExtensionInfo {
return []ExtensionInfo{
{
Name:        "security",
Version:     "1.0.0",
Description: "SAST scanning, secret detection, vulnerability checks",
Author:      "Agenticide Team",
Repository:  "https://github.com/ivikasavnish/agenticide-go",
Tags:        []string{"security", "sast", "scanner", "secrets"},
Category:    "security",
InstallURL:  "",
},
{
Name:        "code-analyzer",
Version:     "1.0.0",
Description: "Code complexity metrics, dead code detection, AST analysis",
Author:      "Agenticide Team",
Repository:  "https://github.com/ivikasavnish/agenticide-go",
Tags:        []string{"code", "analysis", "metrics", "ast"},
Category:    "development",
InstallURL:  "",
},
{
Name:        "project-runner",
Version:     "1.0.0",
Description: "Auto-detect and run any project type (Node, Go, Python, Rust)",
Author:      "Agenticide Team",
Repository:  "https://github.com/ivikasavnish/agenticide-go",
Tags:        []string{"runner", "execution", "automation"},
Category:    "development",
InstallURL:  "",
},
{
Name:        "web-search",
Version:     "1.0.0",
Description: "Multi-engine web search with content extraction",
Author:      "Agenticide Team",
Repository:  "https://github.com/ivikasavnish/agenticide-go",
Tags:        []string{"search", "web", "research"},
Category:    "research",
InstallURL:  "",
},
{
Name:        "db-analytics",
Version:     "1.0.0",
Description: "Database query builder, terminal charts, data visualization",
Author:      "Agenticide Team",
Repository:  "https://github.com/ivikasavnish/agenticide-go",
Tags:        []string{"database", "analytics", "charts"},
Category:    "data",
InstallURL:  "",
},
{
Name:        "deployment",
Version:     "1.0.0",
Description: "Deploy to AWS, GCP, Azure, Vercel with one command",
Author:      "Agenticide Team",
Repository:  "https://github.com/ivikasavnish/agenticide-go",
Tags:        []string{"deploy", "cloud", "ci-cd"},
Category:    "deployment",
InstallURL:  "",
},
{
Name:        "cost-controller",
Version:     "1.0.0",
Description: "Track and optimize cloud costs, API usage forecasting",
Author:      "Agenticide Team",
Repository:  "https://github.com/ivikasavnish/agenticide-go",
Tags:        []string{"cost", "optimization", "cloud"},
Category:    "ops",
InstallURL:  "",
},
{
Name:        "monitoring",
Version:     "1.0.0",
Description: "Proactive monitoring, health checks, alerting",
Author:      "Agenticide Team",
Repository:  "https://github.com/ivikasavnish/agenticide-go",
Tags:        []string{"monitoring", "alerts", "observability"},
Category:    "ops",
InstallURL:  "",
},
{
Name:        "ui-design",
Version:     "0.9.0",
Description: "Lovable design system integration, component previews",
Author:      "Agenticide Team",
Repository:  "https://github.com/ivikasavnish/agenticide-go",
Tags:        []string{"design", "ui", "frontend"},
Category:    "design",
InstallURL:  "",
},
{
Name:        "ai-recipes",
Version:     "1.0.0",
Description: "Pre-built AI workflows: code review, bug hunting, refactoring",
Author:      "Agenticide Team",
Repository:  "https://github.com/ivikasavnish/agenticide-go",
Tags:        []string{"ai", "automation", "workflows"},
Category:    "ai",
InstallURL:  "",
},
{
Name:        "git-ops",
Version:     "1.0.0",
Description: "Git operations, PR reviews, commit analysis",
Author:      "Community",
Repository:  "https://github.com/agenticide-extensions/git-ops",
Tags:        []string{"git", "vcs", "collaboration"},
Category:    "development",
InstallURL:  "https://github.com/agenticide-extensions/git-ops/releases/latest/download/git-ops",
},
{
Name:        "test-generator",
Version:     "1.0.0",
Description: "Automatically generate unit tests for any codebase",
Author:      "Community",
Repository:  "https://github.com/agenticide-extensions/test-generator",
Tags:        []string{"testing", "automation", "tdd"},
Category:    "development",
InstallURL:  "https://github.com/agenticide-extensions/test-generator/releases/latest/download/test-generator",
},
}
}
