package time_test

import (
	Time "golang-backend-microservice/container/time"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	expectedRealTime := time.Now()
	gotRealTime := Time.RealTime{}.Now()
	expectedMockTime, _ := time.Parse(Time.DateTimeLayout, "2023-01-01 04:30:18")
	gotMockTime := Time.MockTime{}.Now()

	if expectedRealTime.Format(Time.DateTimeLayout) != gotRealTime.Format(Time.DateTimeLayout) {
		t.Errorf("Expected %v, got %v: incorrect real time", expectedRealTime.Format(Time.DateTimeLayout), gotRealTime.Format(Time.DateTimeLayout))
	}

	if expectedMockTime != gotMockTime {
		t.Errorf("Expected %v, got %v: incorrect mock time", expectedMockTime, gotMockTime)
	}
}
