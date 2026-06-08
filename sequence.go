package tween

import (
	"slices"
	"time"
)

// Sequence represents an ordered chain of Tweens executed one after another.
type Sequence struct {
	Tweens []*Tween

	// Yoyo makes the sequence reverse back to the start after reaching the end.
	Yoyo bool
	// IsReversed runs the sequence backwards when true.
	IsReversed bool

	Index int
	// Loop is the canonical loop count (0 = infinite).
	Loop int
	// LoopRemaining is the number of loops still to execute.
	LoopRemaining int

	Value                 float64
	IsActiveTweenFinished bool
	IsFinished            bool
}

// NewSequence returns a new Sequence containing the provided tweens.
func NewSequence(tweens ...*Tween) *Sequence {
	return &Sequence{
		Tweens:        tweens,
		LoopRemaining: 1,
		Loop:          1,
	}
}

// Add appends one or more Tweens to the end of the Sequence.
func (s *Sequence) Add(tweens ...*Tween) {
	s.Tweens = append(s.Tweens, tweens...)
}

// Remove deletes the Tween at the given index from the Sequence.
func (s *Sequence) Remove(index int) {
	if index >= 0 && index < len(s.Tweens) {
		s.Tweens = slices.Delete(s.Tweens, index, index+1)
	}
}

// Update advances the sequence by dt. It drives the active tween forward.
func (s *Sequence) Update(dt time.Duration) {
	if len(s.Tweens) == 0 {
		s.Value = 0
		s.IsActiveTweenFinished = false
		s.IsFinished = true
		return
	}

	remaining := dt
	completedAny := false

	for {
		activeTween := s.Tweens[s.Index]
		activeTween.Update(remaining)

		if !activeTween.IsFinished() {
			s.Value = activeTween.Value
			s.IsActiveTweenFinished = completedAny
			s.IsFinished = false
			return
		}

		remaining = activeTween.Overflow
		if remaining < 0 {
			remaining = -remaining
		}

		completedAny = true

		nextIndex := s.Index
		if s.IsReversed {
			nextIndex--
		} else {
			nextIndex++
		}

		if s.Yoyo {
			if nextIndex < 0 {
				s.IsReversed = false
				if s.LoopRemaining >= 1 {
					s.LoopRemaining--
				}
				if s.LoopRemaining == 0 || remaining == 0 {
					s.Value = activeTween.Begin
					s.IsActiveTweenFinished = true
					s.IsFinished = true
					return
				}
				nextIndex = 0
			} else if nextIndex >= len(s.Tweens) {
				s.IsReversed = true
				nextIndex = len(s.Tweens) - 1
			}
		} else if nextIndex >= len(s.Tweens) || nextIndex < 0 {
			if s.LoopRemaining >= 1 {
				s.LoopRemaining--
			}
			if s.LoopRemaining == 0 || remaining == 0 {
				if s.IsReversed {
					s.Value = activeTween.Begin
				} else {
					s.Value = activeTween.End
				}
				s.IsActiveTweenFinished = true
				s.IsFinished = true
				return
			}
			if nextIndex >= len(s.Tweens) {
				nextIndex = 0
			} else {
				nextIndex = len(s.Tweens) - 1
			}
		}

		s.Index = nextIndex
		nextTween := s.Tweens[s.Index]
		nextTween.Reversed = s.IsReversed
		nextTween.Reset()
	}
}

// SetReversed changes the playback direction of the sequence.
// It reverses both the overall sequence flow and the currently active tween.
func (s *Sequence) SetReversed(r bool) {
	s.Tweens[s.Index].Reversed = r
	s.IsReversed = r
}

// ActiveTween returns active *Tween
func (s *Sequence) ActiveTween() *Tween {
	return s.Tweens[s.Index]
}

// TotalDuration returns the total animation duration by summing each tween's
// TotalDuration (Delay + TotalDuration).
func (s *Sequence) TotalDuration() time.Duration {
	var total time.Duration
	for _, t := range s.Tweens {
		total += t.TotalDuration()
	}
	return total
}

// SetIndex resets the current tween and moves the active index.
func (s *Sequence) SetIndex(index int) {
	s.Tweens[s.Index].Reversed = s.IsReversed
	s.Tweens[s.Index].Reset()
	s.Index = index
}

// SetLoop sets both the canonical loop count and resets the remaining counter.
func (s *Sequence) SetLoop(amount int) {
	s.Loop = amount
	s.LoopRemaining = amount
}

// Reset resets the sequence and all contained tweens to their initial state.
func (s *Sequence) Reset() {
	s.LoopRemaining = s.Loop
	s.Index = 0
	s.IsFinished = false
	s.IsActiveTweenFinished = false

	for _, t := range s.Tweens {
		t.Reset()
	}
}

// HasTweens reports whether the Sequence contains any Tweens.
func (s *Sequence) HasTweens() bool {
	return len(s.Tweens) > 0
}
