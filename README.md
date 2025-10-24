# 🧠 Piscine-Go Shell & Command Line Quests 1 Solutions

This repository contains my completed solutions for several shell scripting and command line challenges from the **01 Edu Piscine Go** training series.  
Each task focuses on building Unix command fluency, data parsing, and problem-solving logic.

---

## 🔍 `look`
**Objective:**  
Search from the current directory and its subfolders for all `.sh` files and display their names (with the `.sh` extension included).

**Tools used:**  
`find`, basic shell redirection

**Final solution:**
```bash
#!/bin/bash
find . \( \
  -name 'a*' \
  -o \( -type f -name '*z' \) \
  -o \( -type f -name 'z*a!' \) \
\)
```

# Shell Scripting Solutions

## 1. myfamily.sh - "someone familiar"

**Task:** Extract relatives information for a superhero by ID from a JSON API.

**File:** `myfamily.sh`
```bash
#!/bin/bash

# Fetch the JSON data and extract relatives for the hero with ID from HERO_ID
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | \
  jq --arg hero_id "$HERO_ID" '.[] | select(.id == ($hero_id | tonumber)) | .connections.relatives' | \
  sed 's/^"//; s/"$//'
```

**Usage:**
```bash
export HERO_ID=1
./myfamily.sh
```

**Key Points:**
- Uses `curl -s` to silently fetch JSON
- Uses `jq` to parse JSON and filter by ID
- Path is `.connections.relatives` (not just `.relatives`)
- Removes surrounding quotes with `sed` but keeps internal `\n` as literal text

---

## 2. lookagain.sh - "keep looking..."

**Task:** Find all `.sh` files, show only filenames (no path, no extension), sorted descending.

**File:** `lookagain.sh`
```bash
#!/bin/bash

# Find all .sh files, get only the filename, remove the .sh extension, and sort in descending order
find . -type f -name "*.sh" -exec basename {} \; | cut -d'.' -f1 | sort -r
```

**Usage:**
```bash
./lookagain.sh | cat -e
```

**Key Points:**
- `find . -type f -name "*.sh"` finds all shell scripts
- `basename` extracts just the filename (no path)
- `cut -d'.' -f1` removes the `.sh` extension
- `sort -r` sorts in reverse (descending) order

---

## 3. countfiles.sh - "Now, do your inventory"

**Task:** Count total number of files and folders (including current directory).

**File:** `countfiles.sh`
```bash
#!/bin/bash

# Count all regular files and directories including the current directory
find . | wc -l
```

**Usage:**
```bash
./countfiles.sh | cat -e
```

**Key Points:**
- `find .` lists everything including `.` (current directory)
- `wc -l` counts the lines (total items)

---

## 4. Special Filename - "Be accurate"

**Task:** Create a file named `"\?$*'ChouMi'*$?\"` containing `01`

**Command (for Linux/Mac/WSL):**
```bash
echo -n "01" > '"\?$*'ChouMi'*$?\"'
```

**Alternative Methods:**
```bash
# Using printf
printf "01" > '"\?$*'ChouMi'*$?\"'

# Using cat
cat > '"\?$*'ChouMi'*$?\"' << 'EOF'
01
EOF
```

## 5. skip.sh - "pick your equipment"

**Task:** Print `ls -l` output, showing every other line starting with the first.

**File:** `skip.sh`
```bash
#!/bin/bash

# Print ls -l output, skipping 1 line out of 2, starting with the first one
ls -l | awk 'NR % 2 == 1'
```

**Alternative using sed:**
```bash
#!/bin/bash

ls -l | sed -n 'p;n'
```

**Usage:**
```bash
./skip.sh
```

**Key Points:**
- `ls -l` lists files in long format
- `awk 'NR % 2 == 1'` prints only odd-numbered lines (1, 3, 5, ...)
- `NR` is the line number in awk
- `NR % 2 == 1` means "line number modulo 2 equals 1" (odd lines)

---
## 6. 🕵️‍♂️ the-final-cl-test (Murder Mystery)

Objective:
Solve the Terminal City murder mystery using Linux commands and data from the mystery/ directory.

🧩 Steps

Find clues in mystery/crimescene:

Suspect is a tall male (6’+)

Found wallet with memberships:

AAA

Delta SkyMiles

Local Library

Museum of Bash History

Witness Annabel saw a blue Honda with a plate starting with L337 and ending with 9

Combine data:

Intersect membership lists to find all people in those four clubs.

Filter for males ≥ 6 ft using mystery/people.

Match L337…9 plates and blue Hondas in mystery/vehicles.

Find suspects:

Candidates: Brian Boyer, Erika Owens, Joe Germuska, Mike Bostock, Matt Waite, etc.

Cross-checking revealed Dartey Henv as the only tall male wallet holder owning a blue Honda.

Key witness:

```From mystery/interviews/interview-699607:

“Interviewed Ms. Church… Describes a blue Honda with a plate starting with ‘L337’ and ending with ‘9’.”


Witness: Annabel Church
```
Conclusion:
The murderer is Dartey Henv — tall male, blue Honda, L337…9, wallet matches clues.

`my_answer.sh`

```bash
#!/bin/bash

echo "Dartey Henv"
```
---
## 7. 🧾 `explain.sh`

**Objective:**  
Summarize how the case was solved, showing:

1. The key witness’s name  
2. Her interview number  
3. The car color and make of the main suspect  
4. The 3 other suspects not arrested (in alphabetical order of last name)

**Final output:**
```bash
#!/bin/bash

echo "Annabel Church"
echo "699607"
echo "Blue Honda"
echo "Joe Germuska"
echo "Hellen Maher"
echo "Erika Owens"
echo ""
```
---
## 8.🧑‍🏫 `teacher.sh`

