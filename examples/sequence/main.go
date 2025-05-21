package main

import (
	"fmt"

	"github.com/setanarut/tween"
	"github.com/setanarut/tween/ease"
)

func main() {
	var s = tween.NewSequence(
		tween.NewTween(0, 4, 4, ease.Linear),
		tween.NewTween(4, 0, 10, ease.OutInQuad),
	)
	s.SetLoop(-1)
	for range 20 {
		s.Update(1)
		fmt.Println(s.Value(), s.IsActiveTweenFinished(), s.IsFinished())
	}
}
