[![GoDoc](https://godoc.org/github.com/setanarut/tween?status.svg)](https://pkg.go.dev/github.com/setanarut/tween)
[![Go Report Card](https://goreportcard.com/badge/github.com/setanarut/tween)](https://goreportcard.com/report/github.com/setanarut/tween)
[![Coverage Status](https://coveralls.io/repos/github/setanarut/tween/badge.svg?branch=main)](https://coveralls.io/github/setanarut/tween?branch=main)

# Tween

Tween is a small library to perform [tweening](https://en.wikipedia.org/wiki/Tweening) in Go. It has a minimal
interface, and it comes with several easing functions.

## Quick start

Tween usage

```Go
func main() {
	// tween from 0 to 1 in 3 seconds
	tw := tween.NewTween(0, 1, 3*time.Second, tween.Linear, false)

	// advance by 1.5 seconds
	tw.Update(time.Millisecond * 1500)

	// get tween value at 1.5 seconds
	fmt.Println(tw.Value) // 0.5

	// merge multiple tweens into a sequence
	sequence := tween.NewSequence(
		tween.NewTween(0, 100, 3*time.Second, tween.InCubic, false),
		tween.NewTween(100, 40, 2*time.Second, tween.OutCubic, false),
		tween.NewTween(4, 100, 20*time.Second, tween.InOutBounce, false),
	)

	// advance by 7.5 seconds
	sequence.Update(time.Millisecond * 7500)

	// get sequence value at 7.5 seconds
	fmt.Println(sequence.Value) // 5.3125
}

```

See [_examples](./_examples/) folder for more examples.

## Easing functions

Easing functions are functions that express how slow/fast the interpolation happens in tween.

![tween-families](https://github.com/user-attachments/assets/b364ff8d-bc7b-4b35-82ac-d89bf0eec933)

The easing functions can be divided into several families:

* `linear` is the simplest easing function, straight from one value to the other.
* `quad`, `cubic`, `quart`, `quint`, `expo`, `sine` and `circle` are all "smooth" curves that will make transitions look natural.
* The `back` family starts by moving the interpolation slightly "backwards" before moving it forward.
* The `bounce` family simulates the motion of an object bouncing.
* The `elastic` family simulates inertia in the easing, like an elastic gum.

Each family (except `linear`) has 4 variants:
* `In` starts slow, and accelerates at the end
* `Out` starts fast, and decelerates at the end
* `InOut` starts and ends slow, but it's fast in the middle
* `OutIn` starts and ends fast, but it's slow in the middle

| family      | in        | out        | inOut        | outIn        |
| ----------- | --------- | ---------- | ------------ | ------------ |
| **Linear**  | Linear    | Linear     | Linear       | Linear       |
| **Quad**    | InQuad    | OutQuad    | InOutQuad    | OutInQuad    |
| **Cubic**   | InCubic   | OutCubic   | InOutCubic   | OutInCubic   |
| **Quart**   | InQuart   | OutQuart   | InOutQuart   | OutInQuart   |
| **Quint**   | InQuint   | OutQuint   | InOutQuint   | OutInQuint   |
| **Expo**    | InExpo    | OutExpo    | InOutExpo    | OutInExpo    |
| **Sine**    | InSine    | OutSine    | InOutSine    | OutInSine    |
| **Circ**    | InCirc    | OutCirc    | InOutCirc    | OutInCirc    |
| **Back**    | InBack    | OutBack    | InOutBack    | OutInBack    |
| **Bounce**  | InBounce  | OutBounce  | InOutBounce  | OutInBounce  |
| **Elastic** | InElastic | OutElastic | InOutElastic | OutInElastic |
