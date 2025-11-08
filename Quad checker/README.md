
## 🎯 What is this program supposed to do?

You're building a **pattern detector**! 

There are 5 different "quad" functions (quadA, quadB, quadC, quadD, quadE) that draw rectangles using special characters. Your job is to:
1. Read a pattern (rectangle) from input
2. Figure out which quad function(s) could have made it
3. Tell the user which one(s) match

## 📚 Simple Example

Imagine quadA draws this when given 3 width and 3 height:
```
o-o
| |
o-o
```

When someone runs `./quadA 3 3 | go run .`, your program reads that pattern and says: "Hey! This looks like quadA with size 3x3!"

## 🔍 Understanding the Requirements

**Input:** A text pattern (like a rectangle made of characters)

**Output:** 
- If it matches a quad: `[quadA] [3] [3]`
- If multiple quads match: `[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]`
- If nothing matches: `Not a quad function`

## 💡 Step-by-Step Solution Plan

1. **Read the input** (the pattern)
2. **Figure out the dimensions** (width and height)
3. **Generate patterns** using each quad function with those dimensions
4. **Compare** your generated patterns with the input
5. **Print the matches**

## 📝 Simple Code Structure

Here's a beginner-friendly version with lots of comments:

## 🔑 Key Concepts Explained Simply

### 1. **Reading Input** 📖
```go
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
    lines = append(lines, scanner.Text())
}
```
This reads the pattern line by line from the pipe (when you do `./quadA 3 3 | go run .`)

### 2. **Calculating Dimensions** 📏
```go
height := len(lines)        // How many lines?
width := len(lines[0])      // How long is each line?
```

### 3. **Comparing Patterns** 🔍
```go
if generated == input {
    // They match!
}
```

## ⚠️ What You Need To Do

1. **Find your quad implementations** - You should already have quadA, quadB, quadC, quadD, quadE functions from previous exercises
2. **Copy the correct logic** - Replace my example implementations with your actual quad patterns
3. **Test it** - Run the examples from the instructions

## 🧪 Testing

```bash
# Test 1: Single match
./quadA 3 3 | go run .
# Expected: [quadA] [3] [3]

# Test 2: Multiple matches
./quadC 1 1 | go run .
# Expected: [quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
```

## 🤔 Common Mistakes to Avoid

1. **Forgetting the newline** - Make sure each line ends with `\n`
2. **Wrong order** - Results must be alphabetical (quadA, quadB, quadC...)
3. **Width/height confusion** - First number is width, second is height




## 🎨 Visual Understanding

### How quadA draws a 3x3 rectangle:

```
Let me show you step by step:

When x=3, y=3:

Row 0 (top):
  col=0: corner? YES → 'o'
  col=1: corner? NO → '-'
  col=2: corner? YES → 'o'
  Result: "o-o\n"

Row 1 (middle):
  col=0: edge? YES → '|'
  col=1: edge? NO → ' '
  col=2: edge? YES → '|'
  Result: "| |\n"

Row 2 (bottom):
  col=0: corner? YES → 'o'
  col=1: corner? NO → '-'
  col=2: corner? YES → 'o'
  Result: "o-o\n"

Final result: "o-o\n| |\n"o-o\n"
```

### 🔄 How the program flows:

```
START
  ↓
Read input from pipe (scanner)
  ↓
Store each line in a list
  ↓
Check: Do we have lines? NO → Print "Not a quad function"
  ↓ YES
Calculate width (length of first line)
Calculate height (number of lines)
  ↓
Check: All lines same width? NO → Print "Not a quad function"
  ↓ YES
Join all lines into one string
  ↓
Try quadA(width, height) → Does it match input?
Try quadB(width, height) → Does it match input?
Try quadC(width, height) → Does it match input?
Try quadD(width, height) → Does it match input?
Try quadE(width, height) → Does it match input?
  ↓
Collect all matches
  ↓
Any matches? NO → Print "Not a quad function"
  ↓ YES
Print all matches joined with " || "
  ↓
END
```

## 🧩 Key Go Concepts Explained:

### 1. **Variables**
```go
var result string  // Creates an empty text box
x := 5            // Creates a number and puts 5 in it
```

### 2. **Loops**
```go
for i := 0; i < 3; i++ {
    // This runs 3 times
    // i will be: 0, then 1, then 2
}
```

### 3. **If Statements**
```go
if x == 5 {
    // Do this if x equals 5
} else {
    // Do this if x is NOT 5
}
```

