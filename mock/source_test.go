package mock

import (
	"testing"
)

func Benchmark1MData(b *testing.B) {
	InitRNG() // just in case
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < 1000000; i++ {
		gender := Rand.Float64() < .5
		fname, lname := GetName(gender)
		addr := GetStreetAddress()
		loc := GetLocation()
		b.Logf("%d : %s %s / %s / %s, %s %s / %s", i, fname, lname, addr, loc.City, loc.State, loc.Zipcode, GetPhoneNumber())
	}
	b.StopTimer()
}
