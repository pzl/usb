package gusb

import (
	"github.com/apex/log"
)

var lg *log.Logger

func SetLogger(l *log.Logger) { lg = l }
