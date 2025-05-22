package main

import (
	"fmt"

	"github.com/setanarut/tween"
	"github.com/setanarut/tween/ease"
)

func main() {
	fmt.Println(1. / 25)
	fmt.Println(1. / 60.)
	fmt.Println(1. / 120.)
	var s = tween.NewSequence(
		tween.NewTween(2, 90, 20, ease.InCubic),
		tween.NewTween(90, 2, 20, ease.OutCubic),
	)
	for range int(s.Tweens[0].Duration() + s.Tweens[1].Duration()) {
		s.Update(1)
		if s.IsFinished() {
			break
		}
		n := int(s.Value())
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
