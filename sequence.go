package tween

import "slices"

// Sequence represents a sequence of Tweens, executed one after the other.
type Sequence struct {
	Tweens []*Tween
	// Yoyo makes the sequence "Yoyo" back to the beginning after it reaches the end
	Yoyo bool

	// Reversed runs the sequence backwards when true
	Reversed bool

	index int
	// loop is the initial number of loops for this sequence to make
	loop int
	// loopRemaining is the remaining number of times to loop through the sequence
	loopRemaining int

	value                 float64
	isActiveTweenFinished bool
	isFinished            bool
}

// NewSequence returns a new Sequence object.
func NewSequence(tweens ...*Tween) *Sequence {
	seq := &Sequence{
		Tweens:        tweens,
		Yoyo:          false,
		Reversed:      false,
		loopRemaining: 1,
		loop:          1,
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
func (s *Sequence) Update(dt float64) {
	if !s.HasTweens() {
		s.value = 0
		s.isActiveTweenFinished = false
		s.isFinished = true
		return
	}
	var completed []int
	remaining := dt

	for {
		if s.Yoyo {
			if s.index < 0 {
				// Out of bounds at beginnning, loop
				s.Reversed = false
				s.index = s.clampIndex(s.index)
				if s.loopRemaining >= 1 {
					s.loopRemaining--
				}
				if s.loopRemaining == 0 || remaining == 0 {
					s.value = s.Tweens[s.index].Begin
					s.isActiveTweenFinished = len(completed) > 0
					s.isFinished = true
					return
				}
				s.Tweens[s.index].Reversed = s.IsReversed()
				s.Tweens[s.index].Reset()
			}
			if s.index >= len(s.Tweens) {
				// Out of bounds at end, yoyo
				s.Reversed = true
				s.index = s.clampIndex(s.index)
				s.Tweens[s.index].Reversed = s.IsReversed()
				s.Tweens[s.index].Reset()
			}
		} else if s.index >= len(s.Tweens) || s.index <= -1 {
			// out of bounds at either end, loop
			if s.loopRemaining >= 1 {
				s.loopRemaining--
			}
			if s.loopRemaining == 0 || remaining == 0 {
				if s.Reversed {
					s.value = s.Tweens[s.clampIndex(s.index)].Begin
					s.isActiveTweenFinished = len(completed) > 0
					s.isFinished = true
					return

				}
				s.value = s.Tweens[s.clampIndex(s.index)].End
				s.isActiveTweenFinished = len(completed) > 0
				s.isFinished = true
				return
			}
			s.index = s.wrapIndex(s.index)
			s.Tweens[s.index].Reversed = s.IsReversed()
			s.Tweens[s.index].Reset()
		}
		s.Tweens[s.index].Update(remaining)
		if !s.Tweens[s.index].IsFinished() {
			s.value = s.Tweens[s.index].Value()
			s.isActiveTweenFinished = len(completed) > 0
			s.isFinished = false
			return
		}
		remaining = s.Tweens[s.index].overflow
		completed = append(completed, s.index)
		if remaining < 0 {
			remaining *= -1
		}
		if s.Reversed {
			s.index--
		} else {
			s.index++
		}
		// On the way back, tweens need to be configured to not go forward
		if s.index < len(s.Tweens) && s.index >= 0 {
			s.Tweens[s.index].Reversed = s.IsReversed()
			s.Tweens[s.index].Reset()
		}
	}
}

// Value returns the current value of the Sequence, which is the value of the currently active Tween.
func (s *Sequence) Value() float64 {
	return s.value
}

// Duration calculates and returns the total duration of the Sequence by summing the durations of all Tweens.
func (s *Sequence) Duration() float64 {
	if s.HasTweens() {
		var total float64
		for _, tween := range s.Tweens {
			total += tween.Duration
		}
		return total
	} else {
		return 0
	}
}

// IsActiveTweenFinished returns whether the currently active Tween is finished.
func (s *Sequence) IsActiveTweenFinished() bool {
	return s.isActiveTweenFinished
}

// IsFinished returns whether the entire Sequence is finished.
// This is true when all Tweens in the Sequence are finished and the Sequence has no remaining loops.
// If the Sequence is set to loop infinitely, this will always return false.
func (s *Sequence) IsFinished() bool {
	return s.isFinished
}

// Index returns the current index of the Sequence. Note that this can exceed the number of Tweens in the Sequence.
func (s *Sequence) Index() int {
	return s.index
}

// SetIndex sets the current index of the Sequence, influencing which Tween is active at any given time.
func (s *Sequence) SetIndex(index int) {
	s.Tweens[s.index].Reversed = s.IsReversed()
	s.Tweens[s.index].Reset()
	s.index = index
}

// SetLoop sets the default loop and the current remaining loops
//
// -1 means infinite loops
//
// 0 means no loops
//
// 1 means one loop
func (s *Sequence) SetLoop(amount int) {
	s.loop = amount
	s.loopRemaining = s.loop
}

// Reset resets the Sequence, resetting all Tweens and setting the Sequence's index back to 0.
func (seq *Sequence) Reset() {
	seq.loopRemaining = seq.loop
	for _, tween := range seq.Tweens {
		tween.Reset()
	}
	seq.index = 0
}

// HasTweens returns whether the Sequence is populated with Tweens or not.
func (s *Sequence) HasTweens() bool {
	return len(s.Tweens) > 0
}

// IsReversed returns whether the Sequence currently running in reverse.
func (s *Sequence) IsReversed() bool {
	return s.Reversed
}

// SetReversed sets whether the Sequence will start running in reverse.
func (s *Sequence) SetReversed(r bool) {
	if s.index >= len(s.Tweens) || s.index < 0 {
		s.index = s.clampIndex(s.index)
	}
	s.Tweens[s.index].Reversed = r
	s.Reversed = r
}

// clampIndex clamps the provided index to the bounds of the Tweens slice
func (s *Sequence) clampIndex(index int) int {
	if index < 0 {
		return 0
	}
	if index >= len(s.Tweens) {
		return len(s.Tweens) - 1
	}
	return index
}

// wrapIndex wraps the provided index when it is out of bounds, otherwise returns index.
func (s *Sequence) wrapIndex(index int) int {
	if index >= len(s.Tweens) {
		return 0
	}
	if index < 0 {
		return len(s.Tweens) - 1
	}
	return index
}
