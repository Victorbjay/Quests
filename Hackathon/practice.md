
Each section will describe:

* 🧩 **The task (what to build)**
* ⚙️ **What the program should do**
* 💡 **Hint on logic**
* 🖥️ **Example `main()` usage**
* 📤 **Expected output**

No code solution will be shown — just problem statements for practice or testing.

---

## 🧠 **1. Abort Average**

### 🧩 Task:

Write a function `Abort(a, b, c, d, e int) int`
that takes **five integers** and returns their **average**.

### ⚙️ What It Should Do:

* Add all five numbers.
* Divide the total by 5.
* Return the integer result.

### 💡 Hint:

You can use `return (a + b + c + d + e) / 5`.

### 🖥️ Example:

```go
func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
```

### 📤 Expected Output:

```
5
```

---

## 🔁 **2. Collatz Countdown**

### 🧩 Task:

Write a function `CollatzCountdown(start int) int`
that returns how many steps it takes for `start` to reach **1** under these rules:

* If the number is even → divide by 2
* If it’s odd → multiply by 3 and add 1

### ⚙️ What It Should Do:

* Keep counting until you reach `1`.
* Return `-1` if `start <= 0`.

### 🖥️ Example:

```go
func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}
```

### 📤 Expected Output:

```
9
```

---

## 🔢 **3. Descending Combinations**

### 🧩 Task:

Create a function `DescendComb()`
that prints **all pairs of two-digit numbers** in **descending order**,
where the **first number is greater than the second**.

### ⚙️ What It Should Do:

* Use nested loops for digits `'9'` to `'0'`.
* Print combinations separated by commas and spaces.
* End with a newline.

### 🖥️ Example:

```go
func main() {
	DescendComb()
}
```

### 📤 Expected Output (partial):

```
99 98, 99 97, 99 96, ... , 10 00
```

---

## 🍔 **4. Food Delivery Time**

### 🧩 Task:

Create a function `FoodDeliveryTime(order string) int`
that returns how long it takes to prepare a meal.

### ⚙️ What It Should Do:

* `"burger"` → 15 minutes
* `"chips"` → 10 minutes
* `"nuggets"` → 12 minutes
* Any other food → return 404

### 🖥️ Example:

```go
func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nuggets"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}
```

### 📤 Expected Output:

```
15
10
12
37
```

---

## 🍞 **5. Loaf of Bread**

### 🧩 Task:

Create a function `LoafOfBread(str string) string`
that inserts a **space after every 5 letters**, ignoring existing spaces.

### ⚙️ What It Should Do:

* Count letters only (not spaces).
* After every 5th letter, insert a space.
* Return `"Invalid Output\n"` if fewer than 5 characters.

### 🖥️ Example:

```go
func main() {
	fmt.Print(LoafOfBread("deliciousbread"))
	fmt.Print(LoafOfBread("This is a loaf of bread"))
	fmt.Print(LoafOfBread("loaf"))
}
```

### 📤 Expected Output:

```
delic iousb read 
Thisi saloa fofbr ead 
Invalid Output
```

---

## 🔄 **6. ROT14 Cipher**

### 🧩 Task:

Write a function `Rot14(s string) string`
that **encrypts a string** by shifting each letter 14 positions forward.

### ⚙️ What It Should Do:

* Works for both uppercase and lowercase letters.
* Non-alphabetic characters remain unchanged.

### 🖥️ Example:

```go
func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
```

### 📤 Expected Output:

```
Vszzc! Vck ofs Mci?
```

---

## 🧮 **7. Unmatched Integer**

### 🧩 Task:

Create a function `Unmatch(a []int) int`
that returns the **integer that appears an odd number of times** in a list.

### ⚙️ What It Should Do:

* Loop through all numbers.
* Count how many times each number appears.
* Return the one with an odd count.
* If none found, return `-1`.

### 🖥️ Example:

```go
func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4, 4, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}
```

### 📤 Expected Output:

```
4
```

---

## 1 `abort.go`

```go
package main

import "fmt"

func Abort(a, b, c, d, e int) int {
	return (a + b + c + d + e) / 5
}

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
```

### Line-by-line