### 4. **Lists (Slices)**
```go
var lines []string           // Empty list
lines = append(lines, "hi")  // Add "hi" to list
len(lines)                   // How many items? Answer: 1
```

### 5. **Maps (Dictionaries)**
```go
ages := map[string]int{
    "Alice": 25,
    "Bob": 30,
}
fmt.Println(ages["Alice"])  // Prints: 25
```

## 💡 Important Operators:

- `==` means "is equal to"
- `!=` means "is NOT equal to"
- `<=` means "less than or equal to"
- `||` means "OR"
- `&&` means "AND"
- `+=` means "add to the end"

## 🎯 What YOU need to do:

1. **Copy this code** to a file called `main.go`
2. **Implement quadB, quadC, quadD, quadE** - Use the same pattern as quadA, but with different characters
3. **Test it**:
```bash
./quadA 3 3 | go run .
```


// package main tells Go this is a standalone program (not a library)
// Think of it like saying "This is the MAIN program, start here!"
package main

// import brings in pre-made code that other people wrote
// It's like borrowing tools from a toolbox
import (
	"bufio"   // This helps us read input line by line (buf = buffer, io = input/output)
	"fmt"     // This helps us print stuff (format and print)
	"os"      // This lets us talk to the operating system (read from keyboard, etc.)
	"strings" // This helps us work with text (join, split, etc.)
)

// func means "function" - a reusable piece of code
// quadA is the name of this function
// (x, y int) means: this function takes two numbers called x and y
// string means: this function gives back (returns) text
func quadA(x, y int) string {
	// This is a safety check!
	// If x or y is 0 or negative, we can't draw a rectangle
	// So we return an empty string "" (nothing)
	if x <= 0 || y <= 0 {
		return ""
	}
	
	// var means "variable" - a box to store something
	// result is the name of our box
	// string means it holds text
	// We start with empty text
	var result string
	
	// This is a LOOP - it repeats code multiple times
	// row starts at 0, keeps going while row < y, adds 1 each time
	// So if y=3, row will be: 0, then 1, then 2 (that's 3 rows!)
	for row := 0; row < y; row++ {
		
		// Another loop! This one goes across (columns)
		// If x=3, col will be: 0, then 1, then 2 (that's 3 columns!)
		for col := 0; col < x; col++ {
			
			// Now we decide WHAT CHARACTER to draw
			// Are we on the top row OR bottom row?
			if row == 0 || row == y-1 {
				// Top or bottom row
				
				// Are we at the left corner OR right corner?
				if col == 0 || col == x-1 {
					result += "o" // += means "add to the end of result"
					// Draw 'o' at corners
				} else {
					result += "-" // Draw '-' for horizontal edges
				}
			} else {
				// We're in a middle row (not top, not bottom)
				
				// Are we on the left edge OR right edge?
				if col == 0 || col == x-1 {
					result += "|" // Draw '|' for vertical edges
				} else {
					result += " " // Draw a space in the middle (empty)
				}
			}
		}
		// After each row is done, add a newline (go to next line)
		result += "\n"
	}
	// Give back the complete pattern
	return result
}

func quadB(x, y int) string {
	// TODO: Implement quadB pattern
	// Example pattern might be different corners/edges
	return "" // Replace with actual implementation
}

func quadC(x, y int) string {
	// TODO: Implement quadC pattern
	if x <= 0 || y <= 0 {
		return ""
	}
	
	var result string
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 {
				// Top row
				if col == 0 {
					result += "A"
				} else if col == x-1 {
					result += "A"
				} else {
					result += "B"
				}
			} else if row == y-1 {
				// Bottom row
				if col == 0 {
					result += "C"
				} else if col == x-1 {
					result += "C"
				} else {
					result += "B"
				}
			} else {
				// Middle rows
				if col == 0 || col == x-1 {
					result += "B"
				} else {
					result += " "
				}
			}
		}
		result += "\n"
	}
	return result
}

func quadD(x, y int) string {
	// TODO: Implement quadD pattern
	if x <= 0 || y <= 0 {
		return ""
	}
	
	var result string
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 {
				if col == 0 {
					result += "A"
				} else if col == x-1 {
					result += "C"
				} else {
					result += "B"
				}
			} else if row == y-1 {
				if col == 0 {
					result += "A"
				} else if col == x-1 {
					result += "C"
				} else {
					result += "B"
				}
			} else {
				if col == 0 || col == x-1 {
					result += "B"
				} else {
					result += " "
				}
			}
		}
		result += "\n"
	}
	return result
}

