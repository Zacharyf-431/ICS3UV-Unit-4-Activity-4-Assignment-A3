/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-12-11
	* This file searches for a word in a given sentence
	*/


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

		reader := bufio.NewReader(os.Stdin)

		// INPUT
		fmt.Print("Please enter a sentence?\n")
		sentence, _ := reader.ReadString('\n')
		sentence = strings.TrimSpace(sentence)

		fmt.Print("Please enter a word to search for in your sentence?\n")
		word, _ := reader.ReadString('\n')
		word = strings.TrimSpace(word)


		// assign variables
		var found = false
		var sentenceLow = strings.ToLower(sentence)
		var wordLow = strings.ToLower(word)
		var maxLen = len(sentenceLow) - len(wordLow)

		// PROCESS
		for i := 0; i <= maxLen; i++ {
			match := true
			for j := 0; j < len(wordLow); j++ {
				if sentenceLow[i+j] != wordLow[j] {
					match = false
				}
			}
			if match {
				found = true
			}
		}

		// OUTPUT
		if found {
			fmt.Println("Hooray, the word", word, "was found in the sentence:", sentence)
		} else {
			fmt.Println("Sorry, the word", word, "was not found in the sentence:", sentence)
		}

		fmt.Println("\nDone.")
}