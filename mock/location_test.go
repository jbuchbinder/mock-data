package mock

import (
	"fmt"
	"github.com/jbuchbinder/mock-data/types"
	"log"
	"testing"
)

func Test_RandomLocations(t *testing.T) {
	orig := types.Location{
		Latitude:  41.716667,
		Longitude: 72.216667,
	}
	log.Printf("Original = %s", showLocation(orig))
	for i := 0; i < 20; i++ {
		l := RandomLocationWithinDistance(orig, 20)
		log.Printf("Location %d: %s", i, showLocation(l))
	}
}

func showLocation(l types.Location) string {
	return fmt.Sprintf("LAT: %f, LON: %f", l.Latitude, l.Longitude)
}
