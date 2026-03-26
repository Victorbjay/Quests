package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ============================================================
// Q1 — joinWithPunctuation
// ============================================================
func joinWithPunctuation(tokens []string) string {
	result := ""
	for i, token := range tokens {
		if i > 0 && !isPunctuation(token) {
			result += " "
		}
		result += token
	}
	return result
}

// ============================================================
// Q2 — binToDecimal
// ============================================================
func binToDecimal(binStr string) (int64, error) {
	result, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// ============================================================
// Q3 — hexToDecimal
// ============================================================
func hexToDecimal(hexStr string) (int64, error) {
	result, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// ============================================================
// Q5 — isPunctuation
// ============================================================
func isPunctuation(s string) bool {
	punctuations := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range punctuations {
		if s == p {
			return true
		}
	}
	return false
}

// ============================================================
// Q6 — fixSingleQuotes
// ============================================================
func fixSingleQuotes(text string) string {
	result := ""
	match := false
	inner := ""

	for _, char := range text {
		if char == '\'' {
			if match {
				result += "'" + strings.TrimSpace(inner) + "'"
				inner = ""
			}
			match = !match
		} else if match {
			inner += string(char)
		} else {
			result += string(char)
		}
	}
	return result
}

// ============================================================
// Q7 — uppercaseLastN
// ============================================================
func uppercaseLastN(words []string, n int) []string {
	start := len(words) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return words
}

// ============================================================
// Q8 — aOrAn
// ============================================================
func aOrAn(nextWord string) string {
	firstLetter := strings.ToLower(nextWord[:1])
	vowelsAndH := "aeiouh"
	for _, v := range vowelsAndH {
		if firstLetter == string(v) {
			return "an"
		}
	}
	return "a"
}

// ============================================================
// Q9 — fixArticles
// ============================================================
func fixArticles(sentence string) string {
	words := strings.Fields(sentence)
	for i := 0; i < len(words)-1; i++ {
		word := words[i]
		nextWord := words[i+1]
		if word == "a" || word == "A" {
			stripped := strings.Trim(nextWord, ",.!?:;")
			if len(stripped) > 0 {
				firstLetter := strings.ToLower(string(stripped[0]))
				vowelsAndH := "aeiouh"
				if strings.Contains(vowelsAndH, firstLetter) {
					if word == "a" {
						words[i] = "an"
					} else {
						words[i] = "An"
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}

// ============================================================
// Q10 — capitaliseFirstLETTER
//

func capitaliseFirstWORD(word string) string {
	lower := strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	return lower
}

// ============================================================
// MAIN — test everything!!
// ============================================================
func main() {

	// Q1
	tokens := []string{"hello", ",", "world", "!"}
	fmt.Printf("Q1: %q\n", joinWithPunctuation(tokens))

	// Q2
	res2, _ := binToDecimal("10")
	fmt.Printf("Q2: %v\n", res2)

	// Q3
	res3, _ := hexToDecimal("1E")
	fmt.Printf("Q3: %v\n", res3)

	// Q5
	fmt.Printf("Q5: %v\n", isPunctuation(","))
	fmt.Printf("Q5: %v\n", isPunctuation("x"))

	// Q6
	fmt.Printf("Q6: %q\n", fixSingleQuotes("' awesome '"))
	fmt.Printf("Q6: %q\n", fixSingleQuotes("' hello world '"))

	// Q7
	words := []string{"this", "is", "so", "exciting"}
	fmt.Printf("Q7: %q\n", uppercaseLastN(words, 2))

	// Q8
	fmt.Printf("Q8: %q\n", aOrAn("apple"))
	fmt.Printf("Q8: %q\n", aOrAn("horse"))
	fmt.Printf("Q8: %q\n", aOrAn("book"))

	// Q9
	fmt.Printf("Q9: %q\n", fixArticles("There is a elephant and A horse"))
	// Q10
	fmt.Printf("Q10: %q\n", capitaliseFirstWORD("hELLO"))
}
//==================================================
//				THEORY
// =================================================

//Q1 explain strings.Fields and its importance in your project
//Q2 explain strconv.ParseInt(s, 2, 64) 
//Q3 explain difference between strings.ToUpper and strings.Title
//Q4 run theough how your program  handles the bin conversion in the input text "it has been "1010" (bin) years since the last event"
