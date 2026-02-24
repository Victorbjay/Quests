# Go Reloaded - Beginner's Complete Guide 🚀

## 👋 Welcome, New Coder!

This guide is for YOU — even if you've never written code before. We'll go slow, explain everything, and use lots of examples.

**By the end, you'll be able to:**
1. Read and understand Go code
2. Build a text transformation tool
3. Teach this to your friends! 🎓

---

## 🧠 Part 1: Understanding the Basics

### What is Programming Actually?

Imagine you have a recipe book 📖

```
Recipe: Make Tea
1. Boil water
2. Put tea bag in cup
3. Pour hot water
4. Wait 3 minutes
5. Remove tea bag
6. Add sugar (optional)
7. Serve
```

**Programming is just writing a recipe for a computer!**

The computer follows each step exactly. No thinking, no guessing — just pure execution.

---

### What is Go?

**Go** (also called Golang) is a programming language created by Google. It's known for being:
- ⚡ **Fast** — runs quickly
- 📝 **Simple** — easy to read
- 🔧 **Powerful** — builds big systems

Think of Go as a tool. Like a hammer builds houses, Go builds software.

---

### Your First Go Code — Line by Line

Let's break down our main.go piece by piece:

```go
package main
```

**What is this?** Every Go file starts with `package main`.

**Think of it like:** Putting a label on a box. This says "This code belongs to the main program."

---

### Variables — The Memory Box 📦

```go
var name string = "Victor"
```

**Translation:** "Hey computer, create a box called 'name'. Put the word 'Victor' inside it."

**Breaking it down:**
- `var` = "I want to create a variable" (a storage box)
- `name` = the box label
- `string` = the type of thing I'm storing (text)
- `= "Victor"` = what I'm putting inside

**Real-world analogy:**
```
┌─────────────────┐
│     📦 BOX      │
│   Label: name   │
│  Contents:      │
│   "Victor"      │
└─────────────────┘
```

---

### More Variable Types

| Type | What it stores | Example |
|------|----------------|---------|
| `string` | Text | `"Hello"` |
| `int` | Whole numbers | `42`, `100` |
| `float64` | Decimal numbers | `3.14` |
| `bool` | True/False | `true`, `false` |

---

### Functions — Reusable Recipes 🍳

```go
func greet(name string) string {
    return "Hello, " + name
}
```

**What this does:**
1. Takes a name (input)
2. Says "Hello, [that name]" (process)
3. Gives back the greeting (output)

**Think of it like a machine:**

```
┌─────────────────────────────────┐
│      GREETING MACHINE           │
│                                 │
│  Input: "Victor"               │
│         ↓                       │
│  [Processing...]                │
│         ↓                       │
│  Output: "Hello, Victor"        │
└─────────────────────────────────┘
```

---

### Loops — Doing Things Again and Again 🔄

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

**Translation:** "Start at 0, count up to 5 (not including 5), and print each number."

**Output:**
```
0
1
2
3
4
```

**Real-world:** Like counting candles on a birthday cake! 🕯️

---

### If/Else — Making Decisions 🧠

```go
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Child")
}
```

**Translation:** "If age is 18 or more, say 'Adult'. Otherwise, say 'Child'."

**Flow chart:**

```
        Is age >= 18?
           /      \
         YES      NO
          ↓        ↓
    "Adult"    "Child"
```

---

## 🏗️ Part 2: Building Go Reloaded

Now that you understand the basics, let's build our text processor!

---

### The Big Picture

Our program does this:

```
INPUT: "it (cap) was the best"
       ↓
[PROCESSING]
       ↓
OUTPUT: "It was the best"
```

We process text in **4 stages**:

```
1. Split text into words
2. Transform markers (hex, bin, up, low, cap)
3. Fix "a" → "an"
4. Fix punctuation and quotes
```

---

### Stage 1: Splitting Text

**Problem:** How do we break "Hello World" into ["Hello", "World"]?

**Solution:** Use `strings.Fields()`

```go
import "strings"

text := "Hello World"
words := strings.Fields(text)
// words is now: ["Hello", "World"]
```

**Visual:**

```
"Hello World"
     ↓
┌───────┬───────┐
│ Hello │ World │
└───────┴───────┘
```

---

### Stage 2: Processing Markers

This is the core of our project. Let's understand each marker:

#### (hex) — Hexadecimal to Decimal

**What is hexadecimal?**
It's a number system with 16 digits: 0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F

**Why do we need it?** Computers use hex for memory addresses, colors, etc.

**Example:**
- Hex: `1E` = Decimal: `30`
- Hex: `FF` = Decimal: `255`

**Code:**

```go
import "strconv"

number := "1E"
decimal, err := strconv.ParseInt(number, 16, 64)
// decimal = 30
```

**The "16" means:** Parse as base-16 (hexadecimal)

---

#### (bin) — Binary to Decimal

**What is binary?** Numbers using only 0 and 1. All computers think in binary!

**Examples:**
- Binary: `10` = Decimal: `2`
- Binary: `1010` = Decimal: `10`

**Code:**

```go
number := "10"
decimal, err := strconv.ParseInt(number, 2, 64)
// decimal = 2
```

---

#### (up), (low), (cap) — Changing Case

**Case = the format of letters**

| Case | Example | Description |
|------|---------|-------------|
| UPPER | HELLO | All capital |
| lower | hello | All small |
| Capitalize | Hello | First letter big |

