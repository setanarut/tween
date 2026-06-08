package main

import (
	"fmt"
	"time"

	"github.com/setanarut/tween"
)

func main() {

	seq := tween.NewSequence(
		tween.NewTween(
			90,
			2,
			time.Second*3,
			0,
			tween.OutCubic,
			false,
		),
		tween.NewTween(
			2,
			90,
			time.Second*3,
			time.Second, // inital delay
			tween.InCubic,
			false,
		),
	)

	lines := 20
	for range lines {
		seq.Update(seq.TotalDuration() / time.Duration(lines))
		for range int(seq.Value) {
			fmt.Print("+")
		}
		for range 90 {
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}

	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// +++++++++++++++++++++++++++++++++++++++++
	// ++++++++++++++++++++++++++
	// +++++++++++++++
	// ++++++++
	// ++++
	// ++
	// ++
	// ++
	// ++
	// ++
	// ++
	// ++
	// ++++
	// ++++++++
	// +++++++++++++++
	// ++++++++++++++++++++++++++
	// +++++++++++++++++++++++++++++++++++++++++
	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
}
