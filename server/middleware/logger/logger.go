// sectfs - logger.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

const timestampFormat = "2006-01-02 15:04:05.000"

// LoadEchoLogger can load logger of Echo.
func LoadEchoLogger(e *echo.Echo) {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: timestampFormat,
		FullTimestamp:   true,
		ForceQuote:      true,
	})
	// log.SetFormatter(&logrus.JSONFormatter{
	// 	TimestampFormat: timestampFormat,
	// })
	middlewareRequestLogger := middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRemoteIP: true,
		LogURI:      true,
		LogStatus:   true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"remote_ip": values.RemoteIP,
				"url":       values.URI,
				"status":    values.Status,
			}).Info("REQUEST")
			return nil
		},
	})
	e.Use(middlewareRequestLogger)
}