**Code:**

```go
word := "hello"

// Uppercase
strings.ToUpper(word)   // "HELLO"

// Lowercase  
strings.ToLower(word)   // "hello"

// Capitalize (first letter big)
func capitalize(w string) string {
    if len(w) == 0 {
        return w
    }
    // Convert first rune to uppercase
    runes := []rune(w)
    runes[0] = unicode.ToUpper(runes[0])
    return string(runes)
}

capitalize(word)  // "Hello"
```

**Why use runes?** Because some letters (like in Nigerian names: "Chidinma", "Nnenna") need special handling!

---

### Stage 3: A → An Rule

**The rule:** "a" becomes "an" before words starting with:
- Vowels: a, e, i, o, u
- Letter 'h' (like "hour")

**Examples:**
- "a car" ✓ (c is not a vowel)
- "an car" ✗
- "an egg" ✓ (e is a vowel)
- "a egg" ✗
- "an hour" ✓ (h)

**Code:**

```go
func processAtoAn(words []string) []string {
    for i := 0; i < len(words)-1; i++ {
        // Check if current word is "a"
        if strings.ToLower(words[i]) == "a" {
            // Get first letter of next word
            nextWord := words[i+1]
            firstChar := strings.ToLower(string(nextWord[0]))
            
            // Check if vowel or 'h'
            if strings.Contains("aeiouh", firstChar) {
                words[i] = "an"  // Change to "an"
            }
        }
    }
    return words
}
```

---

### Stage 4: Punctuation & Quotes

**Punctuation rules:**
- No space before: `. , ! ? : ;`
- `...` stays together
- `'word'` removes extra spaces

**Examples:**
- `"hello , world"` → `"hello, world"`
- `"awesome "` → `"awesome"`

---

## 🔍 Part 3: Understanding Our Code

Let me walk through main.go line by line:

### Package Declaration

```go
package main
```
This tells Go: "This is the main program."

---

### Imports

```go
import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
    "unicode"
)
```

These are **libraries** — pre-written code we can use:

| Import | What it does |
|--------|--------------|
| `fmt` | Print things to screen |
| `os` | Read/write files |
| `regexp` | Pattern matching |
| `strconv` | Convert between number types |
| `strings` | Text manipulation |
| `unicode` | Handle letters (including Unicode) |

---

### The Main Function

```go
func main() {
    // This runs when you run the program!
}
```

Think of it as the "START HERE" sign in a maze.

---

### Reading a File

```go
content, err := os.ReadFile(inputFile)
```

- `content` = the text from the file
- `err` = any error (we check this!)

---

### Processing the Text

```go
result := ProcessText(string(content))
```

We call our transformation function and store the result.

---

## 🎯 Part 4: Running Your Code

### Step 1: Install Go

**On Windows:**
1. Go to https://go.dev/dl
2. Download Windows installer
3. Run it!

**On Mac:**
```bash
brew install go
```

**On Linux:**
```bash
sudo apt install golang-go
```

---

### Step 2: Run Your Program

```bash
# Navigate to folder
cd go-reloaded

# Run with test file
go run . sample.txt result.txt

# See the output
cat result.txt
```

---

## ❓ Frequently Asked Questions

### Q: What if I get an error?

**Don't panic!** Read the error message — it usually tells you:
- Which line has the problem
- What kind of problem it is

### Q: What does "index out of range" mean?

You're trying to access something that doesn't exist.

```go
words := []string{"hello"}
fmt.Println(words[5])  // ❌ Only index 0 exists!
```

### Q: Why use runes instead of strings for letters?

```go
// ❌ Wrong for Unicode
name := "Chidinma"
name[0]  // Returns 67 (byte value)

// ✅ Correct
runes := []rune(name)
runes[0] // Returns 'C' (letter)
```

This matters for names with special characters!

---

## 📝 Practice Exercises

Try these to test your understanding:

### Exercise 1: What does this print?

```go
func main() {
    name := "Victor"
    fmt.Println("Hello, " + name)
}
```

### Exercise 2: Fix the bug

```go
func main() {
    numbers := []int{1, 2, 3}
    for i := 0; i <= 3; i++ {
        fmt.Println(numbers[i])
    }
}
```

### Exercise 3: Write a function

Write a function that takes a name and returns "Good morning, [name]!"

---

## 🧠 Quick Reference Card

| Symbol/Word | Meaning |
|-------------|---------|
| `:=` | Create and assign at same time |
| `=` | Assign to existing variable |
| `==` | Check if equal |
| `!=` | Not equal |
| `func` | Create a function |
| `return` | Give back a value |
| `[]string` | A list of text |
| `len(x)` | How many items in x |

---

## 🎓 Teaching This to Others

When you teach your friends:

1. **Start with analogies** — use real-life examples
2. **Go slow** — don't rush the basics
3. **Let them make mistakes** — debugging teaches more than perfect code
4. **Celebrate small wins** — every working program is an achievement!

---

## 📚 Additional Resources

- [Go Tour](https://tour.golang.org/) — Interactive Go tutorial
- [Go by Example](https://gobyexample.com/) — Code examples
- [Gophercises](https://gophercises.com/) — Practice exercises

---

**Remember:** Every expert was once a beginner. You CAN do this! 💪

*Created: February 2026*
*For: Complete Beginners*
*Teaching: Go Reloaded (01-Edu)*
