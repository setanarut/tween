package main

import (
	"fmt"
	"time"

	"github.com/setanarut/tween"
)

func main() {
	// tween from 0 to 1 in 3 seconds
	tw := tween.NewTween(0, 1, 3*time.Second, tween.Linear, false)

	// advance by 1.5 seconds
	tw.Update(time.Millisecond * 1500)

	// get tween value at 1.5 seconds
	fmt.Println(tw.Value) // 0.5

	// merge multiple tweens into a sequence
	sequence := tween.NewSequence(
		tween.NewTween(0, 100, 3*time.Second, tween.InCubic, false),
		tween.NewTween(100, 40, 2*time.Second, tween.OutCubic, false),
		tween.NewTween(4, 100, 20*time.Second, tween.InOutBounce, false),
	)

	// advance by 7.5 seconds
	sequence.Update(time.Millisecond * 7500)

	// get sequence value at 7.5 seconds
	fmt.Println(sequence.Value) // 5.3125
}
