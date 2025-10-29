### Before starting this Quest, ensure to check on your gittea piscine-go, the `go.mod` file shouldn't have the github.com/z01 package else the checker will fail you. delete that 3rd linw entirely
##Also for those Testing, make sure you don't add and push those test files to your gittea, stop using `git add .` entirely and use `git add filename` instead.
---

# 🧩 **Lesson: `printprogramname`**

---

## 🎯 **Goal**
look at the file to submit info 

**You must create a folder** `printprogramename` and then create the file `main.go`

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
 git add printprogramname/main.go #copy and paste this way
```

---

Excellent 🧠 — we’re now on **Lesson: `printparams`**!
Let’s break it down just like before — full lecture, visual flow, and easy explanation for our learning logs. 🪄

---

# 🧩 **Lesson: `printparams`**

---

## 🎯 **Goal**

We need to write a Go program that **prints every argument** it receives from the command line —
**each argument on a new line**.

Example:

```bash
$ go run . choumi is the best cat
choumi
is
the
best
cat
```

---

## 🧠 **Core Idea**

In Go, when you run a program, **everything typed after the program name** (the words you type in the terminal)
is stored in a special variable called `os.Args`.

Example:

```bash
$ go run . choumi is the best cat
```

Then:

```go
os.Args == ["main", "choumi", "is", "the", "best", "cat"]
```

* `os.Args[0]` → is the program name (we **skip this**)
* `os.Args[1:]` → is a **slice** of the actual arguments we want to print.

---

## 💡 **Code**

```go
package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	// Loop through the command-line arguments, skipping the program name
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]

		// Print each character (rune) in the argument
		for _, r := range arg {
			z01.PrintRune(r)
		}

		// After each word, print a newline
		z01.PrintRune('\n')
	}
}
```

---

## 🧩 **Step-by-Step Breakdown**

| Step | Code Section            | What Happens                                        | Example Output                                   |
| ---- | ----------------------- | --------------------------------------------------- | ------------------------------------------------ |
| 1️⃣  | `os.Args`               | Reads command-line arguments                        | `["main", "choumi", "is", "the", "best", "cat"]` |
| 2️⃣  | `for i := 1; ...`       | Loops through all arguments except the program name | Starts with `"choumi"`                           |
| 3️⃣  | `for _, r := range arg` | Loops through each character (rune)                 | `'c'`, `'h'`, `'o'`, `'u'`, `'m'`, `'i'`         |
| 4️⃣  | `z01.PrintRune(r)`      | Prints one character at a time                      | → `choumi`                                       |
| 5️⃣  | `z01.PrintRune('\n')`   | Adds a new line after each word                     | Output moves to next line                        |

---

## 🎬 **Visualization**

🧱 Command:

```
go run . choumi is the best cat
```

🧭 What happens internally:

```
os.Args → ["main", "choumi", "is", "the", "best", "cat"]
              ↑
         skip index 0
```

🌀 Loop starts:

| i | arg value  | Printed output |
| - | ---------- | -------------- |
| 1 | `"choumi"` | choumi         |
| 2 | `"is"`     | is             |
| 3 | `"the"`    | the            |
| 4 | `"best"`   | best           |
| 5 | `"cat"`    | cat            |

🖨️ **Final terminal output:**

```
choumi
is
the
best
cat
```

---

## ⚙️ **How the Loops Work (Visual Animation)**

Imagine this flow:

```
os.Args = ["main", "choumi", "is", "the", "best", "cat"]
            └── skip
             ↓
          i = 1 → "choumi"
              ↓ print c → h → o → u → m → i → \n
          i = 2 → "is"
              ↓ print i → s → \n
          ...
