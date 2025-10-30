### Task 1-Understanding the Problem

The task is to create a function called `AppendRange` that takes two numbers (integers): a starting number called `min` and an ending number called `max`. The function should give back a list (in Go, we call this a "slice") of all the whole numbers starting from `min` up to but not including `max`.

Imagine you're counting apples in a basket:
- If min is 5 and max is 10, you count: 5, 6, 7, 8, 9 (that's 5 apples, stopping before 10).
- If min is 10 and max is 5, there's nothing to count because 10 is bigger than 5, so you get an empty basket (which in code is like "nil" or nothing).

Important rules:
- Include the min number.
- Exclude the max number (don't add it).
- If min is equal to or bigger than max, return nothing (nil slice).
- You can't use a special command called "make" to create the list. Instead, you have to build it by adding items one by one using "append" (like adding toys to a toy box one at a time).

This is like building a chain of numbers step by step, without cheating by making the whole chain at once.

### The Solution Code

Here's the complete function in Go. I'll explain it line by line right after.

```go
func AppendRange(min, max int) []int {
	var result []int
	if min >= max {
		return nil
	}
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}
```

This code works perfectly for the examples:
- For `AppendRange(5, 10)`, it returns `[5 6 7 8 9]`.
- For `AppendRange(10, 5)`, it returns `[]` (which looks empty, but is actually nil).

### Line-by-Line Explanation

I'll break it down super simply, step by step. Think of this like building a Lego tower: each line adds a brick. I'll use easy examples and "pictures" (described in words, like a story) to help you see it in your mind.

1. **func AppendRange(min, max int) []int {**
   - This is the "name tag" for our function. It's like saying, "Hey, I'm a recipe called AppendRange. I need two ingredients: min and max (both are numbers like 5 or 10). When I'm done, I'll give you back a box of numbers (a slice of ints)."
   - Illustration: Imagine a machine at a candy factory. You put in "min=5" and "max=10" as buttons you press. The machine spits out a bag of candies labeled [5,6,7,8,9]. The curly brace { is like opening the machine's door to see how it works inside.

2. **var result []int**
   - Here, we're creating an empty box called "result" that can hold numbers. Right now, it's totally empty (like a nil slice, which means "nothing yet").
   - Why empty? Because we have to build it ourselves using append, like starting with an empty backpack and adding snacks one by one.
   - Illustration: Picture an empty toy box. If min >= max, we just leave it empty and say "no toys today." If not, we'll fill it with toys numbered from min to max-1.

3. **if min >= max {**
   - This is a checkpoint: "Is the starting number bigger than or equal to the ending number?" Like checking if your start line in a race is after the finish line.
   - If yes (e.g., min=10, max=5), it's impossible to count forward, so we...
   - Illustration: Imagine walking from house 10 to house 5. But houses go up in numbers, so you'd be walking backwards or nowhere. We decide: "Nope, can't do that!"

4. **return nil**
   - ...give back nothing (nil). This closes the checkpoint and ends the function early.
   - In Go, returning nil for a slice means "empty or nothing." When you print it, it shows as [] (empty list).
   - Illustration: You hand back an empty bag and say, "Sorry, no candies this time." This matches the example where AppendRange(10,5) prints [].

5. **}**
   - This closes the checkpoint door. If we didn't hit the "return nil," we keep going to build the list.

6. **for i := min; i < max; i++ {**
   - This is a loop, like a repeating robot arm that grabs numbers one by one.
   - Break it down:
     - "i := min" means start with a counter i set to min (e.g., i=5).
     - "i < max" means keep going as long as i is less than max (e.g., stop when i reaches 10).
     - "i++" means add 1 to i each time (like counting up: 5, then 6, etc.).
   - Illustration: Think of a conveyor belt. It starts at apple #5. As long as the apple number is less than 10, grab it and move to the next (5→6→7→8→9). When it hits 10, stop! No apple #10.

7. **result = append(result, i)**
   - This adds the current number (i) to our box (result).
   - "append" is like gluing a new Lego brick to the tower. It takes the old result, adds i, and makes a new, bigger result.
   - Illustration: Start with empty box. First loop: add 5 (box now [5]). Second: add 6 (box [5,6]). Keep going until [5,6,7,8,9]. It's like stacking blocks: empty → [5] → [5,6] → and so on.

8. **}**
   - Closes the loop. No more adding after i reaches max.

9. **return result**
   - Finally, hand back the full box (result) with all the numbers inside.
   - Illustration: The candy machine beeps "Done!" and gives you the bag [5,6,7,8,9].

10. **}**
    - Closes the whole function. End of the recipe!

### Why This Works (With Visual Examples)

Let's "run" the code in our minds with pictures for two cases.

**Case 1: min=5, max=10 (Good case)**
- Empty box: []
- Check: 5 >= 10? No.
- Loop: i=5 (<10? Yes) → Add 5 → Box: [5]
- i=6 (<10? Yes) → Add 6 → Box: [5,6]
- i=7 → [5,6,7]
- i=8 → [5,6,7,8]
- i=9 → [5,6,7,8,9]
- i=10 (<10? No) → Stop loop.
- Return [5,6,7,8,9]
- Visual: Imagine a number line: 5-----6-----7-----8-----9  (stop before 10, like a fence at 10).

**Case 2: min=10, max=5 (Bad case)**
- Empty box: []
- Check: 10 >= 5? Yes.
- Return nil (empty [] when printed).
- Visual: Number line: 10 is way after 5, like trying to walk uphill backwards. Just give up and return empty hands.

**Another Example: min=0, max=3**
- Box: []
- Check: 0 >= 3? No.
- Loop: i=0 → [0]
- i=1 → [0,1]
- i=2 → [0,1,2]
- i=3 (<3? No) → Stop.
- Return [0,1,2]
- Visual: Counting fingers: Thumb=0, Index=1, Middle=2. Stop before Ring=3.

**Edge Case: min=5, max=5**
- Check: 5 >= 5? Yes.
- Return nil.
- Visual: Start and end at the same spot—no distance to walk, so empty.

### Documentation (Important Notes)

This function is efficient because:
- It only loops as many times as needed (e.g., 5 times for 5 to 10).
- No extra space wasted—append grows the slice smartly in Go.
- Handles negative numbers too: AppendRange(-2, 1) → [-2, -1, 0] (like counting up from below zero).

If the numbers are huge (like min=1, max=1000000), it might use a lot of memory, but that's okay for this exercise.

In Go, slices are like magic lists that can grow. Append is the spell to add to them without "make" (which is like pre-building a fixed-size list).

If you test it with the main function provided:
- It prints [5 6 7 8 9]
- Then []

Perfect match!

---
`makerange.go`

### Challenge: `MakeRange(min, max int) []int`

| Rule | What it means |
|------|----------------|
| `min` included | `min` goes into the slice |
| `max` excluded | Stop **before** `max` |
| `min >= max` | Return `nil` |
| **Cannot use `append`** | No `append()` allowed! |
| **Can use `make`** | We’re allowed to pre-create the slice |

---

## Final Solution 

```go
package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	// Step 1: Count how many numbers we need
	size := max - min // e.g., 10 - 5 = 5 numbers: 5,6,7,8,9

	// Step 2: Make a slice of exactly that size
	result := make([]int, size)

	// Step 3: Fill it one by one
	for i := 0; i < size; i++ {
		result[i] = min + i
	}

	return result
}
```

---

## Line-by-Line Explanation 

Imagine you're building a **train with exactly the right number of cars** — no more, no less.

---

### 1. `package piscine`
> This says: "This code belongs in the `piscine` folder."  
> Like labeling your toy box: **"Piscine’s Toys"**

---

### 2. `func MakeRange(min, max int) []int {`
> The function name is `MakeRange`.  
> It takes two numbers: `min` and `max`.  
> It returns a **train** (slice) of numbers.

---

### 3. `if min >= max { return nil }`
> **Safety check!**  
> If starting point is same or past the end → no train!  
> Example: `min=10`, `max=5` → 10 is after 5 → return nothing.

**Visual**:  
```
Start: 10    End: 5
       X <--- can't go backwards!
→ Return empty (nil)
```

---

### 4. `size := max - min`
> How many train cars do we need?

| min | max | size |
|-----|-----|------|
| 5   | 10  | 5    |
| 1   | 4   | 3    |
| -2  | 1   | 3    |

> So from 5 to 10 → **5 numbers**: 5,6,7,8,9

**Visual**:  
```
min=5     max=10
  ↓         ↓
[5][6][7][8][9]  ← 5 cars
```

---

### 5. `result := make([]int, size)`
> This is the **magic**!  
> `make` builds a slice with **exactly `size` empty slots**, all ready.

**Before**:
```
result = [ _ _ _ _ _ ]  ← 5 empty boxes
```

**Visual**:  
```
[ ] [ ] [ ] [ ] [ ]  ← train cars waiting to be filled
```

---

### 6. `for i := 0; i < size; i++ {`
> Loop through each car (0 to 4)

| i | What we do |
|---|-----------|
| 0 | Fill car 0 |
| 1 | Fill car 1 |
| ... | ... |
| 4 | Fill car 4 |

---

### 7. `result[i] = min + i`
> Put the **correct number** in each spot!

| i | min + i | Put in result[i] |
|---|--------|------------------|
| 0 | 5 + 0 = 5 | → result[0] = 5 |
| 1 | 5 + 1 = 6 | → result[1] = 6 |
| 2 | 5 + 2 = 7 | → result[2] = 7 |
| 3 | 5 + 3 = 8 | → result[3] = 8 |
| 4 | 5 + 4 = 9 | → result[4] = 9 |

**After loop**:
```
result = [5][6][7][8][9]
```

**Visual Train**:
```
[5] → [6] → [7] → [8] → [9]
```

---

### 8. `return result`
> Give back the fully loaded train!

---

## Test with Example

```go
fmt.Println(piscine.MakeRange(5, 10))  // [5 6 7 8 9]
fmt.Println(piscine.MakeRange(10, 5))  // []
```

**Output**:
```
[5 6 7 8 9]
[]
```

**Perfect!**

---

## Bonus: Negative Numbers?

```go
MakeRange(-2, 1)
→ size = 1 - (-2) = 3
→ [ -2, -1, 0 ]
```

**Visual**:
```
[-2] → [-1] → [0]
```

---

## Summary Table

| Input (min, max) | size | result |
|------------------|------|--------|
| 5, 10            | 5    | [5,6,7,8,9] |
| 10, 5            | —    | nil → [] |
| 0, 3             | 3    | [0,1,2] |
| -1, 2            | 3    | [-1,0,1] |

---

## Key Difference from `AppendRange`

| Feature         | `AppendRange`       | `MakeRange`         |
|----------------|---------------------|---------------------|
| Use `append`?  | YES                 | NO (not allowed)    |
| Use `make`?    | NO (not allowed)    | YES                 |
| How it grows   | Add one by one      | Pre-build full size |

**Visual**:
```
AppendRange:  [] → [5] → [5,6] → [5,6,7]...
MakeRange:    [ _ _ _ _ _ ] → fill → [5,6,7,8,9]
```

---

## Final File: `makerange.go`

```go
package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	size := max - min
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = min + i
	}
	return result
}
```

---

**You did it!**  
You now know **two ways** to make a range in Go:

- `append` → build slowly
- `make` → build all at once

Like choosing between **adding bricks one by one** or **building the whole wall first and painting numbers**.

---
# `ConcatParams` 

---

**Quest Name:** `ConcatParams`  
**Level:** 14  
**XP Reward:** 3.06 kB of Brain Power!  
**Goal:** Turn a list of words into one long message with **newlines** between them!  
Your mission? Write a spell (`ConcatParams`) that takes a **bag of strings** and returns **one big string** with each word on a new line.

---

## The Magic Spell (Function) You Must Write

```go
func ConcatParams(args []string) string {

}
```

> **Your job:** Fill in the `{ }` so it works like magic!

---

## Let’s Break It Down – Line by Line (Super Simple!)

| Line | Explanation | Fun Analogy |
|------|-----------|-------------|
| `func` | This tells Go: “Hey! I’m starting a function!” | Like saying: “Abracadabra! A new spell begins!” |
| `ConcatParams` | The **name** of your function. Call it like this: `ConcatParams(...)` | Your spell’s name: **“StringGlue!”** |
| `(args []string)` | This is the **input**: a **slice** (like a list) of strings. | Imagine a **bag of notes** – each note has a word. |
| `string` | This is what the function **returns** (gives back). | The final **magic scroll** with all words written down! |
| `{ }` | The **body** – where the magic happens! | The cauldron where you mix everything! |

---

## Example Input & Output (The Test!)

```go
test := []string{"Hello", "how", "are", "you?"}
fmt.Println(piscine.ConcatParams(test))
```

**Output:**
```
Hello
how
are
you?
```

> Each word is on its own line!  
> That’s because we used `\n` → the **newline spell**!

---

## Step-by-Step: How to Solve It (Like a Game!)

Let’s walk through it like a **level in a video game**!

---

### Level 1: Understand the Input

```go
args := []string{"Hello", "how", "are", "you?"}
```

| Index | Value |
|-------|-------|
| 0 | "Hello" |
| 1 | "how" |
| 2 | "are" |
| 3 | "you?" |

> This is a **slice** – like a train with 4 cars, each carrying a word.

---

### Level 2: We Need to Glue Words with `\n`

We want:
```
Hello\nhow\nare\nyou?
```

Which prints as:
```
Hello
how
are
you?
```

> `\n` = **"Press Enter"** – it makes a new line!

---

### Level 3: Use a Loop to Visit Each Word

We’ll use a `for` loop to open each "note" in the bag.

```go
result := ""
for i := 0; i < len(args); i++ {
    result = result + args[i] + "\n"
}
```

Let’s see it in action!

| Step | `i` | `args[i]` | `result` becomes |
|------|-----|-----------|------------------|
| 1 | 0 | "Hello" | `"" + "Hello" + "\n" = "Hello\n"` |
| 2 | 1 | "how" | `"Hello\n" + "how" + "\n" = "Hello\nhow\n"` |
| 3 | 2 | "are" | `"Hello\nhow\n" + "are" + "\n"` |
| 4 | 3 | "you?" | `"Hello\nhow\nare\n" + "you?" + "\n"` |

Final: `"Hello\nhow\nare\nyou?\n"`

> Uh-oh! Extra `\n` at the end?

---

### Level 4: Fix the Extra Newline! (Boss Fight!)

We don’t want a blank line at the end.

**Bad:**
```
Hello
how
are
you?

<-- empty line!
```

**Good:**
```
Hello
how
are
you?
```

---

### Pro Trick: Manual Loop (No Imports Allowed!)

Since `strings.Join` is **forbidden** (cheating alert! 🚨), we build it by hand!

---

## Final Working Code

```go
package piscine

func ConcatParams(args []string) string {
    result := ""
    for i, word := range args {
        result += word
        if i < len(args)-1 {
            result += "\n"
        }
    }
    return result
}
```

> **BOOM!** Quest Complete! (And 100% Legal!)

---

## Visual Magic: See It Happen!

```
Input:  ["Hello", "how", "are", "you?"]

       ┌───────┐
       │ Hello │
       ├───────┤
       │ how   │
       ├───────┤
       │ are   │
       ├───────┤
       │ you?  │
       └───────┘

       ↓↓↓ Manual Loop + "\n" (except last) ↓↓↓

Output: "Hello\nhow\nare\nyou?"

Printed:
Hello
how
are
you?
```

---

## Interactive Challenge for Students! (Class Game Time!)

### Game: "Build the String Tower!"

**Rules:**
1. Teacher says a list: `["Hi", "Go", "is", "fun"]`
2. Students **shout** the output line by line!
3. First to say the **correct full string** wins a star!

**Example Round:**
> Teacher: `["Go", "is", "cool"]`  
> Student 1: "Go\nis\ncool" → **Correct!**  
> Student 2: "Gois\ncool" → Try again!

---

## Why No `strings.Join`?

| Problem | Fix |
|--------|-----|
| `import "strings"` | **Banned!** (`--allow-builtin` only allows basics) |
| Cheating Error | `illegal-import strings` |
| Solution | **Manual loop** – pure Go power! |

---

## Summary: What You Learned!

| Skill | You Mastered It! |
|------|------------------|
| Functions | `func Name(input) output` |
| Slices | `[]string` = list of strings |
| Loops | `for i, word := range args` |
| String concat | `+` to glue strings |
| Newlines | `\n` = line break |
| Clean code | Only add `\n` between words! |

---

## Final Test: Can You Solve This?

```go
words := []string{"I", "love", "coding"}
// What does ConcatParams return?
```

**Answer (shout it!):**
```
I\nlove\ncoding
→ Prints:
I
love
coding
```
---

## Files to Submit

**File:** `concatparams.go`

```go
package piscine

func ConcatParams(args []string) string {
    result := ""
    for i, word := range args {
        result += word
        if i < len(args)-1 {
            result += "\n"
        }
    }
    return result
}
```

---

**Quest Complete!**  
Now go run:

```bash
go run .
```

And watch the magic!

```
Hello
how
are
you?
```

---

**Teacher Tip:**  
> Always ask:  
> - “What goes in?”  
> - “What comes out?”  
> - “How do I get from A to B?”  

That’s how **every** coding quest is solved!

---
# `SplitWhiteSpaces`
---

**Quest Name:** `SplitWhiteSpaces`   
**Goal:** **Chop a sentence into words** using spaces, tabs, and newlines as **scissors**!  

Your mission? Take a **messy string** and **slice it cleanly** into a **list of words**!

> No `strings.Split` allowed!  
> Only **pure Go** and `--allow-builtin`!  
> (No imports! No cheating!)

---

## The Magic Spell (Function) You Must Write

```go
func SplitWhiteSpaces(s string) []string {

}
```

> **Your job:** Fill in the `{ }` so it **cuts perfectly**!

---

## Let’s Break It Down – Line by Line (Super Simple!)

| Line | Explanation | Fun Analogy |
|------|-----------|-------------|
| `func` | Starts a function | “Let the slicing begin!” |
| `SplitWhiteSpaces` | Name of your ninja move | **“WordChop!”** |
| `(s string)` | Input: one long string | A **giant sushi roll** of text |
| `[]string` | Output: a **slice** of strings | A **plate of neatly cut sushi pieces** |
| `{ }` | Where the **chopping happens** | The **cutting board** |

---

## Example Input & Output (The Test!)

```go
fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
```

**Output:**
```go
[]string{"Hello", "how", "are", "you?"}
```

> 4 perfect words!  
> No extra spaces!  
> No empty strings!

---

## What Are "White Spaces"?

| Symbol | Name | What It Does |
|--------|------|--------------|
| ` ` | Space | Normal space |
| `\t` | Tab | Big space (like pressing Tab) |
| `\n` | Newline | Line break |

> **Your job:** Treat **all three** as **word separators**!

---

## Step-by-Step: How to Solve It (Like a Game!)

Let’s play **"Word Chopper"**!

---

### Level 1: Understand the Input

```go
s := "Hello   how\tare\nyou?"
```

| Part | Meaning |
|------|--------|
| `"Hello"` | First word |
| `   ` | 3 spaces → **separator** |
| `"how"` | Second word |
| `\t` | Tab → **separator** |
| `"are"` | Third word |
| `\n` | Newline → **separator** |
| `"you?"` | Last word |

---

### Level 2: We Need to Find Words

We **skip** white spaces.  
We **collect** letters until we hit a space.

---

### Level 3: Use a Loop to Scan Character by Character

We’ll walk through the string like a **robot vacuum**:

```go
for each character in s:
    if it's a letter → add to current word
    if it's space/tab/newline → finish current word, start new one
```

---

### Level 4: Store Words in a Slice

We’ll use a **slice** (`[]string`) to collect words.

```go
words := []string{}
current := ""
```

---

## Final Working Code

```go
package piscine

func SplitWhiteSpaces(s string) []string {
    words := []string{}
    current := ""

    for _, char := range s {
        if char == ' ' || char == '\t' || char == '\n' {
            if current != "" {
                words = append(words, current)
                current = ""
            }
        } else {
            current += string(char)
        }
    }

    // Don't forget the last word!
    if current != "" {
        words = append(words, current)
    }

    return words
}
```
---

## Visual Magic: See It Happen!

```
Input: "Hello   how\tare\nyou?"

Step 1: Scan char by char
       H e l l o [   ] [   ] h o w [	] a r e [ \n ] y o u ?

Step 2: Build words
       current = "Hello" → space → save! → current = ""
       current = "how"   → tab  → save! → current = ""
       current = "are"   → \n   → save! → current = ""
       current = "you?"  → end  → save!

Step 3: Final words
       ["Hello", "how", "are", "you?"]

Output: []string{"Hello", "how", "are", "you?"}
```

---

## Interactive Challenge for you! (Class Game Time!)

### Game: "Human Word Chopper!"

**Rules:**
1. Teacher writes a messy string on the board:
   ```
   "Go\tis  fun\n  yes!"
   ```
2. Students **shout** the words one by one!
3. First team to list **all 4 words** wins!

**Answer:** `"Go"`, `"is"`, `"fun"`, `"yes!"`

---

## Why This Code Works (Line-by-Line Magic!)

| Line | What It Does | Why It Matters |
|------|-------------|----------------|
| `words := []string{}` | Empty plate for words | Start fresh! |
| `current := ""` | Current word being built | Like a basket |
| `for _, char := range s` | Loop over **each letter** | We scan everything! |
| `if char == ' ' || ...` | Is it a separator? | Yes → time to cut! |
| `if current != ""` | Only save if we have a word | Avoid empty strings! |
| `append(words, current)` | Add word to plate | Collect it! |
| `current = ""` | Reset basket | Ready for next word |
| `else` | Not a space? | Keep building word |
| `current += string(char)` | Add letter | Grow the word |
| Final `if current != ""` | Last word? | Don’t forget it! |

---

## Common Bugs & Fixes

| Bug | Fix |
|-----|-----|
| Empty strings in result | Only `append` if `current != ""` |
| Missing last word | Add final `if current != ""` |
| Using `strings.Fields` | **BANNED!** (import needed) |
| Using `strings.Split` | **BANNED!** |

---

## Summary: What You Learned!

| Skill | You Mastered It! |
|------|------------------|
| Loops with `range` | Scan every character |
| `rune` vs `string` | `char` is a `rune` |
| `append()` | Grow slices dynamically |
| Conditionals | `if char == ' '` etc. |
| No imports! | Pure Go power! |

---

## Final Test: Can You Solve This?

```go
s := "  Go\tlang  \nrocks!  "
// What does SplitWhiteSpaces return?
```

**Answer (shout it!):**
```go
[]string{"Go", "lang", "rocks!"}
```

> No leading/trailing spaces!  
> No empty strings!

---

## Files to Submit

**File:** `splitwhitespaces.go`

```go
package piscine

func SplitWhiteSpaces(s string) []string {
    words := []string{}
    current := ""

    for _, char := range s {
        if char == ' ' || char == '\t' || char == '\n' {
            if current != "" {
                words = append(words, current)
                current = ""
            }
        } else {
            current += string(char)
        }
    }

    if current != "" {
        words = append(words, current)
    }

    return words
}
```

---

**Quest Complete!**  
Now go run:

```bash
go run .
```

And see the magic:

```go
[]string{"Hello", "how", "are", "you?"}
```

---
# `PrintWordsTables` 

---

**Quest Name:** `PrintWordsTables`  
**Level:** 14 (Optional – Extra XP!)  
**XP Reward:** 2.45 kB of Brain Power!  
**Goal:** **Print each word on its own line** using **only `z01.PrintRune`**!  
  
Your mission? Take a **list of words** and **print them one per line** – like a magical typewriter!

> **Allowed Magic:**  
> - `z01.PrintRune` (to print **one letter at a time**)  
> - `--allow-builtin`  
> - **No `fmt.Println`!**  
> - **No `strings` package!**

---

## The Magic Spell (Function) You Must Write

```go
func PrintWordsTables(a []string) {

}
```

> **Your job:** Fill in the `{ }` so it **prints perfectly**!

---

## Let’s Break It Down – Line by Line (Super Simple!)

| Line | Explanation | Fun Analogy |
|------|-----------|-------------|
| `func` | Start the spell | “Let’s cast PrintWords!” |
| `PrintWordsTables` | Name of your function | **“WordPrinter 3000!”** |
| `(a []string)` | Input: a **slice** of words | A **stack of note cards** |
| `{ }` | Where the **printing happens** | The **typewriter station** |

---

## Example Input & Output (The Test!)

```go
a := piscine.SplitWhiteSpaces("Hello how are you?")
piscine.PrintWordsTables(a)
```

**Output:**
```
Hello
how
are
you?
```

> Each word on its **own line**!  
> Powered by `z01.PrintRune`!

---

## What is `z01.PrintRune`?

| Function | What It Does |
|---------|---------------|
| `z01.PrintRune('A')` | Prints **one letter** (`A`) |
| `z01.PrintRune('\n')` | Prints a **newline** (line break) |

> You are a **letter-by-letter printer**!

---

## Step-by-Step: How to Solve It (Like a Game!)

Let’s play **"Typewriter Challenge"**!

---

### Level 1: Understand the Input

```go
a := []string{"Hello", "how", "are", "you?"}
```

| Index | Word |
|-------|------|
| 0 | "Hello" |
| 1 | "how" |
| 2 | "are" |
| 3 | "you?" |

---

### Level 2: Print Each Word

For **each word**:
1. Print **each letter** using `z01.PrintRune`
2. Print a **newline** `\n`

---

### Level 3: Loop Through Words & Letters

```go
for each word in a:
    for each letter in word:
        z01.PrintRune(letter)
    z01.PrintRune('\n')
```

---

## Final Working Code

```go
package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
    for _, word := range a {
        for _, char := range word {
            z01.PrintRune(char)
        }
        z01.PrintRune('\n')
    }
}
```

> **BOOM!** Quest Complete!  
> **No `fmt`!**  
> **Only `z01.PrintRune`!**

---

## Visual Magic: See It Happen!

```
Input: ["Hello", "how", "are", "you?"]

Step 1: Loop over words
       word = "Hello" → print H e l l o → \n
       word = "how"   → print h o w     → \n
       word = "are"   → print a r e     → \n
       word = "you?"  → print y o u ?   → \n

Output:
Hello
how
are
you?
```

---

## Interactive Challenge for Students! (Class Game Time!)

### Game: "Human Typewriter!"

**Rules:**
1. Teacher says a word: `"Go"`
2. Students **shout** each letter **one by one**:
   - Student 1: `G`
   - Student 2: `o`
   - Student 3: **NEWLINE!**
3. Next word: `"lang"` → repeat!

**Winner:** Team that prints **all words perfectly**!

---

## Why This Code Works (Line-by-Line Magic!)

| Line | What It Does | Why It Matters |
|------|-------------|----------------|
| `for _, word := range a` | Loop over **each word** | Visit every note card |
| `for _, char := range word` | Loop over **each letter** | Print one by one |
| `z01.PrintRune(char)` | Print **one rune** | Like typing on a keyboard |
| `z01.PrintRune('\n')` | **New line** after word | Each word gets its own line |

---

## Common Bugs & Fixes

| Bug | Fix |
|-----|-----|
| No newline at end | You **want** it! (Each word on own line) |
| Using `fmt.Println` | **BANNED!** |
| Forgetting inner loop | Only prints first letter! |
| Using `string(char)` | Wrong! `char` is already a `rune` |

---

## Summary: What You Learned!

| Skill | You Mastered It! |
|------|------------------|
| Nested loops | `for` inside `for` |
| `z01.PrintRune` | Print **one character** |
| `rune` type | `char` in `range` = `rune` |
| Slices | `[]string` = list of words |
| No `fmt`! | Pure `z01` power! |

---

## Final Test: Can You Predict This?

```go
words := []string{"I", "love", "Go"}
PrintWordsTables(words)
```

**Output (shout it!):**
```
I
love
Go
```

---

## Badge Earned!

```
TYPEWRITER WIZARD – LEVEL 14
"You print letters like a pro!"
```

---

## Files to Submit

**File:** `printwordstables.go`

```go
package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
    for _, word := range a {
        for _, char := range word {
            z01.PrintRune(char)
        }
        z01.PrintRune('\n')
    }
}
```

---

**Quest Complete!**  
Now go run:

```bash
go run .
```

And watch the magic:

```
Hello
how
are
you?
```

---

**Teacher Tip:**  
> Always ask:  
> - “How do I print **one letter**?” → `z01.PrintRune`  
> - “How do I go to the **next line**?” → `'\n'`  
> - “How do I visit **every word and letter**?” → **nested loops**  
---
# `Split` 
---

**Quest Name:** `Split`  
**Level:** 14 (Optional – Extra XP!)  
**XP Reward:** 3.06 kB of Brain Power!  
**Goal:** **Chop a string into pieces** using a **custom separator** (like a magic knife)!  
Your mission?  
> Take a string and a **separator**, and **cut** the string **every time** the separator appears!

> **Allowed Magic:**  
> - `make`  
> - `--allow-builtin`  
> - **No `strings.Split`!**  
> - **No imports!**

---

## The Magic Spell (Function) You Must Write

```go
func Split(s, sep string) []string {

}
```

> **Your job:** Fill in the `{ }` so it **slices perfectly**!

---

## Let’s Break It Down – Line by Line (Super Simple!)

| Line | Explanation | Fun Analogy |
|------|-----------|-------------|
| `func` | Start the spell | “Let the cutting begin!” |
| `Split` | Name of your function | **“StringSlicer!”** |
| `(s, sep string)` | Two inputs: **text** and **separator** | A **cake** and a **knife pattern** |
| `[]string` | Output: **list of pieces** | **Plates of cake slices** |
| `{ }` | Where the **slicing happens** | The **kitchen counter** |

---

## Example Input & Output (The Test!)

```go
s := "HelloHAhowHAareHAyou?"
fmt.Printf("%#v\n", piscine.Split(s, "HA"))
```

**Output:**
```go
[]string{"Hello", "how", "are", "you?"}
```

> `"HA"` is the **separator**!  
> Every time `"HA"` appears → **cut**!

---

## Visual Magic: See the Cuts!

```
Original: HelloHAhowHAareHAyou?

Knife:    "HA"

Cuts at:      ↑     ↑     ↑

Result:   ["Hello", "how", "are", "you?"]
```

---

## Step-by-Step: How to Solve It (Like a Game!)

Let’s play **"Find the Knife!"**

---

### Level 1: Understand the Input

```go
s   = "HelloHAhowHAareHAyou?"
sep = "HA"
```

We need to **find every `"HA"`** in `s` and **split there**.

---

### Level 2: Scan the String

We’ll walk through `s` **character by character**, looking for **full matches** of `sep`.

> Can’t just split on `'H'` or `'A'` — must be **both together**!

---

### Level 3: Use a Loop + Index Tracking

```go
result := []string{}
start := 0

for i := 0; i < len(s); i++ {
    if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
        // Found separator!
        word := s[start:i]
        result = append(result, word)
        i += len(sep) - 1  // Skip the separator
        start = i + 1
    }
}

// Don't forget the last piece!
if start < len(s) {
    result = append(result, s[start:])
}
```

---

## Final Working Code (The Spell is Complete!)

```go
package piscine

func Split(s, sep string) []string {
    result := []string{}
    start := 0
    seplen := len(sep)

    for i := 0; i < len(s); i++ {
        if seplen > 0 && i+seplen <= len(s) && s[i:i+seplen] == sep {
            // Found full separator
            word := s[start:i]
            result = append(result, word)
            i += seplen - 1 // Skip the entire sep
            start = i + 1
        }
    }

    // Add the last part
    if start < len(s) {
        result = append(result, s[start:])
    }

    return result
}
```

> **BOOM!** Quest Complete!  
> **No imports!**  
> **Handles any separator!**

---

## Visual Step-by-Step (With Example!)

```go
s   = "HelloHAhowHAareHAyou?"
sep = "HA"
```

| `i` | `s[i:i+2]` | Match? | Action |
|-----|------------|--------|-------|
| 0 | "He" | No | — |
| 1 | "el" | No | — |
| ... | ... | ... | ... |
| 5 | "HA" | YES! | Cut: `"Hello"` → add to result |
| 6 | "Ah" | No | — |
| ... | ... | ... | ... |
| 10 | "HA" | YES! | Cut: `"how"` |
| 14 | "HA" | YES! | Cut: `"are"` |
| 17 | "yo" | No | — |
| End | — | — | Add last: `"you?"` |

**Final result:** `["Hello", "how", "are", "you?"]`

---

## Interactive Challenge for Students! (Class Game Time!)

### Game: "Knife Hunt!"

**Rules:**
1. Teacher writes: `AppleXXBananaXXCherry`
2. Separator: `XX`
3. Students **shout** the cuts:
   - “Cut at position 5!”
   - “Cut at position 12!”
4. Final answer: `["Apple", "Banana", "Cherry"]`

**Winner:** First team to get all pieces!

---

## Edge Cases (Boss Fights!)

| Case | Input | Expected | Code Handles? |
|------|-------|----------|---------------|
| Empty string | `""`, `"HA"` | `[]` | YES |
| No separator | `"Hello"`, `"HA"` | `["Hello"]` | YES |
| Separator at start | `"HAHello"`, `"HA"` | `["", "Hello"]` | YES |
| Separator at end | `"HelloHA"`, `"HA"` | `["Hello", ""]` | YES |
| Multiple sep | `"HAHAA"`, `"HA"` | `["", "", "A"]` | YES |
| Empty sep | `"abc"`, `""` | **Don’t crash!** | We check `seplen > 0` |

---

## Why This Code Works (Line-by-Line Magic!)

| Line | What It Does | Why It Matters |
|------|-------------|----------------|
| `result := []string{}` | Empty list | Start collecting pieces |
| `start := 0` | Where current word begins | Track start of slice |
| `seplen := len(sep)` | Length of separator | Avoid recalculating |
| `i+seplen <= len(s)` | Bounds check | Prevent panic |
| `s[i:i+seplen] == sep` | Full match? | Only split on **exact** sep |
| `i += seplen - 1` | Skip entire sep | Don’t recheck same letters |
| Final `if start < len(s)` | Last piece | Don’t forget it! |

---

## Summary: What You Learned!

| Skill | You Mastered It! |
|------|------------------|
| String slicing | `s[i:j]` |
| Bounds checking | `i+len <= len(s)` |
| Index tracking | `start` and `i` |
| `append()` | Grow slice |
| Edge cases | Empty, start/end, multiple |
| No imports! | Pure Go! |

---

## Final Test: Can You Predict This?

```go
Split("++A++B++", "++")
```

**Answer (shout it!):**
```go
[]string{"", "A", "B", ""}
```

> Empty at start, middle, and end!

---

## Badge Earned!

```
STRING SLICER MASTER – LEVEL 14
"You cut strings like a sushi chef!"
```

---

## Files to Submit

**File:** `split.go`

```go
package piscine

func Split(s, sep string) []string {
    result := []string{}
    start := 0
    seplen := len(sep)

    for i := 0; i < len(s); i++ {
        if seplen > 0 && i+seplen <= len(s) && s[i:i+seplen] == sep {
            word := s[start:i]
            result = append(result, word)
            i += seplen - 1
            start = i + 1
        }
    }

    if start < len(s) {
        result = append(result, s[start:])
    }

    return result
}
```

---

**Quest Complete!**  
Now go run:

```bash
go run .
```

And see the magic:

```go
[]string{"Hello", "how", "are", "you?"}
```

---

**Teacher Tip:**  
> Always ask:  
> - “Where does the **separator start**?”  
> - “Did I **skip** the whole separator?”  
> - “Did I add the **last piece**?”  
---
# `ConvertBase` 

---

**Quest Name:** `ConvertBase`  
**Level:** 14 (BONUS – Ultimate XP!)  
**XP Reward:** 3.67 kB of Brain Power!  
**Goal:** **Transform numbers from ANY base to ANY base** – like a **number wizard**!  

Your mission?  
> Take a number in **one base** (like binary `101011`)  
> Convert it to **another base** (like decimal `43`)  
> **No negative numbers**  
> **Only valid bases**

> **Allowed Magic:**  
> - `--allow-builtin`  
> - **No imports!**  
> - **No `strconv`**  
> - **Pure Go logic only!**

---

## The Magic Spell (Function) You Must Write

```go
func ConvertBase(nbr, baseFrom, baseTo string) string {

}
```

> **Your job:** Fill in the `{ }` so it **transmutes perfectly**!

---

## Let’s Break It Down – Line by Line (Super Simple!)

| Line | Explanation | Fun Analogy |
|------|-----------|-------------|
| `func` | Start the spell | “Let the alchemy begin!” |
| `ConvertBase` | Name of your function | **“BaseTransmuter!”** |
| `(nbr, baseFrom, baseTo string)` | 3 inputs | **Potion**, **Old Flask**, **New Flask** |
| `string` | Output: number in **new base** | **Final Elixir** |
| `{ }` | Where the **conversion happens** | The **cauldron** |

---

## Example Input & Output (The Test!)

```go
result := piscine.ConvertBase("101011", "01", "0123456789")
fmt.Println(result)
```

**Output:**
```
43
```

> `"101011"` in **base 2 (binary)** → `43` in **base 10 (decimal)**!

---

## The 3-Step Alchemy Formula

| Step | Name | What You Do |
|------|------|-------------|
| 1 | **Parse** | Turn `nbr` (in `baseFrom`) → **decimal** |
| 2 | **Convert** | Turn **decimal** → `baseTo` |
| 3 | **Return** | Give back as **string** |

---

## Step 1: Parse – From Any Base → Decimal

```go
value := 0
for each digit in nbr:
    value = value * len(baseFrom) + position_of(digit)
```

### Example: `"101011"` in base `"01"`

| Digit | Position in `"01"` | Calculation |
|-------|---------------------|-----------|
| `1` | 1 | `0*2 + 1 = 1` |
| `0` | 0 | `1*2 + 0 = 2` |
| `1` | 1 | `2*2 + 1 = 5` |
| `0` | 0 | `5*2 + 0 = 10` |
| `1` | 1 | `10*2 + 1 = 21` |
| `1` | 1 | `21*2 + 1 = 43` |

→ **Decimal = 43**

---

## Step 2: Convert – From Decimal → Any Base

```go
result := ""
while value > 0:
    digit = value % len(baseTo)
    result = baseTo[digit] + result
    value = value / len(baseTo)
```

### Example: `43` → base `"0123456789"`

| Step | `value` | `digit = 43 % 10` | `baseTo[digit]` | `result` |
|------|--------|-------------------|------------------|----------|
| 1 | 43 | 3 | `'3'` | `"3"` |
| 2 | 4 | 4 | `'4'` | `"43"` |
| 3 | 0 | — | — | **Done!** |

→ **Result = `"43"`**

---

## Final Working Code

```go
package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
    // Step 1: Parse to decimal
    decimal := 0
    for i := 0; i < len(nbr); i++ {
        digit := nbr[i]
        pos := 0
        for j := 0; j < len(baseFrom); j++ {
            if baseFrom[j] == digit {
                pos = j
                break
            }
        }
        decimal = decimal*len(baseFrom) + pos
    }

    // Step 2: Convert decimal to baseTo
    if decimal == 0 {
        return "0"
    }

    result := ""
    for decimal > 0 {
        digit := decimal % len(baseTo)
        result = string(baseTo[digit]) + result
        decimal /= len(baseTo)
    }

    return result
}
```

> **BOOM!** Quest Complete!  
> **No imports!**  
> **Works for ANY valid base!**

---

## Visual Magic: Full Conversion Flow

```
Input:
nbr       = "101011"
baseFrom  = "01"        (base 2)
baseTo    = "0123456789" (base 10)

Step 1: Parse → Decimal
"1" → 1
"0" → 0 ×2 +1 = 2
"1" → 1 ×2 +0 = 5
"0" → 0 ×2 +1 = 10
"1" → 1 ×2 +0 = 21
"1" → 1 ×2 +1 = 43

Step 2: Convert 43 → base 10
43 % 10 = 3 → '3'
43 / 10 = 4
4  % 10 = 4 → '4'
4  / 10 = 0 → stop

Result: "43"
```

---

## Interactive Challenge for you (Class Game Time!)

### Game: "Base Wizard Duel!"

**Rules:**
1. Teacher says: `Convert "FF" from "0123456789ABCDEF" to "0123456789"`
2. Students **race** to shout:
   - Step 1: `"FF" = 15×16 + 15 = 255`
   - Step 2: `255 → 2×100 + 5×10 + 5 = "255"`
3. First correct answer wins a **potion**!

---

## Edge Cases (Boss Fights!)

| Case | Input | Expected | Code Handles? |
|------|-------|----------|---------------|
| Zero | `"0"`, any bases | `"0"` | YES |
| Base 2 → 16 | `"1111"`, `"01"`, `"0123456789ABCDEF"` | `"F"` | YES |
| Base 10 → 2 | `"10"`, `"0123456789"`, `"01"` | `"1010"` | YES |
| Single digit | `"A"`, `"0123456789ABCDEF"`, `"0123456789"` | `"10"` | YES |

---

## Why This Code Works 

| Line | What It Does | Why It Matters |
|------|-------------|----------------|
| `decimal := 0` | Start at zero | Clean slate |
| Outer loop | Go through each digit | Left to right |
| Inner loop | Find digit position | `baseFrom[j] == digit` |
| `decimal = decimal*len + pos` | **Horner’s method** | Efficient parsing |
| `if decimal == 0` | Special case | Return `"0"` |
| `decimal % len(baseTo)` | Get digit | Remainder |
| `string(baseTo[digit])` | Get symbol | `'A'`, `'B'`, etc. |
| Prepend | `result = char + result` | Build **right to left** |

---

## Summary: What You Learned!

| Skill | You Mastered It! |
|------|------------------|
| Base conversion | Any → Decimal → Any |
| String indexing | `baseFrom[j]` |
| Nested loops | Find digit position |
| Modulo & division | `%` and `/` |
| String building | `+` prepend |
| Edge cases | Zero, single digit |

---

## Final Test: Can You Solve This?

```go
ConvertBase("1010", "01", "0123456789ABCDEF")
```

**Answer (shout it!):**
```
A
```

> `1010` in binary = `10` in decimal = `A` in hex!

---

## Badge Earned!

```
BASE ALCHEMIST SUPREME – LEVEL 14
"You transmute numbers like a god!"
```

---

## Files to Submit

**File:** `convertbase.go`

```go
package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
    // Step 1: Parse to decimal
    decimal := 0
    for i := 0; i < len(nbr); i++ {
        digit := nbr[i]
        pos := 0
        for j := 0; j < len(baseFrom); j++ {
            if baseFrom[j] == digit {
                pos = j
                break
            }
        }
        decimal = decimal*len(baseFrom) + pos
    }

    // Step 2: Convert to baseTo
    if decimal == 0 {
        return "0"
    }

    result := ""
    for decimal > 0 {
        digit := decimal % len(baseTo)
        result = string(baseTo[digit]) + result
        decimal /= len(baseTo)
    }

    return result
}
```

---

**Quest Complete!**  
Now go run:

```bash
go run .
```

And see the magic:

```
43
```

---

**Teacher Tip:**  
> Always ask:  
> - “How do I **read** the number?” → **Parse to decimal**  
> - “How do I **write** the number?” → **Divide & mod**  
> - “What if it’s **zero**?” → **Special case!**  
