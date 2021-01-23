package rollinglog

import (
	"github.com/mmszhezhi/rollinglog"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"

	"path/filepath"
)
var Logger *logrus.Logger

func Init(format string,path string) {
	levels := []logrus.Level{logrus.InfoLevel,logrus.ErrorLevel,logrus.WarnLevel,logrus.FatalLevel,logrus.PanicLevel}
	if(path==""){
		path = "."
	}
	if(format==""){
		format = "%time% - %lvl%: %msg% \n"
	}
	hook, err := rollinglog.NewTimeBasedRollingFileHook("ff",
		levels,
		&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       format,
		},
		filepath.Join(path,"%Y-%m-%d-%H.log"))

	if err != nil {
		panic(err)
	}
	// Create a new logrus.Logger
	Logger = logrus.New()
	Logger.Hooks.Add(hook)
	Logger.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       format,
	})

}