**Objective:**  
Automate the process of identifying and displaying the key evidence used to solve the Terminal City mystery, while adapting to different `mystery` folder locations in each training dataset.

**Steps performed by the script:**

1. **Step 1:**  
   Locate the key interview file containing the phrase related to the clue (`L337...9`) anywhere within a `mystery/interviews` directory.  
   Extract and store only the interview number (digits) into the environment variable `KEY_INTERVIEW`.

2. **Step 2:**  
   Print the value of the environment variable `KEY_INTERVIEW`.

3. **Step 3:**  
   Display the full contents of the interview file.

4. **Step 4:**  
   Print the value of the environment variable `MAIN_SUSPECT` (set by the environment or the grading system).

---

**Final Output Example:**
```
699607
Interviewed Ms. Church at 2:04 pm. Witness stated that she did not see anyone she could identify as the shooter, that she ran away as soon as the shots were fired.

However, she reports seeing the car that fled the scene. Describes it as a blue Honda, with a license plate that starts with "L337" and ends with "9"
Dartey Henv
```

---

**Final Solution:**
```bash
#!/usr/bin/env bash

# Step 1: Isolate the key interview number into an environment variable
export INTERVIEW=$(grep -h "SEE INTERVIEW" mystery/streets/* | grep -oE '[0-9]+')

# Step 2: Print the newly created environment variable
echo "$INTERVIEW"

# Step 3: Print what the interview contains
cat mystery/interviews/interview-"$INTERVIEW"

# Step 4: Print the content of the environment variable MAIN_SUSPECT
echo "$MAIN_SUSPECT"
```
---

## How to Push to GitHub
```bash
# Make all scripts executable
chmod +x myfamily.sh lookagain.sh countfiles.sh skip.sh

# Add all files
git add myfamily.sh lookagain.sh countfiles.sh skip.sh

# Commit with a descriptive message
git commit -m "Add shell scripting solutions: myfamily, lookagain, countfiles, skip"

# Push to GitHub
git push origin main
```

---

## Quick Reference Commands

### Make scripts executable:
```bash
chmod +x *.sh
```

### Test individual scripts:
```bash
# Test myfamily.sh
export HERO_ID=1
./myfamily.sh

# Test lookagain.sh
./lookagain.sh

# Test countfiles.sh
./countfiles.sh

# Test skip.sh
./skip.sh
```

### Clean up test files:
```bash
rm -f file_name # Remove the test file created 
```

---

## Troubleshooting

### Issue: "Permission denied"
**Solution:**
```bash
chmod +x script_name.sh
```

### Issue: jq command not found
**Solution:**
```bash
# Ubuntu/Debian
sudo apt-get install jq

```

### Issue: Special character filename won't create on Windows
**Solution:** Use WSL or test on Linux system. Windows filesystem doesn't support `\`, `?`, `*`, `"` in filenames.

---

## Summary

| Script | Purpose | Key Tools |
|--------|---------|-----------|
| myfamily.sh | Extract superhero relatives from JSON API | curl, jq, sed |
| lookagain.sh | List .sh files without extension, sorted | find, basename, cut, sort |
| countfiles.sh | Count all files and directories | find, wc |
| Special file | Create file with special characters | echo, special quoting |
| skip.sh | Show every other line from ls -l | ls, awk |

---

**All scripts are ready to push! 🚀**
# 🧠 Piscine-Go Quests 2 Solutions
  
All work is submitted from the root repo: `piscine-go`.

---

## 🔧 One-Time Setup (do this once in `piscine-go`)
```bash
go mod init piscine
go get github.com/01-edu/z01
go mod tidy
```
### 🧼 Formatting Before Every Push (do this for each exercise)

Run one of these in the exercise folder (e.g., printalphabet/ or printreversealphabet/):
```bash
gofmt -w .
```
or
```bash
gofumpt -w .
```
Then commit & push.
## Let's Start!
## 🔤 `printalphabet`

**Goal:** print the lowercase Latin alphabet on one line, then a newline.

**File to submit:** `printalphabet/main.go`
**Allowed function:** `github.com/01-edu/z01.PrintRune`

# Solution:
```go
package main

import "github.com/01-edu/z01"

func main() {
	for char := 'a'; char <= 'z'; char++ {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

```
### ▶️ Usage
```
$ go run .
abcdefghijklmnopqrstuvwxyz
$
```
---
Then push only these files and remember to format(check above for the command):

```go
go.mod
printalphabet/main.go
```
---

Note:
All exercises (like printalphabet, printdigits, etc.) must be in their own folders inside `piscine-go/` and submitted from there.
If Go command not found install, then install it 
```bash
sudo apt update
#quick update
sudo apt install golang
#this installs total packages
go version #confirms the version installed
```
---
## 🧩 printreversealphabet
### 🧰 Instructions

Write a program that prints the Latin alphabet in lowercase in reverse order (from z to a) on a single line.

A line is a sequence of characters preceding the end-of-line character ('\n').
### Code
```go
package main

import "github.com/01-edu/z01"

func main() {
	for ch := 'z'; ch >= 'a'; ch-- {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

```
### Next: remember to format and push only the files created:
**File to submit:** `printreversealphabet/main.go`
---
## Printdigits:
## Explanation:

`for digit := '0'; digit <= '9'; digit++` - Loop through ASCII characters from '0' to '9'
`z01.PrintRune(digit)` - Print each digit character using the allowed function <br>
`z01.PrintRune('\n')` - Print a newline at the end
```go
package main

import "github.com/01-edu/z01"

func main() {
	for ch := '0'; ch <= '9'; ch++ {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

```
`go run printdigits/main.go`
*Output:*
0123456789
### Next: remember to format and push only the files created:
**File to submit:** `printdigits/main.go`
---
# isnegative

