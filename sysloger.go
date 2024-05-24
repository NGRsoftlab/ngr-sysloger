// Copyright 2020-2024 NGR Softlab
package sysloger

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"time"

	syslog "github.com/RackSec/srslog"

	logging "github.com/NGRsoftlab/ngr-logging"
)

// SyslogParams syslog dial params
type SyslogParams struct {
	Level    int    `json:"level"`    // syslog level info/error/fatal
	Host     string `json:"host"`     // host to send
	Port     int    `json:"port"`     // port to send
	Protocol string `json:"protocol"` // tcp\udp

	Priority syslog.Priority
	Tag      string `json:"tag"` // syslog tag

	NeedTls bool
	TlsConf *tls.Config
}

// NewSyslogWriter create new syslog writer with params and custom formatter
func NewSyslogWriter(params SyslogParams, formatter syslog.Formatter) (*syslog.Writer, error) {
	var sysLogger *syslog.Writer
	var err error

	if params.NeedTls {
		if params.TlsConf == nil {
			logging.Logger.Error("ERROR nil TlsConf")
			return nil, errors.New("nil TlsConf")
		}

		sysLogger, err = syslog.DialWithTLSConfig(params.Protocol,
			fmt.Sprintf("%s:%d", params.Host, params.Port), params.Priority, params.Tag, params.TlsConf)
	} else {
		sysLogger, err = syslog.Dial(params.Protocol,
			fmt.Sprintf("%s:%d", params.Host, params.Port), params.Priority, params.Tag)
	}

	if err != nil {
		logging.Logger.Errorf("ERROR failed to dial syslog: %s", err.Error())
		return nil, err
	}

	sysLogger.SetFormatter(formatter)

	return sysLogger, nil
}

// NewSyslogWriterWithTimeout create new syslog writer with params and custom formatter + timeout
func NewSyslogWriterWithTimeout(params SyslogParams, formatter syslog.Formatter, timeout time.Duration) (*syslog.Writer, error) {
	var sysLogger *syslog.Writer
	var err error
	var dial func(network, addr string) (net.Conn, error)
	netDialer := new(net.Dialer)
	netDialer.Timeout = timeout

	if params.NeedTls {
		if params.TlsConf == nil {
			logging.Logger.Error("ERROR nil TlsConf")
			return nil, errors.New("nil TlsConf")
		}

		dial = func(network, addr string) (net.Conn, error) {
			// cannot use "network" here as it'll simply be "custom" which will fail
			return tls.DialWithDialer(netDialer, params.Protocol, addr, params.TlsConf)
		}
	} else {
		dial = func(network, addr string) (net.Conn, error) {
			// cannot use "network" here as it'll simply be "custom" which will fail
			return net.DialTimeout(params.Protocol, addr, timeout)
		}
	}

	// param 'network' is contained in dial function (params.Protocol)
	sysLogger, err = syslog.DialWithCustomDialer("custom", fmt.Sprintf("%s:%d", params.Host, params.Port),
		params.Priority, params.Tag, dial)
	if err != nil {
		logging.Logger.Errorf("ERROR failed to dial syslog: %s", err.Error())
		return nil, err
	}

	sysLogger.SetFormatter(formatter)

	return sysLogger, nil
}
