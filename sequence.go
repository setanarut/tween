package tween

import (
	"slices"
	"time"
)

// Sequence represents a sequence of Tweens, executed one after the other.
type Sequence struct {
	Tweens []*Tween
	// Yoyo makes the sequence "Yoyo" back to the beginning after it reaches the end
	Yoyo bool

	// IsReversed runs the sequence backwards when true
	IsReversed bool

	Index int
	// Loop is the initial number of loops for this sequence to make
	Loop int
	// LoopRemaining is the remaining number of times to loop through the sequence
	LoopRemaining int

	Value                 float64
	IsActiveTweenFinished bool
	IsFinished            bool
}

// NewSequence returns a new Sequence object.
func NewSequence(tweens ...*Tween) *Sequence {
	seq := &Sequence{
		Tweens:        tweens,
		Yoyo:          false,
		IsReversed:    false,
		LoopRemaining: 1,
		Loop:          1,
	}
	return seq
}

// Add adds one or more Tweens in order to the Sequence.
func (s *Sequence) Add(tweens ...*Tween) {
	s.Tweens = append(s.Tweens, tweens...)
}

// Remove removes a Tween of the specified index from the Sequence.
func (s *Sequence) Remove(index int) {
	if index >= 0 && index < len(s.Tweens) {
		s.Tweens = slices.Delete(s.Tweens, index, index+1)
	}
}

// Update updates the currently active Tween in the Sequence; once that Tween is done, the Sequence moves onto the next one.
func (s *Sequence) Update(dt time.Duration) {
	if !s.HasTweens() {
		s.Value = 0
		s.IsActiveTweenFinished = false
		s.IsFinished = true
		return
	}
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
				if s.IsReversed {
					s.Value = s.Tweens[s.clampIndex(s.Index)].Begin
					s.IsActiveTweenFinished = len(completed) > 0
					s.IsFinished = true
					return

				}
				s.Value = s.Tweens[s.clampIndex(s.Index)].End
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
		completed = append(completed, s.Index)
		if remaining < 0 {
			remaining *= -1
		}
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

// Duration calculates and returns the total duration of the Sequence by summing the durations of all Tweens.
func (s *Sequence) Duration() time.Duration {
	if s.HasTweens() {
		var total time.Duration
		for _, t := range s.Tweens {
			total += t.Duration
		}
		return total
	}
	return 0
}

// SetIndex sets the current index of the Sequence.
func (s *Sequence) SetIndex(index int) {
	s.Tweens[s.Index].Reversed = s.IsReversed
	s.Tweens[s.Index].Reset()
	s.Index = index
}

// SetLoop sets the default loop and the current remaining loops.
func (s *Sequence) SetLoop(amount int) {
	s.Loop = amount
	s.LoopRemaining = s.Loop
}

// Reset resets the Sequence, resetting all Tweens and setting the Sequence's index back to 0.
func (seq *Sequence) Reset() {
	seq.LoopRemaining = seq.Loop
	for _, tween := range seq.Tweens {
		tween.Reset()
	}
	seq.Index = 0
}

// HasTweens returns whether the Sequence is populated with Tweens or not.
func (s *Sequence) HasTweens() bool {
	return len(s.Tweens) > 0
}

// SetReversed sets whether the Sequence will start running in reverse.
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
