# Beginner's Guide to Programming Exercises

> A comprehensive guide to help you understand and complete programming exercises involving Unix file permissions and Go programming.

## 📚 Table of Contents

1. [Introduction](#introduction)
2. [Exercise 1: Unix File Permissions](#exercise-1-unix-file-permissions)
3. [Exercise 2: Go Programming - Print Alphabet](#exercise-2-go-programming---print-alphabet)
4. [Additional Resources](#additional-resources)

---

## Introduction

This guide covers two fundamental computer science concepts:
- **Unix File Permissions**: Understanding how operating systems control access to files
- **Go Programming**: Learning basic programming with the Go language

Don't worry if you're new to programming - we'll explain everything step by step!

---

## Exercise 1: Unix File Permissions

### 🎯 What You'll Learn

- How Unix/Linux manages file access
- Reading and writing file permissions
- Converting between symbolic and octal notation
- Using terminal commands to manipulate files

### 📖 Understanding File Permissions

#### The Basics

Every file and directory in Unix/Linux has **permissions** that control:
- **Read (r)**: View the file contents
- **Write (w)**: Modify the file
- **Execute (x)**: Run the file as a program

These permissions apply to three groups:
1. **Owner (u)**: The person who created the file
2. **Group (g)**: Users in the same group
3. **Others (o)**: Everyone else

#### Reading Permission Strings

When you run `ls -l`, you see something like:
```
-rwxr-xr--
```

Let's break it down:
```
-  rwx  r-x  r--
│   │    │    │
│   │    │    └─── Others: read only
│   │    └──────── Group: read + execute
│   └───────────── Owner: read + write + execute
└───────────────── File type (- = file, d = directory)
```

#### Converting to Octal (Numbers)

Permissions can be represented as numbers:
- **Read (r) = 4**
- **Write (w) = 2**
- **Execute (x) = 1**
- **Nothing (-) = 0**

**How to calculate:**
Add up the values for each group!

**Example 1:** `rwx` = 4 + 2 + 1 = **7**
**Example 2:** `r-x` = 4 + 0 + 1 = **5**
**Example 3:** `r--` = 4 + 0 + 0 = **4**

So `rwxr-xr--` becomes **754**:
- Owner: rwx = 7
- Group: r-x = 5
- Others: r-- = 4

#### Common Permission Patterns

| Octal | Symbolic | Meaning |
|-------|----------|---------|
| 777 | rwxrwxrwx | Everyone can do everything (dangerous!) |
| 755 | rwxr-xr-x | Owner can edit, others can read/run |
| 644 | rw-r--r-- | Owner can edit, others can only read |
| 600 | rw------- | Only owner can read/write |
| 700 | rwx------ | Only owner has all permissions |

### 🛠️ Essential Commands

#### Viewing Permissions
```bash
ls -l filename          # See permissions for a file
ls -ld directory        # See permissions for a directory
stat filename           # Detailed file information
```

#### Changing Permissions
```bash
chmod 755 filename      # Set permissions using octal
chmod u+x filename      # Add execute for owner
chmod go-w filename     # Remove write for group and others
```

#### Creating Files and Directories
```bash
touch filename          # Create empty file
mkdir dirname           # Create directory
```

#### Setting Timestamps
```bash
touch -t 202501011200 file   # Set time to Jan 1, 2025, 12:00
# Format: YYYYMMDDHHmm
```

### 📝 Exercise Solution Walkthrough

**Goal:** Create files with specific permissions and timestamps, then package them into a tar archive.

**Step-by-step:**

1. **Create directory `0`:**
   ```bash
   mkdir 0
   chmod 501 0                    # dr-x-----x
   touch -t 198601050000 0        # Jan 5, 1986, 00:00
   ```

2. **Create file `1`:**
   ```bash
   touch 1
   chmod 402 1                    # -r------w-
   touch -t 198611130001 1        # Nov 13, 1986, 00:01
   ```

3. **Create symbolic link `3` → `0`:**
   ```bash
   ln -s 0 3                      # Create link
   touch -h -t 199002160011 3     # Feb 16, 1990, 00:11
   ```

4. **Create tar archive:**
   ```bash
   tar -cf done.tar 0 1 2 3 4 5 6 7 8 9 A
   ```

5. **Verify:**
   ```bash
   TZ=utc ls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'
   ```

### 🔍 Quick Reference Table

| File | Permissions | Octal | Calculation |
|------|-------------|-------|-------------|
| 0 | dr-x-----x | 501 | 5(r-x) + 0(---) + 1(--x) |
| 1 | -r------w- | 402 | 4(r--) + 0(---) + 2(-w-) |
| 2 | -rw----r-- | 604 | 6(rw-) + 0(---) + 4(r--) |
| 4 | -r-x--x--x | 511 | 5(r-x) + 1(--x) + 1(--x) |
| 5 | -r--rw---- | 460 | 4(r--) + 6(rw-) + 0(---) |

---

## Exercise 2: Go Programming - Print Alphabet

### 🎯 What You'll Learn

- Basic Go syntax
- Using loops in Go
- Working with the z01 package
- Understanding runes (characters)

### 📖 Go Programming Basics

#### What is Go?

Go (or Golang) is a programming language created by Google. It's:
- **Fast**: Compiles quickly and runs efficiently
- **Simple**: Clean syntax that's easy to read
- **Modern**: Built for today's computing needs

Companies like Google, Uber, Netflix, and Docker use Go!

#### Basic Go Program Structure

```go
package main                    // Every Go program starts here

import "fmt"                    // Import packages we need

func main() {                   // Main function - program starts here
    fmt.Println("Hello World")  // Print something
}
```

#### Understanding Runes

In Go, a **rune** is a single character. Examples:
- `'a'` is a rune
- `'Z'` is a rune
- `'1'` is a rune
- `'\n'` is a newline rune

Runes are written in **single quotes** (`'a'`), not double quotes!

#### For Loops in Go

The basic loop structure:
```go
for i := 0; i < 10; i++ {
    // Do something
}
```

Breaking it down:
- `i := 0` - Start with i = 0
- `i < 10` - Keep going while i is less than 10
- `i++` - Add 1 to i each time

**Looping through letters:**
```go
for char := 'a'; char <= 'z'; char++ {
    // char will be 'a', then 'b', then 'c'... up to 'z'
}
```

### 🛠️ The z01.PrintRune Function

The `z01.PrintRune` function prints a single character:
```go
z01.PrintRune('a')    // Prints: a
z01.PrintRune('B')    // Prints: B
z01.PrintRune('\n')   // Prints: newline (moves to next line)
```

**Why use z01.PrintRune instead of fmt.Println?**
- It's part of the 01-edu curriculum
- Teaches you to work with low-level output
- You learn to build strings character by character

### 📝 Exercise Solution Walkthrough

**Goal:** Print the lowercase alphabet on one line.

**Complete solution:**
```go
package main

import "github.com/01-edu/z01"

func main() {
    // Loop from 'a' to 'z'
    for char := 'a'; char <= 'z'; char++ {
        z01.PrintRune(char)
    }
    // Print newline at the end
    z01.PrintRune('\n')
}
```

**How it works:**
1. Start with `char = 'a'`
2. Print `'a'` using `z01.PrintRune(char)`
3. Move to next letter (`char++` makes it `'b'`)
4. Repeat until we reach `'z'`
5. Print a newline at the end

**Output:**
```
abcdefghijklmnopqrstuvwxyz
```

### 🚀 Setting Up Your Go Project

#### 1. Install Go
```bash
# Check if Go is installed
go version

# If not installed, visit: https://go.dev/dl/
```

#### 2. Create Project Directory
```bash
mkdir printalphabet
cd printalphabet
```

#### 3. Initialize Go Module
```bash
go mod init printalphabet
```

#### 4. Get the z01 Package
```bash
go get github.com/01-edu/z01
```

#### 5. Create main.go
```bash
nano main.go
# Or use any text editor
```

#### 6. Run Your Code
```bash
go run .
```

#### 7. Format Your Code (Important!)
```bash
# Install gofumpt
go install mvdan.cc/gofumpt@latest

# Format your code
gofumpt -w .
```

### 🎨 Go Formatting Rules

Go is very strict about formatting! Common issues:

**❌ Wrong:**
```go
func main()  {        // Extra space before {
    fmt.Println("hi")
}                     // No newline at end of file
```

**✅ Correct:**
```go
func main() {         // No extra space
    fmt.Println("hi")
}
                      // Newline at end
```

**Pro tip:** Always run `gofumpt -w .` before submitting!

---

## Additional Resources

### 📚 Learning Resources

#### Unix/Linux File Permissions
- [Linux File Permissions Explained](https://www.redhat.com/en/blog/linux-file-permissions-explained) - Comprehensive guide with examples
- [chmod Command Tutorial](https://linuxcommand.org/lc3_lts0090.php) - Understanding chmod with visual examples
- [Octal Permissions Calculator](https://chmod-calculator.com/) - Interactive tool to practice

#### Go Programming
- [Official Go Tutorial](https://go.dev/doc/tutorial/getting-started) - Start here for official docs
- [Go by Example](https://gobyexample.com/) - Hands-on Go examples
- [W3Schools Go Tutorial](https://www.w3schools.com/go/) - Beginner-friendly with interactive examples
- [GeeksforGeeks Go Tutorial](https://www.geeksforgeeks.org/golang-tutorial-learn-go-programming-language/) - Comprehensive Go guide

### 🛠️ Essential Tools

#### Terminal/Command Line
- **Windows**: Git Bash, PowerShell, or WSL (Windows Subsystem for Linux)
- **Mac**: Terminal (built-in)
- **Linux**: Terminal (Ctrl+Alt+T)

#### Text Editors
- **VS Code** (Free) - Best for Go development
- **GoLand** (Paid) - Professional Go IDE
- **nano** - Simple terminal-based editor
- **vim** - Powerful terminal editor

#### Go Tools
```bash
go run .          # Run your program
go build          # Compile your program
go fmt            # Format your code
gofumpt -w .      # Format with strict rules
go mod init       # Initialize a module
go get            # Download packages
```

### 💡 Tips for Success

1. **Practice Daily**: Even 30 minutes helps
2. **Read Error Messages**: They tell you what's wrong
3. **Google is Your Friend**: Search "[your error] golang" or "[your question] linux"
4. **Use Version Control**: Learn Git early
5. **Ask for Help**: Communities are friendly!
6. **Read Others' Code**: GitHub is full of examples
7. **Build Projects**: Theory + Practice = Mastery

### 🐛 Common Mistakes to Avoid

#### Unix Permissions
- ❌ Forgetting that `chmod 777` is dangerous
- ❌ Not using `sudo` when needed
- ❌ Confusing symbolic (`rwx`) with octal (`7`)

#### Go Programming
- ❌ Using double quotes for characters: `"a"` instead of `'a'`
- ❌ Forgetting the newline: `\n`
- ❌ Not formatting code with `gofumpt`
- ❌ Importing packages you don't use
- ❌ Using `fmt.Println` when you should use `z01.PrintRune`

### 🎓 Next Steps

After mastering these basics:

1. **Unix/Linux**:
   - Learn file manipulation (`cp`, `mv`, `rm`)
   - Understand processes and process management
   - Explore shell scripting with bash

2. **Go Programming**:
   - Functions and parameters
   - Slices and arrays
   - Structs and methods
   - Goroutines (concurrency)

3. **General Programming**:
   - Data structures (arrays, maps, structs)
   - Algorithms (sorting, searching)
   - Version control with Git
   - Testing and debugging

---

## 📞 Getting Help

### When You're Stuck

1. **Read the error message carefully**
2. **Google it**: "[error message] golang" or "[question] linux"
3. **Check documentation**: `man chmod` or go.dev
4. **Ask in forums**: 
   - Stack Overflow
   - Reddit (r/golang, r/linux)
   - Discord/Slack programming communities

### Useful Commands for Debugging

```bash
# Check what directory you're in
pwd

# List all files (including hidden)
ls -la

# Check Go version
go version

# See environment variables
go env

# Get help on a command
man chmod
man ls
go help
```

---

## 🎉 You've Got This!

Remember: **Every expert was once a beginner**. Take it one step at a time, practice regularly, and don't be afraid to make mistakes - that's how we learn!

Happy coding! 🚀

---

*Last updated: October 2025*
*Created for: 01-edu programming exercises*
