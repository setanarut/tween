package main

import (
	"time"

	"github.com/setanarut/tween"
)

func main() {
	t := tween.NewTween(0, 90, time.Second*20, tween.InBounce, false)

	for range 20 {
		t.Update(time.Second)
		n := int(t.Value)
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
