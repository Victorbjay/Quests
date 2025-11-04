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

## 🧩 Step 1: Understand the Task

We are told to create a Go program inside a folder named `point`, and to make this code (given by the school platform) **work correctly**.

The **expected output** when we run it is:

```
x = 42, y = 21
```

But if we copy the given code directly, it won’t run yet — because:

* The type `point` is not defined.
* The `fmt` package is not imported.

---

## 🧠 Step 2: What’s Happening in the Code

Let’s look at what it’s trying to do:

```go
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}
```

Here:

* `setPoint` is a function that receives a pointer (`*point`).
* It sets the `x` and `y` values **inside that point** to `42` and `21`.

So we need to define what a `point` is.

---

## 🧱 Step 3: Define the Structure

We define a `struct` (structure) to hold two integers — `x` and `y`.

```go
type point struct {
	x int
	y int
}
```

This means a `point` looks like:

```
+--------+
|  x: ?  |
|  y: ?  |
+--------+
```

---

## 🧭 Step 4: Full Working Code

📄 **point/main.go**

```go
package main

import "fmt"

// Step 1: Define the 'point' structure
type point struct {
	x int
	y int
}

// Step 2: Function to set the values of a point
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// Step 3: The main function
func main() {
	// Create a new point in memory and get its pointer
	points := &point{}

	// Call the function that sets the x and y values
	setPoint(points)

	// Print the result
	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
```

---

## 🧠 Step 5: Why Each Line Matters 

| Line                                 | What it does                                     | Why it’s needed                                       |
| ------------------------------------ | ------------------------------------------------ | ----------------------------------------------------- |
| `package main`                       | Tells Go this is the main executable file.       | Every Go program starts with it.                      |
| `import "fmt"`                       | We need `fmt` to print things.                   | Without it, `fmt.Printf` won’t work.                  |
| `type point struct { x int; y int }` | Defines a box with two integer slots (`x`, `y`). | We need this “blueprint” for our points.              |
| `func setPoint(ptr *point)`          | Function takes a *pointer* to a point.           | The star `*` means “give me the address, not a copy.” |
| `ptr.x = 42; ptr.y = 21`             | Changes values inside the real point.            | Because `ptr` points to the original memory.          |
| `points := &point{}`                 | Creates a new point and gets its address.        | We need a pointer to pass into the function.          |
| `fmt.Printf(...)`                    | Displays values on screen.                       | To show `x = 42, y = 21`.                             |

---

## 🧩 Step 6: Visual Diagram

```
      ┌────────────┐
      │   points   │
      │ (pointer)  │────────────┐
      └────────────┘            │
                                ▼
                       ┌───────────────────┐
                       │   point struct    │
                       │  x: 42            │
                       │  y: 21            │
                       └───────────────────┘
```

✅ `setPoint(points)` changes the values **directly** inside that box.

---

## 🧪 Step 7: How to Run It

1. Open your terminal.
2. Create and enter the folder:

   ```bash
   mkdir point
   cd point
   ```
3. Create `main.go` and paste the full code above.
4. Run:

   ```bash
   go run .
   ```

✅ Expected Output:

```
x = 42, y = 21
```

---


