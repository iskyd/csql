package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/iskyd/csql/pkg/csql"
)

func main() {
	// Define command-line flags
	sqlFile := flag.String("file", "", "Path to the SQL file")
	outputFileName := flag.String("output", "", "Path to the output CSS file")
	flag.Parse()

	if *sqlFile == "" {
		log.Fatal("Please provide the path to the SQL file using the -file flag.")
	}

	file, err := os.Open(*sqlFile)
	if err != nil {
		log.Fatal("File does not exists.")
	}
	defer file.Close()

	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		log.Fatal("Cannot create file.")
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		row := scanner.Text()

		cssCode, err := csql.Convert(string(row))
		if err != nil {
			log.Fatalf("Failed to convert SQL to CSS: %v", err)
		}

		fmt.Fprintln(writer, cssCode)
	}

	if err := writer.Flush(); err != nil {
		log.Fatalf("Error while writer flushing: %v", err)
	}

	fmt.Println("Conversion completed successfully.")
}
