
# 🧩 Project: `quadchecker` — Beginner Guide

---

## 🧠 Step 1: What It Does (In Plain English)

The **`quadchecker`** program reads a **shape** (the output of one of your quad functions: A, B, C, D, E).
It then **analyzes** that shape and says:

> “Hey, this shape looks like quadA (3×3)”

If it looks like more than one shape (for example, C, D, and E sometimes look the same for very small sizes),
it shows all the matches in alphabetical order.

If it doesn’t match any → it says **“Not a quad function”**.

---

## 🧱 Step 2: How Data Moves (Visual Diagram)

Let’s visualize how data flows when you test the program.

---

### 🧩 1️⃣ You *pipe* the shape into quadchecker

```
+---------+        pipe ( | )       +---------------+
| quadA   |  ---->  sends output →  |  quadchecker  |
| (3, 3)  |                         | (reads input) |
+---------+                         +---------------+
```

You type in the terminal:

```bash
./quadA 3 3 | go run .
```

---

### 🧩 2️⃣ `quadchecker` reads the shape

```
o--o
|  |
o--o
```

---

### 🧩 3️⃣ It counts:

```
Width (x)  = 4 characters
Height (y) = 3 lines
```

---

### 🧩 4️⃣ It generates all 5 shapes internally

and compares them to the one it received:

```
quadA(4,3) ? matches ✅
quadB(4,3) ? no
quadC(4,3) ? no
quadD(4,3) ? no
quadE(4,3) ? no
```

---

### 🧩 5️⃣ It prints:

```
[quadA] [4] [3]
```

---

## 🧠 Step 3: How To Setup in VS Code

We’ll use a Go project structure that matches what the checker expects.

Open your terminal in VS Code and type:

```bash
git clone your repo
cd quadchecker
go mod init quadchecker
```

Now create this file:

```
quadchecker/
 ├── go.mod
 └── main.go
```

Open `main.go` and paste this full code 👇

---

## ✅ Final Code (`main.go`)

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func quadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 || row == y {
				if col == 1 || col == x {
					res.WriteRune('o')
				} else {
					res.WriteRune('-')
				}
			} else {
				if col == 1 || col == x {
					res.WriteRune('|')
				} else {
					res.WriteRune(' ')
				}
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}

func quadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				res.WriteRune('/')
			} else if row == 1 && col == x {
				res.WriteRune('\\')
			} else if row == y && col == 1 {
				res.WriteRune('\\')
			} else if row == y && col == x {
				res.WriteRune('/')
			} else if row == 1 || row == y || col == 1 || col == x {
				res.WriteRune('*')
			} else {
				res.WriteRune(' ')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}

func quadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && (col == 1 || col == x) {
				res.WriteRune('A')
			} else if row == y && (col == 1 || col == x) {
				res.WriteRune('C')
			} else if row == 1 || row == y || col == 1 || col == x {
				res.WriteRune('B')
			} else {
				res.WriteRune(' ')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}

func quadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 && col == 1) || (row == y && col == 1) {
				res.WriteRune('A')
			} else if (row == 1 && col == x) || (row == y && col == x) {
				res.WriteRune('C')
			} else if row == 1 || row == y || col == 1 || col == x {
				res.WriteRune('B')
			} else {
				res.WriteRune(' ')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}

func quadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && (col == 1 || col == x) {
				res.WriteRune('A')
			} else if row == y && (col == 1 || col == x) {
				res.WriteRune('C')
			} else if row == 1 || row == y {
				res.WriteRune('B')
			} else if col == 1 || col == x {
				res.WriteRune('B')
			} else {
				res.WriteRune(' ')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}

