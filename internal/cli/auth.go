package cli

import (
"fmt"

"github.com/spf13/cobra"

"github.com/ivikasavnish/agenticide-go/internal/auth"
"github.com/ivikasavnish/agenticide-go/internal/ui"
"github.com/ivikasavnish/agenticide-go/pkg/license"
)

type AuthCommands struct{}

func NewAuthCommands() *AuthCommands {
return &AuthCommands{}
}

func (ac *AuthCommands) SignupCommand() *cobra.Command {
return &cobra.Command{
Use:   "signup",
Short: "Register for Agenticide",
Long: `Register for Agenticide by providing your contact information.
Your registration will be reviewed and you'll receive a license key via email.`,
Run: func(cmd *cobra.Command, args []string) {
form, err := auth.PromptSignup()
if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Signup failed: %v", err)))
return
}

// Submit signup
requestID, err := auth.SubmitSignup(form)
if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Failed to submit: %v", err)))
return
}

fmt.Println()
fmt.Println(ui.Success("Registration submitted successfully!"))
fmt.Println()
fmt.Printf("Request ID: %s\n", ui.Highlight(requestID))
fmt.Println()
fmt.Println(ui.Info("Next steps:"))
fmt.Println("  1. Check your email for approval notification (24-48 hours)")
fmt.Println("  2. Once approved, you'll receive a license key")
fmt.Println("  3. Activate with: agenticide activate <license-key>")
fmt.Println()
fmt.Println(ui.Muted("Questions? Email: support@agenticide.dev"))
},
}
}

func (ac *AuthCommands) ActivateCommand() *cobra.Command {
return &cobra.Command{
Use:   "activate <license-key>",
Short: "Activate your license",
Args:  cobra.ExactArgs(1),
Run: func(cmd *cobra.Command, args []string) {
licenseKey := args[0]

fmt.Println(ui.Title("🔑 Activating License"))
fmt.Println()

if err := license.Activate(licenseKey); err != nil {
fmt.Println(ui.Error(fmt.Sprintf("Activation failed: %v", err)))
return
}

fmt.Println(ui.Success("License activated successfully!"))
fmt.Println()
fmt.Println(ui.Info("You can now use all Agenticide features."))
fmt.Println()
fmt.Println("Try:")
fmt.Println("  agenticide chat")
fmt.Println("  agenticide ext list")
fmt.Println("  agenticide task list")
},
}
}

func (ac *AuthCommands) StatusCommand() *cobra.Command {
return &cobra.Command{
Use:   "status",
Short: "Check license status",
Run: func(cmd *cobra.Command, args []string) {
fmt.Println(ui.Title("📋 License Status"))
fmt.Println()

lic, err := license.Check()
if err != nil {
fmt.Println(ui.Error(fmt.Sprintf("No active license: %v", err)))
fmt.Println()
fmt.Println(ui.Info("To get started:"))
fmt.Println("  1. Register: agenticide signup")
fmt.Println("  2. Activate: agenticide activate <license-key>")
return
}

fmt.Printf("License Key: %s\n", ui.Highlight(lic.Key))
fmt.Printf("Email: %s\n", lic.Email)
fmt.Printf("Tier: %s\n", ui.Badge(lic.Tier))
fmt.Printf("Status: %s\n", ui.Success(lic.Status))
},
}
}
