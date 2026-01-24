package tween

import (
	"testing"
	"time"
)

func TestSequence_BasicFlow(t *testing.T) {
	s := NewSequence(
		NewTween(0, 10, 1*time.Second, "Linear", false),
		NewTween(10, 20, 1*time.Second, "Linear", false),
	)

	s.Update(1500 * time.Millisecond)

	if s.Index != 1 {
		t.Errorf("expected index 1, got %v", s.Index)
	}
	if s.Value != 15 {
		t.Errorf("expected value 15, got %v", s.Value)
	}
}

func TestSequence_Loop(t *testing.T) {
	s := NewSequence(NewTween(0, 10, 1*time.Second, "Linear", false))
	s.SetLoop(2)

	s.Update(1500 * time.Millisecond)
	if s.IsFinished {
		t.Errorf("sequence should not be finished yet")
	}

	s.Update(1 * time.Second)
	if !s.IsFinished {
		t.Errorf("expected sequence to be finished after 2 loops")
	}
}

func TestSequence_Reverse(t *testing.T) {
	s := NewSequence(
		NewTween(0, 10, 1*time.Second, "Linear", false),
		NewTween(10, 20, 1*time.Second, "Linear", false),
	)
	s.SetReversed(true)
	s.Index = 1
	s.Tweens[1].Reset()

	s.Update(500 * time.Millisecond)
	if s.Value != 15 {
		t.Errorf("expected value 15 in reverse, got %v", s.Value)
	}
}

func TestSequence_Yoyo(t *testing.T) {
	s := NewSequence(NewTween(0, 10, 1*time.Second, "Linear", false))
	s.Yoyo = true

	s.Update(1500 * time.Millisecond)
	if !s.IsReversed {
		t.Errorf("expected sequence to be reversed due to yoyo")
	}
	if s.Value != 5 {
		t.Errorf("expected value 5 on yoyo return, got %v", s.Value)
	}
}

func TestSequence_LargeOverflow(t *testing.T) {
	s := NewSequence(
		NewTween(0, 10, 1*time.Second, "Linear", false),
		NewTween(10, 20, 1*time.Second, "Linear", false),
		NewTween(20, 30, 1*time.Second, "Linear", false),
	)

	s.Update(2500 * time.Millisecond)
	if s.Index != 2 {
		t.Errorf("expected index 2, got %v", s.Index)
	}
	if s.Value != 25 {
		t.Errorf("expected value 25, got %v", s.Value)
	}
}
