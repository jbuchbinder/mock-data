package mock

import (
	"fmt"
	"github.com/jbuchbinder/mock-data/data"
	"github.com/jbuchbinder/mock-data/types"
	"math/rand"
	"strings"
	"time"
)

var (
	MaxHouseNumber = 10000
	Suffixes       = []string{"Ave", "Blvd", "Ct", "Ctr", "Dr", "Jct", "Ln", "Pt", "Rd", "St", "Sq", "Way"}
	Rand           *rand.Rand
)

func init() {
	InitRNG()
}

func InitRNG() {
	Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetName(isFemale bool) (string, string) {
	if isFemale {
		return data.NameGivenFemale[Rand.Intn(len(data.NameGivenFemale))], data.NameSurname[Rand.Intn(len(data.NameSurname))]
	} else {
		return data.NameGivenMale[Rand.Intn(len(data.NameGivenMale))], data.NameSurname[Rand.Intn(len(data.NameSurname))]
	}
}

func GetStreetAddress() string {
	return fmt.Sprintf("%d %s %s", Rand.Intn(MaxHouseNumber), strings.ToUpper(data.Streets[Rand.Intn(len(data.Streets))]), strings.ToUpper(Suffixes[Rand.Intn(len(Suffixes))]))
}

func GetLocation() types.Location {
	return data.Locations[Rand.Intn(len(data.Locations))]
}
