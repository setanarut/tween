package tween

import (
	"testing"

	"github.com/setanarut/tween/ease"
)

func TestNew(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)

	if tw.begin != 0 {
		t.Errorf("expected begin to be %v, got %v", 0, tw.begin)
	}
	if tw.end != 10 {
		t.Errorf("expected end to be %v, got %v", 10, tw.end)
	}
	if tw.change != 10 {
		t.Errorf("expected change to be %v, got %v", 10, tw.change)
	}
	if tw.duration != 10 {
		t.Errorf("expected duration to be %v, got %v", 10, tw.duration)
	}
	if tw.time != 0 {
		t.Errorf("expected time to be %v, got %v", 0, tw.time)
	}
	if tw.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tw.overflow)
	}
	if tw.Reverse {
		t.Errorf("expected Reverse to be false, got %v", tw.Reverse)
	}
}

func TestTween_Set(t *testing.T) {
	tween := NewTween(0, 10, 10, ease.Linear)
	tween.Set(2)
	if tween.Current() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tween.Current())
	}
	if tween.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tween.overflow)
	}
	if tween.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
	tween.Set(11)
	if tween.Current() != 10 {
		t.Errorf("expected Current() to be %v, got %v", 10, tween.Current())
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
	tw.Set(2)
	if tw.current != 2 {
		t.Errorf("expected current to be %v, got %v", 2, tw.current)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
	tw.Set(-1)
	if tw.current != 0 {
		t.Errorf("expected current to be %v, got %v", 0, tw.current)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_SetReverse(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Reverse = true
	tw.Set(2)
	if tw.Current() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tw.Current())
	}
	if tw.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tw.overflow)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
	tw.Set(11)
	if tw.Current() != 10 {
		t.Errorf("expected Current() to be %v, got %v", 10, tw.Current())
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
	tw.Reverse = true
	tw.Set(2)
	if tw.Current() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tw.Current())
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
	tw.Set(-1)
	if tw.Current() != 0 {
		t.Errorf("expected Current() to be %v, got %v", 0, tw.Current())
	}
	if !tw.IsFinished() {
		t.Errorf("expected IsFinished() to be true")
	}
}

func TestTween_Reset(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Set(2)
	if tw.Current() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tw.Current())
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
	tw.Set(2)
	tw.Reverse = true
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
	if tw.Current() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tw.Current())
	}
	if tw.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tw.overflow)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
	tw.Update(9)
	if tw.Current() != 10 {
		t.Errorf("expected Current() to be %v, got %v", 10, tw.Current())
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
	if tw.Current() != 2 {
		t.Errorf("expected Current() to be %v, got %v", 2, tw.Current())
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_UpdateNeg(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Update(2)
	tw.Update(-1)
	if tw.Current() != 1 {
		t.Errorf("expected Current() to be %v, got %v", 1, tw.Current())
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_UpdateNegReverse(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Update(2)
	tw.Reverse = true
	tw.Update(-1)
	if tw.Current() != 3 {
		t.Errorf("expected Current() to be %v, got %v", 3, tw.Current())
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_Defaults_Forward(t *testing.T) {
	tween := NewTween(0, 10, 10, ease.Linear)
	if tween.Reverse {
		t.Errorf("expected Reverse to be false, got %v", tween.Reverse)
	}
}

func TestTween_CanReverse(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Update(8)
	tw.Reverse = true
	tw.Update(2)
	if tw.Current() != 6 {
		t.Errorf("expected Current() to be %v, got %v", 6, tw.Current())
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
	tw.Reverse = true
	tw.Update(2)
	if tw.Current() != 8 {
		t.Errorf("expected Current() to be %v, got %v", 8, tw.Current())
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished() to be false")
	}
}

func TestTween_CanReverseFromStart(t *testing.T) {
	tw := NewTween(0, 10, 10, ease.Linear)
	tw.Reverse = true
	tw.Update(0)
	if !tw.IsFinished() {
		t.Errorf("expected IsFinished() to be true")
	}
	if tw.Current() != 0 {
		t.Errorf("expected Current() to be %v, got %v", 0, tw.Current())
	}
	if tw.overflow != 0 {
		t.Errorf("expected overflow to be %v, got %v", 0, tw.overflow)
	}
	tw.Update(1)
	if !tw.IsFinished() {
		t.Errorf("expected IsFinished() to be true")
	}
	if tw.Current() != 0 {
		t.Errorf("expected Current() to be %v, got %v", 0, tw.Current())
	}
	if tw.overflow != -1.0 {
		t.Errorf("expected overflow to be %v, got %v", -1.0, tw.overflow)
	}
}
