
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
### 🧩 Visualization — how the code runs


**Goal:** Calculate `n! = 1 × 2 × 3 × ... × n`

**Example:** `4! = 1 × 2 × 3 × 4`

| Loop | i | result (before) | result × i | result (after) |
|------|---|-----------------|------------|----------------|
| 1st  | 1 | 1               | 1 × 1      | 1              |
| 2nd  | 2 | 1               | 1 × 2      | 2              |
| 3rd  | 3 | 2               | 2 × 3      | 6              |
| 4th  | 4 | 6               | 6 × 4      | **24** ✅      |

**Final answer:** `4! = 24`


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

## 🎨 Step 7: Visualization 

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
let’s go through **`RecursivePower`** in full detail, just like before.

---

## 🧩 Task Summary

We need to write a recursive Go function that calculates:
[
nb^{power} = nb \times nb \times nb \times ... \text{(power times)}
]

---

## ✅ Full Solution (`recursivepower.go`)

```go
package piscine

// RecursivePower calculates nb raised to the power of 'power' using recursion.
// If power is negative, it returns 0 (not defined for integers).
// No loops (for) are used as per instructions.
func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return nb * RecursivePower(nb, power-1)
}
```

---

## 🧠 Line-by-Line Explanation

| Line                                           | Code                                                                                                                                                | Explanation |
| ---------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| `func RecursivePower(nb int, power int) int {` | Starts defining a function named `RecursivePower` that takes two integers: `nb` (the base) and `power` (the exponent), returning an integer result. |             |
| `if power < 0 { return 0 }`                    | Handles **negative powers** — the question says to return `0` if the power is negative (since we’re only working with integers).                    |             |
| `if power == 0 { return 1 }`                   | **Base case** — any number to the power of `0` equals `1`. This stops the recursion.                                                                |             |
| `return nb * RecursivePower(nb, power-1)`      | **Recursive step:** multiply `nb` by the result of calling `RecursivePower` again with one less power. This keeps going until power = 0.            |             |
| `}`                                            | Ends the function.                                                                                                                                  |             |

---

## 🔁 Visual Recursive Flow (for `RecursivePower(4, 3)`)

| Step | Function Call          | Action                     | Return       |
| ---- | ---------------------- | -------------------------- | ------------ |
| 1    | `RecursivePower(4, 3)` | `4 * RecursivePower(4, 2)` | waits        |
| 2    | `RecursivePower(4, 2)` | `4 * RecursivePower(4, 1)` | waits        |
| 3    | `RecursivePower(4, 1)` | `4 * RecursivePower(4, 0)` | waits        |
| 4    | `RecursivePower(4, 0)` | returns `1` (base case)    | ✅            |
| 5    | Back to step 3         | `4 * 1 = 4`                | returns `4`  |
| 6    | Back to step 2         | `4 * 4 = 16`               | returns `16` |
| 7    | Back to step 1         | `4 * 16 = 64`              | returns `64` |

🎯 Final Answer = **64**

---

## 🧩 Memory Stack Visualization

```
RecursivePower(4,3)
 ├── RecursivePower(4,2)
 │    ├── RecursivePower(4,1)
 │    │    ├── RecursivePower(4,0)
 │    │    └── returns 1
 │    └── returns 4*1 = 4
 └── returns 4*4 = 16
returns 4*16 = 64 ✅
```

---

