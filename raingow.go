// Wholesale borrowed from https://github.com/busyloop/lolcat.
package raingow

import (
	"math"
	"strings"

	"github.com/gookit/color"
)

func colourCode(freq, spread float64) color.RGBColor {
	red   := uint8(math.Sin(freq * spread + 0) * 127 + 128)
	green := uint8(math.Sin(freq * spread + 2 * math.Pi/3) * 127 + 128)
	blue  := uint8(math.Sin(freq * spread + 4 * math.Pi/3) * 127 + 128)
	return color.RGB(red, green, blue)
}

// TODO: thread colour-index state across lines so a gradient can span
// multiple calls.
func RaingowLine(source string, span float64) string {
	var out strings.Builder
	out.Grow(len(source) * 20)
	j := 0
	for _, r := range source {
		c := colourCode(0.1, float64(j)*span)
		out.WriteString(c.Sprint(string(r)))
		j++
	}
	return out.String()
}