## 📋 Instructions
Write a function that prints `'T'` (true) on a single line if the `int` passed as parameter is negative, otherwise it prints `'F'` (false).

**Level:** 6  
**Files to submit:** `isnegative.go`  
**Allowed functions:** `github.com/01-edu/z01.PrintRune`, `--allow-builtin`

## 📝 Expected Function
```go
func IsNegative(nb int) {
}
```

## 💡 Solution
```go
package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
```

## 🔍 Explanation
1. **`if nb < 0`** - Check if the number is negative
2. **`z01.PrintRune('T')`** - Print 'T' for true (number is negative)
3. **`z01.PrintRune('F')`** - Print 'F' for false (number is zero or positive)
4. **`z01.PrintRune('\n')`** - Print a newline character after the result

## 🧪 Usage Example (To test locally) !!!! Pls skip this if you are not interested in testing and jump to 'How to Submit below'
```bash
mkdir -p test #create a test folder
cd test #moves you into the test folder 📁 created so you can create the main file below
nano main.go # paste the test code below into this file
```
```go
package main

import "piscine"

func main() {
	piscine.IsNegative(1)   // Not negative
	piscine.IsNegative(0)   // Not negative (zero)
	piscine.IsNegative(-1)  // Negative
}
```
`cd test` to make sure you are seeing `/piscine-go/test ` now Run it
`go run .` <br>
*📤 Expected Output*
```
F
F
T
```

## 🎯 Test Cases
| Input | Output | Reason |
|-------|--------|--------|
| `1` | `F` | 1 is positive |
| `0` | `F` | 0 is not negative |
| `-1` | `T` | -1 is negative |
| `42` | `F` | 42 is positive |
| `-999` | `T` | -999 is negative |

## 📌 Key Points
- Zero (0) is **not** considered negative, so it returns 'F'
- Only numbers less than 0 return 'T'
- Each result must be followed by a newline character
- Must use `github.com/01-edu/z01.PrintRune` for output
- Function must be in the `piscine` package

## 🚀 How to Submit (remember to format)
if you tested, make sure you go back to the `piscine-go` folder 📁 using this command `cd ..`
Confirm you are back `~piscine-go` if so, you're good 2 go 👍 
```bash
git add isnegative.go
git commit -m "Add isnegative solution"
git push
```
---
# Printcomb

## 📋 Instructions
Write a function that prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.
These combinations are separated by a comma and a space.

**Files to submit:** `printcomb.go`  

## 📝 Expected Function
```go
func PrintComb () {
}
```

## 💡 Solution
```go
package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				
				// Don't print comma and space after the last combination (789)
				if i != '7' || j != '8' || k != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

```

## 🔍 Explanation
1. **First loop:** `i := '0'; i <= '7'` - First digit ranges from 0 to 7
2. **Second loop:** `j := i + 1; j <= '8'` - Second digit starts from i+1 (ensuring i < j) up to 8
3. **Third loop:** `k := j + 1; k <= '9'` - Third digit starts from j+1 (ensuring j < k) up to 9
4. **Print combination:** Print all three digits
5. **Comma and space:** Add ", " after each combination except the last one (789)
6. **Newline:** Print '\n' at the end
## 🧪 Usage Example (To test locally) !!!! Pls skip this if you are not interested in testing and jump to 'How to Submit below'
replace your previous test file `main.go` with this:
```go
package main

package main

import "piscine"

func main() {
	piscine.PrintComb()
}
```
`cd test` and Run it
`go run .` <br>
*📤 check Expected Output on your dashboard*
## 🚀 How to Submit (remember to format)
if you tested, make sure you go back to the `piscine-go` folder using this command `cd ..`
```bash
By now you should have mastered the art of adding a file, commiting with a message and pushing.
I think you gat this, thumbs up! 👍 😎
If you see any error occurred check again and again, you probably made a mistake terminal is your friend.
```
---
# 🧩 printcomb2

### 🧰 Instructions
Write a function that prints **all possible combinations of two different two-digit numbers** (`00 01`, `00 02`, …, `98 99`) in **ascending order** on a single line.

Each combination must be separated by a comma and a space, and the final output should end with a newline.

**Allowed function:**  
`github.com/01-edu/z01.PrintRune`  
**Casting is not allowed.**

---

### ✅ Code (`printcomb2.go`)
```go
package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var i rune
	var j rune
	for i = 0; i <= 98; i++ {
		for j = i + 1; j <= 99; j++ {
			z01.PrintRune('0' + i/10)
			z01.PrintRune('0' + i%10)
			z01.PrintRune(' ')
			z01.PrintRune('0' + j/10)
			z01.PrintRune('0' + j%10)
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

```
### ▶️ Usage Example
This is the test file below (replace main.go):
```go
package main

import "piscine"

func main() {
	piscine.PrintComb2()
}

```
**Expected Output:**
```
$ go run .
00 01, 00 02, 00 03, ..., 98 99
$
```
*Always format your code before you push:*
```bash
gofmt -w .
# or, if installed:
gofumpt -w .
```
**File to push:** `printcomb2/printcomb2.go`
---
## Here's the solution for the `printnbr.go` 

📋 **Explanation:**

1. **Handle minimum int:** The most negative int value (-9223372036854775808) cannot be converted to positive by simply negating it (it would overflow). So we handle it specially by printing '-' and '9', then recursively printing the rest.

2. **Handle negative numbers:** If n is negative, print '-' and make n positive.

3. **Recursive printing:** 
   - If n >= 10, recursively call PrintNbr with n/10 to print all digits except the last
   - Print the last digit using `'0' + n%10`

## 🧪 Code