## 🧪 Test Program

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RecursivePower(4, 3))  // Expected: 64
	fmt.Println(piscine.RecursivePower(2, 5))  // Expected: 32
	fmt.Println(piscine.RecursivePower(5, 0))  // Expected: 1
	fmt.Println(piscine.RecursivePower(3, -2)) // Expected: 0
}
```
### 🗂️ Files to Submit

✅ Only submit: `recursivepower.go`

---
We’re now moving into the classic **Fibonacci** problem — a perfect example of recursion!

---

## 🧩 TASK SUMMARY

We’re writing a recursive function called `Fibonacci(index int)` that returns the **Fibonacci number** at position `index`.

🧮 The **Fibonacci sequence** goes like this:

```
Index:  0  1  2  3  4  5  6 ...
Value:  0, 1, 1, 2, 3, 5, 8 ...
```

Formula (recursively defined):

[
Fib(n) = Fib(n-1) + Fib(n-2)
]
with:

* `Fib(0) = 0`
* `Fib(1) = 1`

---

## ✅ FILE TO SUBMIT

**Filename:**
`fibonacci.go`

**Allowed functions:**
✅ Built-ins only (no `for` loops, no external libraries)

---

## ✅ FULL CODE (Beginner-friendly)

```go
package piscine

// Fibonacci returns the Fibonacci number at the given index using recursion.
// - If index < 0, returns -1 (invalid).
// - If index == 0, returns 0.
// - If index == 1, returns 1.
// - Otherwise, it returns Fibonacci(index-1) + Fibonacci(index-2).
func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
```

---

## 🧠 LINE-BY-LINE EXPLANATION

| Line                                             | Code                                                                                             | Explanation |
| ------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ----------- |
| `if index < 0 { return -1 }`                     | If someone gives a negative index, we return `-1` because it doesn’t make sense in the sequence. |             |
| `if index == 0 { return 0 }`                     | The first number in Fibonacci is 0.                                                              |             |
| `if index == 1 { return 1 }`                     | The second number in Fibonacci is 1.                                                             |             |
| `return Fibonacci(index-1) + Fibonacci(index-2)` | The recursive “magic” 🪄 happens here: to get Fib(n), we add the two previous Fibonacci values.  |             |

---

## 🔁 VISUAL FLOW (for `Fibonacci(4)`)

| Step | Function Call                                             | What Happens        | Returns |
| ---- | --------------------------------------------------------- | ------------------- | ------- |
| 1    | `Fibonacci(4)`                                            | = `Fib(3) + Fib(2)` | waits   |
| 2    | `Fibonacci(3)`                                            | = `Fib(2) + Fib(1)` | waits   |
| 3    | `Fibonacci(2)`                                            | = `Fib(1) + Fib(0)` | waits   |
| 4    | `Fibonacci(1)`                                            | returns 1           | ✅       |
| 5    | `Fibonacci(0)`                                            | returns 0           | ✅       |
| 6    | Back to step 3 → `Fib(2)` = `1 + 0 = 1`                   |                     |         |
| 7    | Back to step 2 → `Fib(3)` = `1 (Fib(2)) + 1 (Fib(1)) = 2` |                     |         |
| 8    | `Fibonacci(2)` again (right side of step 1) → returns 1   |                     |         |
| 9    | Finally: `Fib(4)` = `2 + 1 = 3 ✅`                         |                     |         |

🎯 Final Answer = **3**

---

## 🧱 MEMORY STACK VISUALIZATION

```
Fibonacci(4)
 ├── Fibonacci(3)
 │    ├── Fibonacci(2)
 │    │    ├── Fibonacci(1) -> 1
 │    │    └── Fibonacci(0) -> 0
 │    └── returns 1 + 0 = 1
 │    ├── Fibonacci(1) -> 1
 │    └── returns 1 + 1 = 2
 ├── Fibonacci(2)
 │    ├── Fibonacci(1) -> 1
 │    └── Fibonacci(0) -> 0
 └── returns 2 + 1 = 3 ✅
```

---

## 🧪 TEST PROGRAM

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Fibonacci(4))  // 3
	fmt.Println(piscine.Fibonacci(0))  // 0
	fmt.Println(piscine.Fibonacci(1))  // 1
	fmt.Println(piscine.Fibonacci(6))  // 8
	fmt.Println(piscine.Fibonacci(-3)) // -1
}
```

---

## 🧾 STEPS TO PASS THE CHECKER ✅

1. Create file:
   `fibonacci.go` inside your `piscine` folder.
2. Write the function exactly as shown.
3. No `for` loops! (Recursion only)
4. Run:

   ```bash
   go run .
   ```
