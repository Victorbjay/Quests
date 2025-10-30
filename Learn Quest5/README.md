
## Do not be in a hurry to finish, make sure you rehearse and understand the code!


This section is string handling — specifically **runes**, which are very important in Go for dealing with Unicode characters.

Let’s go through it **step-by-step**, in the same clear format as quest4 before — with explanation, visuals, and test code.

---

## 🧩 TASK NAME: `FirstRune`

### 🧠 What does “rune” mean?

A **rune** in Go is basically a **Unicode character**.
Unlike normal bytes, a rune can represent letters from *any* language (not just English).

👉 Example:

```go
'A'   → rune for letter A
'س'   → rune for Arabic letter seen
'你'  → rune for Chinese character "you"
```

Go uses **UTF-8 encoding**, so one “letter” in a string can be made up of multiple bytes.
That’s why using **runes** is safer when working with characters.

---

## ⚙️ What the task is asking

We must write a function:

```go
func FirstRune(s string) rune
```

✅ It takes a **string**
✅ It returns the **first rune (character)** from it

---

## ✅ FULL CODE (firstrune.go)

```go
package piscine

// FirstRune returns the first rune (Unicode character) of a string.
// It correctly handles multi-byte characters like emojis or non-Latin letters.
func FirstRune(s string) rune {
	for _, r := range s {
		return r // the first iteration gives the first rune
	}
	return 0 // if string is empty
}
```

---

## 🧠 LINE-BY-LINE EXPLANATION

| Line                  | Code                                                                                  | Explanation |
| --------------------- | ------------------------------------------------------------------------------------- | ----------- |
| `for _, r := range s` | This loop goes through each **rune** in the string `s`. The `_` ignores the index.    |             |
| `return r`            | On the first loop, we immediately return the first rune. No need to continue looping. |             |
| `return 0`            | If the string is empty (no runes), we return 0 (which means “no rune”).               |             |

---

## 🔍 VISUAL EXAMPLE

Let’s say:

```go
s := "Hello!"
```

Internally, Go sees:

| Rune | Unicode | Index |
| ---- | ------- | ----- |
| 'H'  | 72      | 0     |
| 'e'  | 101     | 1     |
| 'l'  | 108     | 2     |
| 'l'  | 108     | 3     |
| 'o'  | 111     | 4     |
| '!'  | 33      | 5     |

👉 The loop goes:

```
1️⃣ r = 'H' → return 'H'
✅ Done!
```

So `FirstRune("Hello!")` → `'H'`

---

## 🧱 VISUAL FLOW DIAGRAM

```
Start
 ↓
for _, r := range s
 ↓
Is this the first rune?
 ↓
Yes → return r
 ↓
End (function stops immediately)
```

---

## 🧪 TEST FILE (main.go)

```go
package main

import (
	"piscine"
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.FirstRune("Hello!")) // H
	z01.PrintRune(piscine.FirstRune("Salut!")) // S
	z01.PrintRune(piscine.FirstRune("Ola!"))   // O
	z01.PrintRune('\n')
}
```

### 🖥️ Expected Output

```
HSO
```

---

## 🗂️ FILES TO SUBMIT

* ✅ `firstrune.go`

---

## 🧩 QUICK SUMMARY

| Concept      | Meaning                            |
| ------------ | ---------------------------------- |
| **rune**     | Unicode character (not just ASCII) |
| **range s**  | iterates over runes, not bytes     |
| **return r** | returns the first rune immediately |
| **return 0** | if string is empty                 |

---

**The reverse** of the previous task — instead of returning the *first* rune, we’ll return the *last* rune of a string.


---

## 🧩 TASK NAME: `LastRune`

### 🧠 What is a “rune” again?

A **rune** in Go is a Unicode character — meaning it can represent **any letter, symbol, or emoji** (not just English characters).

✅ `'A'`, `'é'`, `'ع'`, `'你'`, `'🌍'` — all are **runes**.
Go treats them as **int32** values representing Unicode code points.

---

## ⚙️ What the task is asking

We must write a function:

```go
func LastRune(s string) rune
```

✅ Takes a **string**
✅ Returns the **last rune** in that string
✅ Works correctly with Unicode (non-ASCII) text

---

## ✅ FINAL CODE (lastrune.go)

```go
package piscine

// LastRune returns the last rune (Unicode character) in a string.
// It properly handles multi-byte characters like emojis or non-Latin letters.
func LastRune(s string) rune {
	var last rune
	for _, r := range s {
		last = r // keeps updating until it reaches the last rune
	}
	return last
}
```

---

## 🧠 LINE-BY-LINE EXPLANATION

| Line                  | Code                                                                     | Explanation |
| --------------------- | ------------------------------------------------------------------------ | ----------- |
| `var last rune`       | Create a variable to hold the last rune found. Starts empty.             |             |
| `for _, r := range s` | Loop through each rune (character) in the string. `_` ignores the index. |             |
| `last = r`            | On each loop, update `last` to the current rune. It keeps changing...    |             |
| `return last`         | When the loop ends, `last` contains the **last** rune in the string.     |             |

---

## 🔍 EXAMPLE: `"Hello!"`

| Step | r (current rune) | last (after assignment) |
| ---- | ---------------- | ----------------------- |
| 1    | 'H'              | 'H'                     |
| 2    | 'e'              | 'e'                     |
| 3    | 'l'              | 'l'                     |
| 4    | 'l'              | 'l'                     |
| 5    | 'o'              | 'o'                     |
| 6    | '!'              | '!'                     |

🧩 After the loop finishes → `last = '!'`
✅ The function returns `'!'`

---

## 🧱 VISUAL FLOW DIAGRAM

```
Start
 ↓
var last rune
 ↓
for _, r := range s
 ├─> update last = r
 └─> loop until string ends
 ↓
return last
 ↓
End
```

---

## 🧪 TEST FILE (main.go)

```go
package main

import (
	"piscine"
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.LastRune("Hello!")) // !
	z01.PrintRune(piscine.LastRune("Salut!")) // !
	z01.PrintRune(piscine.LastRune("Ola!"))   // !
	z01.PrintRune('\n')
}
```

### 🖥️ Expected Output

```
!!!
```

---

## 🗂️ FILES TO SUBMIT

✅ `lastrune.go`

---

## 🎨 VISUAL UNDERSTANDING (Rune Travel Map)

```
String: "Ola!"
Runes:   O → l → a → !
                     ↑
          👈 this one is the last rune returned
```

---

## 🧩 QUICK SUMMARY

| Concept         | Meaning                            |
| --------------- | ---------------------------------- |
| **Rune**        | A Unicode character                |
| **range s**     | Loops over runes (not bytes)       |
| **last = r**    | Keeps overwriting until final rune |
| **return last** | Gives us the last character        |

---


## 🧩 TASK NAME: `NRune`

### 🧠 What are we doing?

We need to **get the nth character (rune)** from a given string.

For example:

| String     | n | Output |
| ---------- | - | ------ |
| `"Hello!"` | 3 | `'l'`  |
| `"Salut!"` | 2 | `'a'`  |
| `"Ola!"`   | 4 | `'!'`  |

If it’s **not possible** (for example, if n is negative or larger than the number of characters), we return **0**.

---

## 💡 KEY IDEAS

* A **rune** = a single Unicode character.
  So looping with `range` gives us each rune in the correct order.
* We **count** each rune until we reach the `n`th one.
* If we finish the loop and didn’t reach it → return `0`.

---

## ✅ FINAL CODE (nrune.go)

