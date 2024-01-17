package main

import (
	"bufio"
	"log"
	"os"
)

func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// input := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(file)
	var texts = make([]string, 0, 50000)
	for scanner.Scan() {
		lineText := scanner.Text()
		texts = append(texts, lineText)
	}
	return texts, nil
}

func ReadFirstLine() string {
	open, err := os.Open("test.txt")
	defer open.Close()
	if err != nil {
		log.Print(err)
		return ""
	}
	reader := bufio.NewReader(open)
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Print(err)
		return ""
	}
	return string(line)
}
