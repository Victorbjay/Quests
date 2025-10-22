# 🧠 Piscine-Go Shell & Command Line Quest Solutions

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
🕵️‍♂️ the-final-cl-test (Murder Mystery)

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

From mystery/interviews/interview-699607:

“Interviewed Ms. Church… Describes a blue Honda with a plate starting with ‘L337’ and ending with ‘9’.”

Witness: Annabel Church

Conclusion:
The murderer is Dartey Henv — tall male, blue Honda, L337…9, wallet matches clues.

`my_answer.sh`

```bash
#!/bin/bash

echo "Dartey Henv"
```
---
## 🧾 `explain.sh`

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
rm -f '"'  # Remove the test file created earlier
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

# Mac
brew install jq
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
