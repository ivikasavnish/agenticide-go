package auth

import (
"bufio"
"fmt"
"os"
"regexp"
"strings"

"github.com/ivikasavnish/agenticide-go/internal/ui"
)

type SignupForm struct {
Email    string
Mobile   string
LinkedIn string
UseCase  string
}

func (sf *SignupForm) Validate() error {
// Email validation
emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
if !emailRegex.MatchString(sf.Email) {
return fmt.Errorf("invalid email address")
}

// Mobile validation (basic: +XX XXXXXXXXXX)
mobileRegex := regexp.MustCompile(`^\+\d{1,3}\s?\d{7,15}$`)
if !mobileRegex.MatchString(sf.Mobile) {
return fmt.Errorf("invalid mobile number (format: +XX XXXXXXXXXX)")
}

// LinkedIn validation
if !strings.Contains(sf.LinkedIn, "linkedin.com/in/") {
return fmt.Errorf("invalid LinkedIn URL (must contain linkedin.com/in/)")
}

// Use case validation
if len(sf.UseCase) < 20 {
return fmt.Errorf("use case description too short (minimum 20 characters)")
}

return nil
}

func PromptSignup() (*SignupForm, error) {
scanner := bufio.NewScanner(os.Stdin)

fmt.Println(ui.Title("🔐 Agenticide Registration"))
fmt.Println()
fmt.Println(ui.Muted("Registration is required to use Agenticide."))
fmt.Println(ui.Muted("Your information will be reviewed for approval (24-48 hours)."))
fmt.Println()

form := &SignupForm{}

// Email
fmt.Print("Email address: ")
scanner.Scan()
form.Email = strings.TrimSpace(scanner.Text())

// Mobile
fmt.Print("Mobile number (with country code, e.g. +1 1234567890): ")
scanner.Scan()
form.Mobile = strings.TrimSpace(scanner.Text())

// LinkedIn
fmt.Print("LinkedIn profile URL: ")
scanner.Scan()
form.LinkedIn = strings.TrimSpace(scanner.Text())

// Use case
fmt.Println("Use case (minimum 20 characters):")
fmt.Print("> ")
scanner.Scan()
form.UseCase = strings.TrimSpace(scanner.Text())

fmt.Println()

// Validate
if err := form.Validate(); err != nil {
return nil, fmt.Errorf("validation failed: %w", err)
}

return form, nil
}

func SubmitSignup(form *SignupForm) (string, error) {
// In production, this would submit to API
// For now, save locally and generate pending ID

requestID := fmt.Sprintf("REQ-%d", os.Getpid())

// Save to local storage (would be API call in production)
home, _ := os.UserHomeDir()
signupFile := fmt.Sprintf("%s/.agenticide/signup-%s.txt", home, requestID)

os.MkdirAll(fmt.Sprintf("%s/.agenticide", home), 0755)

content := fmt.Sprintf(`Agenticide Registration Request
Request ID: %s
Status: Pending Approval

Email: %s
Mobile: %s
LinkedIn: %s
Use Case: %s

Submitted: %s

Next Steps:
1. Your request has been submitted for review
2. You will receive an email notification within 24-48 hours
3. Upon approval, you'll receive a license key
4. Activate with: agenticide activate <license-key>

Support: support@agenticide.dev
`, requestID, form.Email, form.Mobile, form.LinkedIn, form.UseCase, 
   fmt.Sprintf("%s", os.Getenv("TZ")))

if err := os.WriteFile(signupFile, []byte(content), 0600); err != nil {
return "", fmt.Errorf("failed to save signup: %w", err)
}

return requestID, nil
}
