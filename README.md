# Tween [![](https://godoc.org/github.com/setanarut/tween?status.svg)](http://godoc.org/github.com/setanarut/tween)

Tween is a small library to perform [tweening](https://en.wikipedia.org/wiki/Tweening) in Go. It has a minimal
interface, and it comes with several easing functions.

## Quick start

Tween usage

```Go
// tween from 0 to 1 in 3 seconds
tw := tween.NewTween(0, 1, 3, ease.Linear)

// advance by 1.5 seconds
tw.Update(1.5)

// get tween value at 1.5 seconds using Value()
fmt.Println(tw.Value()) // 0.5
```
Sequence usage

```Go
// merge multiple tweens into a sequence
sequence := tween.NewSequence(
	tween.NewTween(0, 100, 3, ease.InCubic),
	tween.NewTween(100, 40, 2, ease.OutCubic),
	tween.NewTween(4, 100, 20, ease.InOutBounce),
)

// advance by 7.5 seconds
sequence.Update(7.5)

// get sequence value at 7.5 seconds using Value()
fmt.Println(sequence.Value()) // 5.3125
```

See [examples](./examples/) folder for more examples.

## Easing functions

Easing functions are functions that express how slow/fast the interpolation happens in tween.

![tween-families](https://github.com/user-attachments/assets/b364ff8d-bc7b-4b35-82ac-d89bf0eec933)

The easing functions can be found in the `ease` package.

They can be divided into several families:

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

### Custom easing functions

You are not limited to gween's easing functions; if you pass a function parameter
in the easing, it will be used.

The passed function will need to suite the TweenFunc interface: `func(t, b, c, d float64) float64`

* `t` (time): starts in 0 and usually moves towards duration
* `b` (begin): initial value of the of the property being eased.
* `c` (change): ending value of the property - starting value of the property
* `d` (duration): total duration of the tween

And must return the new value after the interpolation occurs.

Here's an example using a custom easing.

```golang
labelTween := tween.new(0, 300, 4, func(t, b, c, d) float64 {
  return c*t/d + b // linear ease
})
```

## Credits

tween is an independent fork of the [gween](https://github.com/tanema/gween) package. The API is different.
