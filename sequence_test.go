package tween

import (
	"testing"

	"github.com/setanarut/tween/ease"
)

func TestSequenceNew(t *testing.T) {
	s := NewSequence(NewTween(0, 1, 1, ease.Linear))
	s.Update(0.0)
	if s.Value() != 0.0 {
		t.Errorf("expected current to be %v, got %v", 0.0, s.Value())
	}
	if s.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be false, got %v", s.IsActiveTweenFinished())
	}
	if s.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", s.IsFinished())
	}
	if s.Index() != 0 {
		t.Errorf("expected index to be %v, got %v", 0, s.Index())
	}
}

func TestSequence_Update(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
	)

	seq.Update(0.5)
	if seq.Value() != float64(0.5) {
		t.Errorf("expected current to be %v, got %v", float64(0.5), seq.Value())
	}
	if seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be false, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.Index() != 0 {
		t.Errorf("expected index to be %v, got %v", 0, seq.Index())
	}
}

func TestSequence_Reset(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
	)

	seq.Update(1.5)
	seq.Reset()
	if seq.Index() != 0 {
		t.Errorf("expected index to be %v, got %v", 0, seq.Index())
	}
	if seq.Tweens[0].time != float64(0.0) {
		t.Errorf("expected Tweens[0].time to be %v, got %v", float64(0.0), seq.Tweens[0].time)
	}
	if seq.Tweens[0].overflow != float64(0.0) {
		t.Errorf("expected Tweens[0].overflow to be %v, got %v", float64(0.0), seq.Tweens[0].overflow)
	}
	if seq.Tweens[1].time != float64(0.0) {
		t.Errorf("expected Tweens[1].time to be %v, got %v", float64(0.0), seq.Tweens[1].time)
	}
	if seq.Tweens[1].overflow != float64(0.0) {
		t.Errorf("expected Tweens[1].overflow to be %v, got %v", float64(0.0), seq.Tweens[1].overflow)
	}
}

