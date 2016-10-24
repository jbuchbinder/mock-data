package mock

import (
	"github.com/jbuchbinder/mock-data/types"
	"math"
	"math/rand"
)

// RandomLocationWithinDistance generates a random location within a certain
// distance of a location entry
func RandomLocationWithinDistance(l types.Location, miles float64) types.Location {
	// Convert radius from meters to degrees
	radiusInDegrees := float64(MilesToMeters(miles) / 111000)

	u := rand.Float64()
	v := rand.Float64()
	w := radiusInDegrees * math.Sqrt(u)
	t := 2 * math.Pi * v
	x := w * math.Cos(t)
	y := w * math.Sin(t)

	// Adjust the x-coordinate for the shrinking of the east-west distances
	newX := x / math.Cos(l.Longitude)

	foundLongitude := newX + l.Longitude
	foundLatitude := y + l.Latitude

	return types.Location{Latitude: foundLatitude, Longitude: foundLongitude}
}

// MilesToMeters
func MilesToMeters(miles float64) float64 {
	return miles * 1609.34
}

// MilesToRadians determines the number of radians equivalent to a distance
func MilesToRadians(miles float64) float64 {
	return miles / 69.1586137
}

// MetersToRadians determines the number of radians equivalent to a distance
func MetersToRadians(meters float64) float64 {
	return meters / 111300
}

// DegreesToRadians
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// RadiansToDegrees
func RadiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
