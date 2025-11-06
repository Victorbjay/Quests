
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
