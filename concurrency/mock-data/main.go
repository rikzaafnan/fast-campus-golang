package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"mock-data/data"
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

	var mapping map[string]string
	if err := readInput(inputPath, &mapping); err != nil {
		fmt.Printf("failed read input: %s \n", err)
		os.Exit(1)
	}

	if err := validateType(mapping); err != nil {
		fmt.Printf("failed validation data type: %s \n", err)
		os.Exit(1)
	}

	result, err := generateOuput(mapping)
	if err != nil {
		fmt.Printf("failed read input: %s \n", err)
		os.Exit(1)
	}

	if err := writeOutput(outputhPath, result); err != nil {
		fmt.Printf("gagal menulis hasil: %s \n", err)
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

func readInput(path string, mapping *map[string]string) error {
	if path == "" {
		return errors.New("input is empty")
	}

	if mapping == nil {
		return errors.New("mapping not valid")
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileByte, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(fileByte) == 0 {
		return errors.New("input is empty")
	}

	if err := json.Unmarshal(fileByte, &mapping); err != nil {
		return err
	}

	return nil
}

func validateType(mapping map[string]string) error {

	for _, value := range mapping {
		if !data.Supported[value] {
			return errors.New("data type is not supported")
		}
	}

	return nil
}

func generateOuput(mapping map[string]string) (map[string]any, error) {

	resultt := make(map[string]any)

	for key, value := range mapping {
		resultt[key] = data.Generate(value)
	}

	return resultt, nil
}

func writeOutput(path string, result map[string]any) error {

	if path == "" {
		return errors.New("path is empty")
	}

	flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(path, flags, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	resultByte, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		return err
	}

	_, err = file.Write(resultByte)
	if err != nil {
		return err
	}
	return nil
}
