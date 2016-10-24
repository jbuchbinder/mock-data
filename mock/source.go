package mock

import (
	_ "github.com/jbuchbinder/mock-data/data"
	"math"
)

func GetName(isFemale bool) (string, string) {
	if isFemale {
		return NameGivenFemale[math.Intn(len(NameGivenFemale))], NameSurname[math.Intn(len(NameSurname))]
	} else {
		return NameGivenMale[math.Intn(len(NameGivenMale))], NameSurname[math.Intn(len(NameSurname))]
	}
}

func GetLocation() Location {
	return Locations[math.Intn(len(Locations))]
}
