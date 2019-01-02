package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Append tense to each sentence in file
// ***FOR TESTING***
func parseFile(filename string, tense int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Print(text + " : ")
		sentenceTense := globalIdentifier.checkTense(text)
		if tense == sentenceTense {
			fmt.Println("OK")
		} else {
			fmt.Println(sentenceTense)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
