package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	// Output to stdout instead of the default stderr.
	Log.SetOutput(os.Stdout)

	// You can also set up different formatters. For instance, a JSON formatter.
	Log.SetFormatter(&logrus.TextFormatter{
		PadLevelText: true,
		ForceColors:  true,
	})
}