func main() {
	// Step 1: Read from stdin
	reader := bufio.NewReader(os.Stdin)
	input, _ := io.ReadAll(reader)
	content := string(input)

	if strings.TrimSpace(content) == "" {
		fmt.Println("Not a quad function")
		return
	}

	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	y := len(lines)
	x := len([]rune(lines[0]))

	matches := []string{}

	// Step 2: Compare input to each quad
	if content == quadA(x, y) {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", x, y))
	}
	if content == quadB(x, y) {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", x, y))
	}
	if content == quadC(x, y) {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", x, y))
	}
	if content == quadD(x, y) {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", x, y))
	}
	if content == quadE(x, y) {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", x, y))
	}

	// Step 3: Print result
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}
```
---
Here’s the correct setup — step by step — for your folder and Go module.

---

### 🧱 STEP 1: Create the folder

In your terminal:

```bash
gitclone *repoURL.git*
cd quadchecker
```

---

### ⚙️ STEP 2: Initialize your Go module

You should **use `go mod init`**
The command should look like this:

```bash
go mod init quadchecker
```

✅ This creates a `go.mod` file inside your `quadchecker` directory — it tells Go “this is a module” and allows you to run and import things properly.

---

### 📄 STEP 3: Create your main file

In the same folder:

```bash
touch main.go
```

Now you’ll have:

```
quadchecker/
├── go.mod
└── main.go
```

---

### 🧰 STEP 4: Write your code

Open the folder in VS Code:

```bash
code .
```

Then open `main.go` and paste your code there.

---

## 🧩 1. Overview — How the Project Works

When you run:

```bash
./quadA 3 3 | go run .
```

this means:

```
[quadA output] → [pipe →] → [quadchecker reads from stdin]
```

Your `quadchecker` doesn’t take file names — it just reads the shape text and decides which `quadX` it matches.

---

## 🧱 2. Folder Structure

You should have this layout:

```
quadchecker/
├── go.mod
├── main.go
├── quadA.go
│ 
├── quadB.go
│  
├── quadC.go
│  
├── quadD.go
│   
└── quadE.go
```

---

## 🚀 3. Quad Programs (Generators)

Each program takes **two integers** (`width`, `height`) and prints a rectangle using specific corner and border rules.

---

### 🅰️ `quadA.go`

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	w, _ := strconv.Atoi(os.Args[1])
	h, _ := strconv.Atoi(os.Args[2])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if (y == 0 && x == 0) || (y == 0 && x == w-1) {
				fmt.Print("o")
			} else if (y == h-1 && x == 0) || (y == h-1 && x == w-1) {
				fmt.Print("o")
			} else if y == 0 || y == h-1 {
				fmt.Print("-")
			} else if x == 0 || x == w-1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```

---

### 🅱️ `quadB.go`

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	w, _ := strconv.Atoi(os.Args[1])
	h, _ := strconv.Atoi(os.Args[2])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == 0 && (x == 0 || x == w-1) {
				fmt.Print("/")
			} else if y == h-1 && (x == 0 || x == w-1) {
				fmt.Print("\\")
			} else if y == 0 || y == h-1 {
				fmt.Print("*")
			} else if x == 0 || x == w-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```

---

### 🅲 `quadC.go`

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	w, _ := strconv.Atoi(os.Args[1])
	h, _ := strconv.Atoi(os.Args[2])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == 0 && x == 0 {
				fmt.Print("A")
			} else if y == 0 && x == w-1 {
				fmt.Print("A")
			} else if y == h-1 && x == 0 {
				fmt.Print("C")
			} else if y == h-1 && x == w-1 {
				fmt.Print("C")
			} else if y == 0 || y == h-1 {
				fmt.Print("B")
			} else if x == 0 || x == w-1 {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```

---