```go
package piscine

// NRune returns the nth rune (Unicode character) from a string.
// If n is out of range (too small or too big), it returns 0.
func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}

	count := 0
	for _, r := range s {
		count++
		if count == n {
			return r
		}
	}
	return 0
}
```

---

## 🧱 LINE-BY-LINE EXPLANATION

| Line                         | Code                                                    | What it means |
| ---------------------------- | ------------------------------------------------------- | ------------- |
| `if n <= 0 { return 0 }`     | Negative or zero position? Impossible → return 0.       |               |
| `count := 0`                 | We’ll count how many runes we’ve seen.                  |               |
| `for _, r := range s`        | Go through each rune (Unicode character) in the string. |               |
| `count++`                    | Increase our count every time we see a rune.            |               |
| `if count == n { return r }` | If this is the nth rune, return it immediately!         |               |
| `return 0`                   | If we finish and didn’t find it, return 0.              |               |

---

## 🔍 VISUAL EXAMPLE: `"Hello!"`, n = 3

| Step | Rune  | Count | Is Count == n? | Action       |
| ---- | ----- | ----- | -------------- | ------------ |
| 1    | `'H'` | 1     | ❌              | Continue     |
| 2    | `'e'` | 2     | ❌              | Continue     |
| 3    | `'l'` | 3     | ✅              | Return `'l'` |

✅ Function returns `'l'`

---

## 🎨 VISUAL DIAGRAM

```
String: "H e l l o !"
Index:   1 2 3 4 5 6

We want n = 3

Loop:
   H → e → l → stop!
         ↑
         return this rune
```

---

## 🧪 TEST FILE (main.go)

```go
package main

import (
	"piscine"
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.NRune("Hello!", 3)) // l
	z01.PrintRune(piscine.NRune("Salut!", 2)) // a
	z01.PrintRune(piscine.NRune("Bye!", -1))  // 0 (nothing printed)
	z01.PrintRune(piscine.NRune("Bye!", 5))   // 0 (nothing printed)
	z01.PrintRune(piscine.NRune("Ola!", 4))   // !
	z01.PrintRune('\n')
}
```

---

### 🖥️ Expected Output

```
la!
```

---

## 🧩 ERROR & EDGE CASE CHECKS

| Input         | Explanation              | Output |
| ------------- | ------------------------ | ------ |
| `"Hello!", 0` | 0 is invalid             | `0`    |
| `"Hi", 5`     | Too big (only 2 letters) | `0`    |
| `"🙂Go", 1`   | Works fine with emojis   | `'🙂'` |

---

## 🗂️ FILES TO SUBMIT

✅ `nrune.go`

---

## 🧠 QUICK SUMMARY

| Concept      | Meaning                            |
| ------------ | ---------------------------------- |
| **Rune**     | A single Unicode character         |
| **range s**  | Loops through runes                |
| **count**    | Tracks rune positions              |
| **return 0** | When position invalid or not found |

---


## 🧩 **Task Name:** Compare

### 🧠 What does “Compare” mean?

To *compare two strings* means checking if they are:

* **Exactly the same**
* **One is smaller (comes before)**
* **One is larger (comes after)**

In Go, this is similar to how you compare words in a dictionary (lexicographical order).

### 📖 Example

| a        | b        | Result | Why              |
| -------- | -------- | ------ | ---------------- |
| "Hello!" | "Hello!" | 0      | They are equal   |
| "Salut!" | "lut!"   | -1     | "S" < "l"        |
| "Ola!"   | "Ol"     | 1      | "a" > "" (empty) |

👉 The function returns:

* `0` → if both strings are equal
* `-1` → if `a` < `b`
* `1` → if `a` > `b`

---

### ⚙️ What the task is asking

We must write a function:

```go
func Compare(a, b string) int
```

✅ It takes **two strings** (`a` and `b`)
✅ It returns **an integer** showing their order comparison

---

### 🧱 Step-by-Step Breakdown

#### 🪜 Step 1 — Loop through both strings

Compare each rune (character) of `a` and `b` **from start to end**.

#### 🪜 Step 2 — When a difference is found

* If `a[i] < b[i]` → return `-1`
* If `a[i] > b[i]` → return `1`

#### 🪜 Step 3 — If no difference found

* If both strings end together → return `0`
* If one string is longer → the longer one is “greater”.

---

### ✅ **Final Code** (`compare.go`)

```go
package piscine

// Compare compares two strings a and b lexicographically.
// Returns:
//   0 if a == b
//  -1 if a < b
//   1 if a > b
func Compare(a, b string) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	// If we reached here, all compared characters are equal
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}
	return 0
}
```

---

### 💻 **Test File** (`main.go`)

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Compare("Hello!", "Hello!")) // expected: 0
	fmt.Println(piscine.Compare("Salut!", "lut!"))   // expected: -1
	fmt.Println(piscine.Compare("Ola!", "Ol"))       // expected: 1
}
```

---

### 🧩 **Visual Flow Explanation**

Let’s visualize how `Compare("Salut!", "lut!")` runs:

| Step | i | a[i] | b[i] | Comparison           | Action        |
| ---- | - | ---- | ---- | -------------------- | ------------- |
| 1    | 0 | 'S'  | 'l'  | 'S' (83) < 'l' (108) | return **-1** |

So the function stops immediately and returns **-1**.

---

### 🧠 **How Go Compares Strings (ASCII Order)**

Each character has a numeric value (ASCII):

```
A = 65, B = 66, C = 67, ...
a = 97, b = 98, c = 99, ...
```

Uppercase letters come **before** lowercase letters.
So `"Salut!"` < `"lut!"` because `'S' (83)` < `'l' (108)`

---

### 🧾 **Expected Output**

```
$ go run .
0
-1
1
```

---

### 🗂️ **Files to Submit**

* `compare.go`

---
Perfect ✅ Got it — we’ll keep using the **Go Learning Log format** (with full documentation, explanations, diagrammatic step-by-step visual of code flow, and example output).

Here’s your **Go Learning Log entry for `AlphaCount`** 👇

---

## 🧩 `AlphaCount`**

### 🗂️ File to submit:

`alphacount.go`

---

### 🧠 **Purpose**

The function **`AlphaCount`** counts the number of alphabetic letters (A–Z, a–z) in a string.
It ignores numbers, spaces, and special symbols.

---

### 🧩 **Function Definition**

```go
package piscine

func AlphaCount(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			count++
		}
	}
	return count
}
```

---

### 🔍 **Explanation (Line-by-Line)**

| Line                       | Description                                                 |                         |                                                                                       |
| -------------------------- | ----------------------------------------------------------- | ----------------------- | ------------------------------------------------------------------------------------- |
| `count := 0`               | Initializes a counter to store the number of letters found. |                         |                                                                                       |
| `for _, r := range s`      | Iterates through each character (`rune`) in the string `s`. |                         |                                                                                       |
| `if (r >= 'A' && r <= 'Z') |                                                             | (r >= 'a' && r <= 'z')` | Checks if the character is within the ASCII range for uppercase or lowercase letters. |
| `count++`                  | Increases the counter when a valid letter is found.         |                         |                                                                                       |
| `return count`             | Returns the final number of letters in the string.          |                         |                                                                                       |

---

### 🧮 **Flow Diagram (Step-by-Step)**

**Example Input:** `"Hello 78 World!"`

```
s = "Hello 78 World!"
count = 0

