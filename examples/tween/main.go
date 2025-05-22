package main

import (
	"github.com/setanarut/tween"
	"github.com/setanarut/tween/ease"
)

func main() {
	t := tween.NewTween(0, 90, 20, ease.InBounce)

	for range int(t.Duration) {
		t.Update(1)
		n := int(t.Value())
		for range n {
			print("x")
		}
		for j := n; j < int(t.End); j++ {
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