```bash
package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// Handle the special case of the most negative int
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		PrintNbr(223372036854775808)
		return
	}

	// Handle negative numbers
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	// Recursively print digits
	if n >= 10 {
		PrintNbr(n / 10)
	}

	z01.PrintRune(rune('0' + n%10))
}
```
**Always make sure you format immediately after creating your file**
use this command 
```bash
gofmt -w .
```
## 📤 **Expected Output:**
```
-1230123
```

## 🎯 **Test Cases:**
| Input | Output |
|-------|--------|
| `-123` | `-123` |
| `0` | `0` |
| `123` | `123` |
| `-9223372036854775808` | `-9223372036854775808` |
| `9223372036854775807` | `9223372036854775807` |

## ✅ **Key Points:**
- Handles all possible int values including the minimum int (-9223372036854775808)
- Uses recursion to print digits from left to right
- Cannot use int64 conversion
- Uses only z01.PrintRune for output

🚀 Ready to submit!
---
##  `printcombn.go` 
📋 **Explanation:**

1. **PrintCombN(n int):** Main function that initializes and starts the combination generation
2. **printCombinations:** Recursive function that generates all valid combinations
   - `start`: The minimum digit for current position (must be greater than previous digit)
   - `digit <= 9-(n-pos-1)`: Ensures we have enough digits left for remaining positions
3. **printComb:** Prints a single combination and adds comma/space if not the last one

## 🎯 **Key Logic:**
- For n=3: Last combination is 789 (digits 7,8,9)
- For n=9: Last combination is 012345678 (digits 0,1,2,3,4,5,6,7,8)
- Formula for last combination digit at position i: `10 - n + i`

## 🧪 Code

```bash
package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	// Create a slice to hold the current combination
	comb := make([]int, n)

	// Initialize the first combination (0, 1, 2, ..., n-1)
	for i := 0; i < n; i++ {
		comb[i] = i
	}

	// Print combinations
	printCombinations(comb, n, 0)
	z01.PrintRune('\n')
}

func printCombinations(comb []int, n int, pos int) {
	// Base case: we've filled all positions
	if pos == n {
		printComb(comb, n)
		return
	}

	// Determine the starting digit for this position
	start := 0
	if pos > 0 {
		start = comb[pos-1] + 1
	}

	// Try all possible digits for this position
	for digit := start; digit <= 9-(n-pos-1); digit++ {
		comb[pos] = digit
		printCombinations(comb, n, pos+1)
	}
}

func printComb(comb []int, n int) {
	// Print the combination
	for i := 0; i < n; i++ {
		z01.PrintRune(rune('0' + comb[i]))
	}

	// Check if this is the last combination
	isLast := true
	for i := 0; i < n; i++ {
		if comb[i] != 10-n+i {
			isLast = false
			break
		}
	}

	// Print comma and space if not the last combination
	if !isLast {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

```
**Always make sure you format immediately after creating your file**
use this command 
```bash
gofmt -w .
```

## 📤 **Expected Output:**

**For n=1:**
```
0, 1, 2, 3, 4, 5, 6, 7, 8, 9
```

**For n=3:**
```
012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 689, 789
```

**For n=9:**
```
012345678, 012345679, ..., 123456789
```

## ✅ **Efficiency:**
- Uses recursion with proper bounds to avoid unnecessary iterations
- Only generates valid combinations (no backtracking needed)
- Time complexity: O(C(10,n)) where C is combinations

🚀 Ready to submit!
---
### If you encounter Formatting error!!! then you forgot to format the go file
### If you encounter go mod init piscine error!! you skipped the first instruction
### If you encounter go run package error! go and delete the required github library installed inside your `go.mod` file it must only have two lines, delete the 3rd one:
```bash
module piscine

go 1.25.3

require github.com/01-edu/z01 v0.2.0 // indirect #Delete this last line to clear that error
```
## If you are still struggling to understand for your Checkpoint Practice? After copy and pasting, try to Check the ReadME file in the 📁 titled **Learn Quest2** I hope that helps. GoodLuck!

### Also check -README file in 📁 Resource for beginner-friendly guides. 
---
# QUEST-3
Excellent 👏 — this one introduces **pointers** 🧠
Let’s make it super simple so you’ll *never* be afraid of pointers again.

---

## 🧩 What They’re Asking

> “Write a function that takes a **pointer to an int**, and makes that int equal to 1.”

So the function doesn’t return anything —
it just *changes* the number that was given **through a pointer**.

---

### 👶 Imagine this story

You have a box 📦 that contains a number — let’s call the box `n`.
Normally, if you hand someone the **number itself**, they can’t change your box.

But if you hand them the **key to your box** (the *address*),
they can open it and put a new number inside.

That “key” is called a **pointer** 🔑

---

## ✅ Solution Code

```go
package piscine

func PointOne(n *int) {
	*n = 1
}
```
**Now that you ve created the file, format it now**
---

### 🧒 Explanation

* `n *int` → means:
  “I’m not getting an integer directly — I’m getting a *pointer to* an integer.”

* `*n = 1` →
  means: “go to that address and change what’s inside to 1.”

---

### 🧪 Test Program (you can run this locally)

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	n := 0          // normal integer
	piscine.PointOne(&n) // send the address of n (not the value)
	fmt.Println(n)  // prints 1
}
```

### Output

```
$ go run .
1
```

---

### 💡 Key Concepts

| Concept  | Meaning                                   |
| -------- | ----------------------------------------- |
| `*int`   | A pointer to an integer                   |
| `&n`     | “Address of” n — like the box’s key       |
| `*n`     | “Value inside the address” — open the box |
| `*n = 1` | Change what’s inside the box to 1         |

---

### 🧰 What to Push

Just one file:

```
 git add pointone.go