┌──────────────────────────────────────────────┐
│ Loop through each character in the string:   │
│                                              │
│ H → letter? ✅ count = 1                     │
│ e → letter? ✅ count = 2                     │
│ l → letter? ✅ count = 3                     │
│ l → letter? ✅ count = 4                     │
│ o → letter? ✅ count = 5                     │
│ ' ' → ❌ skip                                │
│ 7 → ❌ skip                                  │
│ 8 → ❌ skip                                  │
│ ' ' → ❌ skip                                │
│ W → ✅ count = 6                             │
│ o → ✅ count = 7                             │
│ r → ✅ count = 8                             │
│ l → ✅ count = 9                             │
│ d → ✅ count = 10                            │
│ ! → ❌ skip                                  │
└──────────────────────────────────────────────┘

✅ Final count = 10
```

---

### 🧾 **Example Usage**

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}
```

---

### 💡 **Expected Output**

```
10
```

---

### 🧠 **Key Concept**

* **Runes** represent Unicode characters (not just bytes).
* Checking ASCII ranges ensures only **Latin letters** are counted.
* Efficient since it loops only once through the string.

---

Would you like me to include **a mini ASCII diagram** showing ASCII ranges (`A–Z` = 65–90, `a–z` = 97–122`) at the end of each similar exercise for quick reference?

Perfect ✅ — we’ll go **back to your preferred full learning format**, where I:
1️⃣ Explain what the task means in simple English
2️⃣ Build it step by step (like teaching)
3️⃣ Add a **visual illustration** of how the code runs (with loops and variable updates)
4️⃣ Then finish with the **final full code** and **file to submit** last.

Let’s redo the `Index` task exactly like that 👇

---

## 🧩 Task Name: `Index`

### 🧠 What does *Index* mean?

The **index** of a character or word is its **position** in a string.

Example:
👉 `"Hello!"`
Positions are numbered starting from **0**:

```
H   e   l   l   o   !
0   1   2   3   4   5
```

If we search for `"l"`, the first `"l"` appears at position **2**.

So, `Index("Hello!", "l")` should return **2**.
If we search for something not inside the text — say `"hOl"` — it returns **-1**.

---

### ⚙️ What the task is asking

We must write a function:

```go
func Index(s string, toFind string) int {
}
```

✅ It receives two strings:

* `s` → the main text
* `toFind` → the word we’re looking for

✅ It returns:

* The **starting position** (index) of where `toFind` first appears inside `s`
* Or **-1** if it doesn’t appear at all.

---

### 🧱 Let’s build it step-by-step

#### Step 1 — Define the function

```go
func Index(s string, toFind string) int {
```

We create our function with two string inputs.

---

#### Step 2 — Loop through the main string `s`

We want to look through every position of `s` where `toFind` *could* start.

We’ll loop from the beginning (index `0`) until there’s enough room left for the word to fit.

```go
for i := 0; i <= len(s)-len(toFind); i++ {
```

🧩 Example:
If `s = "Salut!"` (length 6)
and `toFind = "alu"` (length 3)
then the loop goes from `i = 0` to `i = 3`.

That gives us 4 possible start points:

```
i = 0  → "Sal"
i = 1  → "alu"
i = 2  → "lu!"
i = 3  → "u!"
```

---

#### Step 3 — Check if substring matches

Now, at each position, we take a “slice” of the main string that’s the same length as `toFind`:

```go
if s[i:i+len(toFind)] == toFind {
    return i
}
```

💡 Example:

```
i = 0 → s[0:3] = "Sal" ≠ "alu"
i = 1 → s[1:4] = "alu" ✅ MATCH!
```

So we return **i = 1**.

---

#### Step 4 — If nothing found, return -1

If we finish the loop and never find a match:

```go
return -1
```

That tells us the substring doesn’t exist.

---

### ✅ Final Code (`index.go`)

```go
package piscine

// Index returns the first position of toFind in s.
// If not found, it returns -1.
func Index(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
```

---

### 🔍 Visual Simulation (How the loop runs)

Let’s visualize with
`s = "Salut!"`
`toFind = "alu"`

| i | s[i:i+3] | Comparison      | Match? | Return |
| - | -------- | --------------- | ------ | ------ |
| 0 | "Sal"    | "Sal" == "alu"? | ❌      | —      |
| 1 | "alu"    | "alu" == "alu"? | ✅      | 1      |
| 2 | "lu!"    | "lu!" == "alu"? | ❌      | —      |
| 3 | "u!"     | too short       | —      | —      |

➡️ First match found at index **1** → function stops → returns `1`.

---

### 🧪 Test File (`main.go`)

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))   // 2
	fmt.Println(piscine.Index("Salut!", "alu")) // 1
	fmt.Println(piscine.Index("Ola!", "hOl"))   // -1
}
```

---

### 💻 How to Run

Make sure you have:

```
piscine/index.go
main.go
```

Then run:

```bash
$ go run .
```

**Expected output:**

```
2
1
-1
```

---

### 🗂️ File to Submit

> ✅ **index.go**

---

Would you like me to continue this same deep-visual format with the next task (`Concat`)?
Perfect 😎 Let’s go through this **step-by-step, visual, beginner-friendly** like our previous Go lessons — you’ll understand **exactly how `Concat` works** from inside out.

---

## 🧩 Task Name: `Concat`

---

### 🧠 What does *concatenation* mean?

“Concatenation” simply means **joining** two strings together — like putting two pieces of text side by side.

Example:

```
"Hello!" + " How are you?" = "Hello! How are you?"
```

So, the task is asking us to join two given strings and return the result.

---

### ⚙️ What the task is asking

We must create a function:

```go
func Concat(str1 string, str2 string) string {
}
```

✅ It takes two text strings — `str1` and `str2`
✅ It returns a single string that is `str1` followed by `str2`
✅ No need to print inside the function — just **return** the result

---

### 🧱 Step-by-Step Construction

#### Step 1 — Start the function

We begin by defining our function:

```go
func Concat(str1 string, str2 string) string {
```

This means:

* Input: two strings
* Output: one string

---

#### Step 2 — Combine (join) both strings

In Go, we can **join two strings** easily using the `+` operator.

```go
result := str1 + str2
```

🧩 Example:

```
str1 = "Hello!"
str2 = " How are you?"

result = "Hello! How are you?"
```

---

#### Step 3 — Return the result

Finally, we just return it:

```go
return result
```

---

### ✅ Final Code (`concat.go`)

```go
package piscine

// Concat joins two strings (str1 and str2) together and returns the new string.
func Concat(str1 string, str2 string) string {
	result := str1 + str2
	return result
}
```

---

### 🔍 Visual Simulation — How It Works

Let’s visualize what happens when you run this line in the test file:

```go
piscine.Concat("Hello!", " How are you?")
```

| Step | Variable | Value                 | Explanation      |
| ---- | -------- | --------------------- | ---------------- |
| 1    | str1     | "Hello!"              | first argument   |
| 2    | str2     | " How are you?"       | second argument  |
| 3    | result   | "Hello! How are you?" | joined using `+` |
| 4    | return   | "Hello! How are you?" | final output     |

🧩 So the function *glues* both together and sends back the full sentence.

---

### 🧪 Test File (`main.go`)

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Concat("Hello!", " How are you?"))
}
```

---

### 💻 How to Run

Make sure you have:

```
piscine/concat.go
main.go
```

Then in your terminal:

```bash
$ go run .
```

**Expected Output:**

```
Hello! How are you?
```

---

### 🗂️ File to Submit

> ✅ **concat.go**

---
Perfect 👏 Let’s break this **Go Lang `IsUpper` task** down like we always do — with full explanation, diagram-style visualization, and example testing.

---

## 🧩 Task Name: `IsUpper`

---

### 🧠 What does *IsUpper* mean?

When we say something is **uppercase**, we mean **all letters are capital letters** — from **A to Z** only.

Examples:

| String     | Is it uppercase? | Reason                              |
| ---------- | ---------------- | ----------------------------------- |
| `"HELLO"`  | ✅ true           | all letters A–Z                     |
| `"HELLO!"` | ❌ false          | includes `!` (not a letter)         |
| `"GoLang"` | ❌ false          | includes lowercase letters          |
| `""`       | ✅ true           | (empty string = no lowercase found) |

---

### ⚙️ What the task is asking

We must create a function:

```go
func IsUpper(s string) bool
```

✅ Input → a string `s`
✅ Output → `true` if *every character* is uppercase (A–Z), otherwise `false`
✅ We must check each character carefully

---

### 🧱 Step-by-Step Construction

#### Step 1 — Function definition

```go
func IsUpper(s string) bool {
```

We’re building a function that returns a `bool` (true or false).

---

#### Step 2 — Loop through each character (rune)

We’ll use a `for range` loop to go through every letter in the string:

```go
for _, char := range s {
```

👉 `_` means we ignore the index.
👉 `char` represents each character (rune) in the string.

---

#### Step 3 — Check if it’s between `'A'` and `'Z'`

In Go, characters like `'A'`, `'B'`, `'Z'` have **rune (Unicode) values**.
We can compare them directly:

```go
if char < 'A' || char > 'Z' {
	return false
}
```

This means:

* If it’s smaller than `'A'` or larger than `'Z'`, it’s not uppercase.

---

#### Step 4 — If loop finishes with no problems → return true

```go
return true
```

That means every character passed the uppercase test.

---

### ✅ Final Code (`isupper.go`)

```go
package piscine

// IsUpper checks if a string contains only uppercase letters (A–Z).
// Returns true if all are uppercase, otherwise false.
func IsUpper(s string) bool {
	for _, char := range s {
		if char < 'A' || char > 'Z' {
			return false
		}
	}
	return true
}
```

---

### 🔍 Visual Simulation — How It Works

Let’s see how `"HELLO!"` is checked step-by-step 👇

| Step | Character | Unicode | Is between ‘A’ and ‘Z’? | Result           |
| ---- | --------- | ------- | ----------------------- | ---------------- |
| 1    | `'H'`     | 72      | ✅ Yes                   | continue         |
| 2    | `'E'`     | 69      | ✅ Yes                   | continue         |
| 3    | `'L'`     | 76      | ✅ Yes                   | continue         |
| 4    | `'L'`     | 76      | ✅ Yes                   | continue         |
| 5    | `'O'`     | 79      | ✅ Yes                   | continue         |
| 6    | `'!'`     | 33      | ❌ No                    | return **false** |

So the program stops right there and returns `false`.

---

### 🧪 Test File (`main.go`)

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsUpper("HELLO"))  // true
	fmt.Println(piscine.IsUpper("HELLO!")) // false
}
```

---

### 💻 How to Run

Make sure you have:

```
piscine/isupper.go
main.go
```

Then run:

```bash
$ go run .
```

**Expected Output:**

```
true
false
```

---

### 🗂️ File to Submit

> ✅ **isupper.go**

---
Perfect 😎 — now let’s continue your Go learning journey with **`IsLower`**, using the *same step-by-step + visualization* style you requested!

---

## 🧩 Task Name: `IsLower`

---

### 🧠 What does “IsLower” mean?

We want to check if a string is **completely lowercase** — that means **every character must be between `'a'` and `'z'`**.

Examples:

| String     | Result  | Reason                        |
| ---------- | ------- | ----------------------------- |
| `"hello"`  | ✅ true  | all are lowercase letters     |
| `"hello!"` | ❌ false | includes `!`                  |
| `"Hello"`  | ❌ false | has an uppercase `H`          |
| `""`       | ✅ true  | (empty = no violations found) |

---

### ⚙️ What the task asks

We must create a function:

```go
func IsLower(s string) bool
```

✅ Input → a string
✅ Output → `true` if *all characters* are lowercase letters, otherwise `false`.

---

## 🧱 Step-by-Step Construction

### Step 1 — Function definition

```go
func IsLower(s string) bool {
```

We’re defining a function that takes a string and returns a boolean.

---

### Step 2 — Loop through each character

We’ll use a **for-range loop** to go through each character:

```go
for _, char := range s {
```

* `_` → we ignore the index (we don’t need it).
* `char` → represents each character (as a **rune**).

---

### Step 3 — Check if it’s between `'a'` and `'z'`

Just like before, we use Unicode comparisons:

```go
if char < 'a' || char > 'z' {
	return false
}
```

🧩 Explanation:

* `'a'` = 97
* `'z'` = 122
  If any character is outside this range, it’s **not lowercase**.

---

### Step 4 — If all checks passed, return true

If the loop finishes with no problems, we can safely return:

```go
return true
```

---

## ✅ Final Code — `islower.go`

```go
package piscine

// IsLower checks if a string contains only lowercase letters (a–z).
// Returns true if all are lowercase, otherwise false.
func IsLower(s string) bool {
	for _, char := range s {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}
```

---

## 🔍 Visual Simulation — `"hello!"`

Let’s see what happens line by line 👇

| Step | Character | Unicode | Is it between ‘a’ and ‘z’? | Result           |
| ---- | --------- | ------- | -------------------------- | ---------------- |
| 1    | `'h'`     | 104     | ✅ Yes                      | continue         |
| 2    | `'e'`     | 101     | ✅ Yes                      | continue         |
| 3    | `'l'`     | 108     | ✅ Yes                      | continue         |
| 4    | `'l'`     | 108     | ✅ Yes                      | continue         |
| 5    | `'o'`     | 111     | ✅ Yes                      | continue         |
| 6    | `'!'`     | 33      | ❌ No                       | return **false** |

✅ `"hello"` → true
❌ `"hello!"` → false

---

## 🧪 Test File — `main.go`

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsLower("hello"))   // true
	fmt.Println(piscine.IsLower("hello!"))  // false
}
```

---

## 💻 Run the Test

```bash
$ go run .
```

**Output:**

```
true
false
```

---

## 🗂️ File to Submit

> ✅ **islower.go**

---
Excellent 👏 We’re now moving into **Level 12’s** `IsAlpha` — a slightly more interesting function that combines what we learned in **IsLower**, **IsUpper**, and a bit more logic!

Let’s walk through it step-by-step, visually, and in full beginner-friendly detail.

---

## 🧩 Task Name: `IsAlpha`

---

### 🧠 What does “IsAlpha” mean?

> “Alpha” means alphabetic — so **letters (A–Z and a–z)**.
>
> But in this task, they also allow **numbers (0–9)** — so it’s actually **alphanumeric**.

✅ True → if the string contains only **letters and digits**, or is **empty**
❌ False → if it contains **any symbol, space, or punctuation**

---

### 🧾 Examples

| String                  | Result  | Why                         |
| ----------------------- | ------- | --------------------------- |
| `"HelloHowareyou"`      | ✅ true  | only letters                |
| `"Whatsthis4"`          | ✅ true  | contains number 4 (allowed) |
| `"Hello! How are you?"` | ❌ false | has spaces + punctuation    |
| `""`                    | ✅ true  | empty = acceptable          |

---

### 🧱 Step-by-Step Breakdown

---

### Step 1 — Function definition

```go
func IsAlpha(s string) bool {
```

We’re defining the function that receives a string and returns a boolean.

---

### Step 2 — Loop through every character

```go
for _, char := range s {
```

We go through each **rune** (Unicode character) in the string.

---

### Step 3 — Define the valid ranges

We must allow:

* `'a'`–`'z'` → lowercase letters
* `'A'`–`'Z'` → uppercase letters
* `'0'`–`'9'` → digits

Anything else is **not allowed**.

---

### Step 4 — Check if character is *not* alphanumeric

We use conditions:

```go
if !(char >= 'a' && char <= 'z') &&
   !(char >= 'A' && char <= 'Z') &&
   !(char >= '0' && char <= '9') {
	return false
}
```

Let’s break it down:

| Condition                    | Meaning   |
| ---------------------------- | --------- |
| `char >= 'a' && char <= 'z'` | lowercase |
| `char >= 'A' && char <= 'Z'` | uppercase |
| `char >= '0' && char <= '9'` | digit     |

If the character does **not** match any of those ranges, we instantly return `false`.

---

### Step 5 — If loop finishes fine

If we pass all characters, then the string is valid.

```go
return true
```

---

## ✅ Final Code — `isalpha.go`

```go
package piscine

// IsAlpha checks if a string contains only alphanumeric characters (A-Z, a-z, 0-9).
// Returns true if valid or empty, false otherwise.
func IsAlpha(s string) bool {
	for _, char := range s {
		if !(char >= 'a' && char <= 'z') &&
			!(char >= 'A' && char <= 'Z') &&
			!(char >= '0' && char <= '9') {
			return false
		}
	}
	return true
}
```

---

## 🔍 Visual Flow Example — `"What's this 4?"`

| Step | Char  | Unicode | Allowed? | Result           |
| ---- | ----- | ------- | -------- | ---------------- |
| 1    | `'W'` | 87      | ✅ (A–Z)  | continue         |
| 2    | `'h'` | 104     | ✅ (a–z)  | continue         |
| 3    | `'a'` | 97      | ✅        | continue         |
| 4    | `'t'` | 116     | ✅        | continue         |
| 5    | `'`'` | 39      | ❌        | return **false** |

💥 The moment we hit `'` (apostrophe), we exit and return **false** — no need to check the rest.

---

### ✅ Example 2 — `"Whatsthis4"`

| Step | Char | Unicode       | Allowed?        | Result |
| ---- | ---- | ------------- | --------------- | ------ |
| W    | 87   | ✅             | continue        |        |
| h    | 104  | ✅             | continue        |        |
| a    | 97   | ✅             | continue        |        |
| t    | 116  | ✅             | continue        |        |
| s    | 115  | ✅             | continue        |        |
| t    | 116  | ✅             | continue        |        |
| h    | 104  | ✅             | continue        |        |
| i    | 105  | ✅             | continue        |        |
| s    | 115  | ✅             | continue        |        |
| 4    | 52   | ✅ (digit)     | continue        |        |
| —    | —    | ✅ all checked | return **true** |        |

---

## 🧪 Test File — `main.go`

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsAlpha("Hello! How are you?")) // false
	fmt.Println(piscine.IsAlpha("HelloHowareyou"))     // true
	fmt.Println(piscine.IsAlpha("What's this 4?"))     // false
	fmt.Println(piscine.IsAlpha("Whatsthis4"))         // true
}
```

---

## 🧠 Output

```bash
$ go run .
false
true
false
true
```

---

## 🗂️ File to Submit

> ✅ **isalpha.go**

---

Perfect 👍 We’re moving smoothly through Level 12 — this time with **`IsNumeric`**, a close cousin of `IsAlpha`, but much simpler to reason about!

Let’s go through it in your preferred **detailed documentation format** — full explanation, line-by-line logic, and visual walkthrough 👇

---

## 🧩 Task Name: `IsNumeric`

---

### 🧠 Concept & Goal

> We need to check whether the given string contains **only numbers (0–9)**.

That means:

* ✅ `"12345"` → **true**
* ✅ `""` (empty string) → **true**
* ❌ `"123a5"` → **false**
* ❌ `"12,345"` → **false**
* ❌ `"12 34"` → **false**

---

### 🧾 Problem Understanding

We’ll loop through every **character (rune)** in the string and confirm that:

* Each rune is between `'0'` and `'9'`.

If we find even one character that is **not** in that range → the function immediately returns `false`.

---

## ⚙️ Step-by-Step Explanation

---

### Step 1 — Function Definition

```go
func IsNumeric(s string) bool {
```

We’re defining a function that takes a string `s` and returns a boolean result (`true` or `false`).

---

### Step 2 — Loop Through Each Character

```go
for _, char := range s {
```

We loop through every rune (Go’s way of handling Unicode-safe characters) inside the string.

---

### Step 3 — Check Valid Range

```go
if char < '0' || char > '9' {
	return false
}
```

This condition means:

| Condition    | Meaning                               |
| ------------ | ------------------------------------- |
| `char < '0'` | It’s smaller than `'0'` (not a digit) |
| `char > '9'` | It’s greater than `'9'` (not a digit) |

If either happens → 🚫 the character is **not numeric**.
So, we **return false immediately**.

---

### Step 4 — After Loop Ends

If the loop checks every character and finds all valid (digits only), then:

```go
return true
```

✅ Everything is numeric!

---

## ✅ Final Code — `isnumeric.go`

```go
package piscine

// IsNumeric checks if a string contains only numeric characters (0–9).
// Returns true if all characters are digits or if the string is empty.
func IsNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
```

---

## 🔍 Visual Flow Example — `"01,02,03"`

| Step | Char  | Unicode | Check             | Result           |
| ---- | ----- | ------- | ----------------- | ---------------- |
| 1    | `'0'` | 48      | between '0'-'9' ✅ | continue         |
| 2    | `'1'` | 49      | ✅                 | continue         |
| 3    | `','` | 44      | ❌                 | return **false** |

🚫 As soon as the comma is found, the function exits — no need to check the rest.

---

### ✅ Example 2 — `"010203"`

| Step | Char  | Unicode   | Check           | Result   |
| ---- | ----- | --------- | --------------- | -------- |
| 1    | `'0'` | 48        | ✅               | continue |
| 2    | `'1'` | 49        | ✅               | continue |
| 3    | `'0'` | 48        | ✅               | continue |
| 4    | `'2'` | 50        | ✅               | continue |
| 5    | `'0'` | 48        | ✅               | continue |
| 6    | `'3'` | 51        | ✅               | continue |
| ✅    | End   | All valid | return **true** |          |

---

## 🧪 Test File — `main.go`

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsNumeric("010203"))  // true
	fmt.Println(piscine.IsNumeric("01,02,03")) // false
	fmt.Println(piscine.IsNumeric("123abc"))   // false
	fmt.Println(piscine.IsNumeric(""))         // true
}
```

---

## 🧠 Output

```bash
$ go run .
true
false
false
true
```

---

## 🗂️ File to Submit

> ✅ **isnumeric.go**

---

Perfect 🧩
We’re now on **Level 13 — `IsPrintable`**, one of the easiest to understand *conceptually*, but quite powerful in text validation.

Let’s go through it in our **standard learning + visualization format** — full explanation, line-by-line logic, and examples 👇

---

## 🧩 Task Name: `IsPrintable`

---

### 🧠 What Does “Printable” Mean?

In computer terms, **printable characters** are the ones you can *see* — letters, digits, punctuation, spaces, etc.

They **exclude**:

* Special control characters like `\n` (newline), `\t` (tab), `\r` (carriage return), etc.

🧮 Printable characters in ASCII are those between:

```
' ' (space)  → ASCII 32  
'~' (tilde) → ASCII 126
```

So anything in the range **[32, 126]** is printable.

---

### 🧱 What the task is asking

We need to write a function:

```go
func IsPrintable(s string) bool
```

✅ It returns `true` if every character (rune) in the string is printable
❌ It returns `false` if any character is not printable (like `\n` or `\t`)

---

## ⚙️ Step-by-Step Breakdown

---

### Step 1 — Function definition

```go
func IsPrintable(s string) bool {
```

We define the function that takes a string and returns a boolean value.

---

### Step 2 — Loop through every rune

```go
for _, r := range s {
```

We use `range` to safely loop through all characters (runes).
Each rune is assigned to `r`.

---

### Step 3 — Check if rune is printable

```go
if r < 32 || r > 126 {
	return false
}
```

| Condition | Meaning                                                                     |
| --------- | --------------------------------------------------------------------------- |
| `r < 32`  | The character is a **control character** (like `\n`, `\t`, etc.)            |
| `r > 126` | The character is **beyond printable ASCII** (like emoji or Unicode symbols) |

If either is true → ❌ Not printable → return **false** immediately.

---

### Step 4 — If everything passes

```go
return true
```

✅ If the loop finishes and finds no invalid rune, the whole string is printable.

---

## ✅ Final Code — `isprintable.go`

```go
package piscine

// IsPrintable returns true if all characters in the string
// are printable ASCII characters (from space ' ' to '~').
func IsPrintable(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}
```

---

## 🧩 Visualization — How It Works

Let’s walk through the two test examples:

### 🧠 Example 1: `"Hello"`

| Char | ASCII | Check          | Result   |
| ---- | ----- | -------------- | -------- |
| H    | 72    | between 32–126 | ✅        |
| e    | 101   | ✅              | ✅        |
| l    | 108   | ✅              | ✅        |
| l    | 108   | ✅              | ✅        |
| o    | 111   | ✅              | ✅        |
| ✅    | End   | All good       | **true** |

→ ✅ Every character is printable.

---

### 🧠 Example 2: `"Hello\n"`

| Char | ASCII | Check            | Result           |
| ---- | ----- | ---------------- | ---------------- |
| H    | 72    | ✅                | continue         |
| e    | 101   | ✅                | continue         |
| l    | 108   | ✅                | continue         |
| l    | 108   | ✅                | continue         |
| o    | 111   | ✅                | continue         |
| `\n` | 10    | ❌ (less than 32) | **return false** |

→ 🚫 Not printable.

---

## 🧪 Test File — `main.go`

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrintable("Hello"))   // true
	fmt.Println(piscine.IsPrintable("Hello\n")) // false
	fmt.Println(piscine.IsPrintable(" "))       // true
	fmt.Println(piscine.IsPrintable("Hi😊"))    // false (emoji not ASCII printable)
}
```

---

## 🧠 Expected Output

```bash
$ go run .
true
false
true
false
```

---

## 🗂️ File to Submit

> ✅ **isprintable.go**

---

Perfect 👏 We’re moving deeper into **string transformations** now — specifically the `ToUpper` function, which converts lowercase letters into uppercase.

Let’s follow our complete learning pattern again:
**concept → logic → visualization → code → test → explanation of output → file to submit.**

---

## 🧩 Task Name: `ToUpper`

---

### 🧠 What Does “ToUpper” Mean?

👉 “ToUpper” means **convert all lowercase letters (a–z) into uppercase (A–Z)**.

💡 Example:

```
Input:  "Hello! How are you?"
Output: "HELLO! HOW ARE YOU?"
```

🧮 Remember:

* Only the **letters a–z** should change.
* Numbers, spaces, punctuation marks, etc. must **stay the same**.

---

## ⚙️ Step-by-Step Breakdown

---

### Step 1 — Function Definition

```go
func ToUpper(s string) string {
```

This defines our function that takes a **string** and returns a **new string**.

---

### Step 2 — Convert the string to runes

```go
runes := []rune(s)
```

We do this so Go can handle all Unicode characters safely (since strings are immutable and can contain multi-byte characters).

---

### Step 3 — Loop through each rune (character)

```go
for i, r := range runes {
```

We loop through the string character by character using `range`.
`i` is the index, and `r` is the character (rune).

---

### Step 4 — Check if it’s a lowercase letter

```go
if r >= 'a' && r <= 'z' {
	runes[i] = r - 32
}
```

💡 Why `r - 32`?

In ASCII:

| Char | ASCII | Uppercase | Difference |     |
| ---- | ----- | --------- | ---------- | --- |
| a    | 97    | A         | 65         | -32 |
| b    | 98    | B         | 66         | -32 |
| z    | 122   | Z         | 90         | -32 |

So subtracting **32** converts lowercase to uppercase!

🧩 Visualization:

```
Before:  'a' (97)
After:   'A' (65)
```

---

### Step 5 — Convert runes back to string and return

```go
return string(runes)
```

✅ Done! Now we have a new uppercase version of the string.

---

## ✅ Final Code — `toupper.go`

```go
package piscine

// ToUpper converts all lowercase letters in a string to uppercase.
func ToUpper(s string) string {
	runes := []rune(s)

	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = r - 32 // convert to uppercase
		}
	}

	return string(runes)
}
```

---

## 🧩 Visualization — How It Works

### Example: `"Hello! How are you?"`

| Step | Char  | ASCII | Action               | New Char |
| ---- | ----- | ----- | -------------------- | -------- |
| 1    | H     | 72    | Already uppercase    | H        |
| 2    | e     | 101   | lowercase → -32 → 69 | E        |
| 3    | l     | 108   | lowercase → -32 → 76 | L        |
| 4    | l     | 108   | lowercase → -32 → 76 | L        |
| 5    | o     | 111   | lowercase → -32 → 79 | O        |
| 6    | !     | 33    | not a letter         | !        |
| 7    | space | 32    | not a letter         | space    |
| 8    | H     | 72    | already uppercase    | H        |
| 9    | o     | 111   | lowercase → -32 → 79 | O        |
| 10   | w     | 119   | lowercase → -32 → 87 | W        |
| ...  | ...   | ...   | ...                  | ...      |

✅ Final result → `"HELLO! HOW ARE YOU?"`

---

## 🧪 Test File — `main.go`

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ToUpper("Hello! How are you?"))
	fmt.Println(piscine.ToUpper("gOphErS 2025!"))
	fmt.Println(piscine.ToUpper("123 hello WORLD"))
}
```

---

## 🧠 Expected Output

```bash
$ go run .
HELLO! HOW ARE YOU?
GOPHERS 2025!
123 HELLO WORLD
```

---

## 🗂️ File to Submit

> ✅ **toupper.go**

---

Perfect 👏 You’re now doing the **ToLower** task — it’s the *opposite* of `ToUpper`.
We’ll follow the same learning pattern again: **concept → step-by-step logic → code → visualization → test → output**.

---

## 🧩 Task Name: `ToLower`

---

### 🧠 What Does “ToLower” Mean?

👉 “ToLower” means **convert all uppercase letters (A–Z) into lowercase (a–z)**.

💡 Example:

```
Input:  "Hello! How are you?"
Output: "hello! how are you?"
```

✅ Only the **letters A–Z** should change.
Everything else (numbers, spaces, punctuation) stays the same.

---

## ⚙️ Step-by-Step Breakdown

---

### Step 1 — Define the function

```go
func ToLower(s string) string {
```

It takes a **string** as input and returns a **string** as output.

---

### Step 2 — Convert string to rune slice

```go
runes := []rune(s)
```

We do this so we can modify each character (since strings in Go are immutable).

---

### Step 3 — Loop through each rune

```go
for i, r := range runes {
```

`i` is the position, and `r` is the rune (character).

---

### Step 4 — Check if it’s uppercase (A–Z)

```go
if r >= 'A' && r <= 'Z' {
	runes[i] = r + 32
}
```

💡 Why `r + 32`?

| Char | ASCII | Lowercase | Difference |
| ---- | ----- | --------- | ---------- |
| A    | 65    | a         | +32        |
| B    | 66    | b         | +32        |
| Z    | 90    | z         | +32        |

✅ So adding **32** converts uppercase → lowercase.

---

### Step 5 — Convert runes back to string

```go
return string(runes)
```

---

## ✅ Final Code — `tolower.go`

```go
package piscine

// ToLower converts all uppercase letters in a string to lowercase.
func ToLower(s string) string {
	runes := []rune(s)

	for i, r := range runes {
		if r >= 'A' && r <= 'Z' {
			runes[i] = r + 32 // convert to lowercase
		}
	}

	return string(runes)
}
```

---

## 🧩 Visualization — How It Works

### Example: `"Hello! How are you?"`

| Step | Char  | ASCII | Action                | New Char |
| ---- | ----- | ----- | --------------------- | -------- |
| 1    | H     | 72    | uppercase → +32 → 104 | h        |
| 2    | e     | 101   | already lowercase     | e        |
| 3    | l     | 108   | already lowercase     | l        |
| 4    | l     | 108   | already lowercase     | l        |
| 5    | o     | 111   | already lowercase     | o        |
| 6    | !     | 33    | not a letter          | !        |
| 7    | space | 32    | not a letter          | space    |
| 8    | H     | 72    | uppercase → +32 → 104 | h        |
| 9    | o     | 111   | already lowercase     | o        |
| 10   | w     | 119   | already lowercase     | w        |

✅ Final result → `"hello! how are you?"`

---

## 🧪 Test File — `main.go`

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ToLower("Hello! How are you?"))
	fmt.Println(piscine.ToLower("GOPHERS 2025!"))
	fmt.Println(piscine.ToLower("123 HELLO world"))
}
```

---

## 🧠 Expected Output

```bash
$ go run .
hello! how are you?
gophers 2025!
123 hello world
```

---

## 🗂️ File to Submit

> ✅ **tolower.go**

---

Perfect 🎯 This one — **`PrintNbrInOrder`** — is a classic **sorting digits + printing** problem in Go.
We’ll follow our usual pattern again: **concept → breakdown → code → visual flow → output → final file**

---

## 🧩 Task Name: `PrintNbrInOrder`

---

### 🧠 Goal

Write a function that:

* Takes an integer `n`
* Extracts all its digits
* Prints them in **ascending order** using `z01.PrintRune`
* Works for all integers **except negatives**
* Doesn’t convert to `int64`

---

### 🧠 Example

| Input         | Digits    | Sorted    | Printed Output |
| ------------- | --------- | --------- | -------------- |
| `321`         | [3, 2, 1] | [1, 2, 3] | `123`          |
| `0`           | [0]       | [0]       | `0`            |
| `321` (again) | [3, 2, 1] | [1, 2, 3] | `123`          |

🧾 Combined Output (from usage):

```
1230123
```

---

## ⚙️ Step-by-Step Thinking

---

### Step 1 — Handle the special case `n == 0`

If `n` is zero, we directly print `'0'` and stop.

```go
if n == 0 {
	z01.PrintRune('0')
	return
}
```

---

### Step 2 — Extract digits of the number

We’ll repeatedly divide by 10 to get each digit.

Example: `321`

```
digits = [1, 2, 3]
```

Code:

```go
var digits []int
for n > 0 {
	digits = append(digits, n%10) // get last digit
	n /= 10                       // remove last digit
}
```

---

### Step 3 — Sort the digits (ascending order)

We use a **simple sorting algorithm** like **bubble sort**, or Go’s built-in `sort.Ints()` if allowed,
but since this challenge often expects manual logic, let’s use bubble sort (simple and visible).

```go
for i := 0; i < len(digits); i++ {
	for j := i + 1; j < len(digits); j++ {
		if digits[i] > digits[j] {
			digits[i], digits[j] = digits[j], digits[i]
		}
	}
}
```

---

### Step 4 — Print digits in order

Convert each integer digit into its rune (character).

💡 ASCII code for `'0'` is **48**
So `'0' + digit` gives the right character.

```go
for _, d := range digits {
	z01.PrintRune(rune(d + '0'))
}
```

---

## ✅ Full Code — `printnbrinorder.go`

```go
package piscine

