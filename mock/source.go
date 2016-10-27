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
	return items[Rand.Intn(len(items)-1)]
}

func GetEmail(first, last string) string {
	return fmt.Sprintf("%s%s%d@example.com", strings.ToLower(first[0:1]), strings.ToLower(last), Rand.Int())
}

func GetName(isFemale bool) (string, string) {
	if isFemale {
		return data.NameGivenFemale[Rand.Intn(len(data.NameGivenFemale)-1)], data.NameSurname[Rand.Intn(len(data.NameSurname)-1)]
	} else {
		return data.NameGivenMale[Rand.Intn(len(data.NameGivenMale)-1)], data.NameSurname[Rand.Intn(len(data.NameSurname)-1)]
	}
}

func GetNameGenderless() (string, string) {
	return GetName(Rand.Float64() < .5)
}

func GetFullName(isFemale bool) string {
	first, last := GetName(isFemale)
	return first + " " + last
}

func GetFullNameGenderless() string {
	return GetFullName(Rand.Float64() < .5)
}

func GetStreetNumber() string {
	return fmt.Sprintf("%d", Rand.Intn(MaxHouseNumber))
}

func GetStreetName() string {
	return fmt.Sprintf("%s %s", strings.ToUpper(data.Streets[Rand.Intn(len(data.Streets)-1)]), strings.ToUpper(Suffixes[Rand.Intn(len(Suffixes)-1)]))
}

func GetStreetAddress() string {
	return fmt.Sprintf("%d %s %s", Rand.Intn(MaxHouseNumber), strings.ToUpper(data.Streets[Rand.Intn(len(data.Streets)-1)]), strings.ToUpper(Suffixes[Rand.Intn(len(Suffixes)-1)]))
}

func GetLocation() types.Location {
	return data.Locations[Rand.Intn(len(data.Locations)-1)]
}

func GetPhoneNumber() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", getRandomDigit(false), getRandomDigit(false), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true), getRandomDigit(true))
}

func getRandomDigit(includeZeroAndOne bool) string {
	if includeZeroAndOne {
		return NumAll[Rand.Intn(len(NumAll)-1)]
	} else {
		return NumLimited[Rand.Intn(len(NumLimited)-1)]
	}
	return ""
}
