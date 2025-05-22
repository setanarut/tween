package main

import (
	"fmt"

	"github.com/setanarut/tween"
	"github.com/setanarut/tween/ease"
)

func main() {
	// tween from 0 to 1 in 3 seconds
	tw := tween.NewTween(0, 1, 3, ease.Linear)

	// advance by 1.5 seconds
	tw.Update(1.5)

	// get tween value at 1.5 seconds using Value()
	fmt.Println(tw.Value()) // 0.5

	// merge multiple tweens into a sequence
	sequence := tween.NewSequence(
		tween.NewTween(0, 100, 3, ease.InCubic),
		tween.NewTween(100, 40, 2, ease.OutCubic),
		tween.NewTween(4, 100, 20, ease.InOutBounce),
	)

	// advance by 7.5 seconds
	sequence.Update(7.5)

	// get sequence value at 7.5 seconds using Value()
	fmt.Println(sequence.Value()) // 5.3125
}
