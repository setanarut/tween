package main

import (
	"time"

	"github.com/setanarut/tween"
)

func main() {
	tw := tween.NewTween(0, 100, time.Second*3, 0, tween.InBounce, false)
	lines := 20
	for range lines {
		tw.Update(tw.TotalDuration() / time.Duration(lines))
		n := int(tw.Value)
		for range n {
			print("x")
		}
		for j := n; j < int(tw.End); j++ {
			print(" ")
		}
		print("\n")
	}

	/*

		x
		x
		xxxx
		xxxxx
		xx
		xxxxxx
		xxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxx
		xxxxxxxx
		xxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

	*/
}
