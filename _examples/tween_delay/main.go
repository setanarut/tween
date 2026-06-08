package main

import (
	"fmt"
	"time"

	"github.com/setanarut/tween"
)

func main() {
	tw := tween.NewTweenWithDelay(time.Second, 1.345, 100, time.Second*3, tween.Linear, false)

	var dt time.Duration
	fmt.Println(tw.TotalDuration())
	for range 100 {
		dt += (time.Second / 60)
		fmt.Println(tw.Value, tw.IsDelaying())
		tw.Update(dt)
	}
}