```

---

Don’t panic — this one looks scary only because there are more `*`,
but it’s the *same logic* as before, just more boxes 📦 inside boxes.

---

## 🧩UltimatePointOne

> “Write a function that takes a pointer to a pointer to a pointer to an int
> and gives that int the value of 1.”

In plain English:

> “You get a key to a box, inside it there’s another box, and inside *that* one is the number — set it to 1.”

---

## ✅ Solution Code

```go
package piscine

func UltimatePointOne(n ***int) {
	***n = 1
}
```
**If you like no remember to format LOL**
---

## 🧒 Step-By-Step Explanation

Let’s break this down gently:

### 1️⃣ You start with an `int`

```go
a := 0  // box that holds 0
```

### 2️⃣ A pointer to `a`

```go
b := &a  // b holds the address of a
```

### 3️⃣ A pointer to that pointer

```go
n := &b  // n holds the address of b
```

### 4️⃣ The function takes a pointer to *that*

```go
piscine.UltimatePointOne(&n)
```

So the argument `n` inside your function is actually `***int` (pointer to pointer to pointer).

---

### 🧠 Inside the function:

* `n` → pointer to pointer to pointer
* `*n` → pointer to pointer
* `**n` → pointer to int
* `***n` → the integer itself ✅

So `***n = 1` means:

> go three levels deep and change the real number to 1.

---

## 🧪 Test Code

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 0
	b := &a
	n := &b
	piscine.UltimatePointOne(&n)
	fmt.Println(a)
}
```

### Output

```
$ go run .
1
```

---

## 🧩 Visual Analogy (for fun)

| Level      | Variable | What it stores                                       |
| ---------- | -------- | ---------------------------------------------------- |
| 🎯 Level 1 | `a`      | 0                                                    |
| 📦 Level 2 | `b`      | → address of `a`                                     |
| 📦 Level 3 | `n`      | → address of `b`                                     |
| 🔑 Level 4 | `&n`     | → address of `n` (that’s what the function receives) |

So to open all those boxes:
`*` opens 1 layer → `***` opens 3 layers → reveals `a`.

---

## 💡 Key Lesson

| Symbol | Meaning                           |
| ------ | --------------------------------- |
| `*`    | open one box                      |
| `**`   | open two boxes                    |
| `***`  | open three boxes                  |
| `&`    | take the address (make a pointer) |

---

That’s literally all there is to it — same idea as before, just deeper nesting 😄
Beautiful — this one brings together **math + pointers**, so let’s keep the same “baby logic” 🧒🧮

---

## divmod.go-🧩 The Task

> “Write a function that divides two integers (`a` and `b`),
> stores the **quotient** in one pointer and the **remainder** in another.”

---

## ✅ Solution Code

```go
package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}
```
**Hope you remember to format**
---

## 🧒 Explanation (Like Teaching a Child)

You know when you divide **13 ÷ 2** on paper?

```
2 goes into 13 → 6 times, remainder 1
```

So:

* Quotient (the result of division) = `6`
* Remainder (what’s left) = `1`

In Go:

* `a / b` gives **quotient**
* `a % b` gives **remainder**

But the tricky part here is the **pointers**.
We don’t print or return the results — we put them *inside the boxes* (the memory spots) that were given to us.

---

### 💡 Line-by-line Explanation

```go
*div = a / b
```

👉 Go to the memory location `div` is pointing at, and write the quotient there.

```go
*mod = a % b
```

👉 Go to the memory location `mod` is pointing at, and write the remainder there.

---

## 🧪 Test Program

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	piscine.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
```

### 🧾 Output

```
$ go run .
6
1
```

---

## 📦 Quick Summary

| Symbol  | Meaning                                          |
| ------- | ------------------------------------------------ |
| `*div`  | Value inside the pointer (place to store answer) |
| `a / b` | Division result                                  |
| `a % b` | Remainder of the division                        |
| `&div`  | Address of div (we send this to the function)    |

---

## 🧠 How to Remember

Think of it like this:

> “I tell my friend the answers,
> and my friend writes them into two different boxes I gave them.” 📦📦

One box holds **div**, the other holds **mod**.

---

This one is just the *upgraded version* of **DivMod**, but now with **pointers everywhere** 🔁

---

## Ultimatedivmod 🧩 Task Summary

> “You’re given two pointers to integers, `a` and `b`.”
> You must:
>
> * Divide the value inside `a` by the value inside `b`
> * Store the **quotient** back inside `a`
> * Store the **remainder** back inside `b`

---

## ✅ Solution Code

```go
package piscine

func UltimateDivMod(a *int, b *int) {
	quotient := *a / *b
	remainder := *a % *b
	*a = quotient
	*b = remainder
}
```
**gofmt -w .**
---

## 🧒 Explanation (as simple as possible)

Think of:

* `*a` → the number **inside box A**
* `*b` → the number **inside box B**

You don’t want to **return** anything —
you just open both boxes, do your math, and put the results back inside.

---

### Step-by-step logic:

Let’s say:

```go
a = 13
b = 2
```

1️⃣ **Get the numbers inside the boxes**

```go
*a = 13
*b = 2
```

2️⃣ **Do the math**

```go
quotient = 13 / 2  // → 6
remainder = 13 % 2 // → 1
```

3️⃣ **Put them back**

```go
*a = 6
*b = 1
```

So now:

```
a = 6
b = 1
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
	a := 13
	b := 2
	piscine.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
