package main

import (
	"github.com/setanarut/tween"
	"github.com/setanarut/tween/ease"
)

func main() {
	var seq = tween.NewSequence(
		tween.NewTween(2, 90, 20, ease.InCubic),
		tween.NewTween(90, 2, 20, ease.OutCubic),
	)
	for range int(seq.Duration()) {
		seq.Update(1)
		if seq.IsFinished() {
			break
		}
		n := int(seq.Value())
		for range n {
			print("x")
		}
		for j := n; j < 8; j++ {
			print(" ")
		}
		print("\n")
	}

	// out

	/*
		xx
		xx
		xx
		xx
		xxx
		xxxx
		xxxxx
		xxxxxxx
		xxxxxxxxxx
		xxxxxxxxxxxxx
		xxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxx
		xxxxxxxxxxxxx
		xxxxxxxxxx
		xxxxxxx
		xxxxx
		xxxx
		xxx
		xx
		xx
		xx
		xx
	*/

}
