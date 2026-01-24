package tween

import (
	"math"
)

const (
	Linear       string = "Linear"
	InQuad       string = "InQuad"
	OutQuad      string = "OutQuad"
	InOutQuad    string = "InOutQuad"
	OutInQuad    string = "OutInQuad"
	InCubic      string = "InCubic"
	OutCubic     string = "OutCubic"
	InOutCubic   string = "InOutCubic"
	OutInCubic   string = "OutInCubic"
	InQuart      string = "InQuart"
	OutQuart     string = "OutQuart"
	InOutQuart   string = "InOutQuart"
	OutInQuart   string = "OutInQuart"
	InQuint      string = "InQuint"
	OutQuint     string = "OutQuint"
	InOutQuint   string = "InOutQuint"
	OutInQuint   string = "OutInQuint"
	InSine       string = "InSine"
	OutSine      string = "OutSine"
	InOutSine    string = "InOutSine"
	OutInSine    string = "OutInSine"
	InExpo       string = "InExpo"
	OutExpo      string = "OutExpo"
	InOutExpo    string = "InOutExpo"
	OutInExpo    string = "OutInExpo"
	InCirc       string = "InCirc"
	OutCirc      string = "OutCirc"
	InOutCirc    string = "InOutCirc"
	OutInCirc    string = "OutInCirc"
	InElastic    string = "InElastic"
	OutElastic   string = "OutElastic"
	InOutElastic string = "InOutElastic"
	OutInElastic string = "OutInElastic"
	InBack       string = "InBack"
	OutBack      string = "OutBack"
	InOutBack    string = "InOutBack"
	OutInBack    string = "OutInBack"
	InBounce     string = "InBounce"
	OutBounce    string = "OutBounce"
	InOutBounce  string = "InOutBounce"
	OutInBounce  string = "OutInBounce"
)

var EaseMap = map[string]TweenFunc{
	Linear:       LinearFunc,
	InQuad:       InQuadFunc,
	OutQuad:      OutQuadFunc,
	InOutQuad:    InOutQuadFunc,
	OutInQuad:    OutInQuadFunc,
	InCubic:      InCubicFunc,
	OutCubic:     OutCubicFunc,
	InOutCubic:   InOutCubicFunc,
	OutInCubic:   OutInCubicFunc,
	InQuart:      InQuartFunc,
	OutQuart:     OutQuartFunc,
	InOutQuart:   InOutQuartFunc,
	OutInQuart:   OutInQuartFunc,
	InQuint:      InQuintFunc,
	OutQuint:     OutQuintFunc,
	InOutQuint:   InOutQuintFunc,
	OutInQuint:   OutInQuintFunc,
	InSine:       InSineFunc,
	OutSine:      OutSineFunc,
	InOutSine:    InOutSineFunc,
	OutInSine:    OutInSineFunc,
	InExpo:       InExpoFunc,
	OutExpo:      OutExpoFunc,
	InOutExpo:    InOutExpoFunc,
	OutInExpo:    OutInExpoFunc,
	InCirc:       InCircFunc,
	OutCirc:      OutCircFunc,
	InOutCirc:    InOutCircFunc,
	OutInCirc:    OutInCircFunc,
	InElastic:    InElasticFunc,
	OutElastic:   OutElasticFunc,
	InOutElastic: InOutElasticFunc,
	OutInElastic: OutInElasticFunc,
	InBack:       InBackFunc,
	OutBack:      OutBackFunc,
	InOutBack:    InOutBackFunc,
	OutInBack:    OutInBackFunc,
	InBounce:     InBounceFunc,
	OutBounce:    OutBounceFunc,
	InOutBounce:  InOutBounceFunc,
	OutInBounce:  OutInBounceFunc,
}

// backS is a constant that defines the overshoot amount used in Back easing functions.
// It determines how far an easing value will exceed its target before settling, adding a dynamic "back" effect to the transition.
const backS float64 = 1.70158

// TweenFunc provides an interface used for the easing equation. You can use
// one of the provided easing functions or provide your own.
//
//	t = current time
//	b = begin value
//	c = change from begin
//	d = duration
type TweenFunc func(t, b, c, d float64) float64

// LinearFunc is a linear interpolation of some t with respect to a total duration d
// between the values b and b+c
func LinearFunc(t, b, c, d float64) float64 {
	return c*t/d + b
}

// InQuadFunc is a quadratic transition based on the square of t that starts slow
// and speeds up
func InQuadFunc(t, b, c, d float64) float64 {
	return c*math.Pow(t/d, 2) + b
}