```

### Output:

```
$ go run .
6
1
```

---

## 💡 Key Lessons

| Symbol    | Meaning                        |
| --------- | ------------------------------ |
| `*a`      | the value inside the pointer a |
| `*b`      | the value inside the pointer b |
| `*a / *b` | divide values                  |
| `*a % *b` | remainder                      |
| `a *int`  | a is a pointer to an integer   |
| `&a`      | address of a (pointer)         |

---

## 🧠 Easy Way to Remember

> 🧮 **DivMod formula:**
> Store `(a / b)` in `a`,
> Store `(a % b)` in `b`.

Or say it like this:

> “A gets the quotient, B gets the remainder.” 💥

---

— welcome to **PrintStr**, one of the easiest but most *fundamental* exercises in Go strings and runes.

Let’s make this simple and fun 👇

---

## PrintStr 🧩 Task Summary

> Write a function that prints every character of a string one by one on the screen.

---

## ✅ Solution Code

```go
package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
```
**I know you don tire for my reminder sorry**
---

## 🧒 Explanation (Step by Step)

Let’s break it down like storytime 📖:

### 1️⃣ What’s a **string**?

A string is just a list (sequence) of **characters** — like a word:

```
"Hello"
```

is actually
`H`, `e`, `l`, `l`, `o`

---

### 2️⃣ What’s `for _, char := range s`?

That’s Go’s way of saying:

> “Go through every character (rune) inside the string `s`.”

* `_` = ignore the index number
* `char` = the current character (like `'H'`, `'e'`, `'l'`, etc.)

---

### 3️⃣ What’s `z01.PrintRune(char)`?

It **prints** one single character (rune) at a time on the screen —
so when we do it in a loop, it prints the whole string, letter by letter.

---

### 4️⃣ Why not just `fmt.Println(s)`?

Because this challenge wants you to *manually* print every rune using `z01.PrintRune`, not the easy way 😄

You’re training to understand **how iteration works on strings** in Go.

---

## 🧪 Test Program

```go
package main

import "piscine"

func main() {
	piscine.PrintStr("Hello World!")
}
```

### 🧾 Output

```
$ go run . | cat -e
Hello World!$
```

---

## 💡 Quick Recap

| Concept                  | Meaning                                |
| ------------------------ | -------------------------------------- |
| `string`                 | sequence of runes (characters)         |
| `range s`                | loops through each character in string |
| `z01.PrintRune(char)`    | prints one rune at a time              |
| `_`                      | ignore the index                       |
| `for _, char := range s` | go through every character             |

---

## 🧠 Memory Tip

> 🔤 “PrintStr walks through every letter of the string,
> and hands it to z01.PrintRune like giving letters to a printer.” 🖨️

---
## StrLen
Excellent 🔢 — this one teaches you **how to count characters (runes)** in a string manually.
It’s very simple, but it’s a key building block for string manipulation in Go.

---

## ✅ Solution Code

```go
package piscine

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
```
**make I keep mute here!**
---

## 🧠 Step-by-Step Explanation

### 1️⃣ What’s happening here?

You’re given a string — for example:

```go
"Hello World!"
```

You need to **count how many runes (characters)** it has.
So you’re basically counting:
`H` (1), `e` (2), `l` (3), `l` (4), `o` (5), … until the last character.

---

### 2️⃣ Why use `for range`?

When you write:

```go
for range s
```

You’re saying:

> “Go through every character in the string — one at a time.”

And for each one, just increase `count` by 1:

```go
count++
```

When it’s done, `count` holds the **total number of characters**.

---

### 3️⃣ Why not just use `len(s)`?

Good question!
✅ Because `len(s)` counts **bytes**, not **runes**.

So if your string contains special characters like `é`, `ñ`, or emojis 😊
`len(s)` gives the wrong answer (since some runes use multiple bytes).
But `for range` counts each *rune* properly.

---

### 4️⃣ What’s a “rune”?

A **rune** in Go = one Unicode character.
This means even emojis or accented letters count as 1.

---

## 🧪 Test Program

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	l := piscine.StrLen("Hello World!")
	fmt.Println(l)
}
```

### 🧾 Output

```
$ go run .
12
```

---

## 🧩 Quick Recap

| Line           | Meaning                    |
| -------------- | -------------------------- |
| `count := 0`   | start from zero            |
| `for range s`  | go through each character  |
| `count++`      | add 1 for each rune        |
| `return count` | give back the total number |

---

## 🧠 Memory Trick

> “Walk through the string, count every letter you see,
> and when you’re done, return the number to me.” 🧮✨

---

# swap.go - It’s all about **swapping values between two memory boxes** 🧠📦

---

## ✅ Solution Code

```go
package piscine

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
```

---

## 🧒 Step-by-Step Explanation

Imagine you have two boxes:

| Box | Value |
| --- | ----- |
| `a` | 0     |
| `b` | 1     |

Each variable holds a **pointer** — meaning a *reference* to the box, not the box itself.

So:

* `*a` → open box A (get the number inside)
* `*b` → open box B (get the number inside)

---

### 🧠 Step-by-step logic

1️⃣ Store what’s inside `a`

```go
temp := *a   // temp = 0
```

2️⃣ Put what’s inside `b` into `a`

```go
*a = *b      // a = 1
```

3️⃣ Put what’s in temp (old `a`) into `b`

```go
*b = temp    // b = 0
```

Now they’ve swapped!

| Box | New Value |
| --- | --------- |
| `a` | 1         |
| `b` | 0         |

---

## 🧪 Test Program

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 0
	b := 1
	piscine.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
```

### 🧾 Output

```
$ go run .
1
0
```

---

## 🧩 Key Points

| Symbol | Meaning                       |
| ------ | ----------------------------- |
| `*a`   | value stored inside pointer a |
| `*b`   | value stored inside pointer b |
| `&a`   | address of a (pointer to a)   |
| `temp` | temporary storage for swap    |

---

## 💡 Simple Memory Tip

> “Use a temp box to hold one value while you switch the two.” 🔄

Like this in real life:

> You have two cups — one has coffee ☕, one has juice 🧃.
> You need a third empty cup to pour coffee into first,
> so you can swap them without spilling!

---

This one — **StrRev** — teaches you how to *reverse* a string in Go!
That means taking `"Hello World!"` and flipping it backwards to `"!dlroW olleH"`.

---

## ✅ Solution Code

```go
package piscine