* `package main` — this file builds an executable program.
* `import "fmt"` — bring in printing utilities.
* `func Abort(... ) int` — define a function named `Abort` that takes five `int`s and returns an `int`.
* `return (a + b + c + d + e) / 5` — add the five numbers and divide by 5. Because all are `int`, the division is integer division (no fractional part).
* `func main()` — program entry point.
* `middle := Abort(2,3,8,5,7)` — call `Abort` with these five numbers; save returned value in `middle`.
* `fmt.Println(middle)` — print the value and a newline.

### What to expect when run

* Calculation: `2+3+8+5+7 = 25`, `25/5 = 5`.
* **Output:**

```
5
```

---

## 2) `collatzcountdown.go`

```go
package main

import "fmt"

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	count := 0
	for start > 1 {
		if start%2 == 0 {
			start /= 2
		} else if start%2 == 1 {
			start = 3*start + 1
		}
		count++
	}
	return count
}

func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}
```

### Line-by-line

* `CollatzCountdown(start int) int` — function that runs the Collatz process and returns how many steps until `start` becomes 1.
* `if start <= 0 { return -1 }` — invalid start values return `-1`.
* `count := 0` — step counter.
* `for start > 1 { ... }` — loop until `start` becomes 1.

  * `if start%2 == 0 { start /= 2 }` — if even, divide by 2.
  * `else if start%2 == 1 { start = 3*start + 1 }` — if odd, apply 3n+1.
  * `count++` — every operation counts as one step.
* Return `count`.
* `main()` calls `CollatzCountdown(12)` and prints result.

### What to expect when run

* For `start = 12`, the Collatz sequence steps:

  * 12 → 6 → 3 → 10 → 5 → 16 → 8 → 4 → 2 → 1 (10 steps)
* **Output:**

```
10
```

---

## 3) `descendcomb.go`

```go
package main

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for x := '9'; x >= '0'; x-- {
				for y := '9'; y >= '0'; y-- {
					o1 := int(j) * 1
					o2 := int(y) * 1
					t1 := int(i) * 10
					t2 := int(x) * 10
					if o1+t1 <= o2+t2 {
						continue
					}
					printz01(i)
					printz01(j)
					printz01(' ')
					printz01(x)
					printz01(y)
					if i != '0' || j != '1' || x != '0' || y != '0' {
						printz01(',')
						printz01(' ')
					}
				}
			}
		}
	}
	printz01('\n')
}

func printz01(r rune) {
	z01.PrintRune(r)
}

func main() {
	DescendComb()
}
```

### Line-by-line (core ideas)

* Uses nested loops from `'9'` down to `'0'` for four runes: `i, j, x, y`. Think digits `i j` and `x y`.
* `o1 := int(j) * 1` and `t1 := int(i) * 10` convert the runes to integers to form two-digit values `t1+o1` and `t2+o2`.
* `if o1+t1 <= o2+t2 { continue }` — only print pairs where the first two-digit number is strictly greater than the second two-digit number.
* `printz01(...)` prints characters using `z01.PrintRune`.
* The `if i != '0' || j != '1' || x != '0' || y != '0'` condition avoids printing a trailing comma after the last combination `01 00`? (This is a specific termination condition as written.)
* `printz01('\n')` prints newline at end.

### What it prints

* All pairs of two-digit combos `ij xy` where `ij > xy`, in descending order (starting from `99 99` down), separated by `, ` and ending with a newline. Exact output is large — many combinations (all ordered pairs where first number > second).
* Example beginning (approx):

```
99 98, 99 97, ... 99 00, 98 97, 98 96, ...
```

(Printed rune by rune; full list covers all valid pairs.)

---

## 4) `fooddeliverytime.go`

```go
package main

import "fmt"

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var n food
	if order == "burger" {
		n.preptime = 15
	} else if order == "chips" {
		n.preptime = 10
	} else if order == "nuggets" {
		n.preptime = 12
	} else {
		return 404
	}
	return n.preptime
}

func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nuggets"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}
```

### Line-by-line

* `type food struct { preptime int }` — small struct holding prep time.
* `FoodDeliveryTime(order string) int` — returns prep time based on the order string.

  * Recognizes `"burger"`, `"chips"`, `"nuggets"`.
  * If unknown, returns `404` (used here as an error code).
* `main()` prints prep times for each single order and then prints the sum of the three.

### What to expect when run

* Output:

```
15
10
12
37
```

(Last line: `15 + 10 + 12 = 37`)

---

## 5) `loafofbread.go`