```

---

## 🚫 **Common Mistakes**

| Mistake                     | Result                          | Why                                   |
| --------------------------- | ------------------------------- | ------------------------------------- |
| Printing `os.Args` directly | `[main choumi is the best cat]` | It prints the slice, not each word    |
| Starting loop at `0`        | Includes program name           | You’ll get `main` as the first output |
| Forgetting newline          | All words appear on one line    | Missing `\n` after each argument      |

---

## ✅ **Expected Output**

```
$ go run . choumi is the best cat
choumi
is
the
best
cat
```

---

## 🗂️ **File to Submit**

```
printparams/main.go
```

---

## 🧾 **Summary Table**

| Concept           | Meaning                     |
| ----------------- | --------------------------- |
| `os.Args`         | All command-line inputs     |
| `os.Args[0]`      | Program name                |
| `os.Args[1:]`     | Actual arguments            |
| `z01.PrintRune()` | Prints one rune (character) |
| `\n`              | Newline character           |

---

Perfect 🧠 — now we move to **Lesson: `revparams`** (Level 14)!
Just like before — full teaching lecture, visualization, explanation, and final file info for our logbook 📘✨

---

# 🧩 **Lesson: `revparams`**

---

## 🎯 **Goal**

We will write a program that **prints all command-line arguments in reverse order**,
**one per line**, **excluding** the program name itself.

Example:

```bash
$ go run . choumi is the best cat
cat
best
the
is
choumi
```

---

## 🧠 **Concept Recap**

Earlier in `printparams`, we used:

```go
os.Args[1:]
```

to get all arguments *except the program name*.

Now, we’ll **loop backward** through that slice.

In Go, to loop backward:

```go
for i := len(os.Args) - 1; i > 0; i-- {
	// code
}
```

---

## 💡 **Code**

```go
package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	// Loop from the last argument to the first (skip index 0 because it’s the program name)
	for i := len(os.Args) - 1; i > 0; i-- {
		arg := os.Args[i]

		// Print each character in the argument
		for _, r := range arg {
			z01.PrintRune(r)
		}

		// Print newline after each argument
		z01.PrintRune('\n')
	}
}
```

---

## 🧩 **Step-by-Step Breakdown**

| Step | Code Section                          | Explanation                         | Example                                          |
| ---- | ------------------------------------- | ----------------------------------- | ------------------------------------------------ |
| 1️⃣  | `os.Args`                             | Captures all command-line arguments | `["main", "choumi", "is", "the", "best", "cat"]` |
| 2️⃣  | `len(os.Args) - 1`                    | Gets last index in the list         | `5` (for "cat")                                  |
| 3️⃣  | `for i := len(os.Args)-1; i > 0; i--` | Starts loop backward                | Goes 5 → 4 → 3 → 2 → 1                           |
| 4️⃣  | `arg := os.Args[i]`                   | Picks each word                     | `"cat"`, `"best"`, `"the"`, `"is"`, `"choumi"`   |
| 5️⃣  | `for _, r := range arg`               | Loops through each character        | `'c'`, `'a'`, `'t'`, ...                         |
| 6️⃣  | `z01.PrintRune(r)`                    | Prints each rune                    | prints letter-by-letter                          |
| 7️⃣  | `z01.PrintRune('\n')`                 | Prints new line                     | moves to next word                               |

---

## 🎬 **Visualization**

### Command:

```
go run . choumi is the best cat
```

### Inside Go:

```
os.Args = ["main", "choumi", "is", "the", "best", "cat"]
Indexes =   0          1        2      3        4        5
```

### Loop flow:

| i | Value    | Printed |
| - | -------- | ------- |
| 5 | "cat"    | cat     |
| 4 | "best"   | best    |
| 3 | "the"    | the     |
| 2 | "is"     | is      |
| 1 | "choumi" | choumi  |

✅ Each word is printed on a new line.

---

## ⚙️ **Loop Flow Visualization (Diagram)**

```
           ┌────────────────────────────┐
           │  os.Args = [main, choumi,  │
           │           is, the, best,   │
           │           cat]             │
           └────────────┬───────────────┘
                        │
              len(os.Args) = 6
                        │
         Start i = 5 (last arg: "cat")
                        │
      ↓ Print "cat"
      ↓ i--
         i = 4 ("best")
      ↓ Print "best"
      ↓ i--
         i = 3 ("the")
      ↓ Print "the"
      ↓ i--
         i = 2 ("is")
      ↓ Print "is"
      ↓ i--
         i = 1 ("choumi")
      ↓ Print "choumi"
      ↓ i--
         i = 0 → STOP (don’t print program name)
```

---

## 🚫 **Common Mistakes**

| Mistake                | Problem                      |
| ---------------------- | ---------------------------- |
| Starting loop from `0` | Includes program name        |
| Forgetting `i--`       | Infinite loop!               |
| Not printing newline   | Words all appear on one line |

---

## ✅ **Expected Output**

```
$ go run . choumi is the best cat
cat
best
the
is
choumi
```

---

## 🗂️ **File to Submit**

```
 git add revparams/main.go
```

---

## 🧾 **Summary Table**

| Concept                               | Meaning                                          |
| ------------------------------------- | ------------------------------------------------ |
| `os.Args`                             | List of all command-line arguments               |
| `len(os.Args)`                        | Total number of arguments including program name |
| `for i := len(os.Args)-1; i > 0; i--` | Looping backward                                 |
| `z01.PrintRune()`                     | Prints one character (rune)                      |
| `\n`                                  | Adds newline after each argument                 |

---