import "github.com/01-edu/z01"

// PrintNbrInOrder prints digits of an integer in ascending order
func PrintNbrInOrder(n int) {
	// Handle zero directly
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Extract digits into a slice
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	// Sort digits in ascending order (simple bubble sort)
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}

	// Print digits as characters
	for _, d := range digits {
		z01.PrintRune(rune(d + '0'))
	}
}
```

---

## 🔍 Visualization — Example: `n = 321`

| Step | Operation      | Result      | Comment                 |
| ---- | -------------- | ----------- | ----------------------- |
| 1    | Extract digits | [1, 2, 3]   | We collect by `% 10`    |
| 2    | Sort digits    | [1, 2, 3]   | Sorted ascending        |
| 3    | Print each     | '1' '2' '3' | Use ASCII `'0'` + digit |

✅ Output → `123`

If input was `0`, we directly print `0`.

---

## 🧪 Test File — `main.go`

```go
package main

import "piscine"

func main() {
	piscine.PrintNbrInOrder(321)
	piscine.PrintNbrInOrder(0)
	piscine.PrintNbrInOrder(321)
}
```

---

## 💻 Expected Output

```bash
$ go run . | cat -e
1230123$
$
```

---

## 🗂️ File to Submit

> ✅ **printnbrinorder.go**

---

Excellent 👌 — this one, **`TrimAtoi`**, is a step up from previous string–number exercises.
It’s about **parsing digits** buried in a messy string — just like a **mini number extractor**.
Let’s go through it carefully, with our usual **full explanation, logic, and visualization**.

---

## 🧩 Task Name: `TrimAtoi`

---

### 🧠 Goal

You must:

1. Find all **digits** in a string and turn them into a single integer.
2. Detect a **negative sign `-`** **before** any digits — if it exists, the number becomes negative.
3. Return **0** if no digits exist.
4. Ignore every other character (letters, symbols, spaces, etc.).

---

### 🧠 Example Behavior

| Input            | Extracted Digits   | Sign                         | Final Output |
| ---------------- | ------------------ | ---------------------------- | ------------ |
| `"12345"`        | `12345`            | +                            | `12345`      |
| `"str123ing45"`  | `12345`            | +                            | `12345`      |
| `"012 345"`      | `012345` → `12345` | +                            | `12345`      |
| `"Hello World!"` | *(none)*           | +                            | `0`          |
| `"sd+x1fa2W3s4"` | `1234`             | +                            | `1234`       |
| `"sd-x1fa2W3s4"` | `1234`             | -                            | `-1234`      |
| `"sdx1-fa2W3s4"` | `1234`             | + (because - is after digit) | `1234`       |

---

## ⚙️ Step-by-Step Plan

### 🪜 Step 1 — Initialize values

We’ll need:

* `result` (the final integer)
* `sign` (1 or -1)
* `foundNumber` (boolean to check if digits exist)

---

### 🪜 Step 2 — Loop through each character in the string

We check each rune:

* If it’s a **‘-’** and **no digit has been found yet**, mark `sign = -1`
* If it’s a **digit (‘0’–‘9’)**:

  * Mark `foundNumber = true`
  * Add it to `result` by multiplying the previous result by 10 and adding the new digit.

Example:

```
Input: "sd-x1fa2W3s4"

