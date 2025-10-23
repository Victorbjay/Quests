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
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			z01.PrintRune(rune('0' + i/10))
			z01.PrintRune(rune('0' + i%10))
			z01.PrintRune(' ')
			z01.PrintRune(rune('0' + j/10))
			z01.PrintRune(rune('0' + j%10))
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
---

Keep refreshing!!! solutions not yet uploaded!

Now check README file in 📁 Resource for beginner-friendly guides. 
