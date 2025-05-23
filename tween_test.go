package tween

import (
	"testing"

	"github.com/setanarut/tween/ease"
)

func TestNew(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)

	if tw.Begin != 0 {
		t.Errorf("expected begin to be %v, got %v", 0, tw.Begin)
	}
	if tw.End != 10 {
		t.Errorf("expected end to be %v, got %v", 10, tw.End)
	}
	if tw.change() != 10 {
		t.Errorf("expected change to be %v, got %v", 10, tw.change())
	}
	if tw.Duration != 10 {
		t.Errorf("expected duration to be %v, got %v", 10, tw.Duration)
	}
	if tw.time != 0 {
		t.Errorf("expected time to be %v, got %v", 0, tw.time)
	}
	if tw.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tw.overflow)
	}
	if tw.Reversed {
		t.Errorf("expected Reverse to be false, got %v", tw.Reversed)
	}
}

func TestTween_Set(t *testing.T) {
	tween := NewTween(0, 10, 10, ease.Linear)
	tween.SetTime(2)
	if tween.Value() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tween.Value())
	}
	if tween.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tween.overflow)
	}
	if tween.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
	tween.SetTime(11)
	if tween.Value() != 10 {
		t.Errorf("expected Current() to be %v, got %v", 10, tween.Value())
	}
	if tween.overflow != 1 {
		t.Errorf("expected overflow to be %v, got %v", 1, tween.overflow)
	}
	if !tween.IsFinished() {
		t.Errorf("expected IsFinished() to be true")
	}
}

func TestTween_SetNeg(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.SetTime(2)
	if tw.value != 2 {
		t.Errorf("expected current to be %v, got %v", 2, tw.value)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
	tw.SetTime(-1)
	if tw.value != 0 {
		t.Errorf("expected current to be %v, got %v", 0, tw.value)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_SetReverse(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Reversed = true
	tw.SetTime(2)
	if tw.Value() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tw.Value())
	}
	if tw.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tw.overflow)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
	tw.SetTime(11)
	if tw.Value() != 10 {
		t.Errorf("expected Current() to be %v, got %v", 10, tw.Value())
	}
	if tw.overflow != 1 {
		t.Errorf("expected overflow to be %v, got %v", 1, tw.overflow)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_SetNegReverse(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Reversed = true
	tw.SetTime(2)
	if tw.Value() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tw.Value())
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
	tw.SetTime(-1)
	if tw.Value() != 0 {
		t.Errorf("expected Current() to be %v, got %v", 0, tw.Value())
	}
	if !tw.IsFinished() {
		t.Errorf("expected IsFinished() to be true")
	}
}

func TestTween_Reset(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.SetTime(2)
	if tw.Value() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tw.Value())
	}
	if tw.time != 2 {
		t.Errorf("expected time to be %v, got %v", 2, tw.time)
	}
	if tw.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tw.overflow)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
	tw.Reset()
	if tw.time != 0 {
		t.Errorf("expected time to be %v, got %v", 0, tw.time)
	}
	if tw.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tw.overflow)
	}
}

func TestTween_ResetReverse(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.SetTime(2)
	tw.Reversed = true
	tw.Reset()
	if tw.time != 10 {
		t.Errorf("expected time to be %v, got %v", 10, tw.time)
	}
	if tw.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tw.overflow)
	}
}

func TestTween_Update(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Update(2)
	if tw.Value() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tw.Value())
	}
	if tw.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tw.overflow)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
	tw.Update(9)
	if tw.Value() != 10 {
		t.Errorf("expected Current() to be %v, got %v", 10, tw.Value())
	}
	if tw.overflow != 1 {
		t.Errorf("expected overflow to be %v, got %v", 1, tw.overflow)
	}
	if !tw.IsFinished() {
		t.Errorf("expected IsFinished() to be true")
	}
}

func TestTween_UpdateZero(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Update(2)
	tw.Update(0)
	if tw.Value() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tw.Value())
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_UpdateNeg(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Update(2)
	tw.Update(-1)
	if tw.Value() != 1 {
		t.Errorf("expected Current() to be %v, got %v", 1, tw.Value())
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_UpdateNegReverse(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Update(2)
	tw.Reversed = true
	tw.Update(-1)
	if tw.Value() != 3 {
		t.Errorf("expected Current() to be %v, got %v", 3, tw.Value())
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_Defaults_Forward(t *testing.T) {
	tween := NewTween(0, 10, 10, ease.Linear)
	if tween.Reversed {
		t.Errorf("expected Reverse to be false, got %v", tween.Reversed)
	}
}

func TestTween_CanReverse(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Update(8)
	tw.Reversed = true
	tw.Update(2)
	if tw.Value() != 6 {
		t.Errorf("expected Current() to be %v, got %v", 6, tw.Value())
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_CanReverseFromFinished(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Update(10)
	if !tw.IsFinished() {
		t.Errorf("expected IsFinished() to be true")
	}
	tw.Reversed = true
	tw.Update(2)
	if tw.Value() != 8 {
		t.Errorf("expected Current() to be %v, got %v", 8, tw.Value())
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_CanReverseFromStart(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Reversed = true
	tw.Update(0)
	if !tw.IsFinished() {
		t.Errorf("expected IsFinished() to be true")
	}
	if tw.Value() != 0 {
		t.Errorf("expected Current() to be %v, got %v", 0, tw.Value())
	}
	if tw.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tw.overflow)
	}
	tw.Update(1)
	if !tw.IsFinished() {
		t.Errorf("expected IsFinished() to be true")
	}
	if tw.Value() != 0 {
		t.Errorf("expected Current() to be %v, got %v", 0, tw.Value())
	}
	if tw.overflow != -1.0 {
		t.Errorf("expected overflow to be %v, got %v", -1.0, tw.overflow)
	}
}