→ s (ignore)
→ d (ignore)
→ - (sign = -1)
→ x (ignore)
→ 1 (result = 1)
→ f (ignore)
→ a (ignore)
→ 2 (result = 12)
→ W (ignore)
→ 3 (result = 123)
→ s (ignore)
→ 4 (result = 1234)
```

---

### 🪜 Step 3 — Apply the sign at the end

After the loop:

```go
return result * sign
```

If no digits were found, result remains 0.

---

## ✅ Full Code — `trimatoi.go`

```go
package piscine

// TrimAtoi converts all digits inside a string into an integer
// If a '-' sign appears before any digit, the number is negative.
// Otherwise, it returns 0 if no digits exist.
func TrimAtoi(s string) int {
	sign := 1       // Default: positive
	result := 0     // The integer we’ll build
	foundNumber := false

	for _, char := range s {
		if char == '-' && !foundNumber {
			sign = -1 // Negative sign before digits
		}
		if char >= '0' && char <= '9' {
			foundNumber = true
			result = result*10 + int(char-'0') // Build number
		}
	}

	return result * sign
}
```

---

## 🔍 Visualization — Example: `"sd-x1fa2W3s4"`

| Step    | Char | Action                   | result    | sign |
| ------- | ---- | ------------------------ | --------- | ---- |
| 1       | `s`  | ignored                  | 0         | +    |
| 2       | `d`  | ignored                  | 0         | +    |
| 3       | `-`  | negative (before digits) | 0         | -    |
| 4       | `x`  | ignored                  | 0         | -    |
| 5       | `1`  | add digit                | 1         | -    |
| 6       | `f`  | ignored                  | 1         | -    |
| 7       | `a`  | ignored                  | 1         | -    |
| 8       | `2`  | add digit                | 12        | -    |
| 9       | `W`  | ignored                  | 12        | -    |
| 10      | `3`  | add digit                | 123       | -    |
| 11      | `s`  | ignored                  | 123       | -    |
| 12      | `4`  | add digit                | 1234      | -    |
| ✅ Final | —    | Multiply by sign         | **-1234** | —    |

---

## 🧪 Test File — `main.go`

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.TrimAtoi("12345"))
	fmt.Println(piscine.TrimAtoi("str123ing45"))
	fmt.Println(piscine.TrimAtoi("012 345"))
	fmt.Println(piscine.TrimAtoi("Hello World!"))
	fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sdx1+fa2W3s4"))
}
```

