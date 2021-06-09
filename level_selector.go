// Copyright 2020 NGR Softlab
//
package sysloger

import (
	"errors"
	syslog "github.com/RackSec/srslog"
)

// Syslog levels
const (
	SimpleLevel = iota // just write
	DebugLevel
	InfoLevel
	WarningLevel
	ErrorLevel
	AlertLevel
	CriticalLevel
)

func doSuitableMethod(writer *syslog.Writer, level int, msg string) (int, error) {
	var err error
	var n int

	if writer == nil {
		return n, errors.New("nil writer")
	}

	switch level {
	case SimpleLevel:
		n, err = writer.Write([]byte(msg))
	case DebugLevel:
		err = writer.Debug(msg)
	case InfoLevel:
		err = writer.Info(msg)
	case WarningLevel:
		err = writer.Warning(msg)
	case ErrorLevel:
		err = writer.Err(msg)
	case AlertLevel:
		err = writer.Alert(msg)
	case CriticalLevel:
		err = writer.Crit(msg)
	default:
		err = errors.New("no such level")
	}

	return n, err
}
