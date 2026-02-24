# Go Reloaded - Complete Teaching Guide

## 📚 Overview

**Go Reloaded** is a text processing tool that transforms input text based on special markers. It's the first graded project after Go-Piscine, meaning it's corrected by other students (auditors).

### Learning Objectives

By the end of this project, you will master:
- Go's **file system (fs)** package
- **String manipulation** in Go
- **Regular expressions** (regex)
- Working with **runes** and Unicode
- Building a **text transformation pipeline**

---

## 🏗️ Project Structure

```
go-reloaded/
├── main.go       # Your solution
├── sample.txt    # Test input
└── README.md    # Project instructions
```

---

## 🔧 Key Concepts

### 1. Reading & Writing Files

```go
import "os"

// Read file
content, err := os.ReadFile("input.txt")

// Write file
err = os.WriteFile("output.txt", []byte(result), 0644)
```

**Why this matters:** Go's `os` package provides low-level file operations. The `fs` package (io/fs) is for more advanced filesystem operations.

---

### 2. String Handling

```go
import "strings"

// Split text into words (preserves whitespace info separately)
words := strings.Fields(text)

// Join words back
result := strings.Join(words, " ")

// String transformations
strings.ToUpper(word)
strings.ToLower(word)
strings.Replace(text, old, new, n)

// Check substring
strings.Contains(word, "(hex)")
```

---

### 3. Runes & Unicode

```go
import "unicode"

// Capitalize first letter (NOT strings.ToUpper!)
func capitalize(word string) string {
    if len(word) == 0 {
        return word
    }
    runes := []rune(word)
    runes[0] = unicode.ToUpper(runes[0])
    return string(runes)
}
```

**Why runes?** Go strings are byte arrays. For Unicode (especially non-ASCII like Nigerian characters), you need `rune` (int32) to handle multi-byte characters properly.

---

### 4. Regular Expressions

```go
import "regexp"

// Match pattern with capture group
upMatch := regexp.MustCompile(`\(up,\s*(\d+)\)`).FindStringSubmatch(word)
if len(upMatch) > 0 {
    n, _ := strconv.Atoi(upMatch[1])  // Extract captured number
}

// Replace pattern
cleanWord := regexp.MustCompile(`\(up,\s*\d+\)`).ReplaceAllString(word, "")
```

**Patterns used:**
| Pattern | Meaning |
|---------|---------|
| `\(up,\s*(\d+)\)` | Match `(up, 2)` - captures the number |
| `\s+` | One or more whitespace |
| `[^a-zA-Z]` | Any non-letter character |

---

## 🔄 The Transformation Pipeline

The project processes text in **4 stages**:

```
Input Text
    ↓
1. Split into words (strings.Fields)
    ↓
2. Process markers (hex, bin, up, low, cap)
    ↓
3. Convert "a" → "an"
    ↓
4. Join and fix punctuation & quotes
    ↓
Output Text
```

---

## 📝 Detailed Marker Logic

### (hex) - Hexadecimal to Decimal

```go
if strings.Contains(word, "(hex)") {
    // Convert previous word from hex to decimal
    prevWord := result[len(result)-1]
    if hexNum, err := strconv.ParseInt(prevWord, 16, 64); err == nil {
        result[len(result)-1] = fmt.Sprintf("%d", hexNum)
    }
}
```

**Example:** `1E (hex)` → `30`

**How it works:**
- `strconv.ParseInt(word, 16, 64)` parses base-16 (hex)
- Base 16 uses digits 0-9 and letters A-F/a-f

---

### (bin) - Binary to Decimal

```go
if strings.Contains(word, "(bin)") {
    prevWord := result[len(result)-1]
    if binNum, err := strconv.ParseInt(prevWord, 2, 64); err == nil {
        result[len(result)-1] = fmt.Sprintf("%d", binNum)
    }
}
```

**Example:** `10 (bin)` → `2`

---