func quadE(x, y int) string {
	// TODO: Implement quadE pattern
	if x <= 0 || y <= 0 {
		return ""
	}
	
	var result string
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 {
				if col == 0 {
					result += "A"
				} else if col == x-1 {
					result += "C"
				} else {
					result += "B"
				}
			} else if row == y-1 {
				if col == 0 {
					result += "C"
				} else if col == x-1 {
					result += "A"
				} else {
					result += "B"
				}
			} else {
				if col == 0 || col == x-1 {
					result += "B"
				} else {
					result += " "
				}
			}
		}
		result += "\n"
	}
	return result
}

// func main is SPECIAL - it's where your program starts!
// When you run "go run .", Go looks for main() and starts there
func main() {
	// ============================================
	// STEP 1: READ THE INPUT
	// ============================================
	
	// bufio.NewScanner creates a "scanner" object
	// os.Stdin means "standard input" (the text coming from the pipe |)
	// Think of scanner as a reading machine
	scanner := bufio.NewScanner(os.Stdin)
	
	// This will hold all the lines we read
	// []string means "a list of strings" (text)
	// We start with an empty list
	var lines []string
	
	// This loop keeps reading until there's no more input
	// scanner.Scan() reads one line and returns true if successful
	for scanner.Scan() {
		// scanner.Text() gets the line we just read
		// append adds it to our lines list
		lines = append(lines, scanner.Text())
	}
	
	// ============================================
	// STEP 2: CHECK IF WE GOT ANY INPUT
	// ============================================
	
	// len(lines) tells us how many lines we have
	// If we have 0 lines, there's nothing to check!
	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return // Stop the program here
	}
	
	// ============================================
	// STEP 3: FIGURE OUT THE DIMENSIONS
	// ============================================
	
	// Height = number of lines
	// If we have 3 lines, height is 3
	height := len(lines)
	
	// Width = length of the first line
	// We assume all lines have the same length (we'll check this later)
	width := 0
	if height > 0 {
		width = len(lines[0])
	}
	
	// ============================================
	// STEP 4: VALIDATE ALL LINES HAVE SAME WIDTH
	// ============================================
	
	// This loop goes through each line
	// range lines gives us each line one by one
	// _ means "I don't care about the index"
	for _, line := range lines {
		// If any line has a different length, it's not a valid rectangle!
		if len(line) != width {
			fmt.Println("Not a quad function")
			return // Stop here
		}
	}
	
	// ============================================
	// STEP 5: BUILD ONE BIG STRING FROM ALL LINES
	// ============================================
	
	// strings.Join puts all lines together with \n between them
	// Example: ["ABC", "DEF"] becomes "ABC\nDEF"
	input := strings.Join(lines, "\n")
	
	// Add a final newline at the end
	if len(input) > 0 {
		input += "\n"
	}
	
	// ============================================
	// STEP 6: TRY EACH QUAD FUNCTION
	// ============================================
	
	// This list will hold all the matching quad functions
	var matches []string
	
	// This is a MAP - it's like a dictionary
	// Key: name of the function (text)
	// Value: the actual function itself
	// Think of it like: "quadA" points to the quadA function
	quads := map[string]func(int, int) string{
		"quadA": quadA,
		"quadB": quadB,
		"quadC": quadC,
		"quadD": quadD,
		"quadE": quadE,
	}
	
	// We need to check them in alphabetical order
	// So we make a list of names in order
	quadNames := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	
	// Loop through each quad name
	for _, name := range quadNames {
		// Get the actual function from our map
		// quads[name] looks up the function by name
		quadFunc := quads[name]
		
		// Call the function with our width and height
		// This generates what THIS quad would look like
		generated := quadFunc(width, height)
		
		// Compare: does the generated pattern match what we read?
		if generated == input {
			// YES! This quad matches!
			// Create a string like "[quadA] [3] [3]"
			// %s means "put a string here"
			// %d means "put a number here"∆
			match := fmt.Sprintf("[%s] [%d] [%d]", name, width, height)
			
			// Add this match to our list of matches
			matches = append(matches, match)
		}
	}
	
	// ============================================
	// STEP 7: PRINT THE RESULTS
	// ============================================
	
	// Did we find any matches?
	if len(matches) == 0 {
		// No matches found
		fmt.Println("Not a quad function")
	} else {
		// We found matches!
		// Join them with " || " between them
		// Example: ["[quadA] [3] [3]", "[quadB] [3] [3]"] 
		//       becomes "[quadA] [3] [3] || [quadB] [3] [3]"
		result := strings.Join(matches, " || ")
		
		// Print it!
		fmt.Println(result)
	}
}
```