5. Expected output for index `4` → `3`

---
Perfect 👍 Let’s go through this one in your preferred full structured documentation format — with complete explanation, visualization, and code.

---

## 🧩 **Task Name:** Sqrt

### 🧠 What does “Square Root” mean?

The **square root** of a number is a value that, when multiplied by itself, gives the original number.

For example:

| Number | Square Root | Why                                                 |
| ------ | ----------- | --------------------------------------------------- |
| 4      | 2           | 2 × 2 = 4                                           |
| 9      | 3           | 3 × 3 = 9                                           |
| 16     | 4           | 4 × 4 = 16                                          |
| 3      | —           | no whole number gives 3 when squared → return **0** |

👉 So if the square root is **not a whole number**, we return **0**.

---

### ⚙️ What the task is asking

We must write a function that:

* Takes an integer `nb`
* Returns its **integer square root** (if it exists)
* Returns **0** otherwise

Function header:

```go
func Sqrt(nb int) int
```

✅ If `nb = 4` → return `2`
✅ If `nb = 3` → return `0`
✅ If `nb = 9` → return `3`
✅ If `nb = 10` → return `0`

---

### 🚧 Restrictions

We are **not allowed to use** any built-in math functions.
So we must find the square root **manually**, using a **loop**.

---

### 🧱 Step-by-Step Plan

#### Step 1 — Check invalid cases

If `nb <= 0`, return `0` (no square root for negative or zero).

#### Step 2 — Try every number starting from 1

We test each integer `i`:

* Multiply `i * i`
* If it equals `nb`, we found the square root.
* If it becomes greater than `nb`, stop and return `0`.

#### Step 3 — Return the result

Return the found integer if perfect, otherwise `0`.

---

### ✅ **Final Code** (`sqrt.go`)

```go
package piscine

// Sqrt returns the integer square root of nb if it exists,
// otherwise returns 0.
func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}

	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
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
	fmt.Println(piscine.Sqrt(4))  // expected: 2
	fmt.Println(piscine.Sqrt(3))  // expected: 0
	fmt.Println(piscine.Sqrt(9))  // expected: 3
	fmt.Println(piscine.Sqrt(16)) // expected: 4
	fmt.Println(piscine.Sqrt(10)) // expected: 0
}
```

---

### 🧩 **Visual Step-by-Step Flow**

Let’s take `Sqrt(9)` as example:

| Step | i | i × i | Comparison | Action     |
| ---- | - | ----- | ---------- | ---------- |
| 1    | 1 | 1     | 1 < 9      | continue   |
| 2    | 2 | 4     | 4 < 9      | continue   |
| 3    | 3 | 9     | 9 == 9     | ✅ return 3 |

👉 Loop stops as soon as `i*i == nb`.

Now for `Sqrt(3)`:

| Step | i | i × i | Comparison | Action          |
| ---- | - | ----- | ---------- | --------------- |
| 1    | 1 | 1     | 1 < 3      | continue        |
| 2    | 2 | 4     | 4 > 3      | stop & return 0 |

✅ Because 4 passed 3, we know 3 has no integer square root.

---

### ⚡ **Quick Visual Diagram**

```
Start nb = 9
↓
i = 1 → 1*1 = 1 → not equal
↓
i = 2 → 2*2 = 4 → not equal
↓
i = 3 → 3*3 = 9 → ✅ equal
↓
Return 3
```

For nb = 3:

```
Start nb = 3
↓
i = 1 → 1*1 = 1 → not equal
↓
i = 2 → 2*2 = 4 → greater
↓
Return 0
```

---

### 🧾 **Expected Output**

```
$ go run .
2
0
3
4
0
```

---

### 🗂️ **Files to Submit**

* `sqrt.go`

---
Excellent 👌 Let’s dive into this one in your **full preferred structured format** — with explanations, visualization, table flow, and example code.

---

## 🧩 **Task Name:** IsPrime

---

### 🧠 What does “Prime Number” mean?

