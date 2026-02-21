package license

import (
"crypto/rand"
"encoding/base64"
"fmt"
"os"
"path/filepath"
"strings"
"time"
)

type License struct {
Key       string
Email     string
IssuedAt  time.Time
ExpiresAt time.Time
Tier      string // "free", "pro", "enterprise"
Status    string // "active", "suspended", "expired"
}

func GenerateLicenseKey(email string) string {
// Generate cryptographically secure random key
bytes := make([]byte, 24)
rand.Read(bytes)

encoded := base64.URLEncoding.EncodeToString(bytes)
encoded = strings.ReplaceAll(encoded, "-", "")
encoded = strings.ReplaceAll(encoded, "_", "")

// Format: AGNT-XXXX-XXXX-XXXX-XXXX
key := fmt.Sprintf("AGNT-%s-%s-%s-%s",
encoded[0:4],
encoded[4:8],
encoded[8:12],
encoded[12:16],
)

return strings.ToUpper(key)
}

func Activate(licenseKey string) error {
// Validate format
if !strings.HasPrefix(licenseKey, "AGNT-") {
return fmt.Errorf("invalid license key format")
}

parts := strings.Split(licenseKey, "-")
if len(parts) != 5 {
return fmt.Errorf("invalid license key format")
}

// In production, this would verify with API
// For now, save locally

home, _ := os.UserHomeDir()
licenseDir := filepath.Join(home, ".agenticide")
licenseFile := filepath.Join(licenseDir, "license.txt")

os.MkdirAll(licenseDir, 0755)

license := &License{
Key:       licenseKey,
Email:     "user@example.com", // Would come from API
IssuedAt:  time.Now(),
ExpiresAt: time.Now().AddDate(1, 0, 0), // 1 year
Tier:      "free",
Status:    "active",
}

content := fmt.Sprintf(`Agenticide License
Key: %s
Email: %s
Tier: %s
Status: %s
Issued: %s
Expires: %s
`, license.Key, license.Email, license.Tier, license.Status,
license.IssuedAt.Format("2006-01-02"),
license.ExpiresAt.Format("2006-01-02"))

if err := os.WriteFile(licenseFile, []byte(content), 0600); err != nil {
return fmt.Errorf("failed to save license: %w", err)
}

return nil
}

func Check() (*License, error) {
home, _ := os.UserHomeDir()
licenseFile := filepath.Join(home, ".agenticide", "license.txt")

if _, err := os.Stat(licenseFile); os.IsNotExist(err) {
return nil, fmt.Errorf("no license found - please run: agenticide signup")
}

// Read and parse license
content, err := os.ReadFile(licenseFile)
if err != nil {
return nil, fmt.Errorf("failed to read license: %w", err)
}

// Simple parsing (in production, would decrypt/verify)
lines := strings.Split(string(content), "\n")
license := &License{
Status: "active",
}

for _, line := range lines {
if strings.HasPrefix(line, "Key: ") {
license.Key = strings.TrimPrefix(line, "Key: ")
} else if strings.HasPrefix(line, "Email: ") {
license.Email = strings.TrimPrefix(line, "Email: ")
} else if strings.HasPrefix(line, "Tier: ") {
license.Tier = strings.TrimPrefix(line, "Tier: ")
}
}

return license, nil
}

func IsValid() bool {
license, err := Check()
if err != nil {
return false
}

return license.Status == "active"
}