---

## 🧾 Expected Output

```
12345
12345
12345
0
1234
-1234
1234
1234
```

---

## 🗂️ File to Submit

> ✅ **trimatoi.go**

---
Perfect — let’s continue your **Go learning log** the right way — full explanation, visualization, and a ready-to-submit file.
This lesson covers **`BasicJoin`**, a function that joins multiple strings together.

---

## 🧩 **Task: BasicJoin**

### 🎯 Goal

Write a function that takes a **slice of strings** (`[]string`) and returns **one long string** that combines (concatenates) them all.

---

## 🧠 **Concept Explanation**

* In Go, strings can be joined by **using the `+` operator** or **`strings.Join()`**.
* However, for this exercise, you’ll **manually build the string** using simple concatenation.
* You’ll loop through all the elements in the slice and **add each string to a final result** variable.

---

## 🧱 **Step-by-Step Code (with Visualization)**

```go
package piscine

// BasicJoin takes a slice of strings and returns them joined together
func BasicJoin(elems []string) string {
	// Step 1: Create an empty string to store the result
	result := ""

	// Step 2: Loop through each string in the slice
	for i := 0; i < len(elems); i++ {
		// Step 3: Add the current string to the result
		result += elems[i]
	}

	// Step 4: Return the final concatenated string
	return result
}
```

