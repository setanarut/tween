package main

import (
	"fmt"
	"time"

	"github.com/setanarut/tween"
)

func main() {
	var seq = tween.NewSequence(
		tween.NewTween(2, 90, time.Second, "InCubic", false),
		tween.NewTween(90, 2, time.Second, "OutCubic", false),
	)
	dt := 50 * time.Millisecond
	for {
		seq.Update(dt)
		n := int(seq.Value)
		for range n {
			fmt.Print("x")
		}
		for range 90 {
			fmt.Print(" ")
		}
		fmt.Print("\n")
		if seq.IsFinished {
			break
		}
	}
}

// xx
// xx
// xx
// xx
// xxx
// xxxx
// xxxxx
// xxxxxxx
// xxxxxxxxxx
// xxxxxxxxxxxxx
// xxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxx
// xxxxxxxxxxxxx
// xxxxxxxxxx
// xxxxxxx
// xxxxx
// xxxx
// xxx
// xx
// xx
// xx
// xx
// xx
