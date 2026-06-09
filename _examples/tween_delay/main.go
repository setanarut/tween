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

	// Total Duration:  4s
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.345 Delaying
	// 1.8930819960099998
	// 2.44116530742
	// 2.9892486188299996
	// 3.5373319302399997
	// 4.08541524165
	// 4.63349855306
	// 5.18158186447
	// 5.729665175879999
	// 6.277748487289999
	// 6.8258317986999995
	// 7.37391511011
	// 7.92199842152
	// 8.47008173293
	// 9.01816504434
	// 9.566248355750002
	// 10.114331667160002
	// 10.662414978570002
	// 11.210498289980002
	// 11.758581601390002
	// 12.306664912800002
	// 12.85474822421
	// 13.402831535620003
	// 13.95091484703
	// 14.49899815844
	// 15.047081469850001
	// 15.59516478126
	// 16.14324809267
	// 16.69133140408
	// 17.23941471549
	// 17.7874980269
	// 18.33558133831
	// 18.88366464972
	// 19.43174796113
	// 19.97983127254
	// 20.52791458395
	// 21.07599789536
	// 21.624081206769997
	// 22.17216451818
	// 22.720247829589997
}
