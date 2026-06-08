package tween

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	duration := 10 * time.Second
	tw := NewTween(1.3, 10, duration, "Linear", false)

	if tw.Begin != 1.3 {
		t.Errorf("expected begin to be 0, got %v", tw.Begin)
	}
	if tw.End != 10 {
		t.Errorf("expected end to be 10, got %v", tw.End)
	}
	if tw.Change() != 8.7 {
		t.Errorf("expected change to be 8.7, got %v", tw.Change())
	}
	if tw.Duration != duration {
		t.Errorf("expected duration to be %v, got %v", duration, tw.Duration)
	}
	if tw.Time != 0 {
		t.Errorf("expected Time to be 0, got %v", tw.Time)
	}
	if tw.Overflow != 0 {
		t.Errorf("expected Overflow to be 0, got %v", tw.Overflow)
	}
	if tw.Reversed {
		t.Errorf("expected Reverse to be false, got %v", tw.Reversed)
	}
}

func TestTween_Set(t *testing.T) {
	tw := NewTween(0, 10, 10*time.Second, "Linear", false)

	tw.SetTime(2 * time.Second)
	if tw.Value != 2 {
		t.Errorf("expected Value to be 2, got %v", tw.Value)
	}
	if tw.Overflow != 0 {
		t.Errorf("expected Overflow to be 0, got %v", tw.Overflow)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished to be false")
	}

	tw.SetTime(11 * time.Second)
	if tw.Value != 10 {
		t.Errorf("expected Value to be 10, got %v", tw.Value)
	}
	if tw.Overflow != 1*time.Second {
		t.Errorf("expected Overflow to be 1s, got %v", tw.Overflow)
	}
	if !tw.IsFinished() {
		t.Errorf("expected IsFinished to be true")
	}
}

func TestTween_SetNeg(t *testing.T) {
	tw := NewTween(0, 10, 10*time.Second, "Linear", false)
	tw.SetTime(2 * time.Second)

	tw.SetTime(-1 * time.Second)
	if tw.Value != 0 {
		t.Errorf("expected current to be 0, got %v", tw.Value)
	}
	if tw.Overflow != -1*time.Second {
		t.Errorf("expected overflow to be -1s, got %v", tw.Overflow)
	}
}

func TestTween_SetReverse(t *testing.T) {
	tw := NewTween(0, 10, 10*time.Second, "Linear", false)
	tw.Reversed = true

	tw.SetTime(2 * time.Second)
	if tw.Value != 2 {
		t.Errorf("expected Value to be 2, got %v", tw.Value)
	}
	if tw.IsFinished() {
		t.Errorf("expected IsFinished to be false")
	}

	tw.SetTime(11 * time.Second)
	if tw.IsFinished() {
		t.Errorf("expected IsFinished to be false in reverse at end time")
	}
}

func TestTween_Reset(t *testing.T) {
	tw := NewTween(0, 10, 10*time.Second, "Linear", false)
	tw.SetTime(5 * time.Second)

	tw.Reset()
	if tw.Time != 0 {
		t.Errorf("expected Time to be 0, got %v", tw.Time)
	}

	tw.Reversed = true
	tw.Reset()
	if tw.Time != 10*time.Second {
		t.Errorf("expected Time to be 10s, got %v", tw.Time)
	}
}

func TestTween_Update(t *testing.T) {
	tw := NewTween(0, 10, 10*time.Second, "Linear", false)

	tw.Update(2 * time.Second)
	if tw.Value != 2 {
		t.Errorf("expected Value to be 2, got %v", tw.Value)
	}

	tw.Update(9 * time.Second)
	if !tw.IsFinished() {
		t.Errorf("expected IsFinished to be true after 11s total update")
	}
	if tw.Overflow != 1*time.Second {
		t.Errorf("expected Overflow to be 1s, got %v", tw.Overflow)
	}
}

func TestTween_CanReverse(t *testing.T) {
	tw := NewTween(0, 10, 10*time.Second, "Linear", false)
	tw.Update(8 * time.Second)
	tw.Reversed = true
	tw.Update(2 * time.Second)

	if tw.Value != 6 {
		t.Errorf("expected Value to be 6 after reversing 2s from 8s, got %v", tw.Value)
	}
}

func TestTween_Yoyo(t *testing.T) {
	tw := NewTween(0, 10, 10*time.Second, "Linear", true)

	tw.Update(12 * time.Second)
	if !tw.Reversed {
		t.Errorf("expected Reversed to be true after yoyo trigger")
	}
	if tw.Value != 8 {
		t.Errorf("expected Value to be 8 (10-2), got %v", tw.Value)
	}
}

func TestTween_CanReverseFromStart(t *testing.T) {
	tw := NewTween(0, 10, 10*time.Second, "Linear", false)
	tw.Reversed = true
	tw.Update(1 * time.Second)

	if !tw.IsFinished() {
		t.Errorf("expected IsFinished to be true when reversing at time 0")
	}
	if tw.Overflow != -1*time.Second {
		t.Errorf("expected Overflow to be -1s, got %v", tw.Overflow)
	}
}
func TestTween_Delay(t *testing.T) {
	tw := NewTweenWithDelay(time.Second, 1.3, 10, 1*time.Second, "Linear", false)
	tw.Update(time.Second / 2)

	if tw.Value != 1.3 {
		t.Errorf("Begin value is not 1.3")
	}

	if tw.TotalDuration() != time.Second*2 {
		t.Errorf("total duration is not 2 seconds")
	}

}
