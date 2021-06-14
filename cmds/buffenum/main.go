package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	var flagOutput = flag.String("o", "", "output")
	flag.Parse()

	var input io.Reader
	if flag.NArg() > 0 {
		inputBytes, err := os.ReadFile(flag.Arg(0))
		if err != nil {
			log.Fatalln(err)
		}
		input = bytes.NewReader(inputBytes)
	} else {
		input = os.Stdin
	}

	ei, err := ParseEnum(input)
	if err != nil {
		log.Fatalln(err)
	}

	outputBuffer := new(bytes.Buffer)
	if err := ei.Generate(outputBuffer); err != nil {
		log.Fatalln(err)
	}

	if *flagOutput != "" {
		if err := os.WriteFile(*flagOutput, outputBuffer.Bytes(), 0644); err != nil {
			log.Fatalln(err)
		}
	} else {
		io.Copy(os.Stdout, outputBuffer)
	}
}