A **prime number** is a number that is **only divisible by 1 and itself.**
That means:

* It has **exactly two divisors** → `1` and the number itself.

📘 Examples:

| Number | Divisible By | Prime? | Why                            |
| ------ | ------------ | ------ | ------------------------------ |
| 1      | 1            | ❌      | Not prime (only one divisor)   |
| 2      | 1, 2         | ✅      | Only divisible by 1 and itself |
| 3      | 1, 3         | ✅      | Same reason                    |
| 4      | 1, 2, 4      | ❌      | Divisible by 2 also            |
| 5      | 1, 5         | ✅      | Prime                          |
| 6      | 1, 2, 3, 6   | ❌      | Divisible by 2 & 3             |

So `2, 3, 5, 7, 11, 13, 17…` are all **prime numbers**.

---

### ⚙️ What the task is asking

We must write a function that:
✅ Takes an integer `nb`
✅ Returns `true` if `nb` is prime
✅ Returns `false` otherwise

Function header:

```go
func IsPrime(nb int) bool
```

Also:

* 1 is **not** prime
* Negative numbers are **not** prime
* Optimize it (don’t check too many unnecessary numbers)

---

### ⚡ Optimization Logic

Instead of checking all numbers from 1 to nb,
we only check **up to the square root of nb**.

Because if `nb` has a divisor greater than its square root,
the corresponding divisor smaller than it would already have been found.

📘 Example:
For 25 → we only check divisors up to 5
(5×5 = 25, any number > 5 would multiply past 25)

---

### 🧱 Step-by-Step Plan

#### Step 1 — Handle invalid cases

If `nb <= 1`, return `false` (not prime).

#### Step 2 — Loop from 2 to √nb

If `nb` is divisible by any number in that range,
return `false`.

#### Step 3 — Otherwise

Return `true` (it’s prime).

---

### ✅ **Final Code** (`isprime.go`)

```go
package piscine

// IsPrime checks if nb is a prime number.
// It returns true if nb is prime, false otherwise.
func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	// Only check divisibility up to the square root of nb
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
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
	fmt.Println(piscine.IsPrime(5))  // true
	fmt.Println(piscine.IsPrime(4))  // false
	fmt.Println(piscine.IsPrime(11)) // true
	fmt.Println(piscine.IsPrime(25)) // false
	fmt.Println(piscine.IsPrime(2))  // true
	fmt.Println(piscine.IsPrime(1))  // false
}
```

---

### 🧩 **Visual Flow Table**

| Step | nb | i | i×i ≤ nb? | nb % i    | Result   |
| ---- | -- | - | --------- | --------- | -------- |
| 1    | 5  | 2 | 4 ≤ 5 ✅   | 5 % 2 = 1 | continue |
| 2    | 5  | 3 | 9 > 5 ❌   | stop      | ✅ Prime  |

| Step | nb | i | i×i ≤ nb? | nb % i    | Result      |
| ---- | -- | - | --------- | --------- | ----------- |
| 1    | 4  | 2 | 4 ≤ 4 ✅   | 4 % 2 = 0 | ❌ Not Prime |

---

### ⚡ **Quick Diagram**

#### For nb = 5

```
Start nb=5
↓
Check 2 → 5%2 ≠ 0 → continue
↓
No divisors found
↓
Return true ✅
```

#### For nb = 4

```
Start nb=4
↓
Check 2 → 4%2 = 0 → divisible
↓
Return false ❌
```

---

### 🧾 **Expected Output**

```
$ go run .
true
false
true
false
true
false
```

---

### 🗂️ **Files to Submit**

* `isprime.go`

---
Perfect 🎯
Let’s continue in the **same style and structure** — with full explanation, visualization, diagram, table, and runnable example — for your next Go task:

---

## 🧩 **Task Name:** FindNextPrime

---

### 🧠 What does “Next Prime” mean?

You already know a **prime number** is a number that’s only divisible by **1 and itself**.

Now, this time, the question wants us to find the **next prime number that is equal to or greater than** a given number.

