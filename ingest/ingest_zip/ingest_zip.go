package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	Source  = flag.String("source", "../../source/free-zipcode-database-Primary.csv", "Path to data file")
	Package = flag.String("package", "data", "Go package name")
)

func main() {
	flag.Parse()

	fmt.Printf("package %s\n\n", *Package)
	fmt.Printf("import \"github.com/jbuchbinder/mock-data/types\"\n\n")
	fmt.Printf("var (\n")

	// Open and parse
	if !fileExists(*Source) {
		panic("Unable to open source file")
	}
	f, err := os.Open(*Source)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	fmt.Printf("\tLocations = []types.Location{\n")
	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if i == 0 {
			i++
			continue
		} // skip header

		// Skip missing lat/lng
		if record[5] == "" || record[6] == "" {
			i++
			continue
		}

		fmt.Printf("\t\t"+`types.Location{"%s", "%s", "%s", "%s", %s, %s, "%s", %s}, `+"\n", record[0], record[2], record[3], record[4], record[5], record[6], record[7], record[8])
		i++
	}
	fmt.Printf("\t}\n")

	fmt.Printf(")\n\n")
}
