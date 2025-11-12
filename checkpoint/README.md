Here is  grouping all the exercises by **percentage (XP weight)** for the **Checkpoint** —

| Question # | % Weight | Exercises (All in Checkpoint) |
|------------|----------|-------------------------------|
| **Question 1** | **5%** | `only1`, `onlya`, `onlyb`, `onlyf`, `onlyz` |
| **Question 2** | **10%** | `checknumber`, `countalpha`, `countcharacter`, `printif`, `printifnot`, `rectperimeter`, `retainfirsthalf` |
| **Question 3** | **20%** | `cameltosnakecase`, `digitlen`, `firstword`, `fishandchips`, `gcd`, `hashcode`, `lastword`, `repeatalpha` |
| **Question 4** | **35%** | `findprevprime`, `fromto`, `iscapitalized`, `itoa`, `printmemory`, `printrevcomb`, `thirdtimeisacharm`, `weareunique`, `zipstring` |
| **Question 5** | **50%** | `addprimesum`, `canjump`, `chunk`, `concatalternate`, `concatslice`, `fprime`, `hiddenp`, `inter`, `reversestrcap`, `saveandmiss`, `union`, `wdmatch` |
| **Question 6** | **65%** | `fifthandskip`, `notdecimal`, `revconcatalternate`, `slice` |
| **Question 7** | **75%** | `findpairs`, `revwstr`, `rostring`, `wordflip` |
| **Question 8** | **85%** | `itoabase`, `options`, `piglatin`, `romannumbers` |
| **Question 9** | **95%** | `brackets`, `rpncalc` |
| **Question 10** | **100%** | `brainfuck`, `grouping` |


**Total: 57 exercises** (all in Checkpoint)

## ✅ **Checkpoint Question 1| in this category you might be given either onlyf, onlyz,onlyb and onlya**

### 🧠 **Goal**

Print the character `1` and **nothing else** — no newline, no space, no fmt.

---

### ✅ **Solution Code**

```go
package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune("1")
}
```

---

---

### ⚠️ **Important Notes**

* ❌ Do **not** use `fmt.Print` — unless your question allows it
* ❌ No spaces, no newline (`\n`) but if checker complains, then add it to your code.
* ✅ The output must be exactly:

  ```
  1
  ```
---

## ✅ **Question category1: `onlyz`**

---

### ✅ **Solution**

```go
package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune("z")
}
```

---
✅ **This is **Checkpoint Question Category for number 2 – ``** you might be given count aplha, countcharachter etc**
# `CheckNumber`
### 🧩 **Function**

```go
package piscine

func CheckNumber(arg string) bool {
	for _, r := range arg {       // loop through each character (rune)
		if r >= '0' && r <= '9' { // check if it's a digit
			return true
		}
	}
	return false // no digits found
}
```

### ⚙️ **Explanation (short)**

* Loops through every character in the string.
* If any character is between `'0'` and `'9'`, returns **true**.
* If none found, returns **false**.

✅ **Output example**

```
false
true
```
# CountAlpha Solution

Here's the solution with line-by-line explanation:## Line-by-Line Explanation:

```go
package piscine

func CountAlpha(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
		}
	}
	return count
}
```
```go
func CountAlpha(s string) int {
```
- Declares a function named `CountAlpha`
- Takes one parameter: `s` which is a string
- Returns an integer (the count)

```go
    count := 0
```
- Creates a variable `count` and initializes it to 0
- This will store the number of alphabetic characters we find

