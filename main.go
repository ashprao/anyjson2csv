package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var debugMode, all bool
var source, destination string

func init() {
	flag.BoolVar(&debugMode, "debug", false, "Enable debug mode")
	flag.BoolVar(&all, "all", false, "Includes json object and list columns")
	flag.StringVar(&destination, "output", "output.csv", "Destination file name")

	// Replace the default Usage function with a custom implementation
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <source file>\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "Only one argument referencing the source json file needs to be provided.")
	}
}

func debugLog(message string) {
	if debugMode {
		log.Println(message)
	}
}

func main() {
	flag.Parse()

	log.SetOutput(os.Stdout)
	log.SetPrefix("[json2csv] ")
	log.SetFlags(log.Ldate | log.Ltime)

	debugLog("In main")
	debugLog(fmt.Sprintf("destination =%s", destination))

	args := flag.Args()
	debugLog(fmt.Sprintf("args: = %v", args))

	switch len(args) {
	case 0:
		fmt.Println("Missing source json file name.")
		flag.Usage()
		return
	case 1:
		source = args[0]
		break
	default:
		fmt.Println("Too many arguments provided.")
		flag.Usage()
		return
	}

	err := convertAnyJSONToCSV(source, destination)
	if err != nil {
		log.Fatal(err.Error())
	}

}
