package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	// Output to stdout instead of the default stderr.
	Log.SetOutput(os.Stdout)

	// You can set the log level. For development, DebugLevel can be useful.
	Log.SetLevel(logrus.DebugLevel)

	// You can also set up different formatters. For instance, a JSON formatter.
	Log.SetFormatter(&logrus.TextFormatter{
		PadLevelText: true,
	})
}
