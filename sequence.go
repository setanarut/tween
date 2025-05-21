package tween

import "slices"

// Sequence represents a sequence of Tweens, executed one after the other.
type Sequence struct {
	Tweens []*Tween
	// YoyoEnabled makes the sequence "Yoyo" back to the beginning after it reaches the end
	YoyoEnabled bool

	// isReversed runs the sequence backwards when true
	isReversed bool

	index int
	// loop is the initial number of loops for this sequence to make
	loop int
	// loopRemaining is the remaining number of times to loop through the sequence
	loopRemaining int

	value                 float32
	isActiveTweenFinished bool
	isFinished            bool
}

// NewSequence returns a new Sequence object.
func NewSequence(tweens ...*Tween) *Sequence {
	seq := &Sequence{
		Tweens:        tweens,
		YoyoEnabled:   false,
		isReversed:    false,
		loopRemaining: 1,
		loop:          1,
	}
	return seq
}

// Add adds one or more Tweens in order to the Sequence.
func (seq *Sequence) Add(tweens ...*Tween) {
	seq.Tweens = append(seq.Tweens, tweens...)
}

// Remove removes a Tween of the specified index from the Sequence.
func (seq *Sequence) Remove(index int) {
	if index >= 0 && index < len(seq.Tweens) {
		seq.Tweens = slices.Delete(seq.Tweens, index, index+1)
	}
}

// Update updates the currently active Tween in the Sequence; once that Tween is done, the Sequence moves onto the next one.
func (s *Sequence) Update(dt float32) {
	if !s.HasTweens() {
		s.value = 0
		s.isActiveTweenFinished = false
		s.isFinished = true
		return
	}
	var completed []int
	remaining := dt

	for {
		if s.YoyoEnabled {
			if s.index < 0 {
				// Out of bounds at beginnning, loop
				s.isReversed = false
				s.index = s.clampIndex(s.index)
				if s.loopRemaining >= 1 {
					s.loopRemaining--
				}
				if s.loopRemaining == 0 || remaining == 0 {
					s.value = s.Tweens[s.index].begin
					s.isActiveTweenFinished = len(completed) > 0
					s.isFinished = true
					return
				}
				s.Tweens[s.index].Reverse = s.IsReversed()
				s.Tweens[s.index].Reset()
			}
			if s.index >= len(s.Tweens) {
				// Out of bounds at end, yoyo
				s.isReversed = true
				s.index = s.clampIndex(s.index)
				s.Tweens[s.index].Reverse = s.IsReversed()
				s.Tweens[s.index].Reset()
			}
		} else if s.index >= len(s.Tweens) || s.index <= -1 {
			// out of bounds at either end, loop
			if s.loopRemaining >= 1 {
				s.loopRemaining--
			}
			if s.loopRemaining == 0 || remaining == 0 {
				if s.isReversed {
					s.value = s.Tweens[s.clampIndex(s.index)].begin
					s.isActiveTweenFinished = len(completed) > 0
					s.isFinished = true
					return

				}
				s.value = s.Tweens[s.clampIndex(s.index)].end
				s.isActiveTweenFinished = len(completed) > 0
				s.isFinished = true
				return
			}
			s.index = s.wrapIndex(s.index)
			s.Tweens[s.index].Reverse = s.IsReversed()
			s.Tweens[s.index].Reset()
		}
		s.Tweens[s.index].Update(remaining)
		if !s.Tweens[s.index].IsFinished() {
			s.value = s.Tweens[s.index].Current()
			s.isActiveTweenFinished = len(completed) > 0
			s.isFinished = false
			return
		}
		remaining = s.Tweens[s.index].overflow
		completed = append(completed, s.index)
		if remaining < 0 {
			remaining *= -1
		}
		if s.isReversed {
			s.index--
		} else {
			s.index++
		}
		// On the way back, tweens need to be configured to not go forward
		if s.index < len(s.Tweens) && s.index >= 0 {
			s.Tweens[s.index].Reverse = s.IsReversed()
			s.Tweens[s.index].Reset()
		}
	}
}

// Value returns the current value of the Sequence, which is the value of the currently active Tween.
func (seq *Sequence) Value() float32 {
	return seq.value
}

// IsActiveTweenFinished returns whether the currently active Tween is finished.
func (seq *Sequence) IsActiveTweenFinished() bool {
	return seq.isActiveTweenFinished
}

// IsFinished returns whether the entire Sequence is finished.
// This is true when all Tweens in the Sequence are finished and the Sequence has no remaining loops.
// If the Sequence is set to loop infinitely, this will always return false.
func (seq *Sequence) IsFinished() bool {
	return seq.isFinished
}

// Index returns the current index of the Sequence. Note that this can exceed the number of Tweens in the Sequence.
func (seq *Sequence) Index() int {
	return seq.index
}

// SetIndex sets the current index of the Sequence, influencing which Tween is active at any given time.
func (seq *Sequence) SetIndex(index int) {
	seq.Tweens[seq.index].Reverse = seq.IsReversed()
	seq.Tweens[seq.index].Reset()
	seq.index = index
}

// SetLoop sets the default loop and the current remaining loops
func (seq *Sequence) SetLoop(amount int) {
	seq.loop = amount
	seq.loopRemaining = seq.loop
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
func (seq *Sequence) HasTweens() bool {
	return len(seq.Tweens) > 0
}

// IsReversed returns whether the Sequence currently running in reverse.
func (seq *Sequence) IsReversed() bool {
	return seq.isReversed
}

// SetReverse sets whether the Sequence will start running in reverse.
func (seq *Sequence) SetReverse(r bool) {
	if seq.index >= len(seq.Tweens) || seq.index < 0 {
		seq.index = seq.clampIndex(seq.index)
	}
	seq.Tweens[seq.index].Reverse = r
	seq.isReversed = r
}

// clampIndex clamps the provided index to the bounds of the Tweens slice
func (seq *Sequence) clampIndex(index int) int {
	if index < 0 {
		return 0
	}
	if index >= len(seq.Tweens) {
		return len(seq.Tweens) - 1
	}
	return index
}

// wrapIndex wraps the provided index when it is out of bounds, otherwise returns index.
func (seq *Sequence) wrapIndex(index int) int {
	if index >= len(seq.Tweens) {
		return 0
	}
	if index < 0 {
		return len(seq.Tweens) - 1
	}
	return index
}