---

## 🧩 **Visualization: How it Works**

| Step | `i` | `elems[i]` | `result` after iteration |
| ---- | --- | ---------- | ------------------------ |
| 1    | 0   | `"Hello!"` | `"Hello!"`               |
| 2    | 1   | `" How"`   | `"Hello! How"`           |
| 3    | 2   | `" are"`   | `"Hello! How are"`       |
| 4    | 3   | `" you?"`  | `"Hello! How are you?"`  |

🎬 **Final Output:**
👉 `"Hello! How are you?"`

---

## 🧮 **Example Program to Test**

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.BasicJoin(elems))
}
```

### 🖥️ **Output**

```
Hello! How are you?
```

---

## 💡 **Why This Works**

* Go treats strings like arrays of bytes — when you do `result += elems[i]`, it builds up the string step by step.
* The function is simple but efficient for small slices.
* For large lists, you might later learn to use **`strings.Builder`** (more memory-friendly).

---

## 🧭 **Mini Visualization (Animated Concept)**

Imagine boxes being combined:

```
[ "Hello!" ] + [ " How" ] + [ " are" ] + [ " you?" ]
                ↓
         "Hello! How are you?"
```

---

## 🗂️ **File to Submit**

📄 **Filename:** `basicjoin.go`

```go
package piscine

