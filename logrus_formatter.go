// Spice up your log messages with unexpected control characters and
// readability impairments!

package raingow

import "github.com/sirupsen/logrus"


type RainbowFormatter struct {
	logrus_formatter *logrus.TextFormatter
}

func New() *logrus.Logger {
	logrus_formatter := logrus.New()
	logrus_formatter.Formatter = MakeRainbowFormatter()
	return logrus_formatter
}

func MakeRainbowFormatter() *RainbowFormatter {
	return &RainbowFormatter {
		logrus_formatter: &logrus.TextFormatter{},
	}
}

func (f *RainbowFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	entry.Message = Raingow_line(entry.Message, 3.0)

 	return f.logrus_formatter.Format(entry)
}
