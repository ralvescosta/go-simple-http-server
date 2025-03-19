// Package logger provides utilities for initializing and configuring the logging system.
//
// This package utilizes Logrus as the underlying logging framework, allowing for flexible
// logging configurations based on the application's environment. It supports various log
// formats and levels, making it suitable for different deployment scenarios such as
// development, staging, and production.
//
// The logging behavior can be customized through environment variables, enabling developers
// to easily switch between formats and log levels as needed.
package logger

import (
	"strings"

	"github.com/sirupsen/logrus"

	"githib.com/ralvescosta/go-simple-http-server/pkg/configs"
)

// formatterEnvMap maps environments to their corresponding log formatters.
// JSON format is used for production-like environments, while the default formatter is used for testing and local development.
var formatterEnvMap = map[string]logrus.Formatter{
	"test":  logrus.StandardLogger().Formatter,
	"local": logrus.StandardLogger().Formatter,
	"dev":   &logrus.JSONFormatter{},
	"stg":   &logrus.JSONFormatter{},
	"hml":   &logrus.JSONFormatter{},
	"prd":   &logrus.JSONFormatter{},
}

// SetupLogger initializes the logger with a specific log format and level based on the provided environment and log level.
//
// Parameters:
//   - envs: A pointer to a configs.EnvVars struct that contains the current environment settings,
//     including the environment name (e.g., "dev", "prd") and the desired logging level (e.g., "info", "warn").
//
// The function selects the appropriate log formatter based on the environment. If no formatter
// is found for the specified environment, it falls back to the standard formatter. Additionally,
// it parses the log level and sets it accordingly, defaulting to INFO level if the provided level
// is invalid. If timezone logging is enabled, it adds a timezone hook to the logger.
//
// This function logs the final logging level set for the application.
//
// - `env`: Specifies the current environment (e.g., "dev", "prd").
//
// - `level`: Specifies the logging level (e.g., "info", "warn").
//
// If no formatter is found for the environment, a warning is logged, and the standard formatter is used.
//
// If the log level is invalid, a warning is logged, and the default level INFO is used.
func SetupLogger(envs *configs.EnvVars) {
	formatter, ok := formatterEnvMap[envs.Env]
	if !ok {
		logrus.Warnf("No formatter found for environment '%s'. Falling back to standard formatter", envs.Env)
		formatter = logrus.StandardLogger().Formatter
	}

	logrus.SetFormatter(formatter)

	logLevel, err := logrus.ParseLevel(envs.LogLevel)
	if err != nil {
		logrus.WithError(err).Warn("Failed to parse log level. Falling back to INFO level")
		logLevel = logrus.InfoLevel
	}

	logrus.SetLevel(logLevel)

	if envs.UseTimezoneLogHook {
		logrus.AddHook(NewTimezoneHook(envs.Timezone))
	}

	logrus.Infof("Logging level set to %s", strings.ToUpper(logLevel.String()))
}
