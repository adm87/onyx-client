package engine

import "time"

const (
	maxAccumulator = time.Second / 4
	maxDeltaTime   = time.Second / 4
)

type Time struct {
	deltaTime   time.Duration
	fixedTime   time.Duration
	accumulator time.Duration
	last        time.Time
	steps       int
}

func NewTime(fps int) *Time {
	if fps <= 0 {
		panic("FPS must be greater than 0")
	}
	return &Time{
		fixedTime: time.Second / time.Duration(fps),
	}
}

func (t *Time) Update() {
	if t.last.IsZero() {
		t.last = time.Now()
		return
	}

	now := time.Now()
	t.deltaTime = now.Sub(t.last)
	if t.deltaTime > maxDeltaTime {
		t.deltaTime = maxDeltaTime
	} else if t.deltaTime < 0 {
		t.deltaTime = 0
	}
	t.last = now

	t.accumulator = min(t.accumulator+t.deltaTime, maxAccumulator)
	t.steps = 0

	for t.accumulator >= t.fixedTime {
		t.accumulator -= t.fixedTime
		t.steps++
	}
}

func (t *Time) DeltaTime() float64 {
	return t.deltaTime.Seconds()
}

func (t *Time) FixedTime() float64 {
	return t.fixedTime.Seconds()
}

func (t *Time) Steps() int {
	return t.steps
}