So if you pass a number `nb`, you should return:

* `nb` itself **if it’s already prime**,
* or the **next** number that is prime.

📘 Examples:

| nb | Next Prime        |
| -- | ----------------- |
| 5  | 5 (already prime) |
| 4  | 5 (next one)      |
| 8  | 11                |
| 14 | 17                |
| 1  | 2                 |
| 0  | 2                 |

---

### ⚙️ What the task is asking

We must write a function:

```go
func FindNextPrime(nb int) int
```

✅ It takes one integer `nb`
✅ It returns the first **prime number ≥ nb**
✅ It must be **optimized** to avoid timeouts
✅ We can use a helper function like `IsPrime` (from your last task!)

---

### 🧱 **Step-by-Step Logic**

#### Step 1 — Handle bad inputs

If `nb <= 2`, the next prime is **2** (smallest prime).

#### Step 2 — Loop starting from nb

While the number is **not prime**, increase it by 1.

#### Step 3 — Use your `IsPrime()` function

Check if each candidate is prime until one passes.

#### Step 4 — Return it

Return the first number that is prime.

---

### ⚙️ **Helper Function Reminder: `IsPrime()`**

You can reuse your previous optimized function:

```go
func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
```

---

### ✅ **Final Code — `findnextprime.go`**

```go
package piscine

// FindNextPrime returns the first prime number greater than or equal to nb.
func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	for {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
}
```

---

### 💻 **Test File — `main.go`**

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FindNextPrime(5))  // 5 (already prime)
	fmt.Println(piscine.FindNextPrime(4))  // 5
	fmt.Println(piscine.FindNextPrime(8))  // 11
	fmt.Println(piscine.FindNextPrime(14)) // 17
	fmt.Println(piscine.FindNextPrime(0))  // 2
}
```

---

### 🧩 **Visual Flow Table**

| Step | nb (start) | Prime? | Action               | Result    |
| ---- | ---------- | ------ | -------------------- | --------- |
| 1    | 5          | ✅      | stop                 | return 5  |
| 2    | 4          | ❌      | nb = 5 → check again | return 5  |
| 3    | 8          | ❌      | nb = 9, 10, 11 → ✅   | return 11 |

---

### ⚡ **Visualization**

#### Example: `FindNextPrime(8)`

```
Start nb = 8
↓
IsPrime(8)? → false
↓
nb = 9
↓
IsPrime(9)? → false
↓
nb = 10
↓
IsPrime(10)? → false
↓
nb = 11
↓
IsPrime(11)? → true ✅
↓
Return 11
```

---

### 🧾 **Expected Output**

```
$ go run .
5
5
11
17
2
```

---

### 🗂️ **Files to Submit**

* `findnextprime.go`

---

### ⚡ Extra Tip for Understanding

If you imagine a **number line**, your program “walks” forward from `nb` checking each number:

```
nb → nb+1 → nb+2 → nb+3 → ... until a prime is found ✅
```

That’s how it ensures it always returns the **closest prime ≥ nb**.

---
# Task **`EightQueens()`** (the legendary *Eight Queens puzzle*) step-by-step — with full **explanation**

---

## 🧩 Problem Summary

We must place **8 queens** on a **8×8 chessboard** such that:

* No two queens attack each other.
* That means no two queens share:

  * The same **row** 🟥
  * The same **column** 🟦
  * The same **diagonal** 🟩

We’ll print **all valid arrangements**, each as a line of 8 digits —
each digit shows the **row position** of the queen in that column (from left to right).

Example output:

```
15863724
16837425
17468253
...
```

---

## 🧠 Approach (Recursive Backtracking)

We’ll use **recursion** to try placing a queen column by column.

### Step Flow:

| Step | Description                                                          |
| ---- | -------------------------------------------------------------------- |
| 1️⃣  | Start from column 0.                                                 |
| 2️⃣  | Try placing a queen in each row (1–8).                               |
| 3️⃣  | Check if that position is **safe** (no attacks).                     |
| 4️⃣  | If safe → place queen and **recurse** to next column.                |
| 5️⃣  | If no safe position → **backtrack** (remove queen and try next row). |
| 6️⃣  | When column == 8 → print the current valid solution.                 |

---

## 🧮 Visual Example

Say we try to place queens one by one:

```
Column 1: place at row 1  ✅
Column 2: cannot place at 1 (same row), or 2 (diagonal), ...
           place at row 5 ✅
