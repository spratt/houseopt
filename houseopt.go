package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/jonas-p/go-shp"
)

const (
	ErrorExitCode = 1
	MinArgs = 1
)

func Usage() {
	CommandLine := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
    fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
    flag.PrintDefaults()
    fmt.Fprint(CommandLine.Output(), "The first positional argument is the filename of the SHP broadband input.\n")
}

func main() {
	flag.Parse()
	if flag.NArg() < MinArgs {
		Usage()
		os.Exit(ErrorExitCode)
	}
	filename := flag.Arg(0)
	
	// open a shapefile for reading
	shape, err := shp.Open(filename)
	if err != nil { log.Fatal(err) } 
	defer shape.Close()
	
	// fields from the attribute table (DBF)
	fields := shape.Fields()
	
	// loop through all features in the shapefile
	for shape.Next() {
		n, p := shape.Shape()
		
		// print feature
		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())
		
		// print attributes
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Printf("\t%v: %v\n", f, val)
		}
		fmt.Println()
		os.Exit(ErrorExitCode)
	}
}
