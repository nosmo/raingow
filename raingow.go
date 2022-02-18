//package raingow
package main

import (
	"fmt"
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


func Raingow_line(source_str string, offset int, reverse_palette bool, span float64) string {
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


func main() {


	// TODO why does span lead to the same colour after a few repetitions when the spam is len(colours)?
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 1.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 2.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 3.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 4.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 5.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 6.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 7.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 8.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 9.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line reversed", 0, true, 3.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line offset 1", 1, false, 3.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line offset 2", 2, false, 3.0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line offset 3", 3, false, 3.0))
}