func StrRev(s string) string {
	runes := []rune(s)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}
```
**format pls**
---

## 🧒 Step-by-Step Explanation

### 1️⃣ What’s happening?

We’re taking a string, and flipping it so the last character becomes the first.

For example:

```
Before: H e l l o
After:  o l l e H
```

---

### 2️⃣ Why convert to `[]rune`?

In Go, a **string** is a list of **bytes**, not characters.
Some letters (like `é` or emojis 😊) take *multiple bytes*,
so we convert it to a **rune slice** first to avoid breaking them:

```go
runes := []rune(s)
```

Now `runes` is a proper list of characters that we can swap easily.

---

### 3️⃣ The reversing trick

This part is pure magic 🔁

```go
for i := 0; i < length/2; i++ {
    runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
}
```

We:

* start from the **beginning** (`i`)
* swap it with the **end** (`length-1-i`)
* keep doing that until we reach the middle

It’s like turning the word around in a mirror!

---

### 4️⃣ Convert back to string

After swapping, we turn the rune slice back to a string:

```go
return string(runes)
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
	s := "Hello World!"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
```

### 🧾 Output

```
$ go run .
!dlroW olleH
```

---

## 🧩 Key Concepts Recap

| Concept         | Explanation                               |
| --------------- | ----------------------------------------- |
| `[]rune(s)`     | converts string → list of characters      |
| `len(runes)`    | gives the total number of runes           |
| swapping        | `runes[i], runes[j] = runes[j], runes[i]` |
| `string(runes)` | converts runes back → string              |

---

## 🧠 Memory Trick

> “To reverse a string — walk from both ends to the middle, swapping as you go.” 🔄

Like two people exchanging seats from the edges toward the center of a row 🎭

---

## Basicatoi— this is a very important one: it teaches you how to **convert a string of digits into an integer manually**, just like Go’s `strconv.Atoi()` does internally. 🔢

Let’s break it down gently 👇

---

## ✅ Solution Code

```go
package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, r := range s {
		result = result*10 + int(r-'0')
	}
	return result
}
```

---

## 🧒 Step-by-Step Explanation

### 1️⃣ What is “Atoi”?

`Atoi` means **ASCII to Integer**.

So, `"1234"` (string) → `1234` (int)

Your job: take every digit (as a character) and build up the actual number.

---

### 2️⃣ How do we get digits from characters?

Each character `'0'` to `'9'` has an ASCII code (like 48 to 57).
To turn `'5'` into the actual number 5, we subtract `'0'`:

```go
int('5' - '0') // gives 5
```

That’s the secret conversion.

---

### 3️⃣ Building the full number

We start from 0 and go through each character:

Example: `"123"`

| Step  | Current digit | Calculation | Result |
| ----- | ------------- | ----------- | ------ |
| Start | —             | —           | 0      |
| 1     | `'1'`         | 0×10 + 1    | 1      |
| 2     | `'2'`         | 1×10 + 2    | 12     |
| 3     | `'3'`         | 12×10 + 3   | 123    |

So we just keep multiplying by 10, then adding the next digit.

---

### 4️⃣ The loop in action

```go
for _, r := range s {
    result = result*10 + int(r-'0')
}
```

* `range s` → loop through every character in the string
* `r-'0'` → convert rune (like `'5'`) to number
* multiply by 10 each time (to “shift left” one place)

---

### 5️⃣ Return the result

Finally:

```go
return result
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
	fmt.Println(piscine.BasicAtoi("12345"))
	fmt.Println(piscine.BasicAtoi("0000000012345"))
	fmt.Println(piscine.BasicAtoi("000000"))
}
```

### 🧾 Output

```
$ go run .
12345
12345
0
```

---

## 🧩 Key Concepts

| Concept        | Meaning                                      |
| -------------- | -------------------------------------------- |
| `'0'`          | base ASCII for digits                        |
| `r-'0'`        | converts a rune to its actual number         |
| `result*10`    | shifts existing number one place to the left |
| `+ int(r-'0')` | adds the new digit to the end                |

---

## 🧠 Memory Trick

> “Start at zero, multiply by ten, and add each new digit again and again.” 🧮✨

---

Basicatoi2 is a **smarter version of `BasicAtoi`** 👨🏽‍🏫

It’s the same idea — converting `"12345"` into the integer `12345` —
but **this time**, you must check for *invalid characters* too (like spaces or letters).

If the string contains anything that’s not a digit (`0–9`), you must return **0**.

---

## ✅ Solution Code

```go
package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		result = result*10 + int(r-'0')
	}
	return result
}
```

---

## 🧒 Step-by-Step Explanation

### 1️⃣ Loop through each character

We check one by one every character (rune) in the string.

```go
for _, r := range s {
```

### 2️⃣ Verify that the rune is a digit

If the character isn’t between `'0'` and `'9'`,
that means it’s not a number — return **0** immediately.

```go
if r < '0' || r > '9' {
    return 0
}
```

This line filters out:

* letters like `'A'` or `'h'`
* spaces `' '`
* punctuation `'.'`, `'!'`, etc.

---

### 3️⃣ Convert each valid digit

We use the same logic from `BasicAtoi`:

```go
result = result*10 + int(r-'0')
```

So `"123"` becomes:

* Step 1: 0×10 + 1 = 1
* Step 2: 1×10 + 2 = 12
* Step 3: 12×10 + 3 = 123 ✅

---

### 4️⃣ Return the result

When the loop ends, return your final number:

```go
return result
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
	fmt.Println(piscine.BasicAtoi2("12345"))
	fmt.Println(piscine.BasicAtoi2("0000000012345"))
	fmt.Println(piscine.BasicAtoi2("012 345"))
	fmt.Println(piscine.BasicAtoi2("Hello World!"))
}
```

### 🧾 Output

```
$ go run .
12345
12345
0
0
```

---

## 🧩 Key Takeaways

| Concept                    | Description                    |          |                              |
| -------------------------- | ------------------------------ | -------- | ---------------------------- |
| `r < '0'                   |                                | r > '9'` | ensures character is a digit |
| `r-'0'`                    | converts rune digit to integer |          |                              |
| `result = result*10 + ...` | builds the number gradually    |          |                              |
| `return 0`                 | rejects invalid strings        |          |                              |

---

## 🧠 Memory Trick

> “If you find even one bad apple (non-digit), throw away the whole basket (return 0).” 🍎

---

`atoi.go` — this one’s the **final boss** of the Atoi family 💪

You now have to handle **signs (+ / -)** *and* **invalid characters**.
Let’s break it down cleanly 👇

---

## ✅ Solution Code

```go
package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	sign := 1
	result := 0

	for i, r := range s {
		// Handle signs at the beginning
		if i == 0 && (r == '+' || r == '-') {
			if r == '-' {
				sign = -1
			}
			continue
		}

		// Return 0 if not a digit
		if r < '0' || r > '9' {
			return 0
		}

		result = result*10 + int(r-'0')
	}

	return result * sign
}
```

---

## 🧒 Step-by-Step Explanation

### 1️⃣ Check if string is empty

If the string is empty (`""`), it’s invalid — just return `0`.

---

### 2️⃣ Handle `+` and `-` signs

```go
if i == 0 && (r == '+' || r == '-') {
	if r == '-' {
		sign = -1
	}
	continue
}
```

✅ This means:

* if the **first character** is `+`, keep `sign = 1`
* if it’s `-`, make `sign = -1`
* skip over the sign (don’t treat it as a digit)

---

### 3️⃣ Validate digits only

If any character is not between `'0'` and `'9'`, it’s invalid:

```go
if r < '0' || r > '9' {
	return 0
}
```

So `"12a3"` or `"01 23"` → `0`

---

### 4️⃣ Build the number

The magic line again ✨

```go
result = result*10 + int(r-'0')
```

---

### 5️⃣ Apply the sign at the end

After the loop, multiply by the sign:

```go
return result * sign
```

So `"-1234"` becomes `-1234`.

---

## 🧪 Test Program

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Atoi("12345"))
	fmt.Println(piscine.Atoi("0000000012345"))
	fmt.Println(piscine.Atoi("012 345"))
	fmt.Println(piscine.Atoi("Hello World!"))
	fmt.Println(piscine.Atoi("+1234"))
	fmt.Println(piscine.Atoi("-1234"))
	fmt.Println(piscine.Atoi("++1234"))
	fmt.Println(piscine.Atoi("--1234"))
}
```

