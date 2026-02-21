package components

import (
"fmt"
"strings"

"github.com/ivikasavnish/agenticide-go/internal/ui"
)

type ProgressBar struct {
current int
total   int
width   int
label   string
}

func NewProgressBar(total int, width int) *ProgressBar {
return &ProgressBar{
current: 0,
total:   total,
width:   width,
label:   "",
}
}

func (p *ProgressBar) SetCurrent(current int) *ProgressBar {
p.current = current
return p
}

func (p *ProgressBar) SetLabel(label string) *ProgressBar {
p.label = label
return p
}

func (p *ProgressBar) Increment() *ProgressBar {
p.current++
if p.current > p.total {
p.current = p.total
}
return p
}

func (p *ProgressBar) Render() string {
percent := float64(p.current) / float64(p.total)
filled := int(float64(p.width) * percent)
empty := p.width - filled

bar := ui.ProgressBarFilled.Render(strings.Repeat("█", filled)) +
ui.ProgressBarEmpty.Render(strings.Repeat("░", empty))

percentText := fmt.Sprintf(" %d%% (%d/%d)", int(percent*100), p.current, p.total)

if p.label != "" {
return fmt.Sprintf("%s\n%s%s", p.label, bar, percentText)
}

return bar + percentText
}
