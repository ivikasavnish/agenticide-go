package cli

import (
"context"
"fmt"
"time"

"github.com/spf13/cobra"

"github.com/ivikasavnish/agenticide-go/internal/core/extension"
"github.com/ivikasavnish/agenticide-go/internal/ui"
"github.com/ivikasavnish/agenticide-go/internal/ui/components"
)

type ExtensionCommands struct {
registry extension.Registry
}

func NewExtensionCommands(registry extension.Registry) *ExtensionCommands {
return &ExtensionCommands{
registry: registry,
}
}

func (ec *ExtensionCommands) ListCommand() *cobra.Command {
return &cobra.Command{
Use:   "list",
Short: "List all extensions",
Run: func(cmd *cobra.Command, args []string) {
fmt.Println(ui.Title("📦 Installed Extensions"))
fmt.Println()

extensions := ec.registry.List()
if len(extensions) == 0 {
fmt.Println(ui.Muted("  No extensions installed"))
return
}

table := components.NewTable("Extension", "Version", "Status", "Description")
for _, ext := range extensions {
status := "disabled"
if ec.registry.IsEnabled(ext.Name()) {
status = ui.Success("enabled")
} else {
status = ui.Muted("disabled")
}
table.AddRow(ext.Name(), ext.Version(), status, ext.Description())
}

fmt.Println(table.Render())
fmt.Println()
fmt.Printf(ui.Muted("  Total: %d extensions\n"), len(extensions))
},
}
}

func (ec *ExtensionCommands) EnableCommand() *cobra.Command {
return &cobra.Command{
Use:   "enable <name>",
Short: "Enable an extension",
Args:  cobra.ExactArgs(1),
Run: func(cmd *cobra.Command, args []string) {
name := args[0]

fmt.Printf("Enabling extension %s...\n", ui.Highlight(name))

if err := ec.registry.Enable(name); err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Failed: %v", err)))
return
}

fmt.Println(ui.Success(fmt.Sprintf("Extension '%s' enabled", name)))
},
}
}

func (ec *ExtensionCommands) DisableCommand() *cobra.Command {
return &cobra.Command{
Use:   "disable <name>",
Short: "Disable an extension",
Args:  cobra.ExactArgs(1),
Run: func(cmd *cobra.Command, args []string) {
name := args[0]

fmt.Printf("Disabling extension %s...\n", ui.Highlight(name))

if err := ec.registry.Disable(name); err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Failed: %v", err)))
return
}

fmt.Println(ui.Success(fmt.Sprintf("Extension '%s' disabled", name)))
},
}
}

func (ec *ExtensionCommands) InfoCommand() *cobra.Command {
return &cobra.Command{
Use:   "info <name>",
Short: "Show extension details",
Args:  cobra.ExactArgs(1),
Run: func(cmd *cobra.Command, args []string) {
name := args[0]

ext, err := ec.registry.Get(name)
if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Extension not found: %s", name)))
return
}

fmt.Println(ui.Title(fmt.Sprintf("📦 %s", ext.Name())))
fmt.Println()

panel := components.NewPanel("").
SetContent(fmt.Sprintf(
"Name:        %s\n"+
"Version:     %s\n"+
"Author:      %s\n"+
"Status:      %s\n"+
"Description: %s\n\n"+
"Dependencies: %v\n\n"+
"Commands: %d available",
ext.Name(),
ext.Version(),
ext.Author(),
func() string {
if ec.registry.IsEnabled(ext.Name()) {
return ui.Success("enabled")
}
return ui.Muted("disabled")
}(),
ext.Description(),
ext.Dependencies(),
len(ext.Commands()),
)).
SetWidth(70)

fmt.Println(panel.Render())
},
}
}
