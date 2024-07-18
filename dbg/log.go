package dbg

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Log(msg string, args ...interface{}) {
	str := fmt.Sprintf(msg, args...)
	logrus.Debug(str)
}
