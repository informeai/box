package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/xfrr/goffmpeg/transcoder"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: convert [input file] [output file]")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatalf("Archive: %v not existed.", filepath.Base(inputFile))
		os.Exit(1)
	}

	if _, err := os.Stat(filepath.Dir(outputFile)); os.IsNotExist(err) {
		err := os.Mkdir(filepath.Dir(outputFile), 0755)
		if err != nil {
			log.Fatalf("Not created directory: %v", filepath.Dir(outputFile))
			os.Exit(2)
		}

	}

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputFile, outputFile)

	if err != nil {
		log.Fatalln(err)
	}

	duration := trans.MediaFile().DurationInput()
	fmt.Println(duration)
	done := trans.Run(true)

	progress := trans.Output()
	fmt.Printf("Initialized\n")
	fmt.Printf("INPUT: %v\n", filepath.Base(inputFile))
	fmt.Printf("OUTPUT: %v\n", filepath.Base(outputFile))
	for p := range progress {
		fmt.Printf("\rPROGRESS: %v\tSPEED: %v ", p.CurrentTime, p.Speed[:len(p.Speed)-1])
	}

	err = <-done
	if err != nil && err != io.EOF {
		log.Fatalln(err.Error())
		os.Exit(2)
	} else {
		fmt.Printf("\nFinished\n")
	}
}