```go
    for _, char := range s {
```
- Loops through each character in the string `s`
- `range s` gives us each character one by one
- `_` ignores the index (we don't need it)
- `char` holds the current character (as a rune/Unicode point)

```go
        if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
```
- Checks if the character is alphabetic:
  - `char >= 'a' && char <= 'z'` → checks if it's a lowercase letter (a-z)
  - `char >= 'A' && char <= 'Z'` → checks if it's an uppercase letter (A-Z)
  - `||` means "OR" - true if either condition is true

```go
            count++
```
- If the character is alphabetic, add 1 to our count

```go
    return count
```
- After checking all characters, return the total count

## How It Works:

**Example: `"Hello world"`**
- H → alphabetic ✓ (count = 1)
- e → alphabetic ✓ (count = 2)
- l → alphabetic ✓ (count = 3)
- l → alphabetic ✓ (count = 4)
- o → alphabetic ✓ (count = 5)
- space → not alphabetic ✗
- w → alphabetic ✓ (count = 6)
- o → alphabetic ✓ (count = 7)
- r → alphabetic ✓ (count = 8)
- l → alphabetic ✓ (count = 9)
- d → alphabetic ✓ (count = 10)
- **Result: 10**
---
# CountChar Solution

Here's the solution with line-by-line explanation: ## Line-by-Line Explanation:

```go
package piscine

func CountChar(str string, c rune) int {
	count := 0
	for _, char := range str {
		if char == c {
			count++
		}
	}
	return count
}
```

```go
func CountChar(str string, c rune) int {
```
- Declares a function named `CountChar`
- Takes two parameters:
  - `str` → the string to search in
  - `c` → the character (rune) to count
- Returns an integer (the count)

```go
    count := 0
```
- Creates a variable `count` and sets it to 0
- This will track how many times we find the character

```go
    for _, char := range str {
```
- Loops through each character in the string `str`
- `range str` gives us each character one at a time
- `_` ignores the index (position)
- `char` holds the current character as a rune

```go
        if char == c {
```
- Checks if the current character equals the character we're looking for
- `==` compares the two runes

```go
            count++
```
- If they match, add 1 to the count

```go
    return count
```
- After checking all characters, return the total count

## How It Works:

**Example 1: `CountChar("Hello World", 'l')`**
- H → not 'l' ✗
- e → not 'l' ✗
- l → matches 'l' ✓ (count = 1)
- l → matches 'l' ✓ (count = 2)
- o → not 'l' ✗
- space → not 'l' ✗
- W → not 'l' ✗
- o → not 'l' ✗
- r → not 'l' ✗
- l → matches 'l' ✓ (count = 3)
- d → not 'l' ✗
- **Result: 3**

**Example 2: `CountChar("5  balloons", '5')`**
- The string contains "5" (the digit character)
- We're looking for '5' (rune character)
- Loop finds: '5' → space → space → 'b' → 'a' → 'l' → 'l' → 'o' → 'o' → 'n' → 's'
- First character '5' matches ✓ (count = 1)
- **Result: 1**

**Note:** In the test, it shows `CountChar("5  balloons", 5)` returns 0 because `5` (integer) is different from `'5'` (character rune).

**Example 3: `CountChar("   ", ' ')`**
- Three spaces in the string
- Each space matches ' ' ✓
- **Result: 3**

## Why This Works for Edge Cases:

✅ **Empty string** → loop never runs → returns 0  
✅ **Character not found** → no matches → returns 0  
✅ **Works with any character** → spaces, digits, letters, symbols

---
# PrintIf Solution

Here's the solution with line-by-line explanation:## Line-by-Line Explanation:

```go
package piscine

func PrintIf(str string) string {
	if len(str) == 0 || len(str) >= 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}
```
Here is a more friendlier version :
```go
package piscine

func PrintIf(str string) string {
	if str == "" {
		return "G\n"
	}
	if len(str) >= 3 {
		return "G\n"
	} else {
		return "Invalid Input\n"
	}
}
```
---
```go
func PrintIf(str string) string {
```
- Declares a function named `PrintIf`
- Takes one parameter: `str` which is a string
- Returns a string (either "G\n" or "Invalid Input\n")

```go
    if len(str) == 0 || len(str) >= 3 {
```
- Checks TWO conditions with OR (`||`):
  - `len(str) == 0` → checks if string is empty (length equals 0)
  - `len(str) >= 3` → checks if string length is 3 or more
- If EITHER condition is true, execute the next line

```go
        return "G\n"
```
- Returns the letter "G" followed by a newline character
- `\n` is the newline (moves to next line)

```go
    return "Invalid Input\n"
```
- If the condition was false (length is 1 or 2), return this instead
- Returns "Invalid Input" followed by a newline

## How It Works:

**Example 1: `PrintIf("abcdefz")`**
- Length = 7
- Is it empty? No (0 == 0? No)
- Is it >= 3? Yes (7 >= 3? Yes) ✓
- **Result: "G\n"**

**Example 2: `PrintIf("abc")`**
- Length = 3
- Is it empty? No
- Is it >= 3? Yes (3 >= 3? Yes) ✓
- **Result: "G\n"**

**Example 3: `PrintIf("")`**
- Length = 0
- Is it empty? Yes (0 == 0? Yes) ✓
- **Result: "G\n"**

**Example 4: `PrintIf("14")`**
- Length = 2
- Is it empty? No (2 == 0? No)
- Is it >= 3? No (2 >= 3? No) ✗
- Both conditions false → go to else
- **Result: "Invalid Input\n"**

## Logic Summary:

| Input Length | Condition | Output |
|--------------|-----------|--------|
| 0 (empty) | `len == 0` ✓ | "G\n" |
| 1 | Both false ✗ | "Invalid Input\n" |
| 2 | Both false ✗ | "Invalid Input\n" |
| 3 or more | `len >= 3` ✓ | "G\n" |

The key insight: Empty strings and strings with 3+ characters return "G", while strings with 1 or 2 characters return "Invalid Input".

---
# PrintIfNot Solution

Here's the beginner-friendly solution with line-by-line explanation:## Line-by-Line Explanation:
```go
package piscine

func PrintIfNot(str string) string {
	if len(str) < 3 {
		return "G\n"
	} else {
		return "Invalid Input\n"
	}
}
```

```go
func PrintIfNot(str string) string {
```
- Declares a function named `PrintIfNot`
- Takes one parameter: `str` which is a string
- Returns a string

```go
    if len(str) < 3 {
```
- Checks if the length of the string is **less than 3**
- `len(str)` gets the number of characters in the string
- `< 3` means 0, 1, or 2 characters

```go
        return "G\n"
```
- If length is less than 3, return "G" with a newline
- This includes empty strings (length 0)

```go
    } else {
```
- Otherwise (if length is 3 or more)

```go
        return "Invalid Input\n"
```
- Return "Invalid Input" with a newline

## How It Works:

**Example 1: `PrintIfNot("abcdefz")`**
- Length = 7 characters
- Is 7 < 3? **No** ✗
- Goes to else block
- **Result: "Invalid Input\n"**

**Example 2: `PrintIfNot("abc")`**
- Length = 3 characters
- Is 3 < 3? **No** ✗
- Goes to else block
- **Result: "Invalid Input\n"**

**Example 3: `PrintIfNot("")`**
- Length = 0 characters (empty string)
- Is 0 < 3? **Yes** ✓
- **Result: "G\n"**

**Example 4: `PrintIfNot("14")`**
- Length = 2 characters
- Is 2 < 3? **Yes** ✓
- **Result: "G\n"**

## Simple Logic Table:

| Input Length | Condition (< 3) | Output |
|--------------|-----------------|--------|
| 0 (empty) | Yes ✓ | "G\n" |
| 1 | Yes ✓ | "G\n" |
| 2 | Yes ✓ | "G\n" |
| 3 or more | No ✗ | "Invalid Input\n" |

## Notice: This is the OPPOSITE of PrintIf!

- **PrintIf:** Returns "G" when length is 0 OR >= 3
- **PrintIfNot:** Returns "G" when length is < 3 (0, 1, or 2)

---
# RectPerimeter Solution

Here's the beginner-friendly solution 
```go
package piscine

func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}
	perimeter := 2*w + 2*h
	return perimeter
}
```
with line-by-line explanation:## Line-by-Line Explanation:

```go
func RectPerimeter(w, h int) int {
```
- Declares a function named `RectPerimeter`
- Takes two parameters (both are integers):
  - `w` → width of the rectangle
  - `h` → height of the rectangle
- Returns an integer (the perimeter or -1)

```go
    if w < 0 || h < 0 {
```
- Checks if EITHER width OR height is negative
- `w < 0` → is width negative?
- `||` means "OR"
- `h < 0` → is height negative?
- If either is negative, we need to return -1

```go
        return -1
```
- If width or height is negative, return -1 (invalid input)
- Function stops here

```go
    perimeter := 2*w + 2*h
```
- Calculate the perimeter using the formula
- Perimeter of rectangle = 2 × width + 2 × height
- Or: width + width + height + height (all 4 sides)
- Store the result in `perimeter` variable

```go
    return perimeter
```
- Return the calculated perimeter

## How It Works:

**Example 1: `RectPerimeter(10, 2)`**
- Width = 10, Height = 2
- Is 10 < 0? No
- Is 2 < 0? No
- Both are positive ✓
- Calculate: 2×10 + 2×2 = 20 + 4 = 24
- **Result: 24**

**Example 2: `RectPerimeter(434343, 898989)`**
- Width = 434343, Height = 898989
- Both are positive ✓
- Calculate: 2×434343 + 2×898989
- = 868686 + 1797978
- = 2666664
- **Result: 2666664**

**Example 3: `RectPerimeter(10, -2)`**
- Width = 10, Height = -2
- Is 10 < 0? No
- Is -2 < 0? **Yes** ✓
- Height is negative!
- **Result: -1** (don't calculate, just return -1)

## Visual Example:

```
Rectangle with width=10, height=2:
    
    10
  ┌────┐
2 │    │ 2
  └────┘
    10

Perimeter = 10 + 2 + 10 + 2 = 24
Or: 2×10 + 2×2 = 24
```

## Simple Logic:

1. **First:** Check if any number is negative → return -1
2. **Then:** Calculate perimeter = 2×width + 2×height
3. **Finally:** Return the result

---
# RetainFirstHalf Solution

Here's the beginner-friendly solution 
```go
package piscine

func RetainFirstHalf(str string) string {
	if str == "" {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	halfLength := len(str) / 2
	return str[:halfLength]
}
```
with line-by-line explanation:## Line-by-Line Explanation:

```go
func RetainFirstHalf(str string) string {
```
- Declares a function named `RetainFirstHalf`
- Takes one parameter: `str` which is a string
- Returns a string (the first half)

```go
    if str == "" {
```
- First, check if the string is empty
- `str == ""` compares if the string has no characters

```go
        return ""
```
- If empty, return an empty string

```go
    if len(str) == 1 {
```
- Check if the string has exactly 1 character
- `len(str)` gets the length of the string

```go
        return str
```
- If length is 1, return the whole string (can't split one character in half)

```go
    halfLength := len(str) / 2
```
- Calculate the half length of the string
- `len(str) / 2` divides the length by 2
- **Important:** In Go, dividing integers automatically rounds down
- Example: 11 / 2 = 5 (not 5.5)

```go
    return str[:halfLength]
```
- Return the first half of the string using **slicing**
- `str[:halfLength]` means "from start to halfLength"
- Takes characters from index 0 up to (but not including) halfLength

## How It Works:

**Example 1: `RetainFirstHalf("This is the 1st halfThis is the 2nd half")`**
- Length = 44 characters
- Is it empty? No
- Is length 1? No
- Calculate half: 44 / 2 = 22
- Return first 22 characters: `"This is the 1st half"`
- **Result: "This is the 1st half"**

**Example 2: `RetainFirstHalf("A")`**
- Length = 1 character
- Is it empty? No
- Is length 1? **Yes** ✓
- Return the whole string
- **Result: "A"**

**Example 3: `RetainFirstHalf("")`**
- Length = 0 characters
- Is it empty? **Yes** ✓
- Return empty string
- **Result: ""**

**Example 4: `RetainFirstHalf("Hello World")`**
- Length = 11 characters
- Is it empty? No
- Is length 1? No
- Calculate half: 11 / 2 = 5 (rounds down from 5.5)
- Return first 5 characters: `"Hello"`
- **Result: "Hello"**

## Understanding String Slicing:

```go
str := "Hello"
//     01234  (index positions)

str[:2]  // "He"    (index 0 and 1)
str[:3]  // "Hel"   (index 0, 1, and 2)
str[:5]  // "Hello" (all characters)
```

## Visual Example with "Hello World":

```
String: "Hello World"
Index:   0123456789...

Length = 11
Half = 11 / 2 = 5

str[:5] takes positions 0,1,2,3,4
Result: "Hello"
```

## Logic Summary:

1. **Empty string?** → Return empty string
2. **Length is 1?** → Return the same string
3. **Otherwise:** Calculate half (rounds down automatically), return first half using slicing

---
Thats all for Checkpoint Number 1 and 2
---

---
Lets  go through number 3 - 5 Levels
---

---
# CamelToSnakeCase Solution - will provide 2 seperate approaches for this because of checker

Here's the beginner-friendly solution 
```go
package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	
	if !isValidCamelCase(s) {
		return s
	}
	
	result := ""
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(char)
		} else {
			result += string(char)
		}
	}
	
	return result
}

func isValidCamelCase(s string) bool {
	for i, char := range s {
		if char >= '0' && char <= '9' {
			return false
		}
		
		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
			return false
		}
		
		if char >= 'A' && char <= 'Z' {
			if i == len(s)-1 {
				return false
			}
			
			if i > 0 {
				prevChar := rune(s[i-1])
				if prevChar >= 'A' && prevChar <= 'Z' {
					return false
				}
			}
		}
	}
	
	return true
}
```
with line-by-line explanation:## Line-by-Line Explanation:

### Main Function: `CamelToSnakeCase`

```go
func CamelToSnakeCase(s string) string {
```
- Function that converts camelCase to snake_case
- Takes a string `s` as input
- Returns a string

```go
    if s == "" {
        return ""
    }
```
- Check if the string is empty
- If empty, return empty string

```go
    if !isValidCamelCase(s) {
        return s
    }
```
- Call helper function to check if it's valid camelCase
- `!` means "NOT"
- If it's NOT valid camelCase, return the original string unchanged

```go
    result := ""
```
- Create an empty string to build our result

```go
    for i, char := range s {
```
- Loop through each character in the string
- `i` is the position/index
- `char` is the current character

```go
        if char >= 'A' && char <= 'Z' {
```
- Check if the character is an uppercase letter (A to Z)

```go
            if i > 0 {
                result += "_"
            }
```
- If it's uppercase AND not the first character
- Add an underscore before it
- `i > 0` means "not at position 0"

```go
            result += string(char)
```
- Add the uppercase letter to the result
- `string(char)` converts the character to a string

```go
        } else {
            result += string(char)
        }
```
- If it's a lowercase letter, just add it to the result

```go
    return result
```
- Return the final snake_case string

---

### Helper Function: `isValidCamelCase`

```go
func isValidCamelCase(s string) bool {
```
- Helper function to check if string follows camelCase rules
- Returns `true` if valid, `false` if not

```go
    for i, char := range s {
```
- Loop through each character

```go
        if char >= '0' && char <= '9' {
            return false
        }
```
- Check if character is a number (0-9)
- CamelCase can't have numbers, so return false

```go
        if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
            return false
        }
```
- Check if character is NOT a letter (neither uppercase nor lowercase)
- This catches punctuation, spaces, special characters
- Return false if found

```go
        if char >= 'A' && char <= 'Z' {
```
- If we find an uppercase letter, do more checks

```go
            if i == len(s)-1 {
                return false
            }
```
- Check if uppercase letter is at the end
- `len(s)-1` is the last position
- CamelCase can't end with uppercase, so return false

```go
            if i > 0 {
                prevChar := rune(s[i-1])
                if prevChar >= 'A' && prevChar <= 'Z' {
                    return false
                }
            }
```
- If not the first character, check the previous character
- `s[i-1]` gets the character before current one
- If previous character is also uppercase, return false
- (Two uppercase letters in a row is invalid)

```go
    return true
```
- If all checks passed, it's valid camelCase!

---

## How It Works:

**Example 1: `CamelToSnakeCase("HelloWorld")`**
- Is empty? No
- Is valid camelCase? Check:
  - No numbers ✓
  - No punctuation ✓
  - Doesn't end with uppercase (ends with 'd') ✓
  - No two uppercase in a row ✓
  - Valid! ✓
- Convert:
  - H → uppercase at position 0 → just add "H"
  - e → lowercase → add "e"
  - l → lowercase → add "l"
  - l → lowercase → add "l"
  - o → lowercase → add "o"
  - W → uppercase at position 5 → add "_W"
  - o → lowercase → add "o"
  - r → lowercase → add "r"
  - l → lowercase → add "l"
  - d → lowercase → add "d"
- **Result: "Hello_World"**

**Example 2: `CamelToSnakeCase("hey2")`**
- Is empty? No
- Is valid camelCase? Check:
  - Has number '2' ✗
  - Invalid!
- Return unchanged
- **Result: "hey2"**

**Example 3: `CamelToSnakeCase("CAMELtoSnackCASE")`**
- Is empty? No
- Is valid camelCase? Check:
  - C at position 0
  - A at position 1 (previous is C which is uppercase) ✗
  - Two uppercase in a row = Invalid!
- Return unchanged
- **Result: "CAMELtoSnackCASE"**

**Example 4: `CamelToSnakeCase("camelToSnakeCase")`**
- Valid camelCase ✓
- Convert: c-a-m-e-l-_T-o-_S-n-a-k-e-_C-a-s-e
- **Result: "camel_To_Snake_Case"**

## Visual Example:

```
Input: "helloWorld"
       01234567890

h → lowercase → "h"
e → lowercase → "he"
l → lowercase → "hel"
l → lowercase → "hell"
o → lowercase → "hello"
W → UPPERCASE at position 5 → "hello_W"
o → lowercase → "hello_Wo"
r → lowercase → "hello_Wor"
l → lowercase → "hello_Worl"
d → lowercase → "hello_World"

Output: "hello_World"
```

## Summary of Rules:

✅ **Valid camelCase:**
- Only letters (no numbers, no punctuation)
- Can't end with uppercase
- No two uppercase letters in a row

✅ **Conversion:**
- Add underscore before every uppercase letter (except first character)
- Keep all letters as they are

Now solution 2 -

```go
package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	x := func(c byte) bool {
		return c >= 'a' && c <= 'z'
	}
	y := func(c byte) bool {
		return c >= 'A' && c <= 'Z'
	}

	if s[0] < 'A' || s[0] > 'Z' && s[0] < 'a' || s[0] > 'z' {
		return s
	}
	for i := 0; i < len(s); i++ {
		if !y(s[i]) && !x(s[i]) {
			return s
		}
	}

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 && y(s[i]) {
			return s
		}
		if y(s[i]) && y(s[i+1]) {
			return s
		}
	}

	var a []byte

	for i := 0; i < len(s); i++ {
		c := s[i]
		if y(c) {
			if i > 0 {
				a = append(a, '_')
			}
		}
		a = append(a, c)
	}
	return string(a)
}
```
This function works **perfectly**! It does the same thing

## Line-by-Line Explanation:

```go
func CamelToSnakeCase(s string) string {
    if s == "" {
        return ""
    }
```
- Check if string is empty, return empty if so

```go
    x := func(c byte) bool {
        return c >= 'a' && c <= 'z'
    }
```
- Creates a **helper function** called `x`
- `x` checks if a character is **lowercase** (a-z)
- Returns `true` if lowercase, `false` if not

```go
    y := func(c byte) bool {
        return c >= 'A' && c <= 'Z'
    }
```
- Creates another **helper function** called `y`
- `y` checks if a character is **uppercase** (A-Z)
- Returns `true` if uppercase, `false` if not

---

### Validation Checks (Is it valid camelCase?)

```go
    if s[0] < 'A' || s[0] > 'Z' && s[0] < 'a' || s[0] > 'z' {
        return s
    }
```
- Checks if the **first character** is a letter
- `s[0]` is the first character
- If first character is not a letter, return unchanged

```go
    for i := 0; i < len(s); i++ {
        if !y(s[i]) && !x(s[i]) {
            return s
        }
    }
```
- Loop through **all characters**
- `!y(s[i])` → is it NOT uppercase?
- `!x(s[i])` → is it NOT lowercase?
- If a character is neither uppercase nor lowercase (like a number or punctuation), return unchanged

```go
    for i := 0; i < len(s); i++ {
        if i == len(s)-1 && y(s[i]) {
            return s
        }
```
- Check if the **last character** is uppercase
- `i == len(s)-1` → we're at the last position
- `y(s[i])` → and it's uppercase
- If string ends with uppercase, it's invalid → return unchanged

```go
        if y(s[i]) && y(s[i+1]) {
            return s
        }
    }
```
- Check if **two uppercase letters are next to each other**
- `y(s[i])` → current character is uppercase
- `y(s[i+1])` → next character is also uppercase
- If found, it's invalid → return unchanged

---

### Conversion (Build snake_case string)

```go
    var a []byte
```
- Create an empty **slice** (like an array) to store bytes
- We'll build the result character by character

```go
    for i := 0; i < len(s); i++ {
        c := s[i]
```
- Loop through each character
- Store current character in variable `c`

```go
        if y(c) {
            if i > 0 {
                a = append(a, '_')
            }
        }
```
- If character is uppercase
- AND it's not the first character (`i > 0`)
- Add an underscore `_` to our result

```go
        a = append(a, c)
    }
```
- Add the current character to the result
- `append` adds an item to the slice

```go
    return string(a)
```
- Convert the byte slice back to a string
- Return the final result

---

## Key Differences 

| Aspect |  Solution 2 | Solution 1 |
|--------|---------------|-------------|
| **Helper functions** | Uses inline functions `x` and `y` | Uses separate function `isValidCamelCase` |
| **Character checking** | Uses byte access `s[i]` | Uses `range` loop with runes |
| **Building result** | Uses byte slice `[]byte` | Uses string concatenation |
| **Style** | More compact, uses helper functions | More explicit, separate validation |

## Which is Better?

**solution 2:**
- ✅ More **efficient** (byte slice is faster than string concatenation)
- ✅ More **compact** code
- ✅ Uses inline helper functions (clever!)
- ⚠️ Slightly harder to read for beginners

**solution 1:**
- ✅ More **explicit** and easier to understand
- ✅ Separate validation function is clearer
- ⚠️ Slightly less efficient (string concatenation)

**Both are 100% correct!** This solution is actually **more optimized** because:
1. Using `[]byte` is faster than string concatenation
2. Accessing bytes directly with `s[i]` is faster than `range` with runes
---
✅ **Checkpoint category 3 question: `DigitLen`**

### 🧩 **Function**

```go
package piscine

func DigitLen(n, base int) int {
	if base < 2 || base > 36 { // base must be valid
		return -1
	}
	if n < 0 { // if negative, make it positive
		n = -n
	}
	count := 0
	for {
		count++
		n = n / base
		if n == 0 { // stop when it reaches zero
			break
		}
	}
	return count
}
```

### ⚙️ **Explanation (short)**

* Checks if base is valid (2–36), otherwise returns **-1**.
* Converts `n` to positive if needed.
* Repeatedly divides `n` by `base`, counting how many times until `n` becomes **0**.

✅ **Output**

```
3
7
2
-1
```
✅ **Checkpoint category 3 question: `FirstWord`**

### 🧩 **Function**

```go
package piscine

func FirstWord(s string) string {
	word := ""
	for _, r := range s {
		if r == ' ' { // stop when first space found
			break
		}
		word += string(r)
	}
	return word + "\n"
}
```

### ⚙️ **Explanation (short)**

* Loops through each character.
* Builds `word` until a space `' '` is found.
* Returns the first word followed by a newline `\n`.

✅ **Output**

```
hello

hello
```
✅ **Checkpoint: `FishAndChips`**

### 🧩 **Function**

```go
package piscine

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}
	if n%2 == 0 {
		return "fish"
	}
	if n%3 == 0 {
		return "chips"
	}
	return "error: non divisible"
}
```

---

### 🪄 **Line-by-line Explanation**

1. **`if n < 0 { ... }`**
   → Checks if the number is negative.
   → If true, returns `"error: number is negative"` immediately.

2. **`if n%2 == 0 && n%3 == 0 { ... }`**
   → `%` means *modulus*, it checks for remainder.
   → If divisible by both 2 and 3 (remainder 0 in both), return `"fish and chips"`.

3. **`if n%2 == 0 { ... }`**
   → If divisible only by 2, return `"fish"`.

4. **`if n%3 == 0 { ... }`**
   → If divisible only by 3, return `"chips"`.

5. **`return "error: non divisible"`**
   → If none of the above are true, it means not divisible by 2 or 3.

---

### 🧾 **Example Output**

```
fish
chips
fish and chips
```
✅ **Checkpoint category 3: `Gcd`**

### 🧩 **Function Solution**

```go
package piscine

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```

---

### 🪄 **Line-by-Line Explanation**

1. **`if a == 0 || b == 0 { return 0 }`**
   → If either number is zero, there’s no valid GCD, so return `0` immediately.

2. **`for b != 0 { ... }`**
   → This loop keeps running while `b` is not zero.
   → It uses **Euclid’s algorithm** — a famous method to find the GCD.

3. **`a, b = b, a%b`**
   → Swap values:

   * The new `a` becomes the old `b`.
   * The new `b` becomes the remainder of `a % b`.
     → This gradually reduces the numbers until `b` becomes 0.

4. **`return a`**
   → When the loop ends, `a` holds the **greatest common divisor**.

---

### 🧾 **Example Output**

```
2
6
7
1
```

---

### 🧠 **Quick Tip**

* `%` means *remainder after division*.
* Euclid’s rule:

  ```
  GCD(a, b) = GCD(b, a % b)
  until b = 0
  ```
---
✅ Excellent — this one’s a fun one! Let’s write the `HashCode` function step-by-step and then show you the full working Go code.

---

### 🧠 Logic Breakdown

For a string `dec`, and each character `r` in it:

1. Get the ASCII value:

   ```go
   ascii := int(r)
   ```
2. Add the **length** of the string:

   ```go
   hashed := (ascii + len(dec)) % 127
   ```
3. Check if the new value is *printable* (ASCII 33–126):
   If not, add 33 to make it printable.

   ```go
   if hashed < 33 {
       hashed += 33
   }
   ```
4. Convert back to a rune/character and append to the result string.

---

### 🧩 Full Code

```go
package piscine

func HashCode(dec string) string {
	if len(dec) == 0 {
		return ""
	}

	size := len(dec)
	var result string

	for _, r := range dec {
		hashed := (int(r) + size) % 127
		if hashed < 33 {
			hashed += 33
		}
		result += string(rune(hashed))
	}

	return result
}
```

---

### 🧾 Output

```
B
CD
EDF
Spwwz+bz}wo
```

---
## `LastWord` 

---

### 🧠 Logic Breakdown

We need the **last word** in a given string `s`, where:

* A **word** is separated by **spaces**.
* Ignore **extra spaces** at the beginning or end.
* Return it followed by a newline (`\n`).

---

### 🪜 Steps

1. Trim trailing spaces manually or with logic (not with strings.Trim if not allowed).
2. Move from the **end of the string** backward:

   * Skip spaces until a character is found.
   * Then, collect characters until another space (or start of string) is reached.
3. Reverse the collected runes (since we collected backward).
4. Add `\n` at the end.

---

### ✅ Full Solution

```go
package piscine

func LastWord(s string) string {
	runes := []rune(s)
	end := len(runes) - 1

	// Step 1: skip trailing spaces
	for end >= 0 && runes[end] == ' ' {
		end--
	}

	if end < 0 {
		return "\n"
	}

	// Step 2: find the start of the last word
	start := end
	for start >= 0 && runes[start] != ' ' {
		start--
	}

	// Step 3: extract the word
	word := string(runes[start+1 : end+1])
	return word + "\n"
}
```

---

### 🧾 Expected Output

```
not
lorem,ipsum

```

*(Last line is just an empty line for the string with only spaces)*

---

---
**You're all set!**  
Good luck with your checkpoint — you're going to **ace** it!
