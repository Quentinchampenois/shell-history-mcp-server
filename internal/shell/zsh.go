package shell

import (
	"regexp"
	"strconv"
	"time"
)

// Row history pattern:
// : <timestamp>:<duration in seconds>;<command>
const ZshRowRegexp = `^:\s+(\d+):(\d+);(.*)$`

type ZshHistory struct {
	Timestamp time.Time `json:"timestamp" jsonschema:"the timestamp"`
	Duration  int       `json:"duration" jsonschema:"the duration in seconds"`
	Command   string    `json:"command" jsonschema:"the shell command"`
}

func NewZshHistory(zshRow string) *ZshHistory {
	var regex = regexp.MustCompile(`^:\s+(\d+):(\d+);(.*)$`)
	matches := regex.FindStringSubmatch(zshRow)

	if matches == nil {
		return nil
	}

	timestamp, _ := strconv.ParseInt(matches[1], 10, 64)
	t := time.Unix(timestamp, 0)

	duration, _ := strconv.Atoi(matches[2])
	command := matches[3]

	return &ZshHistory{
		Timestamp: t,
		Duration:  duration,
		Command:   command,
	}
}
