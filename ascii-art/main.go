package main

import (
	"fmt"
	"os"
	"strings"
)

func loadFont(filename string) map[rune][]string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading the file", err)
		return nil
	}

	lines := strings.Split(string(data), "\n")

	charMap := make(map[rune][]string)
	startChar := rune(32)

	for i := 1; i+8 < len(lines); i += 9 {
		char := startChar + rune(i/9)
		charMap[char] = lines[i : i+8]
	}

	return charMap

}

func printArt(input string, charMap map[rune][]string) {
	inputLines := strings.Split(input, "\\n")

	for _, word := range inputLines {
		if word == "" {
			fmt.Println()
			continue
		}

		for row := 0; row < 8; row++ {
			for _, char := range word {
				fmt.Print(charMap[char][row])
			}
			fmt.Println()
		}

	}

}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Use this format: go run main.go <text> <font>")
		fmt.Println("font: standard, shadow or thinkertoy")
		return
	}

	input := os.Args[1]

	font := "standard"
	if len(os.Args) == 3 {
		font = os.Args[2]
	}

	if font != "standard" && font != "shadow" && font != "thinkertoy" {
		fmt.Println("unknown font, choose from: standard, shadow or thinkertoy")
		return
	}

	charMap := loadFont(font + ".txt")
	if charMap == nil {
		return
	}

	printArt(input, charMap)

}
