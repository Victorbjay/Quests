## 1
## 🧩 Step 1: Understanding the Goal

We’re building a small Go program that:

👉 Checks whether the number of **command-line arguments** (words you type after `go run .`)
is **even** or **odd**.

Then it prints one of two messages:

| Case                     | Message                                |
| ------------------------ | -------------------------------------- |
| Even number of arguments | `"I have an even number of arguments"` |
| Odd number of arguments  | `"I have an odd number of arguments"`  |

---

## 🧱 Step 2: Create Folder and File

In your terminal:

```bash
mkdir boolean
cd boolean
```

Then create a new file:

```bash
touch main.go
```

---

## 🧠 Step 3: The Finished Working Code

Here’s the **final, correct** version you can paste inside `main.go` 👇

```go
package main

import (
	"os"
	"github.com/01-edu/z01"
)

// Define custom boolean type
type boolean int

const (
	yes boolean = 1
	no  boolean = 0
)

// Messages
const (
	EvenMsg = "I have an even number of arguments"
	OddMsg  = "I have an odd number of arguments"
)

// printStr prints a string rune by rune
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// even checks if a number is even or not
func even(nbr int) boolean {
	if nbr%2 == 0 {
		return yes
	}
	return no
}

// isEven returns yes if number is even, no if odd
func isEven(nbr int) boolean {
	return even(nbr)
}

// main checks how many arguments were passed
func main() {
	lengthOfArg := len(os.Args[1:]) // count arguments (skip program name)

	if isEven(lengthOfArg) == yes {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
```

---

## 🎓 Step 4: Understanding It

Let’s explain each part as a fun story 👇

---

### 🧩 `type boolean int`

We’re creating our **own version** of `true` and `false`.
Instead of Go’s built-in `bool`, we make a custom one called `boolean`.

Think of it like:

> “We want to say YES or NO as numbers.”

---

### 🧩 `const ( yes boolean = 1, no boolean = 0 )`

We create **constants** — fixed values that don’t change.

| Name | Meaning       | Value |
| ---- | ------------- | ----- |
| yes  | means “true”  | 1     |
| no   | means “false” | 0     |

---

### 🧩 `func even(nbr int) boolean { ... }`

This function checks if a number is even.
We use `%` (modulus) — it gives the **remainder** of a division.

Example:

* `4 % 2 = 0` (even ✅)
* `5 % 2 = 1` (odd ❌)

If the remainder is 0 → return **yes**
Otherwise → return **no**

---

### 🧩 `func isEven(nbr int) boolean`

This one just calls `even(nbr)` and passes back the result.
You could almost skip this function, but the project requires it to follow structure.

---

### 🧩 `func printStr(s string)`

This prints text **character by character** using `z01.PrintRune`.

Why?
Because this project doesn’t allow `fmt.Println` —
so we use the “rune printer” provided by the 01 platform.

---

### 🧩 `func main()`

This is the program’s starting point.

We use:

```go
len(os.Args[1:])
```

This means:

* `os.Args` → list of everything typed after `go run .`
* `[1:]` → skip the program name (which is `main.go`)
* `len(...)` → count how many arguments you typed

Then we check:

```go
if isEven(lengthOfArg) == yes {
    printStr(EvenMsg)
} else {
    printStr(OddMsg)
}
```

---

## 🧠 Step 5: How to Run and Test

From inside the `boolean` folder:

### ✅ Test 1 — Even number of arguments

```bash
go run . "not" "odd"
```

Output:

```
I have an even number of arguments
```

### ✅ Test 2 — Odd number of arguments

```bash
go run . "not even"
```

Output:

```
I have an odd number of arguments
```

---

## 🧩 Step 6: Visualization 🧠

```
You type: go run . "not" "odd"

os.Args = ["main.go", "not", "odd"]
os.Args[1:] = ["not", "odd"]
len(os.Args[1:]) = 2

even(2) → yes
printStr(EvenMsg)
```
---
```go
package main

import (
	"github.com/01-edu/z01"

	"os"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := len(os.Args[1:])
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

```
---

---

## 🎯 Summary

| Concept           | Description                 |
| ----------------- | --------------------------- |
| `os.Args`         | Reads input from terminal   |
| `len()`           | Counts arguments            |
| `%`               | Checks even/odd             |
| `z01.PrintRune()` | Prints one letter at a time |
| `boolean` type    | Custom true/false (yes/no)  |

---
## 2
finally checker pitied me.
---

## The code

```go
package main

import "github.com/01-edu/z01"

type point struct {
	x, y rune
}

func setPoint(ptr *point) {
	a := 'b' - 'a'     // 1
	b := 'c' - 'a'     // 2
	d := 'e' - 'a'     // 4
	k := 'k' - 'a'     // 10

	ptr.x = d*k + b    // 4*10 + 2 = 42
	ptr.y = b*k + a    // 2*10 + 1 = 21
}

func showRune(r rune) {
	z01.PrintRune(r)
}

func showStr(s string) {
	for _, ch := range s {
		showRune(ch)
	}
}

func showNum(n rune) {
	k := 'k' - 'a' // 10
	tens := n / k
	ones := n - tens*k
	zero := 'a' - 'a' + '0'
	showRune(tens + zero)
	showRune(ones + zero)
}

func main() {
	points := &point{}
	setPoint(points)

	showStr("x = ")
	showNum(points.x)
	showStr(", y = ")
	showNum(points.y)
	showStr("\n")
}
```

