
## foreach: What it does:
This function **applies an action to every number in a list**.

---
```go
package piscine

func ForEach(f func(int), a []int) {
	for _, i := range a {
		f(i)
	}
}
```
## Simple breakdown:

**Function name:** `ForEach`

**What it takes in:**
1. `f func(int)` - A function that does something with one number
2. `a []int` - A list of numbers (called a slice in Go)

**What it does:**
```
for _, i := range a {
    f(i)
}
```
- Goes through each number in the list `a`
- For each number, it calls the function `f` with that number
- The `_` means "I don't care about the position, just give me the value"

## Real-world analogy:
Imagine you have a list of people and a rubber stamp. `ForEach` goes through each person and stamps them. You decide what the stamp does (that's the function `f`).

## Example usage:
```go
// Let's say we want to print each number
numbers := []int{1, 2, 3, 4, 5}

ForEach(func(n int) {
    fmt.Println(n)  // This prints the number
}, numbers)

// Output:
// 1
// 2
// 3
// 4
// 5
```

**In short:** ForEach = "Do this action to every item in my list"

---
## Map
```go
package piscine

func Map(f func(int) bool, a []int) []bool {
	var ans []bool
	for _, i := range a {
		ans = append(ans, f(i))
	}
	return ans
}
```

## What it does:
Takes a list of numbers, **transforms each one**, and gives you back a list of the results.

---

## Simple breakdown:

### Function signature:
```go
func Map(f func(int) bool, a []int) []bool
```

**What it takes in:**
1. `f func(int) bool` - A function that takes a number and returns true/false
2. `a []int` - A list of numbers

**What it returns:**
- `[]bool` - A list of true/false values

---

### The code step-by-step:

```go
var ans []bool
```
- Creates an empty list to store the results
- This will hold all the true/false answers

```go
for _, i := range a {
```
- Loop through each number in the input list `a`
- `i` is each number (ignoring the position with `_`)

```go
    ans = append(ans, f(i))
```
- Call the function `f` with the current number `i`
- Whatever `f(i)` returns (true or false), add it to our answer list
- `append` means "add this to the end of the list"

```go
return ans
```
- Give back the complete list of results

---

## Real-world analogy:
Imagine you have a list of students and a test. `Map` gives the test to each student and collects all their pass/fail results in a new list.

---

## Example from the exercise:

```go
a := []int{1, 2, 3, 4, 5, 6}
result := Map(IsPrime, a)
// result = [false, true, true, false, true, false]
```

**What happened:**
- 1 → IsPrime(1) → false (1 is not prime)
- 2 → IsPrime(2) → true (2 is prime)
- 3 → IsPrime(3) → true (3 is prime)
- 4 → IsPrime(4) → false (4 is not prime)
- 5 → IsPrime(5) → true (5 is prime)
- 6 → IsPrime(6) → false (6 is not prime)

---

## Key difference from ForEach:
- **ForEach**: Does something with each item, returns nothing
- **Map**: Transforms each item, returns a new list of results

**In short:** Map = "Test/transform every item and collect all the results"
