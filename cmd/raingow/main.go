// raingow-ise from stdin to stdout
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/nosmo/raingow"
)

func main() {
	span := flag.Float64("span", 3.0, "colour spread per character")
	flag.Parse()

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	for scanner.Scan() {
		fmt.Fprintln(out, raingow.RaingowLine(scanner.Text(), *span))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "raingow:", err)
		os.Exit(1)
	}
}
