package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/ivikasavnish/agenticide-go/internal/core/config"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	promptStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)
)

// FirstRunSetup runs the initial setup wizard
func FirstRunSetup(cfg *config.Config) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println(titleStyle.Render(" Welcome to Agenticide! "))
	fmt.Println()
	fmt.Println("Let's get you set up. This will only take a moment.")
	fmt.Println()

	// Get user info
	fmt.Print(promptStyle.Render("Your email: "))
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print(promptStyle.Render("Your phone (optional): "))
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	// Get AI provider
	fmt.Println()
	fmt.Println("Select your AI provider:")
	fmt.Println("  1) OpenAI (GPT-4)")
	fmt.Println("  2) Anthropic (Claude)")
	fmt.Println("  3) GitHub Copilot")
	fmt.Println("  4) Local (Ollama)")
	fmt.Print(promptStyle.Render("Choice [1-4]: "))
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var provider, apiKey string
	switch choice {
	case "1":
		provider = "openai"
		fmt.Print(promptStyle.Render("OpenAI API Key: "))
		apiKey, _ = reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
	case "2":
		provider = "anthropic"
		fmt.Print(promptStyle.Render("Anthropic API Key: "))
		apiKey, _ = reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
	case "3":
		provider = "copilot"
		fmt.Println("GitHub Copilot will use your GitHub authentication")
	case "4":
		provider = "ollama"
		fmt.Print(promptStyle.Render("Ollama URL [http://localhost:11434]: "))
		url, _ := reader.ReadString('\n')
		url = strings.TrimSpace(url)
		if url == "" {
			url = "http://localhost:11434"
		}
		apiKey = url
	default:
		provider = "openai"
	}

	// Save configuration
	cfg.Set("user.email", email)
	cfg.Set("user.phone", phone)
	cfg.Set("ai.provider", provider)
	cfg.Set("ai.api_key", apiKey)
	cfg.Set("setup_complete", true)

	if err := cfg.Save(); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Println()
	fmt.Println(successStyle.Render("✓ Setup complete!"))
	fmt.Println()
	fmt.Println("You can now use:")
	fmt.Println("  agenticide            - Start AI chat")
	fmt.Println("  agenticide search     - Browse extensions")
	fmt.Println("  agenticide task add   - Manage tasks")
	fmt.Println()

	return nil
}

// CheckSetup checks if setup is complete
func CheckSetup(cfg *config.Config) bool {
	return cfg.GetBool("setup_complete")
}

// RequireSetup runs setup if not complete
func RequireSetup(cfg *config.Config) error {
	if !CheckSetup(cfg) {
		return FirstRunSetup(cfg)
	}
	return nil
}

// GetConfigPath returns the config file path
func GetConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".agenticide", "config.yaml")
}
