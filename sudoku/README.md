# 🚀 Learning Go Language with Sudoku - For Beginners!

## 📚 What is Go?
Go (also called Golang) is a programming language made by Google. It's like giving instructions to a computer, but you have to use special words it understands!

---

## 🎯 Part 1: Understanding the Basics

### What are these weird symbols?

| Symbol | Name | What it means |
|--------|------|---------------|
| `{ }` | Curly braces | "Start" and "End" of a block of code |
| `( )` | Parentheses | Hold information you're giving to something |
| `[ ]` | Square brackets | Used for lists/arrays (like a row of boxes) |
| `;` | Semicolon | End of a statement (like a period in a sentence) |
| `//` | Double slash | A comment (note for humans, computer ignores it) |
| `:=` | Declare | Create a new variable |
| `==` | Equals | Check if two things are the same |
| `!=` | Not equals | Check if two things are different |

---

## 🏗️ Part 2: Building Blocks of Go

### 1. **Package** - Like a toolbox
```go
package main
```
**Translation**: "This is the main toolbox where our program lives"

```go
package piscine
```
**Translation**: "This is a separate toolbox called 'piscine' with helper tools"

---

### 2. **Import** - Borrowing tools
```go
import "fmt"
```
**Translation**: "I want to borrow the 'fmt' tool to print things on the screen"

```go
import "os"
```
**Translation**: "I want to borrow 'os' to read what the user types"

---

### 3. **Variables** - Storage boxes
```go
var grid [9][9]int
```
**Translation**: "Create a box called 'grid' that holds 9 rows of 9 numbers"

Think of it like:
```
[1][2][3][4][5][6][7][8][9]  <- Row 1
[1][2][3][4][5][6][7][8][9]  <- Row 2
[1][2][3][4][5][6][7][8][9]  <- Row 3
... (9 rows total)
```

```go
num := 5
```
**Translation**: "Create a box called 'num' and put the number 5 in it"

---

### 4. **Functions** - Machines that do work
```go
func PrintHello() {
    fmt.Println("Hello!")
}
```
**Translation**: "Create a machine called 'PrintHello' that prints 'Hello!' when you use it"

---

## 🧩 Part 3: Your Sudoku Code Explained SIMPLY

### **File 1: piscine/solver.go** (The Helper Tools)

#### Tool #1: The Safety Checker
```go
func isValid(grid *[9][9]int, row, col, num int) bool {
```
**Translation in Plain English**: 
- "Make a tool called `isValid`"
- "Give it: the grid, a row number, a column number, and a number to test"
- "It will answer TRUE (yes it's safe) or FALSE (no it's not safe)"

**What does `*[9][9]int` mean?**
- `[9][9]int` = A 9×9 grid of whole numbers
- `*` = "Use the REAL grid, not a copy" (like pointing to the actual thing)

Let's break down what happens inside:

```go
for i := 0; i < 9; i++ {
```
**Translation**: "Count from 0 to 8 (that's 9 times). Call each number 'i'"

```go
if grid[row][i] == num || grid[i][col] == num {
    return false
}
```
**Translation**: 
- `grid[row][i]` = "Look at spot 'i' in this row"
- `grid[i][col]` = "Look at spot 'i' in this column"
- `||` means "OR"
- "If the number is already in the row OR column, say FALSE (not allowed!)"

```go
startRow := (row / 3) * 3
startCol := (col / 3) * 3
```
**Translation**: "Find which 3×3 box we're in"

**Example**: If row = 5:
- 5 ÷ 3 = 1 (Go drops the decimal)
- 1 × 3 = 3
- So we start at row 3 (the beginning of the middle section)

```go
for i := startRow; i < startRow+3; i++ {
    for j := startCol; j < startCol+3; j++ {
        if grid[i][j] == num {
            return false
        }
    }
}
```
**Translation**: "Check all 9 spots in the 3×3 box. If the number is there, say FALSE"

```go
return true
```
**Translation**: "All checks passed! Say TRUE (it's safe!)"

---

#### Tool #2: The Puzzle Solver
```go
func SolveSudoku(grid *[9][9]int) bool {
```
**Translation**: "Make a tool called SolveSudoku. It returns TRUE if solved, FALSE if impossible"

```go
for row := 0; row < 9; row++ {
    for col := 0; col < 9; col++ {
```
**Translation**: "Look at every single cell in the grid, row by row"

```go
if grid[row][col] == 0 {
```
**Translation**: "If this cell is empty (0 means empty)..."

```go
for num := 1; num <= 9; num++ {
```
**Translation**: "Try numbers 1, 2, 3, 4, 5, 6, 7, 8, 9"

