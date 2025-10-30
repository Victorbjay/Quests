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
