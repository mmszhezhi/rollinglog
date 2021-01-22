package rollinglog

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"logfile/rlog"
	"path/filepath"
)
var Logger *logrus.Logger

func Init(levels []logrus.Level) {
	hook, err := rlog.NewTimeBasedRollingFileHook("ff",
		levels,
		&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% - %lvl%: %msg% \n",
		},
		filepath.Join(".","%Y-%m-%d-%H.log"))

	if err != nil {
		panic(err)
	}
	// Create a new logrus.Logger
	Logger = logrus.New()
	Logger.Hooks.Add(hook)
	Logger.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "%time% - %lvl%: %msg% \n",
	})

}