```go
if isValid(grid, row, col, num) {
```
**Translation**: "Ask the safety checker: 'Can I put this number here?'"

```go
grid[row][col] = num
```
**Translation**: "Put the number in the cell!"

```go
if SolveSudoku(grid) {
    return true
}
```
**Translation**: "Try to solve the REST of the puzzle. If it works, yay!"

**🤯 MAGIC MOMENT**: The function calls ITSELF! This is called **recursion**. It's like:
- Step 1: Fill one empty spot
- Step 2: Ask yourself to fill the next empty spot
- Step 3: Ask yourself to fill the next one
- Keep going until done!

```go
grid[row][col] = 0
```
**Translation**: "That didn't work. Erase the number (BACKTRACK!)"

```go
return false
```
**Translation**: "None of 1-9 worked. Tell the previous step to try something different"

```go
return true
```
**Translation**: "No empty cells left! Puzzle solved!"

---

#### Tool #3: The Printer
```go
func PrintSudoku(grid [9][9]int) {
```
**Translation**: "Make a tool that prints the grid nicely"

```go
for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
        fmt.Print(grid[i][j])
```
**Translation**: "Go through each row, then each number in that row, and print it"

```go
if j != 8 {
    fmt.Print(" ")
}
```
**Translation**: "Add a space between numbers, but not after the last one"

```go
fmt.Println()
```
**Translation**: "Move to the next line after each row"

---

### **File 2: main.go** (The Main Program)

```go
func main() {
```
**Translation**: "This is where the program starts!"

```go
if len(os.Args) != 10 {
    fmt.Println("Error")
    return
}
```
**Translation**: 
- `os.Args` = What the user typed when running the program
- `len()` = "How many things did they type?"
- "If they didn't type exactly 10 things (program name + 9 rows), show error"

```go
var grid [9][9]int
```
**Translation**: "Create an empty 9×9 grid"

```go
for i := 0; i < 9; i++ {
    row := os.Args[i+1]
```
**Translation**: 
- "Go through arguments 1-9 (argument 0 is the program name)"
- "Each argument is one row of the puzzle"

```go
if len(row) != 9 {
    fmt.Println("Error")
    return
}
```
**Translation**: "Each row must be exactly 9 characters long"

```go
for j := 0; j < 9; j++ {
    if row[j] == '.' {
        grid[i][j] = 0
```
**Translation**: "If you see a dot (.), put 0 in the grid (0 means empty)"

```go
} else if row[j] >= '1' && row[j] <= '9' {
    grid[i][j] = int(row[j] - '0')
```
**Translation**: 
- "If you see '1' to '9', convert it to a real number"
- `row[j] - '0'` = A trick to convert the character '5' into the number 5

```go
if piscine.SolveSudoku(&grid) {
    piscine.PrintSudoku(grid)
} else {
    fmt.Println("Error")
}
```
**Translation**: "Try to solve it. If it works, print the answer. If not, show error"

---

## 🎮 How to Run This Program

1. **Save two files:**
   - `piscine/solver.go` (the helper tools)
   - `main.go` (the main program)

2. **Run it in terminal:**
```bash
go run . "53..7...." "6..195..." ".98....6." "8...6..3." "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79"
```

**What this does:**
- Each string in quotes is one row
- `.` means empty cell
- Numbers are pre-filled cells

---

## 🌟 The Magic in Simple Terms

**The computer is like a detective:**

1. **Finds an empty cell** (marked with 0)
2. **Tries number 1** → Checks if it's allowed → If yes, places it
3. **Moves to next empty cell** → Tries numbers again
4. **Gets stuck?** → Goes BACK and tries a different number (BACKTRACKING!)
5. **Keeps doing this** until all cells are filled!

---

## 🔑 Key Go Words You Need to Know

| Go Word | What It Means |
|---------|---------------|
| `func` | Make a function (a tool that does work) |
| `var` | Create a variable (storage box) |
| `for` | Loop (repeat something many times) |
| `if` | Check a condition (if this is true, do that) |
| `return` | Give back an answer and stop |
| `true/false` | Yes/No answers (called booleans) |
| `int` | Whole number (like 1, 2, 3) |
| `*` | Pointer (use the real thing, not a copy) |
| `&` | Address of (give the location of something) |

---

## 🎯 Practice Exercise

Try to understand these lines:

```go
num := 5
if num > 3 {
    fmt.Println("Big number!")
}
```

**Translation:**
1. Create a box called `num` and put 5 in it
2. If the number is bigger than 3
3. Print "Big number!" on the screen

---

Does this help you understand the Go language better? 😊