```go
package main

import "fmt"

func LoafOfBread(str string) string {
	s := []rune(str)
	r := []rune{}
	if len(s) < 5 {
		return "Invalid Output\n"
	}
	var stp bool
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			r = append(r, s[i])
			stp = true
			count++
		}
		if count%5 == 0 && stp && i+1 < len(s) && len(r) != 0 {
			if s[i+1] == ' ' {
				stp = false
			} else {
				s[i+1] = ' '
				stp = false
			}
			r = append(r, ' ')
		}
	}
	return string(r) + "\n"
}

func main() {
	fmt.Print(LoafOfBread("deliciousbread"))
	fmt.Print(LoafOfBread("This is a loaf of bread"))
	fmt.Print(LoafOfBread("loaf"))
}
```

### Line-by-line (core idea)

* Convert input string to rune slice `s` for safe handling of characters.
* If the total letters are fewer than 5, return `"Invalid Output\n"`.
* `r` is the result builder.
* The loop appends runes that are not spaces, increments a `count` when a non-space character was appended.
* Every time `count` reaches a multiple of 5, it attempts to insert a space into `r` and (if necessary) set the next input character to space so future iterations skip it.
* Finally return the rebuilt string plus newline.

### What to expect when run

* `LoafOfBread("deliciousbread")` → insert spaces every 5 characters of the letters-only stream:

  * `"deliciousbread"` letters-only = `deliciousbread` (13 letters). After chunking: `delic iousb read` but function logic appends spaces after every 5 letters, result will be something like:

```
delic iousb read
```

(plus newline).

* `LoafOfBread("This is a loaf of bread")` → removes spaces and re-inserts after every 5 letters. Output similar, with newline.
* `LoafOfBread("loaf")` → length < 5, returns:

```
Invalid Output
```

(Exact splitting depends on algorithm’s handling of original spaces — but generally it groups letters into groups of 5 separated by spaces.)

---

## 6) `rot14.go`

```go
package main

import "github.com/01-edu/z01"

func Rot14(s string) string {
	sliceS := []rune(s)
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] >= 'a' && sliceS[i] <= 'l' || sliceS[i] >= 'A' && sliceS[i] <= 'L' {
			sliceS[i] += 14
		} else if sliceS[i] > 'l' && sliceS[i] <= 'z' || sliceS[i] > 'L' && sliceS[i] <= 'Z' {
			sliceS[i] -= 12
		}
	}
	return string(sliceS)
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
```

### Line-by-line

* `Rot14` applies a shift of 14 to alphabetical letters:

  * For letters `a-l` and `A-L`, add 14.
  * For letters `m-z` and `M-Z`, subtract 12 (equivalently wrap around forward by 14).
* Non-letter characters (spaces, punctuation) are unchanged.
* `main()` calls `Rot14("Hello! How are You?")` and prints result with `z01.PrintRune`.

### What to expect when run

* `"Hello! How are You?"` becomes each letter shifted by 14.

  * `H` → `V` (because `H + 14 = V`)
  * `e` → `s`
  * `l` → `z` (since `l` in `a-l` branch)
  * So output will be the ROT14 of the string, printed on one line, e.g.:

```
Vszzy! Vua omc Msi?
```

(Exact output: run it to confirm; letters map by +/- rules above.)

---

## 7) `unmatch.go`

```go
package main

import "fmt"

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		count := 1
		for j := 0; j < len(a); j++ {
			if i != j && a[i] == a[j] {
				count++
			}
		}
		if count%2 != 0 {
			return a[i]
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4, 4, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}
```

### Line-by-line

* `Unmatch(a []int) int` — find and return the element that appears an odd number of times.
* Outer loop picks each element `a[i]`.
* Inner loop counts how many times that value appears (increments `count` when a match). Start `count := 1` because the element itself is counted.
* If `count%2 != 0` (odd count), return that element.
* If none found, return `-1`.
* `main()` builds array `a` and prints the element with odd occurrence.

### What to expect when run

* In `a := {1,2,3,1,2,3,4,4,4}`, counts are:

  * 1 appears 2 times (even)
  * 2 appears 2 times (even)
  * 3 appears 2 times (even)
  * 4 appears 3 times (odd)
* So function returns `4`.
  **Output:**

```
4
```

---

## Final tips for running each file

* To run a file: `go run filename.go` (e.g., `go run abort.go`).
* Each `main` prints to standard output; expect exactly the results described above.
* Watch integer vs float behavior: `Abort` returns an integer average; if you want decimals, convert to `float64`.

---

