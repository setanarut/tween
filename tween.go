// Package tween provides the Tween struct that allows an easing function to be
// animated over time. This can be used in tandem with the ease package to provide
// the easing functions.
package tween

import (
	"github.com/setanarut/tween/ease"
)

type (
	// Tween encapsulates the easing function along with timing data. This allows
	// a ease.TweenFunc to be used to be easily animated.
	Tween struct {
		Reverse bool

		time     float32
		begin    float32
		end      float32
		duration float32
		easing   ease.TweenFunc
		overflow float32
		change   float32
		current  float32
	}
)

// NewTween will return a new Tween when passed a beginning and end value, the duration
// of the tween and the easing function to animate between the two values. The
// easing function can be one of the provided easing functions from the ease package
// or you can provide one of your own.
func NewTween(begin, end, duration float32, easing ease.TweenFunc) *Tween {
	return &Tween{
		begin:    begin,
		end:      end,
		change:   end - begin,
		duration: duration,
		easing:   easing,
	}
}

// Set will set the current time along the duration of the tween. It will then return
// the current value as well as a boolean to determine if the tween is finished.
func (t *Tween) Set(time float32) {
	switch {
	case time <= 0.0:
		t.overflow = time
		t.time = 0.0
		t.current = t.begin
	case time >= t.duration:
		t.overflow = time - t.duration
		t.time = t.duration
		t.current = t.end
	default:
		t.overflow = 0.0
		t.time = time
		t.current = t.easing(t.time, t.begin, t.change, t.duration)
	}
}

// Current returns current tween value
func (t *Tween) Current() float32 {
	return t.current
}

// IsFinished will return true if the tween is finished.
func (t *Tween) IsFinished() bool {
	if t.Reverse {
		return t.time <= 0.0
	}
	return t.time >= t.duration
}

// Reset will set the Tween to the beginning of the two values.
func (t *Tween) Reset() {
	if t.Reverse {
		t.Set(t.duration)
	} else {
		t.Set(0.0)
	}
}

// Update will increment the timer of the Tween and ease the value.
func (t *Tween) Update(dt float32) {
	if t.Reverse {
		t.Set(t.time - dt)
	} else {
		t.Set(t.time + dt)
	}
}
