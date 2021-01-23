package rollinglog

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	path := "."
	Init("",path)
	Logger.Info("info")
	Logger.Warn("warning")
	time.Sleep(time.Second)

}
