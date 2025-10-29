Perfect 👌 Let’s do our **standard learning log lecture** — the same clear, child-friendly and visual style we’ve been using.

---

# 🧩 **Lesson: `printprogramname`**

---

## 🎯 **Goal**

We want to write a Go program that prints the **name of the program itself**.

Example:

```bash
$ go build -o Nessy
$ ./Nessy
Nessy
```

---

## 🧠 **Concept Behind It**

When a Go program runs, it can access **the arguments passed to it** (like command-line inputs) through a built-in variable called `os.Args`.

* `os.Args[0]` → always holds **the program’s own name or path**
  Example:

  * If you run `./main` → it’s `"./main"`
  * If you run `./Nessy` → it’s `"./Nessy"`

But we only want the **name part**, not the whole path.
So we’ll remove the “./” or any folder names.

---

## 💡 **Code**

Here’s our full Go code:

```go
package main

import (
	"os"                 // allows us to access command-line arguments
	"github.com/01-edu/z01" // used for printing runes (characters)
)

func main() {
	// Get the full name/path of the program (e.g. "./main")
	name := os.Args[0]

	// Find where the last '/' is so we can remove any folder part
	start := 0
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' { // when we find a '/'
			start = i + 1   // next character is where the name starts
			break
		}
	}

	// Print only the program name, rune by rune
	for _, r := range name[start:] {
		z01.PrintRune(r)
	}

	// Add a new line at the end
	z01.PrintRune('\n')
}
```

---

## 🔍 **Step-by-Step Visualization**

| Step | What happens           | Example value                     | Description                      |
| ---- | ---------------------- | --------------------------------- | -------------------------------- |
| 1️⃣  | Program starts         | `./Nessy`                         | `os.Args[0]` reads the full path |
| 2️⃣  | Loop searches backward | finds `'/'` before `N`            | we locate the last slash `/`     |
| 3️⃣  | Calculate start index  | `start = 2`                       | skip the `./`                    |
| 4️⃣  | Loop prints each rune  | `'N'`, `'e'`, `'s'`, `'s'`, `'y'` | uses `z01.PrintRune`             |
| 5️⃣  | Print newline          | `\n`                              | moves to next line               |

---

## 🎬 **Visual Animation (mental picture)**

```
os.Args[0] = "./Nessy"

            ↓ loop scans backward
  ./Nessy
   ↑     ↑
   / found at index 1 → start printing from index 2

Print order:
N → e → s → s → y → \n
```

🖨️ Output:

```
Nessy
```

---

## ⚙️ **How It Works in Real Time**

| Command                   | Output                             |
| ------------------------- | ---------------------------------- |
| `go build -o Nessy`       | builds the program named **Nessy** |
| `./Nessy`                 | prints `Nessy`                     |
| `go build` (default name) | creates `printprogramname`         |
| `./printprogramname`      | prints `printprogramname`          |

---

## 🚫 **Common Mistakes**

| Mistake                      | Result                          | Why                       |
| ---------------------------- | ------------------------------- | ------------------------- |
| Printing the whole `os.Args` | `[./main]`                      | Prints as slice, not name |
| Printing bytes               | Gibberish like `ð¤¦ð»...`    | Wrong encoding            |
| Forgetting newline           | Name stuck with terminal prompt | Missing `\n`              |

---

## 🧾 **Summary**

| Concept        | Meaning                |
| -------------- | ---------------------- |
| `os.Args[0]`   | Program’s name/path    |
| Loop backwards | Find last `/`          |
| Print runes    | Character by character |
| Add newline    | To look neat in output |

---

## ✅ **Final Output**

```bash
$ go build -o Nessy
$ ./Nessy
Nessy
```

---

## 🗂️ **File to Submit**

```
printprogramname/main.go
```

---

Would you like me to make a small **HTML visual animation version** for this one too (that you can open and host as a single `.html` file)? It would show the path being scanned and the name extracted step-by-step.
