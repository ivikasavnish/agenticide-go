package ui

import (
"github.com/charmbracelet/lipgloss"
)

var (
Primary   = lipgloss.Color("#7D56F4")
Secondary = lipgloss.Color("#FF6B9D")
SuccessColor   = lipgloss.Color("#04B575")
WarningColor   = lipgloss.Color("#FFA500")
ErrorColor     = lipgloss.Color("#FF4444")
InfoColor      = lipgloss.Color("#00BFFF")
MutedColor     = lipgloss.Color("#777777")
TextColor      = lipgloss.Color("#FAFAFA")

TitleStyle = lipgloss.NewStyle().
Bold(true).
Foreground(TextColor).
Background(Primary).
Padding(0, 1).
MarginBottom(1)

SuccessStyle = lipgloss.NewStyle().
Foreground(SuccessColor).
Bold(true)

ErrorStyle = lipgloss.NewStyle().
Foreground(ErrorColor).
Bold(true)

WarningStyle = lipgloss.NewStyle().
Foreground(WarningColor).
Bold(true)

InfoStyle = lipgloss.NewStyle().
Foreground(InfoColor)

MutedStyle = lipgloss.NewStyle().
Foreground(MutedColor)

BorderStyle = lipgloss.NewStyle().
Border(lipgloss.RoundedBorder()).
BorderForeground(Primary).
Padding(1, 2)

HelpStyle = lipgloss.NewStyle().
Foreground(MutedColor).
MarginTop(1)

HighlightStyle = lipgloss.NewStyle().
Background(Primary).
Foreground(TextColor).
Padding(0, 1)

TableHeaderStyle = lipgloss.NewStyle().
Bold(true).
Foreground(TextColor).
Background(Primary).
Padding(0, 1)

TableCellStyle = lipgloss.NewStyle().
Padding(0, 1)

TableRowOddStyle = lipgloss.NewStyle().
Foreground(TextColor)

TableRowEvenStyle = lipgloss.NewStyle().
Foreground(MutedColor)

ProgressBarFilled = lipgloss.NewStyle().
Background(SuccessColor)

ProgressBarEmpty = lipgloss.NewStyle().
Background(MutedColor)

BadgeStyle = lipgloss.NewStyle().
Background(Primary).
Foreground(TextColor).
Padding(0, 1).
MarginRight(1)
)

func Title(text string) string {
return TitleStyle.Render(text)
}

func Success(text string) string {
return SuccessStyle.Render("✓ " + text)
}

func Error(text string) string {
return ErrorStyle.Render("✗ " + text)
}

func Warning(text string) string {
return WarningStyle.Render("⚠ " + text)
}

func Info(text string) string {
return InfoStyle.Render("ℹ " + text)
}

func Muted(text string) string {
return MutedStyle.Render(text)
}

func Border(text string) string {
return BorderStyle.Render(text)
}

func Highlight(text string) string {
return HighlightStyle.Render(text)
}

func Badge(text string) string {
return BadgeStyle.Render(text)
}
