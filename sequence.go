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

	// Delay is the amount of time to wait before the very first tween starts.
	// It is re-armed by Reset. It does NOT fire again on Yoyo reversal.
	Delay        time.Duration
	delayElapsed time.Duration
	delayDone    bool

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

// NewSequenceWithDelay returns a new Sequence with an initial delay before the
// first tween begins.
func NewSequenceWithDelay(delay time.Duration, tweens ...*Tween) *Sequence {
	s := NewSequence(tweens...)
	s.Delay = delay
	s.delayDone = delay <= 0
	return s
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

// Update advances the sequence by dt. It first consumes any remaining
// sequence-level delay, then drives the active tween forward.
func (s *Sequence) Update(dt time.Duration) {
	if !s.HasTweens() {
		s.Value = 0
		s.IsActiveTweenFinished = false
		s.IsFinished = true
		return
	}

	// --- consume sequence-level delay ---
	if !s.delayDone {
		remaining := s.Delay - s.delayElapsed
		if dt < remaining {
			s.delayElapsed += dt
			s.IsFinished = false
			s.IsActiveTweenFinished = false
			return
		}
		dt -= remaining
		s.delayElapsed = s.Delay
		s.delayDone = true
		if dt == 0 {
			s.IsFinished = false
			s.IsActiveTweenFinished = false
			return
		}
	}

	// --- drive tweens ---
	var completed []int
	remaining := dt

	for {
		if s.Yoyo {
			if s.Index < 0 {
				s.IsReversed = false
				s.Index = s.clampIndex(s.Index)
				if s.LoopRemaining >= 1 {
					s.LoopRemaining--
				}
				if s.LoopRemaining == 0 || remaining == 0 {
					s.Value = s.Tweens[s.Index].Begin
					s.IsActiveTweenFinished = len(completed) > 0
					s.IsFinished = true
					return
				}
				s.Tweens[s.Index].Reversed = s.IsReversed
				s.Tweens[s.Index].Reset()
			}
			if s.Index >= len(s.Tweens) {
				s.IsReversed = true
				s.Index = s.clampIndex(s.Index)
				s.Tweens[s.Index].Reversed = s.IsReversed
				s.Tweens[s.Index].Reset()
			}
		} else if s.Index >= len(s.Tweens) || s.Index <= -1 {
			if s.LoopRemaining >= 1 {
				s.LoopRemaining--
			}
			if s.LoopRemaining == 0 || remaining == 0 {
				idx := s.clampIndex(s.Index)
				if s.IsReversed {
					s.Value = s.Tweens[idx].Begin
				} else {
					s.Value = s.Tweens[idx].End
				}
				s.IsActiveTweenFinished = len(completed) > 0
				s.IsFinished = true
				return
			}
			s.Index = s.wrapIndex(s.Index)
			s.Tweens[s.Index].Reversed = s.IsReversed
			s.Tweens[s.Index].Reset()
		}

		s.Tweens[s.Index].Update(remaining)

		if !s.Tweens[s.Index].IsFinished() {
			s.Value = s.Tweens[s.Index].Value
			s.IsActiveTweenFinished = len(completed) > 0
			s.IsFinished = false
			return
		}

		remaining = s.Tweens[s.Index].Overflow
		if remaining < 0 {
			remaining = -remaining
		}
		completed = append(completed, s.Index)

		if s.IsReversed {
			s.Index--
		} else {
			s.Index++
		}
		if s.Index < len(s.Tweens) && s.Index >= 0 {
			s.Tweens[s.Index].Reversed = s.IsReversed
			s.Tweens[s.Index].Reset()
		}
	}
}

// Duration returns the total animation duration by summing each tween's
// TotalDuration (Delay + Duration). The sequence-level Delay is NOT included
// here; use TotalDuration for the full wall-clock length.
func (s *Sequence) Duration() time.Duration {
	var total time.Duration
	for _, t := range s.Tweens {
		total += t.TotalDuration()
	}
	return total
}

// TotalDuration returns the sequence-level Delay plus the summed tween durations.
func (s *Sequence) TotalDuration() time.Duration {
	return s.Delay + s.Duration()
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

// Reset resets the sequence and all contained tweens to their initial state,
// and re-arms the sequence-level delay.
func (s *Sequence) Reset() {
	s.LoopRemaining = s.Loop
	s.Index = 0
	s.IsFinished = false
	s.IsActiveTweenFinished = false

	s.delayDone = s.Delay <= 0
	s.delayElapsed = 0

	for _, t := range s.Tweens {
		t.Reset()
	}
}

// HasTweens reports whether the Sequence contains any Tweens.
func (s *Sequence) HasTweens() bool {
	return len(s.Tweens) > 0
}

// IsDelaying reports whether the sequence-level delay is still counting down.
func (s *Sequence) IsDelaying() bool {
	return !s.delayDone
}

// SetReversed configures the playback direction.
func (s *Sequence) SetReversed(r bool) {
	if s.Index >= len(s.Tweens) || s.Index < 0 {
		s.Index = s.clampIndex(s.Index)
	}
	s.Tweens[s.Index].Reversed = r
	s.IsReversed = r
}

func (s *Sequence) clampIndex(index int) int {
	if index < 0 {
		return 0
	}
	if index >= len(s.Tweens) {
		return len(s.Tweens) - 1
	}
	return index
}

func (s *Sequence) wrapIndex(index int) int {
	if index >= len(s.Tweens) {
		return 0
	}
	if index < 0 {
		return len(s.Tweens) - 1
	}
	return index
}
