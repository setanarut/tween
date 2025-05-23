package main

import (
	"fmt"

	"github.com/setanarut/tween"
	"github.com/setanarut/tween/ease"
)

// Yoyo example
func main() {

	// Create Tween and enable Yoyo
	tw := tween.NewTween(0, 5, 5, ease.Linear).SetYoyo(true)
	for range 20 {
		fmt.Print(tw.Value())
		fmt.Print(" ")
		tw.Update(1)
	}

	// 0 1 2 3 4 5 4 3 2 1 0 1 2 3 4 5 4 3 2 1

	fmt.Print("\n")

	// Disable yoyo and reset
	tw.Yoyo = false // or tw.SetYoyo(false)
	tw.Reset()

	for range 20 {
		fmt.Print(tw.Value())
		fmt.Print(" ")
		tw.Update(1)
	}
	// 0 1 2 3 4 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5

	fmt.Print("\n")
}