### (up), (low), (cap) - Case Transformations

```go
// Single word
strings.ToUpper(prevWord)
strings.ToLower(prevWord)
capitalize(prevWord)

// Multiple words (with number)
start := len(result) - n
if start < 0 { start = 0 }
for j := start; j < len(result); j++ {
    result[j] = strings.ToUpper(result[j])
}
```

**Examples:**
- `go (up)` → `GO`
- `STOP (low)` → `stop`
- `bridge (cap)` → `Bridge`
- `(up, 2)` → Uppercase 2 previous words

---

### A → An Rule

```go
func processAtoAn(words []string) []string {
    for i := 0; i < len(words)-1; i++ {
        if strings.ToLower(words[i]) == "a" {
            firstChar := strings.ToLower(string(words[i+1][0]))
            vowels := "aeiouh"
            if strings.Contains(vowels, firstChar) {
                words[i] = "an"
            }
        }
    }
    return words
}
```

**Logic:** "a" becomes "an" before words starting with:
- Vowels: a, e, i, o, u
- 'h' (silent h): "hour", "honest"

---

### Punctuation Handling

```go
// Remove space before punctuation
punct := []string{".", ",", "!", "?", ":", ";"}
for _, p := range punct {
    text = strings.Replace(text, " "+p, p, -1)
}
```

**Example:** `hello , world !` → `hello, world!`

---

### Smart Quotes

```go
re := regexp.MustCompile(`'([^']+)'`)
text = re.ReplaceAllStringFunc(text, func(match string) string {
    content := match[1 : len(match)-1]
    content = strings.TrimSpace(content)
    return "'" + content + "'"
})
```

**Example:** `' awesome '` → `'awesome'`

---

## 🧪 Testing Your Code

### Test Cases

```bash
# Test 1: Case conversion
echo "it (cap) was the best of times" > test.txt
go run . test.txt out.txt
# Expected: It was the best of times

# Test 2: Hex conversion
echo "1E (hex)" > test.txt
go run . test.txt out.txt
# Expected: 30

# Test 3: Binary conversion
echo "10 (bin)" > test.txt
go run . test.txt out.txt
# Expected: 2

# Test 4: a -> an
echo "a amazing rock" > test.txt
go run . test.txt out.txt
# Expected: an amazing rock
```

---

## ⚠️ Common Pitfalls

### 1. Using string indexing on Unicode
```go
// ❌ WRONG - breaks on Nigerian names like "Chidinma"
name := "Chidinma"
first := name[0]  // Returns 67 (byte), not 'C'

// ✅ CORRECT
runes := []rune(name)
first := runes[0]  // Returns 'C' (rune)
```

### 2. Forgetting to handle edge cases
```go
// ❌ Wrong - crashes on empty result
result[len(result)-1]

// ✅ Correct - check length first
if len(result) > 0 {
    // safe to access
}
```

### 3. Not trimming whitespace in quotes
```go
// ❌ Wrong
return "'" + content + "'"

// ✅ Correct
content = strings.TrimSpace(content)
return "'" + content + "'"
```

---

## 📖 Additional Resources

- [Go strings package](https://pkg.go.dev/strings)
- [Go strconv package](https://pkg.go.dev/strconv)
- [Go regexp package](https://pkg.go.dev/regexp)
- [Unicode in Go](https://go.dev/blog/strings)

---

## 🎯 Quick Reference

| Function | Purpose |
|----------|---------|
| `strings.Fields()` | Split string by whitespace |
| `strings.Join()` | Join strings with separator |
| `strings.Contains()` | Check if substring exists |
| `strings.Replace()` | Replace substrings |
| `strconv.ParseInt(str, base, bits)` | Parse integer |
| `regexp.MustCompile()` | Create regex pattern |
| `unicode.ToUpper()` | Handle Unicode case conversion |

---

*Created: February 2026*
*For: 01-Edu Go-Piscine → Go Reloaded*
