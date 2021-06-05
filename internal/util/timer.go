package util

import "time"

type Timer struct {
	start time.Time
	end   time.Time
}

func (t *Timer) Start() {
	t.start = time.Now()
}

func (t *Timer) Stop() time.Duration {
	t.end = time.Now()
	return t.Duration()
}

func (t *Timer) Duration() time.Duration {
	return t.end.Sub(t.start)
}
