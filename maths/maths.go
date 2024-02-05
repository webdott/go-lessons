package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	return angleToPoint(angle)
}

func minuteHandPoint(t time.Time) Point {
	angle := minutesInRadians(t)
	return angleToPoint(angle)
}

func hourHandPoint(t time.Time) Point {
	angle := hoursInRadians(t)
	return angleToPoint(angle)
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}
