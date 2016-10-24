package mock

import (
	"github.com/jbuchbinder/mock-data/data"
	"github.com/jbuchbinder/mock-data/types"
	"math/rand"
)

func GetName(isFemale bool) (string, string) {
	if isFemale {
		return data.NameGivenFemale[rand.Intn(len(data.NameGivenFemale))], data.NameSurname[rand.Intn(len(data.NameSurname))]
	} else {
		return data.NameGivenMale[rand.Intn(len(data.NameGivenMale))], data.NameSurname[rand.Intn(len(data.NameSurname))]
	}
}

func GetLocation() types.Location {
	return data.Locations[rand.Intn(len(data.Locations))]
}
