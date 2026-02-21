package components

import (
"fmt"
"math"
"strings"

"github.com/ivikasavnish/agenticide-go/internal/ui"
)

type BarChart struct {
data   map[string]int
width  int
maxVal int
}

func NewBarChart(width int) *BarChart {
return &BarChart{
data:   make(map[string]int),
width:  width,
maxVal: 0,
}
}

func (b *BarChart) AddBar(label string, value int) *BarChart {
b.data[label] = value
if value > b.maxVal {
b.maxVal = value
}
return b
}

func (b *BarChart) Render() string {
var sb strings.Builder

for label, value := range b.data {
barWidth := 0
if b.maxVal > 0 {
barWidth = int(float64(value) / float64(b.maxVal) * float64(b.width))
}

bar := ui.ProgressBarFilled.Render(strings.Repeat("█", barWidth))
sb.WriteString(fmt.Sprintf("%-15s %s %d\n", label, bar, value))
}

return sb.String()
}

type SparkLine struct {
data   []int
width  int
height int
}

func NewSparkLine(width, height int) *SparkLine {
return &SparkLine{
data:   make([]int, 0),
width:  width,
height: height,
}
}

func (s *SparkLine) AddValue(value int) *SparkLine {
s.data = append(s.data, value)
if len(s.data) > s.width {
s.data = s.data[1:]
}
return s
}

func (s *SparkLine) Render() string {
if len(s.data) == 0 {
return ""
}

maxVal := s.data[0]
for _, v := range s.data {
if v > maxVal {
maxVal = v
}
}

blocks := []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"}

var sb strings.Builder
for _, v := range s.data {
index := int(math.Floor(float64(v) / float64(maxVal) * float64(len(blocks)-1)))
if index < 0 {
index = 0
}
if index >= len(blocks) {
index = len(blocks) - 1
}
sb.WriteString(ui.SuccessStyle.Render(blocks[index]))
}

return sb.String()
}
