package components

import (
"fmt"
"strings"

"github.com/ivikasavnish/agenticide-go/internal/ui"
)

type List struct {
items    []string
selected int
bullet   string
}

func NewList() *List {
return &List{
items:    make([]string, 0),
selected: -1,
bullet:   "•",
}
}

func (l *List) AddItem(item string) *List {
l.items = append(l.items, item)
return l
}

func (l *List) SetSelected(index int) *List {
if index >= 0 && index < len(l.items) {
l.selected = index
}
return l
}

func (l *List) SetBullet(bullet string) *List {
l.bullet = bullet
return l
}

func (l *List) Render() string {
var sb strings.Builder

for i, item := range l.items {
if i == l.selected {
sb.WriteString(ui.Highlight(fmt.Sprintf(" %s %s ", l.bullet, item)))
} else {
sb.WriteString(fmt.Sprintf("  %s %s", l.bullet, item))
}
if i < len(l.items)-1 {
sb.WriteString("\n")
}
}

return sb.String()
}

func (l *List) RenderWithStatus(statuses []string) string {
var sb strings.Builder

for i, item := range l.items {
status := ""
if i < len(statuses) {
switch statuses[i] {
case "done":
status = ui.RenderSuccess("")
case "pending":
status = ui.RenderMuted("○")
case "in_progress":
status = ui.RenderInfo("◐")
case "error":
status = ui.RenderError("")
default:
status = ui.RenderMuted("○")
}
}

line := fmt.Sprintf("  %s %s", status, item)
if i == l.selected {
sb.WriteString(ui.Highlight(line))
} else {
sb.WriteString(line)
}

if i < len(l.items)-1 {
sb.WriteString("\n")
}
}

return sb.String()
}
