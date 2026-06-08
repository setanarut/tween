package main

import (
	"fmt"
	"time"

	"github.com/setanarut/tween"
)

func main() {
	tw := tween.NewTween(1.345, 100, time.Second*3, time.Second, tween.Linear, false)
	fmt.Println("Total Duration: ", tw.TotalDuration())
	for range 100 {
		if tw.IsDelaying() {
			fmt.Println(tw.Value, "Delaying")
		} else {
			fmt.Println(tw.Value)
		}
		tw.Update(tween.FixedTimeStep)
	}
}
