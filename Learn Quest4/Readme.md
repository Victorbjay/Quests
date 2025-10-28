Learn at a steady pace! stop trying to complete all the tasks in one go, it doesn't matter if you don't Bim at first trial: the idea is to not give up and find a solution to your current bug. That's what makes you a great programmer!
## 🧩 Task name: `IterativeFactorial`

---

### 🧠 What does *factorial* mean?

Factorial of a number (written as `n!`) means:
you multiply all the numbers from **1 up to that number**.

For example:

| n | Formula           | Answer |
| - | ----------------- | ------ |
| 1 | 1                 | 1      |
| 2 | 1 × 2             | 2      |
| 3 | 1 × 2 × 3         | 6      |
| 4 | 1 × 2 × 3 × 4     | 24     |
| 5 | 1 × 2 × 3 × 4 × 5 | 120    |

So,
👉 `4! = 1 × 2 × 3 × 4 = 24`

---

### 🧮 What does *iterative* mean?

It means we use a **loop** (`for` loop) — not recursion.
We’ll repeat a multiplication process several times.

We’ll **start with 1**,
then multiply by 2,
then by 3,
until we reach `nb`.

---

### ⚙️ What the task is asking

We must write a function:

```go
func IterativeFactorial(nb int) int {
}
```

✅ It takes **one integer (`nb`)**
✅ It **returns** an integer (the factorial of nb)
✅ If the number is invalid (like negative or too large), we must return **0**.

---

### 🚨 Handling invalid cases

Why? Because factorials grow very fast —
for example, `13!` is already larger than what fits in a normal Go `int` (on 32-bit).

So:

* If `nb < 0` → return `0`
* If the result is too large (we’ll stop at 20 to be safe) → return `0`

---

### 🧱 Let’s build it step-by-step

#### Step 1 — Start with the function name

```go
func IterativeFactorial(nb int) int {
```

#### Step 2 — Check for bad numbers

```go
	if nb < 0 {
		return 0
	}
```

#### Step 3 — Create a variable to hold the result

Start it as 1 (since multiplying by 1 changes nothing):

```go
	result := 1
```

#### Step 4 — Use a loop to multiply from 1 to nb

```go
	for i := 1; i <= nb; i++ {
		result = result * i
	}
```

#### Step 5 — Return the answer

```go
	return result
}
```

---

### ✅ Final Code (iterativefactorial.go)

```go
package piscine

// IterativeFactorial calculates the factorial of nb using a loop.
// If nb is negative or too large, it returns 0.
func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 { // prevent overflow
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ {
		result = result * i
	}

	return result
}
```

---

### 🧪 Test File (main.go)

Create a separate file to test:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.IterativeFactorial(arg)) // expected: 24
}
```

---

### 💻 How to Run

Make sure you’re in the folder that has both:

```
piscine/iterativefactorial.go
main.go
```

Then run:

```bash
$ go run .
```

Expected output:

```
24
```

---

### 🗂️ Files to Submit

Only **`iterativefactorial.go`**

---
