package raingow

import (
	"regexp"
	"testing"

	"github.com/gookit/color"
)

// string CSI escape sequences for comparison
var stripANSI = regexp.MustCompile(`\x1b\[[0-9;]*[A-Za-z]`)

func TestRaingowLinePreservesVisibleText(t *testing.T) {
	color.ForceOpenColor()

	cases := []string{
		"",
		"hello",
		"a string with spaces",
		"100% safe — even with format verbs %s %d %v",
	}
	for _, in := range cases {
		got := RaingowLine(in, 3.0)
		visible := stripANSI.ReplaceAllString(got, "")
		if visible != in {
			t.Errorf("RaingowLine(%q): visible text = %q, want %q", in, visible, in)
		}
	}
}

func TestRaingowLineEmitsColour(t *testing.T) {
	color.ForceOpenColor()

	got := RaingowLine("x", 3.0)
	if got == "x" {
		t.Fatalf("RaingowLine(%q) returned unchanged input; expected ANSI-wrapped output", "x")
	}
}
