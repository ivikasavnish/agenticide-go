package components

import (
"strings"

"github.com/charmbracelet/lipgloss"
"github.com/ivikasavnish/agenticide-go/internal/ui"
)

type Panel struct {
title   string
content string
width   int
height  int
style   lipgloss.Style
}

func NewPanel(title string) *Panel {
return &Panel{
title:  title,
width:  0,
height: 0,
style:  ui.BorderStyle,
}
}

func (p *Panel) SetContent(content string) *Panel {
p.content = content
return p
}

func (p *Panel) SetWidth(width int) *Panel {
p.width = width
return p
}

func (p *Panel) SetHeight(height int) *Panel {
p.height = height
return p
}

func (p *Panel) Render() string {
style := p.style

if p.width > 0 {
style = style.Width(p.width - 4)
}
if p.height > 0 {
style = style.Height(p.height - 4)
}

var sb strings.Builder
if p.title != "" {
sb.WriteString(ui.Title(p.title))
sb.WriteString("\n")
}
sb.WriteString(style.Render(p.content))

return sb.String()
}
