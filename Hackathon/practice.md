
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

