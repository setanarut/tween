// Package tween provides the Tween struct that allows an easing function to be
// animated over time. This can be used in tandem with the ease package to provide
// the easing functions.
package tween

import (
	"encoding/json"
	"time"
)

// Fixed time step for 60 FPS
const FixedTimeStep = time.Second / 60

// Tween encapsulates the easing function along with timing data. This allows
// a TweenFunc to be used to be easily animated.
type Tween struct {
	Value    float64
	Begin    float64
	End      float64
	Duration time.Duration
	Time     time.Duration
	Overflow time.Duration

	// Delay is the amount of time to wait before the tween begins animating.
	// During the delay period Value holds Begin (or End if Reversed).
	// Delay does NOT repeat on Yoyo direction changes — it only applies once
	// at the very start (or after an explicit Reset).
	Delay        time.Duration
	delayElapsed time.Duration
	delayDone    bool

	Reversed bool
	Yoyo     bool

	// EasingFunc is the easing function to use.
	EasingFunc TweenFunc `json:"-"`
	EaseName   string
}

// NewTween returns a new Tween given begin/end values, a duration, an easing
// function name (from EaseMap), and a yoyo flag.
func NewTween(begin, end float64, duration time.Duration, easeName string, yoyo bool) *Tween {
	fn, ok := EaseMap[easeName]
	if !ok {
		fn = LinearFunc
	}
	return &Tween{
		Value:      begin,
		Begin:      begin,
		End:        end,
		Duration:   duration,
		EasingFunc: fn,
		EaseName:   easeName,
		Yoyo:       yoyo,
	}
}

// NewTweenWithDelay is like NewTween but also sets an initial delay.
func NewTweenWithDelay(delay time.Duration, begin, end float64, duration time.Duration, easeName string, yoyo bool) *Tween {
	t := NewTween(begin, end, duration, easeName, yoyo)
	t.Delay = delay
	return t
}

func (t *Tween) UnmarshalJSON(data []byte) error {
	type Alias Tween
	aux := &Alias{EasingFunc: LinearFunc}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	*t = Tween(*aux)
	if fn, ok := EaseMap[t.EaseName]; ok {
		t.EasingFunc = fn
	} else {
		t.EasingFunc = LinearFunc
	}
	// Restore delayDone based on serialised state: if Time > 0 the delay was
	// already consumed before serialisation.
	if t.Time > 0 {
		t.delayDone = true
		t.delayElapsed = t.Delay
	}
	return nil
}

// SetTime sets the current animation time, bypassing any remaining delay.
// Callers that want delay-aware scrubbing should use Update instead.
func (t *Tween) SetTime(currentTime time.Duration) {
	t.delayDone = true
	t.delayElapsed = t.Delay

	switch {
	case currentTime <= 0:
		t.Overflow = currentTime
		t.Time = 0
		t.Value = t.Begin
	case currentTime >= t.Duration:
		t.Overflow = currentTime - t.Duration
		t.Time = t.Duration
		t.Value = t.End
	default:
		t.Overflow = 0
		t.Time = currentTime
		t.Value = t.EasingFunc(t.Time.Seconds(), t.Begin, t.Change(), t.Duration.Seconds())
	}
}

// Change returns the difference between End and Begin.
func (t *Tween) Change() float64 {
	return t.End - t.Begin
}

// IsFinished returns true when the tween has completed.
func (t *Tween) IsFinished() bool {
	if t.Reversed {
		return t.Time <= 0
	}
	return t.Time >= t.Duration
}

// IsDelaying returns true while the initial delay has not yet elapsed.
func (t *Tween) IsDelaying() bool {
	return !t.delayDone
}

// Reset sets the tween back to its starting position and re-arms the delay.
func (t *Tween) Reset() *Tween {
	t.delayDone = t.Delay <= 0
	t.delayElapsed = 0

	if !t.delayDone {
		// Hold at the correct boundary value while waiting.
		if t.Reversed {
			t.Value = t.End
			t.Time = t.Duration
		} else {
			t.Value = t.Begin
			t.Time = 0
		}
		t.Overflow = 0
		return t
	}

	if t.Reversed {
		t.SetTime(t.Duration)
	} else {
		t.SetTime(0)
	}
	return t
}

// Update advances the tween by dt. It first consumes any remaining delay,
// then advances the animation clock.
func (t *Tween) Update(dt time.Duration) {
	if dt == 0 {
		return
	}

	// --- consume delay ---
	if !t.delayDone {
		remaining := t.Delay - t.delayElapsed
		if dt < remaining {
			t.delayElapsed += dt
			// Still in delay — value stays at boundary, overflow is zero.
			t.Overflow = 0
			return
		}
		// Delay fully consumed; carry the leftover into animation.
		dt -= remaining
		t.delayElapsed = t.Delay
		t.delayDone = true

		if dt == 0 {
			t.Overflow = 0
			return
		}
	}

	// --- advance animation ---
	if t.Reversed {
		t.SetTime(t.Time - dt)
	} else {
		t.SetTime(t.Time + dt)
	}

	if t.Yoyo && t.IsFinished() {
		over := t.Overflow
		if over < 0 {
			over = -over
		}
		t.Reversed = !t.Reversed

		// Yoyo does NOT re-arm the delay on direction changes.
		if t.Reversed {
			t.SetTime(t.Duration - over)
		} else {
			t.SetTime(over)
		}
	}
}

// TotalDuration returns Duration + Delay (the wall-clock time from Reset to
// completion, ignoring Yoyo repetitions).
func (t *Tween) TotalDuration() time.Duration {
	return t.Delay + t.Duration
}
