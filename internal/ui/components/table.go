package components

import (
"strings"

"github.com/charmbracelet/lipgloss"
"github.com/ivikasavnish/agenticide-go/internal/ui"
)

type Table struct {
headers []string
rows    [][]string
widths  []int
}

func NewTable(headers ...string) *Table {
widths := make([]int, len(headers))
for i, h := range headers {
widths[i] = len(h)
}
return &Table{
headers: headers,
rows:    make([][]string, 0),
widths:  widths,
}
}

func (t *Table) AddRow(cells ...string) *Table {
if len(cells) != len(t.headers) {
return t
}

for i, cell := range cells {
if len(cell) > t.widths[i] {
t.widths[i] = len(cell)
}
}

t.rows = append(t.rows, cells)
return t
}

func (t *Table) Render() string {
var sb strings.Builder

headerCells := make([]string, len(t.headers))
for i, h := range t.headers {
style := ui.TableHeaderStyle.Width(t.widths[i])
headerCells[i] = style.Render(padRight(h, t.widths[i]))
}
sb.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, headerCells...))
sb.WriteString("\n")

for idx, row := range t.rows {
cells := make([]string, len(row))
style := ui.TableRowOddStyle
if idx%2 == 1 {
style = ui.TableRowEvenStyle
}

for i, cell := range row {
cellStyle := style.Copy().Width(t.widths[i])
cells[i] = cellStyle.Render(padRight(cell, t.widths[i]))
}
sb.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, cells...))
sb.WriteString("\n")
}

return sb.String()
}

func padRight(s string, width int) string {
if len(s) >= width {
return s
}
return s + strings.Repeat(" ", width-len(s))
}