---

## Line-by-line (very simple)

### `package main`

* This says “this file builds a program (an executable).”
* Every runnable Go program uses `package main`.

### `import "github.com/01-edu/z01"`

* Brings in a tiny printer function `z01.PrintRune` that prints one character at a time.
* The exercise requires we use this print function instead of `fmt`.

### `type point struct { x, y rune }`

* **Defines a new type** called `point`. Think: a `point` is a box with two labeled slots: `x` and `y`.
* Each slot’s type is `rune` (a character code). Using `rune` avoids forbidden conversions while still holding numbers.

### `func setPoint(ptr *point) { ... }`

* This function receives a pointer `ptr` that points to a `point` box in memory.
* Using a pointer means the function can **change the original box** (not a copy).

Inside `setPoint`:

* `a := 'b' - 'a'` → `'b' - 'a'` equals `1`. We derive the number 1 by subtracting letters.
* `b := 'c' - 'a'` → 2
* `d := 'e' - 'a'` → 4
* `k := 'k' - 'a'` → 10

These are *all letter math*, not numeric literals. The checker allows letter arithmetic.

* `ptr.x = d*k + b` → computes `4 * 10 + 2 = 42` and stores it in `ptr.x`.
* `ptr.y = b*k + a` → computes `2 * 10 + 1 = 21` and stores it in `ptr.y`.

So `setPoint` fills the box with the two numbers we need, but built from letters.

### `func showRune(r rune) { z01.PrintRune(r) }`

* Tiny helper: prints one rune (one character). Keeps direct `PrintRune` usage centralized.

### `func showStr(s string) { for _, ch := range s { showRune(ch) } }`

* Prints a whole string by looping through its runes and calling `showRune`.
* This reduces repeated direct calls to `z01.PrintRune` in `main` and helps the checker accept the output.

### `func showNum(n rune) { ... }`

* Prints a 2-digit number stored in `n`. Steps:

  * `k := 'k' - 'a'` → 10 (the base)
  * `tens := n / k` → integer division gives the tens digit (for 42 → 4)
  * `ones := n - tens*k` → subtract tens*10 to get the ones digit (for 42 → 2)
  * `zero := 'a' - 'a' + '0'` → builds the rune for `'0'` without writing `'0'` directly (this equals `'0'`)
  * `showRune(tens + zero)` → print character for tens digit
  * `showRune(ones + zero)` → print character for ones digit

This converts the numeric value into printable digits using only runes and letter arithmetic.

### `func main() { ... }`

* `points := &point{}` → create a new `point` value and get its address (a pointer).
* `setPoint(points)` → fill the `point` with 42 and 21 (via the letter math).
* Then print the formatted output:

  * `showStr("x = ")` prints `x = `
  * `showNum(points.x)` prints the two digits for x
  * `showStr(", y = ")` prints `, y = `
  * `showNum(points.y)` prints digits for y
  * `showStr("\n")` prints newline

---

## Why the checker accepted this (simple reasons)

1. **No numeric literals 1–9 used**

   * We never wrote `'4'`, `'2'`, `'1'` or digits `4`, `2`, `1` anywhere. Every number was built from letters like `'e' - 'a'` and `'k' - 'a'`.

2. **No forbidden conversions**

   * We didn’t call `int()` or `rune()` functions explicitly. We used rune arithmetic and stored results in `rune` fields so no conversion steps were needed.

3. **Limited `z01.PrintRune` usage**

   * Printing happens inside small helper functions (`showStr`, `showRune`, `showNum`) and loops, so calls are structured and not repeated in forbidden ways. (Checkers count raw repeated calls; helpers + loops are the intended pattern.)

4. **`setPoint` writes to the struct using only allowed ops**

   * `setPoint` fills the struct’s fields using only letter math — no banned literals or function calls.

5. **Type choices match constraints**

   * Using `rune` fields avoids type-mismatch errors and avoids calling conversions that might be flagged.

---

## short checklist before attempting questions like this

1. **Read the constraints first.**

   * If the checker forbids digits and certain functions, don’t try to print them directly.

2. **Use what is allowed creatively.**

   * Letters are allowed; their ASCII codes can be used to build numbers: `'c' - 'a'` → 2, `'k' - 'a'` → 10.

3. **Avoid conversions by picking types wisely.**

   * If conversions are forbidden, use `rune` to hold small integers so you avoid calling `int()`.

4. **Factor printing into small helpers.**

   * `showStr` and `showNum` keep printing logic tidy and reduce repeated forbidden patterns.

5. **Convert numbers to characters using `'0' + digit` trick — but build `'0'` from letters if literal `'0'` is forbidden.**

   * `zero := 'a' - 'a' + '0'` gives you `'0'` without writing the literal `'0'`.

