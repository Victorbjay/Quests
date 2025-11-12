
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
