package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	// EnvVars holds the configuration variables for the application.
	// It includes settings for the environment, application name, logging level,
	// timezone, and network configurations such as host and port.
	EnvVars struct {
		Env                string `mapstructure:"environment"`     // The environment in which the application is running (e.g., "local", "dev").
		AppName            string `mapstructure:"appName"`         // The name of the application.
		LogLevel           string `mapstructure:"logLevel"`        // The logging level (e.g., "info", "debug").
		UseLogLevelHook    bool   `mapstructure:"logLevelHook"`    // Indicates whether to use the log level hook.
		UseTimezoneLogHook bool   `mapstructure:"logTimezoneHook"` // Indicates whether to use the timezone log hook.
		Timezone           string `mapstructure:"timezone"`        // The timezone to be used for logging.

		Host        string `mapstructure:"host"`        // The host address for the application.
		Port        int    `mapstructure:"port"`        // The port number for the application.
		GatewayHost string `mapstructure:"gatewayHost"` // The host address for the gateway.
		GatewayPort int    `mapstructure:"gatewayPort"` // The port number for the gateway.
	}
)

var (
	// allowedEnvironments defines the list of valid environments that the application can run in.
	// These include: "test", "local", "dev", "hml", and "prd".
	allowedEnvironments = []string{
		"test",
		"local",
		"dev",
		"hml",
		"stg",
		"prd",
	}

	// ErrInvalidEnvironment is a function that returns an error indicating an invalid environment value.
	//
	// It provides feedback on the valid options available.
	ErrInvalidEnvironment = func(env string) error {
		return fmt.Errorf("'%s' is not a valid environment. Use one of %+v", env, allowedEnvironments)
	}

	Viper  *viper.Viper
	Config *EnvVars
)

// NewConfigs initializes a new EnvVars instance by reading configuration values
// from environment variables and configuration files.
//
// It returns a pointer to the EnvVars instance and an error if any occurs during
// the configuration loading process.
func NewConfigs() (*EnvVars, error) {
	env := os.Getenv("ENVIRONMENT")

	instance := viper.New()

	exePath, _ := os.Executable()
	rootDir := filepath.Dir(exePath)

	instance.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	instance.SetConfigName(mapEnvToPropertiesFilename(env))
	instance.AddConfigPath(rootDir)

	err := instance.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to read configuration")
	}

	instance.AutomaticEnv()

	err = validateEnvironment(instance.GetString("environment"))
	if err != nil {
		return nil, err
	}

	var envVars EnvVars
	if err := instance.Unmarshal(&envVars); err != nil {
		return nil, err
	}

	Viper = instance
	Config = &envVars

	return &envVars, nil
}

// validateEnvironment checks if the provided environment string is valid
// by comparing it against the list of allowed environments.
//
// Parameters:
// - env: The environment string to validate.
//
// Returns:
// An error if the environment is not valid; otherwise, it returns nil.
func validateEnvironment(env string) error {
	if !slices.Contains(allowedEnvironments, env) {
		return ErrInvalidEnvironment(env)
	}

	return nil
}

// mapEnvToPropertiesFilename maps the provided environment string to the corresponding
// properties filename used for configuration.
//
// Parameters:
// - env: The environment string (e.g., "local", "dev").
//
// Returns:
// The name of the properties file associated with the given environment.
func mapEnvToPropertiesFilename(env string) string {
	switch env {
	case "test":
		fallthrough
	case "local":
		return "properties.local"
	case "dev":
		return "properties.dev"
	case "hml":
		fallthrough
	case "stg":
		return "properties.stg"
	case "prd":
		return "properties.prd"
	default:
		return "properties.local"

	}
}