// OutQuadFunc is a quadratic transition based on the square of t that starts fast
// and slows down
func OutQuadFunc(t, b, c, d float64) float64 {
	t /= d
	return -c*t*(t-2) + b
}

// InOutQuadFunc is a quadratic transition based on the square of t that starts and
// ends slow, accelerating through the middle
func InOutQuadFunc(t, b, c, d float64) float64 {
	t = t / d * 2
	if t < 1 {
		return c/2*math.Pow(t, 2) + b
	}
	return -c/2*((t-1)*(t-3)-1) + b
}

// OutInQuadFunc is a quadratic transition based on the square of t that starts and
// ends fast, slowing through the middle
func OutInQuadFunc(t, b, c, d float64) float64 {
	if t < d/2 {
		return OutQuadFunc(t*2, b, c/2, d)
	}
	return InQuadFunc((t*2)-d, b+c/2, c/2, d)
}

// InCubicFunc is a cubic transition based on the cube of t that starts slow and
// speeds up
func InCubicFunc(t, b, c, d float64) float64 {
	return c*math.Pow(t/d, 3) + b
}

// OutCubicFunc is a cubic transition based on the cube of t that starts fast and
// slows down
func OutCubicFunc(t, b, c, d float64) float64 {
	return c*(math.Pow(t/d-1, 3)+1) + b
}

// InOutCubicFunc is a cubic transition based on the cube of t that starts and ends
// slow, accelerating through the middle
func InOutCubicFunc(t, b, c, d float64) float64 {
	t = t / d * 2
	if t < 1 {
		return c/2*t*t*t + b
	}
	t -= 2
	return c/2*(t*t*t+2) + b
}

// OutInCubicFunc is a cubic transition based on the cube of t that starts and ends
// fast, slowing through the middle
func OutInCubicFunc(t, b, c, d float64) float64 {
	if t < d/2 {
		return OutCubicFunc(t*2, b, c/2, d)
	}
	return InCubicFunc((t*2)-d, b+c/2, c/2, d)
}

// InQuartFunc is a quartic transition based on the fourth power of t that starts
// slow and speeds up
func InQuartFunc(t, b, c, d float64) float64 {
	return c*math.Pow(t/d, 4) + b
}

// OutQuartFunc is a quartic transition based on the fourth power of t that starts
// fast and slows down
func OutQuartFunc(t, b, c, d float64) float64 {
	return -c*(math.Pow(t/d-1, 4)-1) + b
}

// InOutQuartFunc is a quartic transition based on the fourth power of t that starts
// and ends slow, accelerating through the middle
func InOutQuartFunc(t, b, c, d float64) float64 {
	t = t / d * 2
	if t < 1 {
		return c/2*math.Pow(t, 4) + b
	}
	return -c/2*(math.Pow(t-2, 4)-2) + b
}

// OutInQuartFunc is a quartic transition based on the fourth power of t that starts
// and ends fast, slowing through the middle
func OutInQuartFunc(t, b, c, d float64) float64 {
	if t < d/2 {
		return OutQuartFunc(t*2, b, c/2, d)
	}
	return InQuartFunc((t*2)-d, b+c/2, c/2, d)
}

// InQuintFunc is a quintic transition based on the fifth power of t that starts
// slow and speeds up
func InQuintFunc(t, b, c, d float64) float64 {
	return c*math.Pow(t/d, 5) + b
}

// OutQuintFunc is a quintic transition based on the fifth power of t that starts
// fast and slows down
func OutQuintFunc(t, b, c, d float64) float64 {
	return c*(math.Pow(t/d-1, 5)+1) + b
}

// InOutQuintFunc is a quintic transition based on the fifth power of t that starts
// and ends slow, accelerating through the middle
func InOutQuintFunc(t, b, c, d float64) float64 {
	t = t / d * 2
	if t < 1 {
		return c/2*math.Pow(t, 5) + b
	}
	return c/2*(math.Pow(t-2, 5)+2) + b
}

// OutInQuintFunc is a quintic transition based on the fifth power of t that starts
// and ends fast, slowing through the middle
func OutInQuintFunc(t, b, c, d float64) float64 {
	if t < d/2 {
		return OutQuintFunc(t*2, b, c/2, d)
	}
	return InQuintFunc((t*2)-d, b+c/2, c/2, d)
}

// InSineFunc is a sinusoidal transition based on the cosine of t that starts slow
// and speeds up
func InSineFunc(t, b, c, d float64) float64 {
	return -c*math.Cos(t/d*(math.Pi/2)) + c + b
}

