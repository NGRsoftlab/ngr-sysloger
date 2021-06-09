// Copyright 2020 NGR Softlab
//
package sysloger

import (
	"fmt"

	syslog "github.com/RackSec/srslog"
	log "github.com/sirupsen/logrus"
)

/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////

// Send list of msgs to syslog
func SendListToSyslog(params SyslogParams, formatter syslog.Formatter, msgList []string) error {
	writer, err := NewSyslogWriter(params, formatter)
	if err != nil {
		log.Error("ERROR failed to create writer:", err)
		return err
	}

	defer func() { _ = writer.Close() }()

	log.Info(fmt.Sprintf("need write: %d msgs to syslog", len(msgList)))

	for i, msg := range msgList {
		_, errW := doSuitableMethod(writer, params.Level, msg)
		if errW != nil {
			log.Warning(fmt.Sprintf("ERROR writing to syslog: %d %v", i, errW))
		}
	}

	return nil
}

// Send single syslog msg
func SendSingleSyslogMsg(params SyslogParams, formatter syslog.Formatter, msg string) error {
	writer, err := NewSyslogWriter(params, formatter)
	if err != nil {
		log.Error("ERROR failed to create writer:", err)
		return err
	}

	defer func() { _ = writer.Close() }()

	log.Info("need write: 1 msg to syslog")

	_, errW := doSuitableMethod(writer, params.Level, msg)
	if errW != nil {
		log.Error("ERROR writing to syslog: ", errW)
		return errW
	}

	return nil
}
