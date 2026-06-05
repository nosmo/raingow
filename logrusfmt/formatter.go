// Spice up your log messages with unexpected control characters and
// readability impairments!

package logrusfmt

import (
	"github.com/nosmo/raingow"
	"github.com/sirupsen/logrus"
)

// amount to span when no argument is given
const DefaultSpan = 3.0

type Formatter struct {
	Inner *logrus.TextFormatter
	Span  float64
}

func New() *Formatter {
	return &Formatter{
		Inner: &logrus.TextFormatter{},
		Span:  DefaultSpan,
	}
}

// return a fresh *logrus.Logger with the raingow formatter
func NewLogger() *logrus.Logger {
	l := logrus.New()
	l.Formatter = New()
	return l
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	span := f.Span
	if span == 0 {
		span = DefaultSpan
	}
	entry.Message = raingow.RaingowLine(entry.Message, span)
	return f.Inner.Format(entry)
}
