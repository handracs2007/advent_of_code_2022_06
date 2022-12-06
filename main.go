package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// hasUniqueCharacters checks whether all characters in s are unique.
func hasUniqueCharacters(s string) bool {
	// Use a map to keep track of the count of each character in the string
	charCount := make(map[rune]int)

	// Iterate over the string, counting each character
	for _, c := range s {
		charCount[c]++
	}

	// Iterate over the map and check if any character has a count greater than 1
	for _, count := range charCount {
		if count > 1 {
			// If a character has a count greater than 1, return false
			return false
		}
	}

	// If we reach here, all characters in the string are unique
	return true
}

// doCheck performs the real logic for this puzzle. It receives a string str and integer uniqueCount that indicates
// how many characters need to be unique.
func doCheck(str string, uniqueCount int) {
	// We will start from the beginning of the string, keep moving to the right. We need to stop moving
	// when we are sure that there is no more enough character to be read that can fulfill the minimum length of
	// uniqueCount.
	for i := 0; i < len(str)-uniqueCount-1; i++ {
		// Let's get the first uniqueCount characters.
		s := str[i : i+uniqueCount]
		if hasUniqueCharacters(s) {
			// We just print it out here and exit the loop as there is nothing else we should do.
			fmt.Println(i + uniqueCount)
			break
		}
	}
}

// getString gets the string from the input file. This string is to be used for this puzzle.
func getString() string {
	// Read the input file.
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %s", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	l, _ := r.ReadString('\n')
	l = strings.TrimSpace(l)

	return l
}

func main() {
	str := getString()
	doCheck(str, 4)
	doCheck(str, 14)
}