Column 3: try rows → only row 8 works ✅
...
```

Result: `15863724`

---

## 🧩 Full Go Code (with explanations)

```go
package piscine

import "github.com/01-edu/z01"

// EightQueens prints all solutions for the 8 queens puzzle
func EightQueens() {
	const size = 8
	var board [size]int // board[col] = row (1-based)

	// check if we can safely place a queen
	var isSafe func(col, row int) bool
	isSafe = func(col, row int) bool {
		for prevCol := 0; prevCol < col; prevCol++ {
			prevRow := board[prevCol]

			// ❌ Same row check
			if prevRow == row {
				return false
			}

			// ❌ Same diagonal check (difference in col == difference in row)
			if abs(prevRow-row) == abs(prevCol-col) {
				return false
			}
		}
		return true
	}

	// recursive backtracking solver
	var solve func(col int)
	solve = func(col int) {
		if col == size {
			// ✅ all queens placed safely → print the board
			for i := 0; i < size; i++ {
				z01.PrintRune(rune(board[i]) + '0')
			}
			z01.PrintRune('\n')
			return
		}

		// try placing queen in each row (1–8)
		for row := 1; row <= size; row++ {
			if isSafe(col, row) {
				board[col] = row
				solve(col + 1) // recurse to next column
			}
		}
	}

	solve(0)
}

// helper absolute value function
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
```

---

## 🧩 Output Example

Running this:

```go
package main

import "piscine"

func main() {
	piscine.EightQueens()
}
```

Will print:

```
15863724
16837425
17468253
...
```

(There are exactly **92** valid solutions.)

---

## 📊 Step-by-Step Flow Diagram (Simplified)

```
solve(0)
 ├── row=1 ✅ → solve(1)
 │     ├── row=1 ❌ same row
 │     ├── row=2 ✅ → solve(2)
 │     │     ├── ...
 │     │     └── full valid board → print
 │     └── backtrack
 ├── row=2 ✅ → solve(1)
 │     ├── ...
 └── row=8 ...
