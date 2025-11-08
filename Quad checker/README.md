
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
mkdir quadchecker
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

### ▶️ STEP 5: Run and test

You can test your program by piping (sending) the output of a quad function (like `quadA`) into your `quadchecker`.

Example:

```bash
./quadA 3 3 | go run .
```

Output should be:

```
[quadA] [3] [3]
```

Or, if you have multiple matches:

```
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
```

If it’s not a valid quad:

```bash
echo "random text" | go run .
```

Output:

```
Not a quad function
```

---


## 🧪 Step 4: Testing It

### 1️⃣ Make sure your `quadA`, `quadB`, etc. binaries exist

If you still have your quad files, compile one like this:

```bash
go run path/to/quadA.go 3 3
```

You’ll see something like:

```
o--o
|  |
o--o
```

---

### 2️⃣ Pipe the output into quadchecker

From inside your `quadchecker` folder:

```bash
./quadA 3 3 | go run .
```

Expected output:

```
[quadA] [3] [3]
```

---

### 3️⃣ Test multiple matches

```bash
./quadC 1 1 | go run .
```

Expected output:

```
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
```

---

### 4️⃣ Test invalid input

```bash
echo "random text" | go run .
```

Output:

```
Not a quad function
```

---

## 📘 Vocabulary

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
