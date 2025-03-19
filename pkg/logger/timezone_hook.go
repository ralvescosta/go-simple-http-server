package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

type TimezoneHook struct {
	// location holds the time.Location for the specified timezone.
	location *time.Location
}

// NewTimezoneHook creates a new TimezoneHook instance for the specified timezone.
//
// Parameters:
// - locationName: A string representing the name of the timezone (e.g., "America/New_York").
// If the timezone cannot be loaded, an error is logged, and the hook will still be created
// with a nil location.
func NewTimezoneHook(locationName string) *TimezoneHook {
	location, err := time.LoadLocation(locationName)
	if err != nil {
		logrus.WithError(err).Errorf("Could not load timezone location: %v", err)
	}

	return &TimezoneHook{location: location}
}

// Fire is called by Logrus when an entry is logged. It adjusts the entry's time to the
// specified timezone.
//
// Parameters:
// - entry: A pointer to the logrus.Entry that contains the log information.
// The entry's time will be converted to the timezone specified in the TimezoneHook.
func (hook *TimezoneHook) Fire(entry *logrus.Entry) error {
	entry.Time = entry.Time.In(hook.location)
	return nil
}

// Levels returns the log levels that this hook is interested in.
//
// This method returns all log levels, indicating that the TimezoneHook will adjust the time
// for any log entry regardless of its level.
func (hook *TimezoneHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
