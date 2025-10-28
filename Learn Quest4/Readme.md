### Learn at a steady pace! 
Stop trying to complete all the tasks in one go, it doesn't matter if you don't Bim at first trial: the idea is to not give up and find a solution to your current bug. That's what makes you a great programmer!
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

### 🚨 Handling invalid cases (if you don't do this, you will get the error `signal killed`)

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

## 🧩 Task: `RecursiveFactorial`
 **iterative thinking** → to **recursive thinking** — and I’ll make this one with diagrams.

---

### 🧠 First — what is *recursion*?

Recursion means a **function calls itself**.

Instead of using a `for` loop, the function **keeps calling itself again and again**, each time with a smaller number, until it reaches the smallest possible number (called the **base case**).

---

### ⚙️ What are we doing?

We want to find the **factorial** of `nb` again.

Like before:

```
4! = 4 × 3 × 2 × 1 = 24
```

But this time, we’re not allowed to use loops (`for` is forbidden 🚫).

We’ll use recursion instead.

---

### 🧩 How recursion will work (conceptually)

Let’s imagine calling:

```
RecursiveFactorial(4)
```

Here’s what happens step by step:

| Step | Call                  | What it returns             |
| ---- | --------------------- | --------------------------- |
| 1    | RecursiveFactorial(4) | `4 * RecursiveFactorial(3)` |
| 2    | RecursiveFactorial(3) | `3 * RecursiveFactorial(2)` |
| 3    | RecursiveFactorial(2) | `2 * RecursiveFactorial(1)` |
| 4    | RecursiveFactorial(1) | returns `1` (base case)     |
| 🧮   | Unfolding back        | 4 × 3 × 2 × 1 = **24**      |

---

### 🎯 Base case

The **base case** is when recursion *stops*.
If `nb` = 0 or `nb` = 1 → factorial is **1**.
If `nb` < 0 → return **0** (invalid).

---

### ✅ Final Code (`recursivefactorial.go`)

```go
package piscine

// RecursiveFactorial returns the factorial of nb using recursion.
// If nb < 0 or too large, it returns 0.
func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 { // prevent overflow
		return 0
	}
	if nb == 0 || nb == 1 { // base case
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
```

---

### 🧪 Test Program (`main.go`)

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.RecursiveFactorial(arg)) // expected: 24
}
```

---

### 💻 How to run

Make sure your files are in this structure:

```
piscine/recursivefactorial.go
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

### 🧭 Step-by-step visual diagram

Let’s **see** how the recursion works when you call `RecursiveFactorial(4)`:

```
Call stack (goes deeper ↓)

RecursiveFactorial(4)
   ↳ 4 * RecursiveFactorial(3)
         ↳ 3 * RecursiveFactorial(2)
               ↳ 2 * RecursiveFactorial(1)
                     ↳ returns 1 (base case)
```

Now it **unwinds** upward (comes back up):

```
RecursiveFactorial(1) = 1
RecursiveFactorial(2) = 2 * 1 = 2
RecursiveFactorial(3) = 3 * 2 = 6
RecursiveFactorial(4) = 4 * 6 = 24 ✅
```

So the recursion **dives down**, reaches the **base**, then **builds up** the result step by step.

---

### 🗂️ Files to Submit

✅ Only submit: `recursivefactorial.go`

---
Let’s unpack **IterativePower** together — step by step, like you’re learning it from scratch. 
---

## 🧩 Step 1: What the Question Is Asking

We need to write a function:

```go
func IterativePower(nb int, power int) int
```

It should return:

> “**nb raised to the power of power**”
> That means: multiply `nb` by itself `power` times.

---

## 📚 Step 2: Understanding “Power”

Let’s break it down with examples:

| Example | Means     | Result |
| :------ | :-------- | :----- |
| 2⁰      | 1         | 1      |
| 2¹      | 2         | 2      |
| 2²      | 2 × 2     | 4      |
| 2³      | 2 × 2 × 2 | 8      |
| 4³      | 4 × 4 × 4 | 64     |

So basically:

```
nb^power = nb × nb × nb × ... (power times)
```

---

## 🚫 Step 3: Special Rules

The question says:

> “Negative powers will return 0.”

That means:

```go
if power < 0 {
    return 0
}
```

And:

> “Overflows do not have to be dealt with.”

So if the result becomes too big, no problem — ignore it.

---

## 🔁 Step 4: Why "Iterative"?

Because this time, we must use **loops** (not recursion).
We’ll use a `for` loop to multiply repeatedly.

---

## ⚙️ Step 5: Logic Plan (Step-by-Step)

1️⃣ Start with a `result` variable = 1 (anything to power 0 is 1).
2️⃣ Repeat `power` times:

* multiply `result` by `nb`
  3️⃣ When the loop ends → return `result`.

Let’s illustrate for `(4, 3)`:

```
nb = 4, power = 3
result = 1 initially

Loop 1 → result = 1 × 4 = 4
Loop 2 → result = 4 × 4 = 16
Loop 3 → result = 16 × 4 = 64

return 64 ✅
```

---

## 💻 Step 6: Full Code

**File to submit:** `iterativepower.go`
**Folder:** `piscine-go/

Here’s your complete, beginner-friendly solution:

```go
package piscine

// IterativePower calculates nb raised to the power of power using a loop
func IterativePower(nb int, power int) int {
    // If the power is negative, return 0
    if power < 0 {
        return 0
    }

    // Any number to power 0 is 1
    if power == 0 {
        return 1
    }

    result := 1 // Start multiplying from 1

    // Multiply nb by itself power times
    for i := 0; i < power; i++ {
        result = result * nb
    }

    return result
}
```

---

## 🎨 Step 7: Visualization (Like a Baby Learns 👶)

Let’s visualize how `(4, 3)` works:

| Step   | i | result before | Multiply by nb | result after |
| ------ | - | ------------- | -------------- | ------------ |
| Start  | — | 1             | —              | 1            |
| Loop 1 | 0 | 1             | × 4            | 4            |
| Loop 2 | 1 | 4             | × 4            | 16           |
| Loop 3 | 2 | 16            | × 4            | 64           |
| End    | — | —             | —              | ✅ 64         |

So final output = **64**

---

## 🧪 Step 8: Test File

You can test it with this (not to submit):

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IterativePower(4, 3)) // 64
	fmt.Println(piscine.IterativePower(2, 5)) // 32
	fmt.Println(piscine.IterativePower(2, 0)) // 1
	fmt.Println(piscine.IterativePower(3, -2)) // 0 (because power < 0)
}
```

Then run:

```bash
go run .
```

Expected output:

```
64
32
1
0
```


---
