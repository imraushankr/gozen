package utils

import (
	"time"
)

const (
	// "1 Jan 2025, 12:30 PM" format
	CompactFormat = "2 Jan 2006, 3:04 PM"
)

// TimeFormatter handles the specific time format
type TimeFormatter struct{}

// NewTimeFormatter creates a new instance
func NewTimeFormatter() *TimeFormatter {
	return &TimeFormatter{}
}

// Format converts time.Time to "1 Jan 2025, 12:30 PM" format
func (tf *TimeFormatter) Format(t time.Time) string {
	return t.Format(CompactFormat)
}

// FormatFromISO converts ISO string directly to target format
func (tf *TimeFormatter) FormatFromISO(isoTime string) (string, error) {
	parsed, err := time.Parse(time.RFC3339Nano, isoTime)
	if err != nil {
		return "", err
	}
	return tf.Format(parsed), nil
}

// Current returns current time in target format
func (tf *TimeFormatter) Current() string {
	return tf.Format(time.Now())
}