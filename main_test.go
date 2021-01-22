package rollinglog

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestInit(t *testing.T) {

	Init([]logrus.Level{logrus.InfoLevel,logrus.ErrorLevel,logrus.WarnLevel})
	Logger.Info("fuckyou")
	Logger.Warn("fuck warning")

}
