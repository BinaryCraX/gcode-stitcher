package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func printUsage() {
	fmt.Println("Usage: " + os.Args[0] + " srcFolder trgFile")
}

func main() {
	flag.Parse()

	srcFolder := flag.Arg(0)
	trgFile := flag.Arg(1)
	if srcFolder == "" || trgFile == "" {
		printUsage()
		return
	}

	origPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	os.Chdir(srcFolder)

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	startFinder, err := regexp.Compile("\nT0\n")
	if err != nil {
		log.Fatal(err)
	}
	endFinder, err := regexp.Compile("\n; layer end\n")
	if err != nil {
		log.Fatal(err)
	}

	out := make([]byte, 0)

	for i, file := range files {
		fmt.Println(file.Name())

		data, err := ioutil.ReadFile(file.Name())
		if err != nil {
			log.Fatal(err)
		}

		startPos := startFinder.FindIndex(data)[0]
		endPos := endFinder.FindIndex(data)[0]

		// use the prefix of the first file and the suffix of the last file
		if i == 0 {
			startPos = 0
		}
		if i == len(files)-1 {
			endPos = len(data)
		}

		out = append(out, data[startPos:endPos]...)
	}

	os.Chdir(origPath)
	ioutil.WriteFile(trgFile, out, os.FileMode(0666))
}