func BasicJoin(elems []string) string {
	result := ""
	for i := 0; i < len(elems); i++ {
		result += elems[i]
	}
	return result
}
```

---
# Task: Join

---

## ✅ What the function must do

`Join(strs []string, sep string) string`
Return all strings from `strs` concatenated together, with `sep` inserted **between** elements (not after the last one). If `strs` is empty, return an empty string.

---

## 🧩 Final code (file to submit: `join.go`)

Save this exact content in `join.go` under `package piscine`.

```go
package piscine

// Join concatenates the strings in 'strs' separated by 'sep'.
func Join(strs []string, sep string) string {
	// If there are no elements, return an empty string immediately.
	if len(strs) == 0 {
		return ""
	}

	// Start with the first element to avoid adding a separator before it.
	result := strs[0]

	// Add each following element prefixed by the separator.
	for i := 1; i < len(strs); i++ {
		result += sep
		result += strs[i]
	}

	return result
}
```

---

## 🔎 Line-by-line explanation (simple)

1. `package piscine`

   * Declares the package name. Your exercises live in package `piscine`.

2. `// Join concatenates the strings in 'strs' separated by 'sep'.`

   * A short comment explaining the function purpose.

3. `func Join(strs []string, sep string) string {`

   * Defines the function `Join`. It accepts:

     * `strs` — a slice of strings to join,
     * `sep` — the separator string to insert between elements,
     * returns a single `string`.

4. `if len(strs) == 0 {`

   * Checks if the input slice is empty.

5. `return ""`

   * If empty, return an empty string — nothing to join.

6. `result := strs[0]`

   * Initialize `result` with the **first** element.
   * This avoids placing a separator before the first element (so we only add separators **before** subsequent elements).

7. `for i := 1; i < len(strs); i++ {`

   * A `for` loop starting at index 1 (the second element), iterating through the remaining strings.

8. `result += sep`

   * Append the separator to `result`. This happens before adding each subsequent element.

9. `result += strs[i]`

   * Append the current element to `result`.

10. `}`

    * End of the loop.

11. `return result`

    * Return the fully concatenated string.

12. `}`

    * End of function.

---

## 🧪 Example (same as your usage)

If you run:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
}
```

Output:

```
Hello!: How: are: you?
```

---

## 💡 Notes / Alternatives

* This implementation is clear and works fine for moderate sizes.
* For very large slices or when performance matters, use `strings.Builder` to avoid repeated allocations. If you want, I can provide a `strings.Builder` version and explain it line-by-line as well.

---

