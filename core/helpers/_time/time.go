package _time

import (
	"github.com/kyaxcorp/go-core/core/helpers/function"
	"time"
)

type SleepStatus struct {
	Duration time.Duration
	Times    uint64
	Round    uint64
	Break    bool // If you want to break from sleep
}

func Sleep(
	// How much time a cycle should sleep
	duration time.Duration,
	// How many cycles it should do
	times uint64,
	// the callback after the time passes out
	callback func(sleepStatus *SleepStatus),
) {
	for i := 0; i < int(times); i++ {
		time.Sleep(duration)
		if function.IsCallable(callback) {
			sleepStatus := &SleepStatus{
				Duration: duration,
				Times:    times,
				Round:    uint64(i),
				Break:    false,
			}
			callback(sleepStatus)
			if sleepStatus.Break {
				break
			}
		}
	}
}