### 🅳 `quadD.go`

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	w, _ := strconv.Atoi(os.Args[1])
	h, _ := strconv.Atoi(os.Args[2])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == 0 && x == 0 {
				fmt.Print("A")
			} else if y == 0 && x == w-1 {
				fmt.Print("C")
			} else if y == h-1 && x == 0 {
				fmt.Print("A")
			} else if y == h-1 && x == w-1 {
				fmt.Print("C")
			} else if y == 0 || y == h-1 {
				fmt.Print("B")
			} else if x == 0 || x == w-1 {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```

---

### 🅴 `quadE.go`

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	w, _ := strconv.Atoi(os.Args[1])
	h, _ := strconv.Atoi(os.Args[2])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == 0 && x == 0 {
				fmt.Print("A")
			} else if y == 0 && x == w-1 {
				fmt.Print("C")
			} else if y == h-1 && x == 0 {
				fmt.Print("C")
			} else if y == h-1 && x == w-1 {
				fmt.Print("A")
			} else if y == 0 || y == h-1 {
				fmt.Print("B")
			} else if x == 0 || x == w-1 {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```

---

## 🧪 4. Build & Test

Inside each quad folder, run:

```bash
go build -o quadA quadA.go
```

Repeat for each (`quadB`, `quadC`, etc.), changing the output name.

Now test one:

```bash
./quadA 3 3
```

Expected:

```
o-o
| |
o-o
```

Then try:

```bash
./quadA 3 3 | go run .
```

Once your `quadchecker` logic is complete, it will detect which quad it matches.
`go build -o quadchecker main.go` -  this creates the quadchecker executable file expected in the output so now you can go ahead and delete all quadA.go - E files.

---

## 🧭 5. VS Code Quick Setup

1. Open the parent folder `quadchecker` in VS Code.
2. Make sure each subfolder has its own `main.go`.
3. Run builds in the **VS Code terminal** using the commands above.
4. To test easily:

   * In VS Code, open a terminal and run:

     ```bash
     ./quadC 1 2 | go run .
     ```


---

# ✅ **MEMORIZE THIS 2-MINUTE EXPLANATION**

---

## 🔹 **What the program does (1 sentence)**

> “The program reads the quad drawing from stdin, figures out its width and height, generates all five quad types of the same size, compares them, and prints which one matches.”

---

# ✅ COMPLETE LINE-BY-LINE BREAKDOWN OF YOUR QUADCHECKER PROGRAM

Let me explain **every single word and symbol** so you understand exactly what's happening!

---

## 🔷 1️⃣ THE PACKAGE DECLARATION

```go
package main
```

**What it means:**
- `package` = a keyword that says "this code belongs to a group"
- `main` = the special package name that tells Go "this is a runnable program, not just a library"
- **Think of it like:** Putting a label on a folder saying "This is the main application"

---

## 🔷 2️⃣ THE IMPORTS

```go
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)
```

**What each import does:**

- `"bufio"` = **Buffered Input/Output** - helps read data efficiently (like reading text line by line)
- `"fmt"` = **Format** - used for printing output to the screen (`fmt.Println`)
- `"io"` = **Input/Output** - basic tools for reading and writing data
- `"os"` = **Operating System** - lets you interact with the computer (like reading from keyboard input)
- `"strings"` = **String tools** - helps manipulate text (like joining, splitting, trimming)

**The parentheses `( )`** = lets you import multiple packages at once instead of writing `import "bufio"` five times

---

## 🔷 3️⃣ FUNCTION: `quadA`

```go
func quadA(x, y int) string {
```

**Breaking it down:**
- `func` = keyword that means "I'm defining a function"
- `quadA` = the name of this function
- `(x, y int)` = this function takes **two inputs**, both are integers (whole numbers)
  - `x` = width (number of columns)
  - `y` = height (number of rows)
- `string` = this function will **return text** as output
- `{` = opening brace, the function's code starts here

---

```go
if x <= 0 || y <= 0 {
    return ""
}
```

**What this checks:**
- `if` = start of a condition
- `x <= 0` = "if x is less than or equal to zero"
- `||` = **OR** operator (either condition can be true)
- `y <= 0` = "if y is less than or equal to zero"
- `return ""` = give back an empty string (nothing) and exit the function
- **Purpose:** If width or height is zero or negative, we can't draw a rectangle, so return nothing

---

```go
var res strings.Builder
```

**What this does:**
- `var` = keyword to declare a variable
- `res` = short for "result", the name we chose
- `strings.Builder` = a special tool that efficiently builds strings piece by piece
- **Think of it like:** An empty canvas where we'll paint our rectangle character by character

---

```go
for row := 1; row <= y; row++ {
```

**Breaking down the loop:**
- `for` = keyword that starts a loop (repeating code)
- `row := 1` = create a variable called `row` and set it to 1 (we start at row 1)
- `;` = separator
- `row <= y` = **condition**: keep looping while row is less than or equal to y
- `;` = separator
- `row++` = **increment**: after each loop, add 1 to row (`row = row + 1`)
- `{` = loop body starts here

**In simple terms:** "Do this code for row 1, then row 2, then row 3... until you reach row y"

---

```go
for col := 1; col <= x; col++ {
```

**Same structure as above but for columns:**
- This is a **nested loop** (a loop inside another loop)
- `col` = column number, starting at 1
- Goes from column 1 to column x

**In simple terms:** "For each row, go through each column from 1 to x"

---

```go
if row == 1 || row == y {
```

**What this checks:**
- `if` = condition starts
- `row == 1` = "if we're on the first row"
- `||` = OR
- `row == y` = "if we're on the last row"
- **In simple terms:** "If we're on the top edge or bottom edge..."

---

```go
if col == 1 || col == x {
    res.WriteRune('o')
```

**Breaking it down:**
- `col == 1` = if we're in the first column
- `col == x` = if we're in the last column
- `res.WriteRune('o')` = write the character 'o' to our result builder
  - `WriteRune` = method to add a single character
  - `'o'` = a single character (the letter o)
- **In simple terms:** "If we're at a corner (top/bottom edge AND left/right edge), draw an 'o'"

---

```go
} else {
    res.WriteRune('-')
}
```

**What this means:**
- `else` = "otherwise" (if the previous condition was false)
- `res.WriteRune('-')` = write a dash character
- **In simple terms:** "If we're on the top or bottom edge but NOT at a corner, draw a '-'"

---

```go
} else {
    if col == 1 || col == x {
        res.WriteRune('|')
```

**What this means:**
- This `else` matches the outer `if` (the one checking if `row == 1 || row == y`)
- **In simple terms:** "If we're NOT on top/bottom edge..."
- Then check: "Are we on the left or right edge?"
- If yes: draw a vertical bar `'|'`

---

```go
} else {
    res.WriteRune(' ')
}
```

**What this means:**
- **In simple terms:** "If we're not on any edge, draw a space (empty space inside the rectangle)"

---

```go
}
res.WriteRune('\n')
```

**What this does:**
- `}` = ends the inner `for` loop (the column loop)
- `res.WriteRune('\n')` = add a newline character
  - `'\n'` = special character that means "go to the next line"
- **In simple terms:** "After finishing all columns in a row, move to the next line"

---

```go
}
return res.String()
```

**What this does:**
- `}` = ends the outer `for` loop (the row loop)
- `return res.String()` = convert the `strings.Builder` to a regular string and give it back
- **In simple terms:** "We're done building the rectangle, now return it as text"

---

## 🔷 4️⃣ FUNCTION: `quadB`

```go
func quadB(x, y int) string {
```

**Same structure as quadA**, but draws a different pattern.

---

```go
if row == 1 && col == 1 {
    res.WriteRune('/')
```

**What this checks:**
- `&&` = **AND** operator (both conditions must be true)
- `row == 1 && col == 1` = "if we're on row 1 AND column 1"
- **In simple terms:** "If we're at the top-left corner, draw a forward slash `/`"

---

```go
} else if row == 1 && col == x {
    res.WriteRune('\\')
```

**What this checks:**
- `else if` = "otherwise, if this condition is true..."
- `row == 1 && col == x` = top-right corner
- `'\\'` = backslash character (we need two `\\` because one `\` is an escape character in Go)
- **In simple terms:** "If we're at the top-right corner, draw a backslash `\`"

---

```go
} else if row == y && col == 1 {
    res.WriteRune('\\')
```

**In simple terms:** "Bottom-left corner gets a backslash `\`"

---

```go
} else if row == y && col == x {
    res.WriteRune('/')
```

**In simple terms:** "Bottom-right corner gets a forward slash `/`"

---

```go
} else if row == 1 || row == y || col == 1 || col == x {
    res.WriteRune('*')
```

**What this checks:**
- **In simple terms:** "If we're on ANY edge but NOT at a corner, draw an asterisk `*`"

---

```go
} else {
    res.WriteRune(' ')
}
```

**In simple terms:** "If we're inside the rectangle (not on any edge), draw a space"

---

## 🔷 5️⃣ FUNCTIONS: `quadC`, `quadD`, `quadE`

These follow the **exact same structure** as quadA and quadB, but use different characters:

- **quadC:** Uses `'A'`, `'B'`, `'C'` characters
- **quadD:** Uses `'A'`, `'B'`, `'C'` characters in different positions
- **quadE:** Uses `'A'`, `'B'`, `'C'` characters in yet another pattern

The logic is identical, just the character placement changes!

---

## 🔷 6️⃣ THE MAIN FUNCTION

```go
func main() {
```

**What this is:**
- `main` = the special function that runs when your program starts
- **Every Go program must have a main function**

---

```go
reader := bufio.NewReader(os.Stdin)
```

**Breaking it down:**
- `reader` = variable name we chose
- `:=` = shorthand for "create variable and assign value"
- `bufio.NewReader` = creates a new buffered reader
- `os.Stdin` = **Standard Input** - the keyboard input
- **In simple terms:** "Create a tool that can read text from the keyboard efficiently"

---

```go
input, _ := io.ReadAll(reader)
```

**Breaking it down:**
- `input` = variable to store what we read
- `_` = underscore means "ignore this value" (in this case, ignoring any error)
- `io.ReadAll(reader)` = read everything from the input until it ends
- **In simple terms:** "Read all the text the user types and store it in `input`"

---

```go
content := string(input)
```

**What this does:**
- `content` = new variable
- `string(input)` = convert the input (which is bytes) into a string (text)
- **In simple terms:** "Convert the raw data into readable text"

---

```go
if strings.TrimSpace(content) == "" {
    fmt.Println("Not a quad function")
    return
}
```

**Breaking it down:**
- `strings.TrimSpace(content)` = remove any spaces, tabs, newlines from the beginning and end
- `== ""` = check if it equals an empty string
- `fmt.Println` = print a line of text to the screen
- `return` = exit the main function (end the program)
- **In simple terms:** "If the user didn't type anything (just empty space), print an error message and quit"

---

```go
lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
```

**Breaking it down (from inside out):**
- `strings.TrimRight(content, "\n")` = remove newlines from the right side of the content
- `strings.Split(..., "\n")` = split the text into pieces wherever there's a newline
- `lines` = variable holding an array (list) of strings
- **In simple terms:** "Break the input into separate lines and store them in a list"

**Example:** If input is:
```
o-o
| |
o-o
```
Then `lines` becomes: `["o-o", "| |", "o-o"]`

---

```go
y := len(lines)
```

**What this does:**
- `len(lines)` = get the length (number of items) in the lines array
- `y` = store that number
- **In simple terms:** "Count how many lines there are - that's the height"

---

```go
x := len([]rune(lines[0]))
```

**Breaking it down:**
- `lines[0]` = get the first line (arrays start at index 0)
- `[]rune(...)` = convert the string to an array of runes (characters)
- `len(...)` = count how many characters
- `x` = store that number
- **In simple terms:** "Count how many characters are in the first line - that's the width"

**Note:** We use `rune` instead of just counting the string length because some characters (like emojis) take up multiple bytes

---

```go
matches := []string{}
```

**What this does:**
- `matches` = variable name
- `[]string` = the type: a slice (list/array) of strings
- `{}` = create an empty slice
- **In simple terms:** "Create an empty list to store which quad functions match the input"

---

```go
if content == quadA(x, y) {
    matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", x, y))
}
```

**Breaking it down:**
- `content == quadA(x, y)` = call the quadA function with width x and height y, then check if it exactly matches the input
- `fmt.Sprintf("[quadA] [%d] [%d]", x, y)` = create a formatted string
  - `Sprintf` = "String print formatted" - creates a string instead of printing
  - `[quadA]` = literal text
  - `[%d]` = placeholder for an integer (decimal number)
  - The first `%d` gets replaced by `x`
  - The second `%d` gets replaced by `y`
  - **Example result:** `"[quadA] [5] [3]"`
- `append(matches, ...)` = add this string to the end of the matches list
- **In simple terms:** "If the input matches what quadA would draw, add 'quadA with these dimensions' to our matches list"

---

The next four `if` statements do **exactly the same thing** but for quadB, quadC, quadD, and quadE.

---

```go
if len(matches) == 0 {
    fmt.Println("Not a quad function")
```

**What this checks:**
- `len(matches)` = count how many items are in the matches list
- `== 0` = equals zero
- **In simple terms:** "If no quad functions matched the input, print an error message"

---

```go
} else {
    fmt.Println(strings.Join(matches, " || "))
}
```

**Breaking it down:**
- `strings.Join(matches, " || ")` = combine all strings in the matches list into one string
  - The second parameter `" || "` is the separator to put between each item
  - **Example:** If matches has `["[quadA] [3] [3]", "[quadC] [3] [3]"]`
  - Result: `"[quadA] [3] [3] || [quadC] [3] [3]"`
- `fmt.Println(...)` = print that combined string
- **In simple terms:** "If we found matches, print them all separated by ' || '"

---

```go
}
```

**Final closing brace** - ends the main function.

---

## 🎯 SUMMARY: WHAT THE WHOLE PROGRAM DOES

1. **Reads input** from the user (a pattern made of characters)
2. **Figures out the dimensions** (width and height)
3. **Tests if the input matches** any of the 5 quad patterns (A, B, C, D, E)
4. **Prints which quad function(s)** created that pattern
5. If **no match is found**, prints "Not a quad function"

**Real-world example:**
- User types:
```
o-o
| |
o-o
```
- Program calculates: width=3, height=3
- Program generates quadA(3,3), quadB(3,3), etc.
- Finds that quadA(3,3) matches!
- Prints: `[quadA] [3] [3]`

---

## 🔑 KEY CONCEPTS YOU LEARNED

1. **Functions** - reusable blocks of code
2. **Loops** - repeating code multiple times
3. **Conditions** - making decisions in code (if/else)
4. **Strings** - working with text
5. **Arrays/Slices** - lists of items
6. **Input/Output** - reading from keyboard, printing to screen

You now understand **every single word** in this program! 🎉---

# ✅ **AUDIT QUICK ANSWERS**

### **Why use rune instead of byte?**

> To correctly count characters in Go’s UTF-8 strings.

### **Why read from stdin?**

> Because quad executables send their drawing through a pipe into quadchecker.

### **Why use strings.Builder?**

> It's faster and avoids creating multiple temporary strings.

### **Why TrimRight?**

> Prevents an extra empty line when splitting by newline.

### **Why compare string to string?**

> Because the exact character layout (shape) determines which quad it is.

---

# ✅ **CRISIS MODE (If auditor asks you to summarize in 10 seconds)**

Just say this:

> “We read the input drawing, calculate its width and height, generate all quads with the same size, compare them to the input, and print whichever ones match. If none match, we print Not a quad function.”

---


## 📘  Vocabulary for Beginners

| Word                     | Meaning                                                          |                                                      |
| ------------------------ | ---------------------------------------------------------------- | ---------------------------------------------------- |
| **Pipe (                 | )**                                                              | Sends the output of one command as input to another. |
| **Executable (./quadA)** | A program file you can run directly.                             |                                                      |
| **Tree**                 | A command that shows folder structure visually.                  |                                                      |
| **Build**                | Compiles your Go code into a runnable program.                   |                                                      |
| **Args (arguments)**     | Extra values you pass to your program when you run it.           |                                                      |
| **stdin (Standard Input)**   | What your program reads when something is typed or piped into it.              |                                                                                    |                                                 |
| **stdout (Standard Output)** | Where your program prints messages (usually your terminal).                    |                                                                                    |                                                 |
| **strings.Builder**          | A Go tool to build long strings efficiently instead of concatenating with `+`. |                                                                                    |                                                 |
| **Rune**                     | A Go “character” type that handles letters, symbols, and Unicode correctly.    |                                                                                    |                                                 |
| **Trim**                     | Removing extra spaces or newline characters at the start/end of a string.      |                                                                                    |                                                 |
| **Module (`go.mod`)**        | File that defines your Go project’s name and dependencies.                     |                                                                                    |                                                 |
| **`io.ReadAll()`**           | Reads everything from input until end (useful for multi-line data).            |                                                                                    |                                                 |
| **Alphabetical order**       | Sorting results by name (A → Z).                                               |                                                                                    |                                                 |
                                        |
| **Import**    | Tells Go what other code you need to use (like `fmt`).                              |                                                                            |                                                                                    |
| **Compile**   | When Go converts your code into an executable program.                              |                                                                            |                                                                                    |
| **Run**       | Executes your Go code directly (`go run .`).                                        |                                                                            |                                                                                    |
| **Directory** | A folder on your computer that holds files like `main.go`.                          |                                                                            |                                                                                    |
| **Tree**      | A command that shows your folder structure neatly.                                  |                                                                            |                                                                                    |

---
