package cli

import (
"fmt"

"github.com/spf13/cobra"

"github.com/ivikasavnish/agenticide-go/internal/ui"
"github.com/ivikasavnish/agenticide-go/internal/ui/components"
"github.com/ivikasavnish/agenticide-go/pkg/marketplace"
)

type MarketplaceCommands struct {
registry *marketplace.Registry
}

func NewMarketplaceCommands() *MarketplaceCommands {
return &MarketplaceCommands{
registry: marketplace.NewRegistry(),
}
}

func (mc *MarketplaceCommands) SearchCommand() *cobra.Command {
return &cobra.Command{
Use:   "search [query]",
Short: "Search for extensions in marketplace",
Args:  cobra.MaximumNArgs(1),
Run: func(cmd *cobra.Command, args []string) {
query := ""
if len(args) > 0 {
query = args[0]
}

fmt.Println(ui.Title("🔍 Extension Marketplace"))
fmt.Println()

extensions, err := mc.registry.Search(query)
if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Search failed: %v", err)))
return
}

if len(extensions) == 0 {
fmt.Println(ui.Muted("No extensions found"))
return
}

table := components.NewTable("Name", "Version", "Category", "Description")
for _, ext := range extensions {
table.AddRow(
ext.Name,
ext.Version,
ui.Badge(ext.Category),
ext.Description,
)
}

fmt.Println(table.Render())
fmt.Println()
fmt.Printf(ui.Muted("  Found %d extensions\n"), len(extensions))
fmt.Println()
fmt.Println(ui.Info("Install with: agenticide install <name>"))
},
}
}

func (mc *MarketplaceCommands) InstallCommand() *cobra.Command {
return &cobra.Command{
Use:   "install <extension>",
Short: "Install extension from marketplace (like npx)",
Args:  cobra.ExactArgs(1),
Run: func(cmd *cobra.Command, args []string) {
name := args[0]

fmt.Println(ui.Title(fmt.Sprintf("📦 Installing %s", name)))
fmt.Println()

// Get extension info
ext, err := mc.registry.Get(name)
if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Extension not found: %s", name)))
fmt.Println()
fmt.Println(ui.Info("Search for extensions: agenticide search"))
return
}

// Show progress
progress := components.NewProgressBar(100, 50)

progress.SetLabel("Downloading...").SetCurrent(30)
fmt.Printf("\r%s", progress.Render())

// Install
if err := mc.registry.Install(name); err != nil {
fmt.Println()
fmt.Println(ui.Error(fmt.Sprintf("Installation failed: %v", err)))
return
}

progress.SetLabel("Installing...").SetCurrent(70)
fmt.Printf("\r%s", progress.Render())

progress.SetLabel("Configuring...").SetCurrent(100)
fmt.Printf("\r%s", progress.Render())

fmt.Println()
fmt.Println()
fmt.Println(ui.Success(fmt.Sprintf("Extension '%s' installed successfully!", name)))
fmt.Println()
fmt.Printf("Name: %s\n", ui.Highlight(ext.Name))
fmt.Printf("Version: %s\n", ext.Version)
fmt.Printf("Author: %s\n", ext.Author)
fmt.Println()
fmt.Println(ui.Info("Use with: agenticide <command>"))
},
}
}

func (mc *MarketplaceCommands) ListCommand() *cobra.Command {
return &cobra.Command{
Use:   "list",
Short: "List installed extensions",
Run: func(cmd *cobra.Command, args []string) {
fmt.Println(ui.Title("📦 Installed Extensions"))
fmt.Println()

installed, err := mc.registry.ListInstalled()
if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Failed to list: %v", err)))
return
}

if len(installed) == 0 {
fmt.Println(ui.Muted("  No extensions installed"))
fmt.Println()
fmt.Println(ui.Info("Search marketplace: agenticide search"))
return
}

table := components.NewTable("Name", "Version", "Category")
for _, ext := range installed {
table.AddRow(ext.Name, ext.Version, ui.Badge(ext.Category))
}

fmt.Println(table.Render())
fmt.Println()
fmt.Printf(ui.Muted("  Total: %d extensions\n"), len(installed))
},
}
}

func (mc *MarketplaceCommands) UninstallCommand() *cobra.Command {
return &cobra.Command{
Use:   "uninstall <extension>",
Short: "Uninstall an extension",
Args:  cobra.ExactArgs(1),
Run: func(cmd *cobra.Command, args []string) {
name := args[0]

fmt.Printf("Uninstalling %s...\n", ui.Highlight(name))

if err := mc.registry.Uninstall(name); err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Failed: %v", err)))
return
}

fmt.Println(ui.Success(fmt.Sprintf("Extension '%s' uninstalled", name)))
},
}
}

func (mc *MarketplaceCommands) InfoCommand() *cobra.Command {
return &cobra.Command{
Use:   "info <extension>",
Short: "Show extension details",
Args:  cobra.ExactArgs(1),
Run: func(cmd *cobra.Command, args []string) {
name := args[0]

ext, err := mc.registry.Get(name)
if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Extension not found: %s", name)))
return
}

fmt.Println(ui.Title(fmt.Sprintf("📦 %s", ext.Name)))
fmt.Println()

panel := components.NewPanel("").
SetContent(fmt.Sprintf(
"Name:        %s\n"+
"Version:     %s\n"+
"Author:      %s\n"+
"Category:    %s\n"+
"Repository:  %s\n\n"+
"Description:\n%s\n\n"+
"Tags: %v",
ext.Name,
ext.Version,
ext.Author,
ui.Badge(ext.Category),
ext.Repository,
ext.Description,
ext.Tags,
)).
SetWidth(70)

fmt.Println(panel.Render())
fmt.Println()
fmt.Println(ui.Info(fmt.Sprintf("Install with: agenticide install %s", ext.Name)))
},
}
}