func TestSequence_CompleteFirst(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
	)

	seq.Update(1.0)
	if seq.Value() != float64(1.0) {
		t.Errorf("expected current to be %v, got %v", float64(1.0), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.index != 1 {
		t.Errorf("expected index to be %v, got %v", 1, seq.index)
	}
}

func TestSequence_OverflowSecond(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
	)

	seq.Update(1.5)
	if seq.Value() != float64(1.5) {
		t.Errorf("expected current to be %v, got %v", float64(1.5), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.index != 1 {
		t.Errorf("expected index to be %v, got %v", 1, seq.index)
	}
}

func TestSequence_OverflowAndComplete(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
		NewTween(2, 3, 1, ease.Linear),
	)

	seq.Update(3.5)
	if seq.Value() != float64(3.0) {
		t.Errorf("expected current to be %v, got %v", float64(3.0), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if !seq.IsFinished() {
		t.Errorf("expected seqFinished to be true, got %v", seq.IsFinished())
	}
	if seq.index != 3 {
		t.Errorf("expected index to be %v, got %v", 3, seq.index)
	}
}

func TestSequence_Loops(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
		NewTween(2, 3, 1, ease.Linear),
	)
	seq.SetLoop(2)
	seq.Update(5.25)
	if seq.Value() != float64(2.25) {
		t.Errorf("expected current to be %v, got %v", float64(2.25), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 1 {
		t.Errorf("expected loopRemaining to be %v, got %v", 1, seq.loopRemaining)
	}
	if seq.index != 2 {
		t.Errorf("expected index to be %v, got %v", 2, seq.index)
	}

	seq.Update(0.75)
	if seq.Value() != float64(3) {
		t.Errorf("expected current to be %v, got %v", float64(3), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if !seq.IsFinished() {
		t.Errorf("expected seqFinished to be true, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 0 {
		t.Errorf("expected loopRemaining to be %v, got %v", 0, seq.loopRemaining)
	}
	if seq.index != 3 {
		t.Errorf("expected index to be %v, got %v", 3, seq.index)
	}
}

func TestSequence_LoopsForever(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
		NewTween(2, 3, 1, ease.Linear),
	)
	seq.SetLoop(-1)
	seq.Update(3*1_000_000 + 2.25)
	if seq.Value() != float64(2.25) {
		t.Errorf("expected current to be %v, got %v", float64(2.25), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != -1 {
		t.Errorf("expected loopRemaining to be %v, got %v", -1, seq.loopRemaining)
	}
	if seq.index != 2 {
		t.Errorf("expected index to be %v, got %v", 2, seq.index)
	}
}

func TestSequence_Yoyos(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
		NewTween(2, 3, 1, ease.Linear),
	)

	seq.YoyoEnabled = true
	seq.Update(5.75)
	if seq.Value() != float64(0.25) {
		t.Errorf("expected current to be %v, got %v", float64(0.25), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 1 {
		t.Errorf("expected loopRemaining to be %v, got %v", 1, seq.loopRemaining)
	}
	if seq.index != 0 {
		t.Errorf("expected index to be %v, got %v", 0, seq.index)
	}

	seq.Update(0.25)
	if seq.Value() != float64(0) {
		t.Errorf("expected current to be %v, got %v", float64(0), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if !seq.IsFinished() {
		t.Errorf("expected seqFinished to be true, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 0 {
		t.Errorf("expected loopRemaining to be %v, got %v", 0, seq.loopRemaining)
	}
	if seq.index != 0 {
		t.Errorf("expected index to be %v, got %v", 0, seq.index)
	}
}

func TestSequence_YoyosAndLoops(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
		NewTween(2, 3, 1, ease.Linear),
	)
	seq.YoyoEnabled = true
	seq.SetLoop(2)
	seq.Update(7.25)
	if seq.Value() != float64(1.25) {
		t.Errorf("expected current to be %v, got %v", float64(1.25), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 1 {
		t.Errorf("expected loopRemaining to be %v, got %v", 1, seq.loopRemaining)
	}
	if seq.index != 1 {
		t.Errorf("expected index to be %v, got %v", 1, seq.index)
	}

	seq.Update(4.75)
	if seq.Value() != float64(0) {
		t.Errorf("expected current to be %v, got %v", float64(0), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if !seq.IsFinished() {
		t.Errorf("expected seqFinished to be true, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 0 {
		t.Errorf("expected loopRemaining to be %v, got %v", 0, seq.loopRemaining)
	}
	if seq.index != 0 {
		t.Errorf("expected index to be %v, got %v", 0, seq.index)
	}
}

func TestSequence_SetReverse(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
		NewTween(2, 3, 1, ease.Linear),
	)
	seq.SetLoop(2)

	// Normal operation
	seq.Update(2.25)
	if seq.Value() != float64(2.25) {
		t.Errorf("expected current to be %v, got %v", float64(2.25), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 2 {
		t.Errorf("expected loopRemaining to be %v, got %v", 2, seq.loopRemaining)
	}
	if seq.index != 2 {
		t.Errorf("expected index to be %v, got %v", 2, seq.index)
	}

	seq.SetReverse(true)

	// Goes in reverse
	seq.Update(2.0)
	if seq.Value() != float64(0.25) {
		t.Errorf("expected current to be %v, got %v", float64(0.25), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 2 {
		t.Errorf("expected loopRemaining to be %v, got %v", 2, seq.loopRemaining)
	}
	if seq.index != 0 {
		t.Errorf("expected index to be %v, got %v", 0, seq.index)
	}
	if !seq.IsReversed() {
		t.Errorf("expected Reverse() to be true, got %v", seq.IsReversed())
	}

	// Consumes a loop at the start!, resets to the end, continues in reverse
	seq.Update(2.0)
	if seq.Value() != float64(1.25) {
		t.Errorf("expected current to be %v, got %v", float64(1.25), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 1 {
		t.Errorf("expected loopRemaining to be %v, got %v", 1, seq.loopRemaining)
	}
	if seq.index != 1 {
		t.Errorf("expected index to be %v, got %v", 1, seq.index)
	}
	if !seq.IsReversed() {
		t.Errorf("expected Reverse() to be true, got %v", seq.IsReversed())
	}

	// Hits the beginning, no more loops, ends
	seq.Update(2.0)
	if seq.Value() != float64(0.0) {
		t.Errorf("expected current to be %v, got %v", float64(0.0), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if !seq.IsFinished() {
		t.Errorf("expected seqFinished to be true, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 0 {
		t.Errorf("expected loopRemaining to be %v, got %v", 0, seq.loopRemaining)
	}
	if seq.index != -1 {
		t.Errorf("expected index to be %v, got %v", -1, seq.index)
	}
	if !seq.IsReversed() {
		t.Errorf("expected Reverse() to be true, got %v", seq.IsReversed())
	}
}

func TestSequence_SetReverseWithYoyo(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
		NewTween(2, 3, 1, ease.Linear),
	)
	seq.YoyoEnabled = true
	seq.SetLoop(2)

	// Standard operation
	seq.Update(2.25)
	if seq.Value() != float64(2.25) {
		t.Errorf("expected current to be %v, got %v", float64(2.25), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 2 {
		t.Errorf("expected loopRemaining to be %v, got %v", 2, seq.loopRemaining)
	}
	if seq.index != 2 {
		t.Errorf("expected index to be %v, got %v", 2, seq.index)
	}

	seq.SetReverse(true)

	// Goes in reverse
	seq.Update(2.0)
	if seq.Value() != float64(0.25) {
		t.Errorf("expected current to be %v, got %v", float64(0.25), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 2 {
		t.Errorf("expected loopRemaining to be %v, got %v", 2, seq.loopRemaining)
	}
	if seq.index != 0 {
		t.Errorf("expected index to be %v, got %v", 0, seq.index)
	}

	// Consumes a loop at the start, despite not reaching the end yet, and continues
	seq.Update(2.0)
	if seq.Value() != float64(1.75) {
		t.Errorf("expected current to be %v, got %v", float64(1.75), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 1 {
		t.Errorf("expected loopRemaining to be %v, got %v", 1, seq.loopRemaining)
	}
	if seq.index != 1 {
		t.Errorf("expected index to be %v, got %v", 1, seq.index)
	}

	// Hits the end, yoyos
	seq.Update(2.0)
	if seq.Value() != float64(2.25) {
		t.Errorf("expected current to be %v, got %v", float64(2.25), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 1 {
		t.Errorf("expected loopRemaining to be %v, got %v", 1, seq.loopRemaining)
	}
	if seq.index != 2 {
		t.Errorf("expected index to be %v, got %v", 2, seq.index)
	}
	if !seq.IsReversed() {
		t.Errorf("expected Reverse() to be true, got %v", seq.IsReversed())
	}

	seq.SetReverse(false) // Go forward instead

	// Hits the end again, yoyos the same
	seq.Update(1.5)
	if seq.Value() != float64(2.25) {
		t.Errorf("expected current to be %v, got %v", float64(2.25), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 1 {
		t.Errorf("expected loopRemaining to be %v, got %v", 1, seq.loopRemaining)
	}
	if seq.index != 2 {
		t.Errorf("expected index to be %v, got %v", 2, seq.index)
	}

	// Consumes a loop at the start like normal, no more loops, end
	seq.Update(2.5)
	if seq.Value() != float64(0.0) {
		t.Errorf("expected current to be %v, got %v", float64(0.0), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if !seq.IsFinished() {
		t.Errorf("expected seqFinished to be true, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 0 {
		t.Errorf("expected loopRemaining to be %v, got %v", 0, seq.loopRemaining)
	}
	if seq.index != 0 {
		t.Errorf("expected index to be %v, got %v", 0, seq.index)
	}
}

func TestSequence_SetReverseAfterComplete(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
		NewTween(2, 3, 1, ease.Linear),
	)
	seq.SetLoop(1)

	// Normal operation
	seq.Update(3.0)
	if seq.Value() != float64(3.0) {
		t.Errorf("expected current to be %v, got %v", float64(3.0), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if !seq.IsFinished() {
		t.Errorf("expected seqFinished to be true, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 0 {
		t.Errorf("expected loopRemaining to be %v, got %v", 0, seq.loopRemaining)
	}
	if seq.index != 3 {
		t.Errorf("expected index to be %v, got %v", 3, seq.index)
	}

	seq.SetReverse(true)
	seq.SetLoop(1)

	// Goes in reverse
	seq.Update(2.0)
	if seq.Value() != float64(1.0) {
		t.Errorf("expected current to be %v, got %v", float64(1.0), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.loopRemaining != 1 {
		t.Errorf("expected loopRemaining to be %v, got %v", 1, seq.loopRemaining)
	}
	if seq.index != 0 {
		t.Errorf("expected index to be %v, got %v", 0, seq.index)
	}
	if !seq.IsReversed() {
		t.Errorf("expected Reverse() to be true, got %v", seq.IsReversed())
	}
}

func TestSequence_Remove(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
		NewTween(2, 3, 1, ease.Linear),
		NewTween(3, 4, 1, ease.Linear),
		NewTween(4, 5, 1, ease.Linear),
	)
	if len(seq.Tweens) != 5 {
		t.Errorf("expected 5 tweens, got %v", len(seq.Tweens))
	}
	seq.Remove(2)
	if len(seq.Tweens) != 4 {
		t.Errorf("expected 4 tweens, got %v", len(seq.Tweens))
	}
	seq.Update(2.5)
	if seq.Value() != float64(3.5) {
		t.Errorf("expected current to be %v, got %v", float64(3.5), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}
	if seq.index != 2 {
		t.Errorf("expected index to be %v, got %v", 2, seq.index)
	}
	seq.Remove(0)
	if len(seq.Tweens) != 3 {
		t.Errorf("expected 3 tweens, got %v", len(seq.Tweens))
	}
	seq.Remove(0)
	if len(seq.Tweens) != 2 {
		t.Errorf("expected 2 tweens, got %v", len(seq.Tweens))
	}
	seq.Remove(0)
	if len(seq.Tweens) != 1 {
		t.Errorf("expected 1 tween, got %v", len(seq.Tweens))
	}
	// Out of bound checking
	seq.Remove(0)
	if len(seq.Tweens) != 0 {
		t.Errorf("expected 0 tweens, got %v", len(seq.Tweens))
	}
	seq.Remove(2)
	if len(seq.Tweens) != 0 {
		t.Errorf("expected 0 tweens, got %v", len(seq.Tweens))
	}
}

func TestSequence_Has(t *testing.T) {
	seq := NewSequence()
	if seq.HasTweens() {
		t.Errorf("expected HasTweens() to be false, got true")
	}
	seq.Add(NewTween(0, 5, 1, ease.Linear))
	if !seq.HasTweens() {
		t.Errorf("expected HasTweens() to be true, got false")
	}
	seq.Remove(0)
	if seq.HasTweens() {
		t.Errorf("expected HasTweens() to be false, got true")
	}
	seq.Update(1)
	if seq.Value() != float64(0) {
		t.Errorf("expected current to be %v, got %v", float64(0), seq.Value())
	}
	if seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be false, got %v", seq.IsActiveTweenFinished())
	}
	if !seq.IsFinished() {
		t.Errorf("expected seqFinished to be true, got %v", seq.IsFinished())
	}
}

func TestSequence_SetIndex(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 1, 1, ease.Linear),
		NewTween(1, 2, 1, ease.Linear),
	)
	seq.SetIndex(1)
	seq.Update(1.5)
	if seq.Value() != float64(2) {
		t.Errorf("expected current to be %v, got %v", float64(2), seq.Value())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if !seq.IsFinished() {
		t.Errorf("expected seqFinished to be true, got %v", seq.IsFinished())
	}
	if seq.index != 2 {
		t.Errorf("expected index to be %v, got %v", 2, seq.index)
	}
}

func TestSequence_RealWorld(t *testing.T) {
	seq := NewSequence(
		NewTween(0, 5, 1, ease.Linear),
		NewTween(5, 0, 1, ease.Linear),
		NewTween(0, 2, 2, ease.Linear),
		NewTween(2, 0, 2, ease.Linear),
		NewTween(0, 1, 100, ease.Linear),
	)

	if len(seq.Tweens) != 5 {
		t.Errorf("expected 5 tweens, got %v", len(seq.Tweens))
	}
	seq.Remove(0)
	seq.Remove(0)
	if len(seq.Tweens) != 3 {
		t.Errorf("expected 3 tweens, got %v", len(seq.Tweens))
	}

	seq.Update(1)
	if seq.Value() != float64(1) {
		t.Errorf("expected current to be %v, got %v", float64(1), seq.Value())
	}
	if seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be false, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}

	seq.Update(1)
	if seq.Value() != float64(2) {
		t.Errorf("expected current to be %v, got %v", float64(2), seq.Value())
	}
	if seq.Index() != 1 {
		t.Errorf("expected index to be %v, got %v", 1, seq.Index())
	}
	if !seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be true, got %v", seq.IsActiveTweenFinished())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}

	seq.Update(2)
	if seq.Index() != 2 {
		t.Errorf("expected index to be %v, got %v", 2, seq.Index())
	}
	if seq.IsFinished() {
		t.Errorf("expected seqFinished to be false, got %v", seq.IsFinished())
	}

	seq.Remove(2)
	seq.Update(1)
	if seq.IsActiveTweenFinished() {
		t.Errorf("expected finishedTween to be false, got %v", seq.IsActiveTweenFinished())
	}
	if !seq.IsFinished() {
		t.Errorf("expected seqFinished to be true, got %v", seq.IsFinished())
	}
}
