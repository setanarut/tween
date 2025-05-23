// Package tween provides the Tween struct that allows an easing function to be
// animated over time. This can be used in tandem with the ease package to provide
// the easing functions.
package tween

import (
	"github.com/setanarut/tween/ease"
)

// Tween encapsulates the easing function along with timing data. This allows
// a ease.TweenFunc to be used to be easily animated.
type Tween struct {
	Begin    float64        // Begin value of the tween
	End      float64        // End value of the tween
	Duration float64        // Duration of the tween
	Easing   ease.TweenFunc // Easing function to use
	Reversed bool           // Reversed will reverse the tween
	Yoyo     bool           // Enable yoyo

	time     float64 // Current time
	overflow float64 // Time overflow is the time that is left over
	value    float64 // Value is the current value of the tween
}

// NewTween will return a new Tween when passed a beginning and end value, the duration
// of the tween and the easing function to animate between the two values. The
// easing function can be one of the provided easing functions from the ease package
// or you can provide one of your own.
func NewTween(begin, end, duration float64, easing ease.TweenFunc) *Tween {
	return &Tween{
		Begin:    begin,
		End:      end,
		Duration: duration,
		Easing:   easing,
	}
}

// SetTime will set the current time along the duration of the tween.
func (t *Tween) SetTime(time float64) {
	switch {
	case time <= 0.0:
		t.overflow = time
		t.time = 0.0
		t.value = t.Begin
	case time >= t.Duration:
		t.overflow = time - t.Duration
		t.time = t.Duration
		t.value = t.End
	default:
		t.overflow = 0.0
		t.time = time
		t.value = t.Easing(t.time, t.Begin, t.change(), t.Duration)
	}
}

// Value returns current tween value
func (t *Tween) Value() float64 {
	return t.value
}

// SetYoyo sets yoyo and returns self
func (t *Tween) SetYoyo(y bool) *Tween {
	t.Yoyo = y
	return t
}

// SetReversed sets Reversed and returns self
func (t *Tween) SetReversed(r bool) *Tween {
	t.Reversed = r
	return t
}

// change is the difference between the end and begin values
func (t *Tween) change() float64 {
	return t.End - t.Begin
}

// IsFinished will return true if the tween is finished.
func (t *Tween) IsFinished() bool {
	if t.Reversed {
		return t.time <= 0.0
	} else {
		return t.time >= t.Duration
	}
}

// Reset will set the Tween to the beginning of the two values.
func (t *Tween) Reset() *Tween {
	if t.Reversed {
		t.SetTime(t.Duration)
	} else {
		t.SetTime(0.0)
	}
	return t
}

// Update will increment the timer of the Tween and ease the value. Unit is seconds.
//
//	25 FPS 1 frame increment = 1/25 = 0.04 dt
//	60 FPS 1 frame increment = 1/60 = 0.016666666666666666 dt
//	120 FPS 1 frame increment = 1/120 = 0.008333333333333333 dt
func (t *Tween) Update(dt float64) {
	if t.Reversed {
		t.SetTime(t.time - dt)
	} else {
		t.SetTime(t.time + dt)
	}

	if t.Yoyo {
		if t.IsFinished() {
			t.Reversed = !t.Reversed
		}
	}
}
