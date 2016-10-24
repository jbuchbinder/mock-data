package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	Path    = flag.String("path", "../../source", "Path to data")
	Package = flag.String("package", "data", "Go package name")
)

func main() {
	flag.Parse()

	fmt.Printf("package %s\n\n", *Package)
	fmt.Printf("var (\n")

	{
		f, err := fileToArray("dist.male.first")
		if err != nil {
			panic(err)
		}
		fmt.Printf(arrayToString("NameGivenMale", f))
	}
	{
		f, err := fileToArray("dist.female.first")
		if err != nil {
			panic(err)
		}
		fmt.Printf(arrayToString("NameGivenFemale", f))
	}
	{
		f, err := fileToArray("dist.all.last")
		if err != nil {
			panic(err)
		}
		fmt.Printf(arrayToString("NameSurname", f))
	}

	fmt.Printf(")\n\n")
}

func fileToArray(fn string) ([]string, error) {
	out := make([]string, 0)

	if !fileExists(*Path + string(os.PathSeparator) + fn) {
		return out, errors.New("Data file not found")
	}
	f, err := os.Open(*Path + string(os.PathSeparator) + fn)
	if err != nil {
		return out, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		path, err := r.ReadString(10) // 0x0A separator = newline
		if err == io.EOF {
			return out, nil
			break
		} else if err != nil {
			return out, err
		}
		if len(path) < 10 {
			return out, nil
			break
		}
		// Only ingest the first word
		out = append(out, path[0:strings.Index(path, " ")])
	}
	return out, nil
}

func arrayToString(name string, items []string) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("\t%s = []string { ", name))
	for i := range items {
		buf.WriteString(fmt.Sprintf(`"%s"`, items[i]))
		if i+1 < len(items) {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(" }\n")
	return buf.String()
}
