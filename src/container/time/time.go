package time

import (
	"time"
)

type Time interface {
	Now() time.Time
}

type RealTime struct{}

func (t RealTime) Now() time.Time {
	return time.Now()
}

type MockTime struct{}

func (m MockTime) Now() time.Time {
	t, _ := time.Parse(DateTimeLayout, "2023-01-01 04:30:18")
	return t
}