// OutSineFunc is a sinusoidal transition based on the sine or cosine of t that
// starts fast and slows down
func OutSineFunc(t, b, c, d float64) float64 {
	return c*math.Sin(t/d*(math.Pi/2)) + b
}

// InOutSineFunc is a sinusoidal transition based on the cosine of t that starts and
// ends slow, accelerating through the middle
func InOutSineFunc(t, b, c, d float64) float64 {
	return -c/2*(math.Cos(math.Pi*t/d)-1) + b
}

// OutInSineFunc is a sinusoidal transition based on the sine or cosine of t that
// starts and ends fast, slowing through the middle
func OutInSineFunc(t, b, c, d float64) float64 {
	if t < d/2 {
		return OutSineFunc(t*2, b, c/2, d)
	}
	return InSineFunc((t*2)-d, b+c/2, c/2, d)
}

// InExpoFunc is a exponential transition based on the 2 to power 10*t that starts
// slow and speeds up
func InExpoFunc(t, b, c, d float64) float64 {
	if t == 0 {
		return b
	}
	return c*math.Pow(2, 10*(t/d-1)) + b - c*0.001
}

// OutExpoFunc is a exponential transition based on the 2 to power 10*t that starts
// fast and slows down
func OutExpoFunc(t, b, c, d float64) float64 {
	if t == d {
		return b + c
	}
	return c*1.001*(-math.Pow(2, -10*t/d)+1) + b
}

// InOutExpoFunc is a exponential transition based on the 2 to power 10*t that
// starts and ends slow, accelerating through the middle
func InOutExpoFunc(t, b, c, d float64) float64 {
	if t == 0 {
		return b
	}
	if t == d {
		return b + c
	}
	t = t / d * 2
	if t < 1 {
		return c/2*math.Pow(2, 10*(t-1)) + b - c*0.0005
	}
	return c/2*1.0005*(-math.Pow(2, -10*(t-1))+2) + b
}

// OutInExpoFunc is a exponential transition based on the 2 to power 10*t that
// starts and ends fast, slowing through the middle
func OutInExpoFunc(t, b, c, d float64) float64 {
	if t < d/2 {
		return OutExpoFunc(t*2, b, c/2, d)
	}
	return InExpoFunc((t*2)-d, b+c/2, c/2, d)
}

// InCircFunc is a circular transition based on the equation for half of a circle,
// taking the square root of t, that starts slow and speeds up
func InCircFunc(t, b, c, d float64) float64 {
	return -c*(math.Sqrt(1-math.Pow(t/d, 2))-1) + b
}

// OutCircFunc is a circular transition based on the equation for half of a circle,
// taking the square root of t, that starts fast and slows down
func OutCircFunc(t, b, c, d float64) float64 {
	return c*math.Sqrt(1-math.Pow(t/d-1, 2)) + b
}

// InOutCircFunc is a circular transition based on the equation for half of a circle,
// taking the square root of t, that starts and ends slow, accelerating through
// the middle
func InOutCircFunc(t, b, c, d float64) float64 {
	t = t / d * 2
	if t < 1 {
		return -c/2*(math.Sqrt(1-t*t)-1) + b
	}
	t -= 2
	return c/2*(math.Sqrt(1-t*t)+1) + b
}

// OutInCircFunc is a circular transition based on the equation for half of a circle,
// taking the square root of t, that starts and ends fast, slowing through the
// middle
func OutInCircFunc(t, b, c, d float64) float64 {
	if t < d/2 {
		return OutCircFunc(t*2, b, c/2, d)
	}
	return InCircFunc((t*2)-d, b+c/2, c/2, d)
}

// InElasticFunc is an elastic transition that wobbles around from the start value,
// extending past start and away from end, and then accelerates towards the end
// value at the end of the transition.
func InElasticFunc(t, b, c, d float64) float64 {
	if t == 0 {
		return b
	}
	t /= d
	if t == 1 {
		return b + c
	}
	p, a, s := calculatePASFunc(c, d)
	t--
	return -(a * math.Pow(2, 10*t) * math.Sin((t*d-s)*(2*math.Pi)/p)) + b
}

// OutElasticFunc is an elastic transition that accelerates quickly away from the
// start and beyond the end value and then wobbles towards the end value at the
// end of the transition.
func OutElasticFunc(t, b, c, d float64) float64 {
	if t == 0 {
		return b
	}
	t /= d
	if t == 1 {
		return b + c
	}
	p, a, s := calculatePASFunc(c, d)
	return a*math.Pow(2, -10*t)*math.Sin((t*d-s)*(2*math.Pi)/p) + c + b
}