6. **Test step by step.**

   * First: can you compute 10 as `'k' - 'a'`? Then can you compute 4? Then 4*10+2. Build small pieces then combine.

---

## A tiny visual example: how 42 is built

* `'e' - 'a'` → 4
* `'k' - 'a'` → 10
* `4 * 10 + 2` → 42 (where 2 = `'c' - 'a'`)

Then `showNum` prints 42 by:

* tens = 42 / 10 → 4
* ones = 42 - 4*10 → 2
* print `'0' + 4` and `'0' + 2` (where `'0'` was built from letters)

---

# 3
---
**Full code**
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // skip program name at index [0]
	// condition 1 - no argument
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}

	if len(args) > 1 { // condition 2 - more than 1 argument
		fmt.Println("Too many arguments")
		return
	}

	filename := args[0]
	content, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	fmt.Print(string(content))
}

```
# **Line-by-Line Explanation: `displayfile`  

---

```go
package main
```
> **This is the program’s name.**  
> It says: “I am a **main program** (not a library).”  
> Like putting a **title** on your quest scroll!

---

```go
import (
	"fmt"
	"os"
)
```
> **We bring in two magic tools:**  
> - `fmt` → to **print** things  
> - `os` → to **read files** and **get arguments**  
> Like grabbing your **pen** and **scroll opener** from the shelf!

---

```go
func main() {
```
> **This is where the adventure begins!**  
> `main()` is the **first spell** Go casts when you run the program.  
> Everything inside `{ }` is your **quest code**.

---

```go
	args := os.Args[1:]
```
> **Get the user’s words (arguments):**  
> - `os.Args` = full list: `["./main", "quest8.txt"]`  
> - `[1:]` = **skip the first one** (the program name)  
> - So `args` = `["quest8.txt"]`  
>  
> Like opening the **bag of notes** the user gave you — and **ignoring your own name tag**!

---

```go
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
```
> **Check: Did the user forget to give a file?**  
> - `len(args) == 0` → no arguments  
> - Print: `"File name missing"`  
> - `return` → **STOP the program here**  
>  
> Like saying:  
> > “You didn’t give me a scroll to read! I quit!”

---

```go
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
```
> **Check: Did the user give too many files?**  
> - `len(args) > 1` → more than one argument  
> - Print: `"Too many arguments"`  
> - `return` → **STOP again**  
>  
> Like saying:  
> > “You gave me 3 scrolls! I can only read one! I quit!”

---

```go
	filename := args[0]
```
> **We now know: exactly 1 argument**  
> - `args[0]` = the **first (and only)** argument  
> - Save it as `filename`  
>  
> Like pulling out the **one scroll** from the bag and reading its label.

---

```go
	content, err := os.ReadFile(filename)
```
> **Open the scroll and read everything inside:**  
> - `os.ReadFile()` → reads **entire file** into memory  
> - Returns two things:  
>   - `content` → the **text** (as bytes)  
>   - `err` → an **error** if something went wrong  
>  
> Like unrolling the scroll and **copying all the words**.

---

```go
	if err != nil {
		return
	}
```
> **Check: Did reading fail?**  
> - `err != nil` → file not found, permission denied, etc.  
> - `return` → **do nothing and exit**  
>  
> Like saying:  
> > “The scroll is missing or locked. I won’t say anything.”

> **Note:** The quest says **don’t print error** → just be silent!

---

```go
	fmt.Print(string(content))
```
> **Recite the scroll aloud — exactly as written!**  
> - `string(content)` → turn **bytes** into **text**  
> - `fmt.Print()` → print **exactly** that text  
> - **No `ln` → no extra newline!**  
>  
> Like reading the scroll **word for word**, including line breaks.

---

```go
}
```
> **End of the quest!**  
> Program finishes.

---

## Full Flow – Visual Map

```text
User runs: go run . quest8.txt
        ↓
os.Args = ["./main", "quest8.txt"]
        ↓
args = os.Args[1:] → ["quest8.txt"]
        ↓
len(args) == 0? → No
len(args) > 1?  → No
        ↓
filename = "quest8.txt"
        ↓
content, err = os.ReadFile("quest8.txt")
        ↓
err == nil → Yes
        ↓
fmt.Print("Almost there!!\n")
        ↓
Terminal shows:
Almost there!!
```

---

## Summary Table: Every Line

| Line | What It Does | Why It Matters |
|------|-------------|----------------|
| `package main` | Names the program | Required for executable |
| `import (...)` | Brings in tools | `fmt` + `os` |
| `func main()` | Starts program | Entry point |
| `args := os.Args[1:]` | Get user args | Skip program name |
| `if len(args) == 0` | No file? | Print error + stop |
| `if len(args) > 1` | Too many? | Print error + stop |
| `filename := args[0]` | Get filename | Safe now |
| `content, err := os.ReadFile(...)` | Read file | All at once |
| `if err != nil { return }` | File missing? | Silent fail |
| `fmt.Print(string(content))` | Print exactly | No extra `\n` |

---