---

### 🧾 Output

```
$ go run .
12345
12345
0
0
1234
-1234
0
0
```

---

## 🧩 Key Concepts

| Concept     | Description                                       |          |                            |
| ----------- | ------------------------------------------------- | -------- | -------------------------- |
| `sign`      | tracks whether the number is positive or negative |          |                            |
| `r-'0'`     | converts rune to numeric value                    |          |                            |
| `r < '0'    |                                                   | r > '9'` | detects invalid characters |
| `continue`  | skips over `+` or `-` sign                        |          |                            |
| `result*10` | shifts number left for next digit                 |          |                            |

---

## 🧠 Memory Trick

> “Read the sign, read the digits, reject the junk.” 🚫
> (And always multiply by the sign at the end!)

---
Final Quest -3 excercise 
 **first sorting algorithm** in Go! 🔢
Let’s make it crystal clear and easy to grasp.

---

## ✅ Solution Code

```go
package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
```

---

## 🧒 Step-by-Step Explanation

### 1️⃣ You are given a slice of integers

Example:

```go
s := []int{5, 4, 3, 2, 1, 0}
```

Goal → reorder it so it becomes:

```go
[0 1 2 3 4 5]
```

---

### 2️⃣ We use two loops

These help compare each number with all the ones after it.

```go
for i := 0; i < len(table)-1; i++ {
	for j := i + 1; j < len(table); j++ {
```

This means:

* `i` picks a number (say the first one)
* `j` looks at every number after `i`
* we compare and swap if needed

---

### 3️⃣ Compare and swap

If the current number is bigger than the next one, we switch them.

```go
if table[i] > table[j] {
	table[i], table[j] = table[j], table[i]
}
```

That’s all Go needs — no temporary variable necessary!
Go allows tuple assignment like that in one line.

---

### 4️⃣ Repeat until sorted

After all comparisons are done, your slice becomes sorted in ascending order.

---

## 🧪 Test Program

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}
```

---

### 🧾 Output

```
$ go run .
[0 1 2 3 4 5]
```

---

## 🧩 Key Concepts

| Concept                                   | Description                   |
| ----------------------------------------- | ----------------------------- |
| `len(table)`                              | gives the number of elements  |
| nested loops                              | compare every pair of numbers |
| `table[i], table[j] = table[j], table[i]` | swap values in Go             |
| in-place sort                             | modifies the slice directly   |

---

## 🧠 Memory Trick

> “Look at every pair — if they’re out of order, swap and keep going.” 🔁

That’s **bubble sort / selection sort logic**, and it’s all you need for this task.

---
Keep refreshing!!! solutions not yet uploaded!
