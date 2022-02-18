package raingow

import (
	"math"
	"github.com/gookit/color"
)

// Wholesale borrowed from https://github.com/busyloop/lolcat
func colour_code(freq float64, spread float64) color.RGBColor {
	red   := uint8(math.Sin(freq * spread + 0) * 127 + 128)
	green := uint8(math.Sin(freq * spread + 2* math.Pi/3) * 127 + 128)
	blue  := uint8(math.Sin(freq * spread + 4* math.Pi/3) * 127 + 128)
	rgb_obj := color.RGB(red, green, blue)
	return rgb_obj
}


func Raingow_line(source_str string, span float64) string {
	// TODO give the ability to pass state of colour_index between
	// iterations (or a different wrapper function) so that we can
	// span colours across multiple lines

	output_str := ""

	for j := 0; j < len(source_str); j++ {
		the_colour := colour_code(0.1, float64(j) * span)
		output_str += the_colour.Sprintf(string(source_str[j]))
	}

	return output_str
}
