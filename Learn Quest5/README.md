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

Would you like me to continue next with the same visual explanation style for **ConcatAlternate** (if that’s your next task)?
