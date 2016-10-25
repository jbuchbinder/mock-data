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
	NumAll         = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	NumLimited     = []string{"2", "3", "4", "5", "6", "7", "8", "9"}
	Suffixes       = []string{"Ave", "Blvd", "Ct", "Ctr", "Dr", "Jct", "Ln", "Pt", "Rd", "St", "Sq", "Way"}
	Rand           *rand.Rand
)

func init() {
	InitRNG()
}

func InitRNG() {
	Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetArbitrary(items []string) string {
	return items[Rand.Intn(len(items))]
}

func GetEmail(first, last string) string {
	return fmt.Sprintf("%s%s%d@example.com", strings.ToLower(first[0:1]), strings.ToLower(last), Rand.Int())
}

func GetName(isFemale bool) (string, string) {
	if isFemale {
		return data.NameGivenFemale[Rand.Intn(len(data.NameGivenFemale))], data.NameSurname[Rand.Intn(len(data.NameSurname))]
	} else {
		return data.NameGivenMale[Rand.Intn(len(data.NameGivenMale))], data.NameSurname[Rand.Intn(len(data.NameSurname))]
	}
}

func GetStreetNumber() string {
	return fmt.Sprintf("%d", Rand.Intn(MaxHouseNumber))
}

func GetStreetName() string {
	return fmt.Sprintf("%s %s", strings.ToUpper(data.Streets[Rand.Intn(len(data.Streets))]), strings.ToUpper(Suffixes[Rand.Intn(len(Suffixes))]))
}

func GetStreetAddress() string {
	return fmt.Sprintf("%d %s %s", Rand.Intn(MaxHouseNumber), strings.ToUpper(data.Streets[Rand.Intn(len(data.Streets))]), strings.ToUpper(Suffixes[Rand.Intn(len(Suffixes))]))
}

func GetLocation() types.Location {
	return data.Locations[Rand.Intn(len(data.Locations))]
}

func GetPhoneNumber() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", getRandomDigit(false), getRandomDigit(false), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true))
}

func getRandomDigit(includeZeroAndOne bool) string {
	if includeZeroAndOne {
		return NumAll[Rand.Intn(len(NumAll))]
	} else {
		return NumLimited[Rand.Intn(len(NumLimited))]
	}
	return ""
}
