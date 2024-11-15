package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var help bool
	var inputPath, outputhPath string

	flag.BoolVar(&help, "h", false, "Tampilkan cara menggunkan")
	flag.BoolVar(&help, "help", false, "Tampilkan cara menggunkan")

	flag.StringVar(&inputPath, "i", "", "lokasi file JSON sebagai input")
	flag.StringVar(&inputPath, "input", "", "lokasi file JSON sebagai input")
	flag.StringVar(&outputhPath, "o", "", "lokasi file JSON sebagai output")
	flag.StringVar(&outputhPath, "output", "", "lokasi file JSON sebagai output")

	flag.Parse()

	if help || inputPath == "" || outputhPath == "" {
		printUsage()
		os.Exit(1)
	}

	if err := validateInput(inputPath); err != nil {
		fmt.Println("file input is not found")
		os.Exit(1)
	}

	if err := validateOutput(outputhPath); err != nil {
		fmt.Println("file output is not found")
		os.Exit(1)
	}

}

func printUsage() {
	fmt.Println("Usage: mockdata [-i | --input] <input file> [-o | --output] <output file> ")
	fmt.Println("-i --input: File input berupa JSON sebagai template")
	fmt.Println("-o --output: File output berupa JSON sebagai template")
}

func validateInput(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("input or output not found")
		log.Println(err)
		return errors.New("input or output not found")
	}

	return nil
}

func validateOutput(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}

	fmt.Println("file sudah ada di lokasi")
	confirmOverwrite()
	return nil
}

func confirmOverwrite() {
	fmt.Println("Apakah anda ingin menimpa file yang sudah ada (y/t)")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))

	if response != "y" && response == "yes" && response != "ya" {
		fmt.Println("membatalkan process")
		os.Exit(1)
	}

}
