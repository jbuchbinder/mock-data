# MOCK-DATA

Golang mock data source.

## EXAMPLES

```
import (
	"github.com/jbuchbinder/mock-data/types"
	"github.com/jbuchbinder/mock-data/mock"
)

// Random location
var t types.Location
t = mock.GetLocation()

// Names
isFemale := false
first, last := types.GetName(isFemale)
```

## SOURCES

| Data | Webpage | Link |
| --- | --- | --- |
| First / Last Names | `https://catalog.data.gov/dataset/names-from-census-1990` | `http://www.census.gov/topics/population/genealogy/data/1990_census/1990_census_namefiles.html` |
| Street Names | `http://inventwithpython.com/blog/2011/09/28/list-of-street-names-list-of-last-names/` | `http://coffeeghost.net/files/city_street_names.zip` |
| Zip Codes + Lat/Lng | `http://federalgovernmentzipcodes.us/` | `http://federalgovernmentzipcodes.us/free-zipcode-database-Primary.csv` |