// InOutElasticFunc is an elastic transition that wobbles around from the start
// value, towards the middle of the transition extending beyond start away from
// end, then rapidly toward, and beyond end value, then wobbling toward end
func InOutElasticFunc(t, b, c, d float64) float64 {
	if t == 0 {
		return b
	}
	t = t / d * 2
	if t == 2 {
		return b + c
	}
	p, a, s := calculatePASFunc(c, d)
	t--
	if t < 0 {
		return -0.5*(a*math.Pow(2, 10*t)*math.Sin((t*d-s)*(2*math.Pi)/p)) + b
	}
	return a*math.Pow(2, -10*t)*math.Sin((t*d-s)*(2*math.Pi)/p)*0.5 + c + b
}

// OutInElasticFunc is an elastic transition that accelerates towards and beyond the
// average of the start and end values, wobbles toward the average, wobbles out
// and slight away from end before accelerating toward the end value
func OutInElasticFunc(t, b, c, d float64) float64 {
	if t < d/2 {
		return OutElasticFunc(t*2, b, c/2, d)
	}
	return InElasticFunc((t*2)-d, b+c/2, c/2, d)
}

// InBackFunc is a much like InQuint, but extends beyond the start away from end
// before snapping quickly to the end
func InBackFunc(t, b, c, d float64) float64 {
	t /= d
	return c*t*t*((backS+1)*t-backS) + b
}

// OutBackFunc is a much like OutQuint, but extends beyond the end away from start
// before easing toward end
func OutBackFunc(t, b, c, d float64) float64 {
	t = t/d - 1
	return c*(t*t*((backS+1)*t+backS)+1) + b
}

// InOutBackFunc is a much like InOutQuint, but extends beyond both start and end
// values on both sides of the transition
func InOutBackFunc(t, b, c, d float64) float64 {
	s := backS * 1.525
	t = t / d * 2
	if t < 1 {
		return c/2*(t*t*((s+1)*t-s)) + b
	}
	t -= 2
	return c/2*(t*t*((s+1)*t+s)+2) + b
}

// OutInBackFunc is a much like OutInQuint, but extends beyond the average of start
// and end during the middle of the transition
func OutInBackFunc(t, b, c, d float64) float64 {
	if t < (d / 2) {
		return OutBackFunc(t*2, b, c/2, d)
	}
	return InBackFunc((t*2)-d, b+c/2, c/2, d)
}

// OutBounceFunc is a bouncing transition that accelerates toward the end value and
// then bounces back slightly in decreasing amounts until coming to reset at end
func OutBounceFunc(t, b, c, d float64) float64 {
	t /= d
	if t < 1/2.75 {
		return c*(7.5625*t*t) + b
	}
	if t < 2/2.75 {
		t -= 1.5 / 2.75
		return c*(7.5625*t*t+0.75) + b
	} else if t < 2.5/2.75 {
		t -= 2.25 / 2.75
		return c*(7.5625*t*t+0.9375) + b
	}
	t -= 2.625 / 2.75
	return c*(7.5625*t*t+0.984375) + b
}

// InBounceFunc is a bouncing transition that slowly bounces away from start at
// increasing amounts before finally accelerating toward end
func InBounceFunc(t, b, c, d float64) float64 {
	return c - OutBounceFunc(d-t, 0, c, d) + b
}

// InOutBounceFunc is a bouncing transition that bounces off of the start value,
// then accelerates toward the average of start and end, then does the opposite
// toward the end value
func InOutBounceFunc(t, b, c, d float64) float64 {
	if t < d/2 {
		return InBounceFunc(t*2, 0, c, d)*0.5 + b
	}
	return OutBounceFunc(t*2-d, 0, c, d)*0.5 + c*.5 + b
}

// OutInBounceFunc is a bouncing transition that accelerates toward the average of
// start and end, bouncing off of the average toward start, then flips and
// bounces off of average toward end in increasing amounts before accelerating
// toward end
func OutInBounceFunc(t, b, c, d float64) float64 {
	if t < d/2 {
		return OutBounceFunc(t*2, b, c/2, d)
	}
	return InBounceFunc((t*2)-d, b+c/2, c/2, d)
}

func calculatePASFunc(c, d float64) (p, a, s float64) {
	p = d * 0.3
	return p, c, p / 4
}