```

---

## 🧩 Key Concepts Learned

✅ Recursion
✅ Backtracking
✅ Diagonal detection logic
✅ Array-based chessboard representation
✅ Problem-solving with pruning

---

Let’s visualize how **recursion + backtracking** works for the **Eight Queens** puzzle — with a clear, **step-by-step chessboard diagram**
---

## 🧠 Recap of the Goal

We must place **8 queens** 👑 on an **8×8 board** so that:

* No two queens share a **row**, **column**, or **diagonal**.

---

## ♟️ Step-by-Step Visual Flow

We’ll walk through the **first valid solution: `15863724`**

Each column (1–8) gets one queen.
Each number in `15863724` means the **row** for the queen in that **column**.

---

### 🧩 Step 1: Place Queen in Column 1, Row 1

```
Col →  1 2 3 4 5 6 7 8
Row 1: 👑 . . . . . . .
Row 2: .  . . . . . . .
Row 3: .  . . . . . . .
Row 4: .  . . . . . . .
Row 5: .  . . . . . . .
Row 6: .  . . . . . . .
Row 7: .  . . . . . . .
Row 8: .  . . . . . . .
```

✅ Safe — continue.

---

### 🧩 Step 2: Try Column 2

* Row 1 ❌ (same row)
* Row 2 ❌ (diagonal)
* Row 3 ❌ (diagonal)
* Row 4 ❌ (diagonal)
* Row 5 ✅ → place queen

```
Col →  1 2 3 4 5 6 7 8
Row 1: 👑 . . . . . . .
Row 2: .  . . . . . . .
Row 3: .  . . . . . . .
Row 4: .  . . . . . . .
Row 5: .  👑 . . . . . .
Row 6: .  . . . . . . .
Row 7: .  . . . . . . .
Row 8: .  . . . . . . .
```

✅ Continue to next column.

---

### 🧩 Step 3: Column 3

* Try rows 1–7, find **Row 8** ✅ safe.

```
Col →  1 2 3 4 5 6 7 8
Row 1: 👑 . . . . . . .
Row 2: .  . . . . . . .
Row 3: .  . . . . . . .
Row 4: .  . . . . . . .
Row 5: .  👑 . . . . . .
Row 6: .  . . . . . . .
Row 7: .  . . . . . . .
Row 8: .  . 👑 . . . . .
```

---

### 🧩 Step 4: Column 4

* After testing, Row 6 ✅ safe.

```
Col →  1 2 3 4 5 6 7 8
Row 1: 👑 . . . . . . .
Row 2: .  . . . . . . .
Row 3: .  . . . . . . .
Row 4: .  . . . . . . .
Row 5: .  👑 . . . . . .
Row 6: .  . . 👑 . . . .
Row 7: .  . . . . . . .
Row 8: .  . 👑 . . . . .
```

---

### 🧩 Step 5: Column 5

* Row 3 ✅ safe.

```
Col →  1 2 3 4 5 6 7 8
Row 1: 👑 . . . . . . .
Row 2: .  . . . . . . .
Row 3: .  . . . 👑 . . .
Row 4: .  . . . . . . .
Row 5: .  👑 . . . . . .
Row 6: .  . . 👑 . . . .
Row 7: .  . . . . . . .
Row 8: .  . 👑 . . . . .
```

---

### 🧩 Step 6: Column 6

* Row 7 ✅ safe.

```
Col →  1 2 3 4 5 6 7 8
Row 1: 👑 . . . . . . .
Row 2: .  . . . . . . .
Row 3: .  . . . 👑 . . .
Row 4: .  . . . . . . .
Row 5: .  👑 . . . . . .
Row 6: .  . . 👑 . . . .
Row 7: .  . . . . 👑 . .
Row 8: .  . 👑 . . . . .
```

---

### 🧩 Step 7: Column 7

* Row 2 ✅ safe.

```
Col →  1 2 3 4 5 6 7 8
Row 1: 👑 . . . . . . .
Row 2: .  . . . . . 👑 .
Row 3: .  . . . 👑 . . .
Row 4: .  . . . . . . .
Row 5: .  👑 . . . . . .
Row 6: .  . . 👑 . . . .
Row 7: .  . . . . 👑 . .
Row 8: .  . 👑 . . . . .
```

---

### 🧩 Step 8: Column 8

* Row 4 ✅ safe → 🎯 **solution found!**

```
Col →  1 2 3 4 5 6 7 8
Row 1: 👑 . . . . . . .
Row 2: .  . . . . . 👑 .
Row 3: .  . . . 👑 . . .
Row 4: .  . . . . . . 👑
Row 5: .  👑 . . . . . .
Row 6: .  . . 👑 . . . .
Row 7: .  . . . . 👑 . .
Row 8: .  . 👑 . . . . .
```

📜 Printed output:

```
15863724
```

---

## 🔁 What Happens Next

After printing, the algorithm **backtracks**:

* Removes the last queen 👑,
* Tries the next row in column 8,
* And continues recursively until **all 92 valid solutions** are found.

---

## 🧩 Summary of the Recursion Tree

```
solve(0)
 ├── row=1
 │   ├── row=5
 │   │   ├── row=8
 │   │   │   ├── row=6
 │   │   │   │   ├── row=3
 │   │   │   │   │   ├── row=7
 │   │   │   │   │   │   ├── row=2
 │   │   │   │   │   │   │   ├── row=4 ✅ (solution)
 │   │   │   │   │   │   │   └── backtrack
```

---
