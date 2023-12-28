package time

import (
	"time"
)

type Time interface {
	Now() time.Time
	Since(t time.Time) time.Duration
}

type RealTime struct{}

func (t RealTime) Now() time.Time {
	return time.Now()
}

func (t RealTime) Since(s time.Time) time.Duration {
	return time.Since(s)
}

type MockTime struct{}

func (m MockTime) Now() time.Time {
	t, _ := time.Parse(DateTimeLayout, "2023-01-01 04:30:18")
	return t
}

func (m MockTime) Since(s time.Time) time.Duration {
	t, _ := time.Parse(DateTimeLayout, "2023-01-01 04:30:18")
	return t.Sub(s)
}
