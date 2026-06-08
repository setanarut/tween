package main

import (
	"fmt"
	"time"

	"github.com/setanarut/tween"
)

// Yoyo example
func main() {

	// Create Tween and enable Yoyo
	tw := tween.NewTween(0, 10, time.Second*10, 0, tween.Linear, true)
	for range 20 {
		fmt.Print(tw.Value)
		fmt.Print(" ")
		tw.Update(time.Second)
	}

	// 0 1 2 3 4 5 6 7 8 9 10 9 8 7 6 5 4 3 2 1

	fmt.Print("\n")

	// Disable yoyo and reset
	tw.Yoyo = false
	tw.Reset()

	for range 20 {
		fmt.Print(tw.Value)
		fmt.Print(" ")
		tw.Update(time.Second)
	}
	// 0 1 2 3 4 5 6 7 8 9 10 10 10 10 10 10 10 10 10 10

	fmt.Print("\n")
}
