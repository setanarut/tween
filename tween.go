// Package tween provides the Tween struct that allows an easing function to be
// animated over time. This can be used in tandem with the ease package to provide
// the easing functions.
package tween

import (
	"encoding/json"
	"time"
)

// Tween encapsulates the easing function along with timing data. This allows
// a TweenFunc to be used to be easily animated.
type Tween struct {
	Value    float64
	Begin    float64
	End      float64
	Duration time.Duration
	Time     time.Duration
	Overflow time.Duration

	Reversed bool
	Yoyo     bool

	// EasingFunc function to use
	EasingFunc TweenFunc `json:"-"`
	EaseName   string
}

// NewTween will return a new Tween when passed a beginning and end value, the duration
// of the tween and the easing function to animate between the two values.
func NewTween(begin, end float64, duration time.Duration, easeName string, yoyo bool) *Tween {
	fn, ok := EaseMap[easeName]
	if !ok {
		fn = LinearFunc
	}
	return &Tween{
		Begin:      begin,
		End:        end,
		Duration:   duration,
		EasingFunc: fn,
		EaseName:   easeName,
		Yoyo:       yoyo,
	}
}

func (t *Tween) UnmarshalJSON(data []byte) error {
	type Alias Tween
	aux := &Alias{
		EasingFunc: LinearFunc,
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	*t = Tween(*aux)
	if fn, ok := EaseMap[t.EaseName]; ok {
		t.EasingFunc = fn
	} else {
		t.EasingFunc = LinearFunc
	}
	return nil
}

// SetTime will set the current time along the duration of the tween.
func (t *Tween) SetTime(currentTime time.Duration) {
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

// Change is the difference between the end and begin values
func (t *Tween) Change() float64 {
	return t.End - t.Begin
}

// IsFinished will return true if the tween is finished.
func (t *Tween) IsFinished() bool {
	if t.Reversed {
		return t.Time <= 0
	}
	return t.Time >= t.Duration
}

// Reset will set the Tween to the beginning of the two values.
func (t *Tween) Reset() *Tween {
	if t.Reversed {
		t.SetTime(t.Duration)
	} else {
		t.SetTime(0)
	}
	return t
}
func (t *Tween) Update(dt time.Duration) {
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

		if t.Reversed {
			t.SetTime(t.Duration - over)
		} else {
			t.SetTime(over)
		}
	}
}
