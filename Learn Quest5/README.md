# Do not be in a hurry to finish, make sure you rehearse and understand the code!
Thissection is string handling — specifically **runes**, which are very important in Go for dealing with Unicode characters.

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
