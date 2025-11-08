
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
├── quadA/
│   └── main.go
├── quadB/
│   └── main.go
├── quadC/
│   └── main.go
├── quadD/
│   └── main.go
└── quadE/
    └── main.go
```

---

## 🚀 3. Quad Programs (Generators)

Each program takes **two integers** (`width`, `height`) and prints a rectangle using specific corner and border rules.

---

### 🅰️ `quadA/main.go`

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

### 🅱️ `quadB/main.go`

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

### 🅲 `quadC/main.go`

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

### 🅳 `quadD/main.go`

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

### 🅴 `quadE/main.go`

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
go build -o ../quadA main.go
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

## 📘 6. Vocabulary for Beginners

| Word                     | Meaning                                                          |                                                      |
| ------------------------ | ---------------------------------------------------------------- | ---------------------------------------------------- |
| **Pipe (                 | )**                                                              | Sends the output of one command as input to another. |
| **Executable (./quadA)** | A program file you can run directly.                             |                                                      |
| **Tree**                 | A command that shows folder structure visually.                  |                                                      |
| **Build**                | Compiles your Go code into a runnable program.                   |                                                      |
| **Args (arguments)**     | Extra values you pass to your program when you run it.           |                                                      |
| **Module (go.mod)**      | A Go file that tells Go how to handle your project dependencies. |                                                      |
| **Run (go run .)**       | Compiles and runs your Go program from the current directory.    |                                                      |

| Word                         | Meaning                                                                        |                                                                                    |                                                 |
| ---------------------------- | ------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------- | ----------------------------------------------- |
| **Pipe (`                    | `)**                                                                           | Sends the *output* of one program as the *input* of another. Example: `./quadA 3 3 | go run .` sends quadA’s drawing to quadchecker. |
| **stdin (Standard Input)**   | What your program reads when something is typed or piped into it.              |                                                                                    |                                                 |
| **stdout (Standard Output)** | Where your program prints messages (usually your terminal).                    |                                                                                    |                                                 |
| **strings.Builder**          | A Go tool to build long strings efficiently instead of concatenating with `+`. |                                                                                    |                                                 |
| **Rune**                     | A Go “character” type that handles letters, symbols, and Unicode correctly.    |                                                                                    |                                                 |
| **Trim**                     | Removing extra spaces or newline characters at the start/end of a string.      |                                                                                    |                                                 |
| **Module (`go.mod`)**        | File that defines your Go project’s name and dependencies.                     |                                                                                    |                                                 |
| **`io.ReadAll()`**           | Reads everything from input until end (useful for multi-line data).            |                                                                                    |                                                 |
| **Alphabetical order**       | Sorting results by name (A → Z).                                               |                                                                                    |                                                 |


| Word          | Meaning                                                                             |                                                                            |                                                                                    |
| ------------- | ----------------------------------------------------------------------------------- | -------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| **Pipe (      | )**                                                                                 | Sends the output of one program as input to another. Example: `./quadA 3 3 | go run .` means “take what quadA prints and give it to the program we’re running.” |
| **Module**    | A Go project with its own `go.mod` file — like a package with its own dependencies. |                                                                            |                                                                                    |
| **Import**    | Tells Go what other code you need to use (like `fmt`).                              |                                                                            |                                                                                    |
| **Compile**   | When Go converts your code into an executable program.                              |                                                                            |                                                                                    |
| **Run**       | Executes your Go code directly (`go run .`).                                        |                                                                            |                                                                                    |
| **Directory** | A folder on your computer that holds files like `main.go`.                          |                                                                            |                                                                                    |
| **Tree**      | A command that shows your folder structure neatly.                                  |                                                                            |                                                                                    |

---
