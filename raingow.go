//package raingow
package main

import (
	"fmt"
	"github.com/fatih/color"
)

func Raingow_line(source_str string, offset int, reverse_palette bool, span int) string {

	//var colours = [7]*color.Color{
	var colours = [6]*color.Color{
		color.New(color.FgYellow),
		color.New(color.FgRed),
		color.New(color.FgGreen),
		color.New(color.FgBlue),
		color.New(color.FgMagenta),
		color.New(color.FgCyan),
		//color.New(color.FgWhite),
	}

	var output_str = ""

	var span_count = 0
	var colour_index = 0

	for i := 0; i < len(source_str); i++ {
		current_char := string(source_str[i])

		if span_count < span {
			span_count += 1
		} else {
			colour_index = (i + offset) % len(colours)
			if reverse_palette {
				colour_index = ( len(colours) - 1) - ((i + offset) % len(colours))
			}
			span_count = 0
		}
		output_str = output_str + colours[colour_index].SprintFunc()(current_char)
	}

	return output_str
}


func main() {
	// TODO why does span lead to the same colour after a few repetitions when the spam is len(colours)?
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 2))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 3))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 4))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 5))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 6))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 7))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 8))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line", 0, false, 9))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line reversed", 0, true, 0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line offset 1", 1, false, 0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line offset 2", 2, false, 0))
	fmt.Println(Raingow_line("My fairly long long long long long long long long long testing line offset 3", 3, false, 0))
}
