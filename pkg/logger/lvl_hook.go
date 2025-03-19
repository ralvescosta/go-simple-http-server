package logger

import "github.com/sirupsen/logrus"

// CustomHook wraps the otellogrus Hook to add log level filtering
type LogLeveFilterlHook struct {
	levelsAllowed    []logrus.Level
	levelsAllowedMap map[logrus.Level]bool
	inner            logrus.Hook
}

// NewLogLevelFilterHook create a new instance of the LogLevelFilterHook struct
//
// @Arg<inner>: Receive the logrus default hooks and the
//
// @Arg<level>: Will be the level that will be filtered, all log levels grater or equal than the
// this arg will be sent the others will be ignored
func NewLogLevelFilterHook(inner logrus.Hook, level logrus.Level) *LogLeveFilterlHook {
	graterOrEqual := uint32(level)
	levelsAllowed := []logrus.Level{}
	levelsAllowedMap := map[logrus.Level]bool{}

	for i := graterOrEqual; i != 0; i-- {
		l := logrus.Level(i)
		levelsAllowed = append(levelsAllowed, l)
		levelsAllowedMap[l] = true
	}

	return &LogLeveFilterlHook{levelsAllowed, levelsAllowedMap, inner}
}

// Levels filters log levels; it only returns levels >= logrus.WarnLevel
func (c *LogLeveFilterlHook) Levels() []logrus.Level {
	return c.levelsAllowed
}

// Fire delegates log entry processing to the wrapped hook
func (c *LogLeveFilterlHook) Fire(entry *logrus.Entry) error {
	if _, ok := c.levelsAllowedMap[entry.Level]; ok {
		return c.inner.Fire(entry)
	}

	return nil
}
