package mock

import (
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/clock"
	"time"
)

var originalMockedTime = time.Now().UTC()

// FakeClockNow mocks clock.Now
func FakeClockNow() {
	clock.Now = func() time.Time {
		return originalMockedTime
	}
}

// FakeClockNowWithDuration moves the clock by with given duration
func FakeClockNowWithExtraTime(d time.Duration) {
	clock.Now = func() time.Time {
		return originalMockedTime.Add(d)
	}
}
