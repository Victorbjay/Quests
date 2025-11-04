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

---

The exercise says:

> Allowed functions: `github.com/01-edu/z01.PrintRune`, `--no-lit=[1-9]`

That means:

* ❌ We **cannot** use `fmt`.
* ✅ We **must use `z01.PrintRune()`** to print everything **character by character**, even numbers.

So… we’ll **rebuild the entire program manually** — step by step like a beginner — and make it work **without fmt**, using only `z01.PrintRune`.

---

## 🧩 Step 1: What we must achieve

Expected output:

```
x = 42, y = 21
```

We need to print that **character by character**:

```
x, space, =, space, 4, 2, ,, space, y, space, =, space, 2, 1, newline
```

---

## 🧱 Step 2: The structure we’ll use

We’ll keep our struct and logic the same — only change the printing part.

---

## ✅ Final Code 
📁 **point/main.go**

```go
package main

import "github.com/01-edu/z01"

// Step 1: Define a struct called 'point'
type point struct {
	x int
	y int
}

// Step 2: Function to set the values of a point
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// Step 3: Function to print a string character by character
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// Step 4: Function to print an integer digit by digit
func printNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+(n%10)))
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

// Step 5: Main function
func main() {
	points := &point{}
	setPoint(points)

	printStr("x = ")
	printNbr(points.x)
	printStr(", y = ")
	printNbr(points.y)
	z01.PrintRune('\n')
}
```

---

## 🧠 Step 3: (why we do each thing)

| Code                                 | Why we need it                                           |
| ------------------------------------ | -------------------------------------------------------- |
| `package main`                       | Every Go executable starts with this.                    |
| `import "github.com/01-edu/z01"`     | We’re told we can only print using this.                 |
| `type point struct { x int; y int }` | Defines our “box” that holds x and y numbers.            |
| `setPoint(ptr *point)`               | Changes the inside of that box to 42 and 21.             |
| `printStr(s string)`                 | Prints a whole word one character at a time.             |
| `printNbr(n int)`                    | Prints numbers by breaking them into digits.             |
| `points := &point{}`                 | Creates an empty box in memory and gives us its address. |
| `setPoint(points)`                   | Fills the box (x=42, y=21).                              |
| `printStr("x = ") ...`               | Prints the output exactly how we want it.                |

---

## 🧩 Step 4: Visualization

```
     ┌──────────────────┐
     │  points (pointer)│────┐
     └──────────────────┘    │
                             ▼
                    ┌────────────────┐
                    │ point struct   │
                    │ x = 42         │
                    │ y = 21         │
                    └────────────────┘
```

Then `printStr` and `printNbr` print everything manually:

```
x = 42, y = 21
```

---

## 🧪 Step 5: Run the Program

1. Create and go into the folder:

   ```bash
   mkdir point
   cd point
   ```
2. Create the file:

   ```bash
   nano main.go
   ```
3. Paste the code above.
4. Run:

   ```bash
   go run .
   ```

✅ Expected Output:

```
x = 42, y = 21
```

---

