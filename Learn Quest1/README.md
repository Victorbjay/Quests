# done.tar task
---

## **Part 1: Understanding File Permissions**

When you do `ls -l`, you see something like:
```
-rw-r--r--  myfile.txt
```

Let me decode this:

### The permission string breakdown:
```
-  rw-  r--  r--
│   │    │    │
│   │    │    └─ Others (everyone else) permissions
│   │    └────── Group permissions  
│   └─────────── Owner (you) permissions
└─────────────── File type (- = file, d = directory, l = link)
```

### Each section has 3 permissions:
- **r** = read (can view the file) = **4**
- **w** = write (can modify the file) = **2**  
- **x** = execute (can run the file) = **1**

### Converting to numbers (octal):
- `rw-` = read(4) + write(2) + nothing(0) = **6**
- `r--` = read(4) + nothing(0) + nothing(0) = **4**
- `r--` = read(4) + nothing(0) + nothing(0) = **4**

So `rw-r--r--` = **644**

---

## **Part 2: Let's decode ONE example from your task**

Take file `1`:
```
-r------w- 1986-11-13 00:01 1
```

Breaking it down:
- `-` = regular file (not a directory)
- `r--` = owner can **read** only = 4
- `---` = group has **no permissions** = 0
- `-w-` = others can **write** only = 2

So the octal permission = **402**

---

## **Part 3: The Commands Explained**

### Creating file 1:
```bash
touch 1              # Creates an empty file named "1"
chmod 402 1          # Sets permissions to -r------w-
touch -t 198611130001 1   # Sets timestamp to Nov 13, 1986, 00:01
```

### The timestamp format:
```
198611130001
│││││││││││└─ Minute (01)
││││││││││└── Hour (00)
│││││││││└─── Day (13)
││││││││└──── Month (11)
│││││││└───── Year (1986)
```

---

## **Part 4: Special cases**

### Directory `0`:
```
dr-------x = d 500 001 = directory with 501 permissions
```
- `d` = it's a directory
- `r--` = owner can read = 400
- `---` = group nothing = 000
- `--x` = others can execute (enter) = 001
- Total: **501**

### Symlink `3`:
```bash
ln -s 0 3    # Creates a link: 3 points to 0
```
Think of it like a shortcut on Windows - file `3` is just a pointer to directory `0`

---

## **Part 5: Now let's do this together!**

Copy and paste this **ONE COMMAND AT A TIME** into your terminal:

### Step 1: Make a fresh directory
```bash
mkdir permissions_task && cd permissions_task
```

### Step 2: Create directory 0
```bash
mkdir 0
chmod 501 0
touch -t 198601050000 0
```

### Step 3: Create file 1
```bash
touch 1
chmod 402 1
touch -t 198611130001 1
```

### Step 4: Create file 2
```bash
touch 2
chmod 604 2
touch -t 198803050010 2
```

### Step 5: Create symlink 3
```bash
ln -s 0 3
touch -h -t 199002160011 3
```

### Step 6: Create file 4
```bash
touch 4
chmod 511 4
touch -t 199010070100 4
```

### Step 7: Create file 5
```bash
touch 5
chmod 460 5
touch -t 199011070101 5
```

### Step 8: Create file 6
```bash
touch 6
chmod 460 6
touch -t 199102080110 6
```

### Step 9: Create file 7
```bash
touch 7
chmod 511 7
touch -t 199103080111 7
```

### Step 10: Create file 8
```bash
touch 8
chmod 604 8
touch -t 199405201000 8
```

### Step 11: Create file 9
```bash
touch 9
chmod 402 9
touch -t 199406101001 9
```

### Step 12: Create directory A
```bash
mkdir A
chmod 501 A
touch -t 199504101010 A
```

### Step 13: Create the tar file
```bash
tar -cf done.tar 0 1 2 3 4 5 6 7 8 9 A
```

### Step 14: Verify
```bash
ls
```
You should see: `0  1  2  3  4  5  6  7  8  9  A  done.tar`

---

## **Quick Reference Table**

| File | Permissions | Octal | Date & Time |
|------|-------------|-------|-------------|
| 0 | dr-------x | 501 | 1986-01-05 00:00 |
| 1 | -r------w- | 402 | 1986-11-13 00:01 |
| 2 | -rw----r-- | 604 | 1988-03-05 00:10 |
| 3 | lrwxrwxrwx | (symlink) | 1990-02-16 00:11 |
| 4 | -r-x--x--- | 511 | 1990-10-07 01:00 |
| 5 | -r--rw---- | 460 | 1990-11-07 01:01 |
| 6 | -r--rw---- | 460 | 1991-02-08 01:10 |
| 7 | -r-x--x--- | 511 | 1991-03-08 01:11 |
| 8 | -rw----r-- | 604 | 1994-05-20 10:00 |
| 9 | -r------w- | 402 | 1994-06-10 10:01 |
| A | dr-------x | 501 | 1995-04-10 10:10 |

---

Alright, let's master octal permissions! This is actually the trickiest part. 💡

---

## **The Permission Letters**

Every file has 3 groups of permissions:

```
-  rwx  rwx  rwx
│   │    │    │
│   │    │    └─── OTHERS (everyone else)
│   │    └──────── GROUP (your team)
│   └───────────── OWNER (you)
└───────────────── file type (ignore for now)
```

Each group has 3 possible permissions:
- **r** = read = **4 points**
- **w** = write = **2 points**
- **x** = execute = **1 point**
- **-** = nothing = **0 points**

---

## **How to Calculate: Add up the points!**

### Example 1: `rw-r--r--`

Break it into 3 groups:

**Group 1 (Owner): `rw-`**
- r = 4
- w = 2
- - = 0
- **Total: 4 + 2 + 0 = 6**

**Group 2 (Group): `r--`**
- r = 4
- - = 0
- - = 0
- **Total: 4 + 0 + 0 = 4**

**Group 3 (Others): `r--`**
- r = 4
- - = 0
- - = 0
- **Total: 4 + 0 + 0 = 4**

**Answer: 644** ✅

---

## **Example 2: `r------w-`** (from file 1 in your task)

**Group 1 (Owner): `r--`**
- r = 4
- - = 0
- - = 0
- **Total: 4**

**Group 2 (Group): `---`**
- - = 0
- - = 0
- - = 0
- **Total: 0**

**Group 3 (Others): `-w-`**
- - = 0
- w = 2
- - = 0
- **Total: 2**

**Answer: 402** ✅

---

## **Example 3: `rwx------`** (full permissions for owner only)

**Group 1 (Owner): `rwx`**
- r = 4
- w = 2
- x = 1
- **Total: 4 + 2 + 1 = 7**

**Group 2 (Group): `---`**
- **Total: 0**

**Group 3 (Others): `---`**
- **Total: 0**

**Answer: 700** ✅

---

## **Example 4: `r-x--x---`** (from file 4 in your task)

**Group 1 (Owner): `r-x`**
- r = 4
- - = 0
- x = 1
- **Total: 4 + 0 + 1 = 5**

**Group 2 (Group): `--x`**
- - = 0
- - = 0
- x = 1
- **Total: 0 + 0 + 1 = 1**

**Group 3 (Others): `---`**
- **Total: 0 + 0 + 0 = 0**

**Answer: 510** ✅

---

## **Quick Memory Trick**

Think of each position as a **light switch**:
- Switch ON = the letter shows (r, w, or x)
- Switch OFF = dash shows (-)

Then add up the "ON" switches:
- **First switch (r) = 4**
- **Second switch (w) = 2**
- **Third switch (x) = 1**

---

## **Practice Problems** 

Try these yourself:

1. `rwxr-xr-x` = ?
2. `rw-------` = ?
3. `r--rw----` = ?
4. `rwxrwxrwx` = ?

<details>
<summary>Click to see answers</summary>

1. `rwxr-xr-x` = **755** (7+5+5)
2. `rw-------` = **600** (6+0+0)
3. `r--rw----` = **460** (4+6+0) - this is file 5 and 6!
4. `rwxrwxrwx` = **777** (7+7+7) - full permissions for everyone

</details>

---

## **All Your Task Files**

Let me calculate ALL of them for you:

| Permissions | Owner | Group | Others | Octal |
|-------------|-------|-------|--------|-------|
| `dr-------x` | r-- = 4 | --- = 0 | --x = 1 | **401** ✅ |
| `-r------w-` | r-- = 4 | --- = 0 | -w- = 2 | **402** ✅ |
| `-rw----r--` | rw- = 6 | --- = 0 | r-- = 4 | **604** ✅ |
| `-r-x--x---` | r-x = 5 | --x = 1 | --- = 0 | **510** ✅ |
| `-r--rw----` | r-- = 4 | rw- = 6 | --- = 0 | **460** ✅ |


---
#Who are you Task
---

## 🎯 **Understanding the Task**

You need to:
1. Fetch data from a URL (JSON file)
2. Find the superhero with `id: 70`
3. Extract only their **name**
4. Save this as a bash script called `who-are-you.sh`

---

## 📚 **Step 1: Understand the Tools**

### **curl** - Downloads content from URLs
```bash
curl https://example.com/data.json
```
This fetches the content from a website.

### **jq** - Processes JSON data
```bash
echo '{"name":"Batman"}' | jq '.name'
```
Output: `"Batman"`

It extracts specific fields from JSON.

---

## 🔍 **Step 2: Explore the Data**

Let's see what the JSON looks like:

```bash
curl https://acad.learn2earn.ng/assets/superhero/all.json
```

You'll see something like:
```json
[
  {
    "id": 1,
    "name": "A-Bomb",
    "powerstats": {...},
    ...
  },
  {
    "id": 2,
    "name": "Abe Sapien",
    ...
  },
  ...
  {
    "id": 70,
    "name": "Some Hero Name",
    ...
  }
]
```

It's an **array** of superhero objects!

---

## 🧩 **Step 3: Build the Solution Piece by Piece**

### **Test 1: Fetch the data**
```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json
```
- `-s` means "silent" (no progress bar)

### **Test 2: Use jq to parse**
```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.'
```
- `'.'` means "show everything"
- This makes the JSON pretty and readable

### **Test 3: Filter by id = 70**
```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[] | select(.id == 70)'
```

**Breaking it down:**
- `.[]` - Loop through each item in the array
- `|` - Pipe (send data to next command)
- `select(.id == 70)` - Keep only items where id equals 70

### **Test 4: Get only the name**
```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[] | select(.id == 70) | .name'
```
- `.name` - Extract just the name field

---

## 📝 **Step 4: Create the Script**

Now let's put it in a file called `who-are-you.sh`:

---

## 🚀 **Step 5: Create and Test the Script**

### **Method 1: Using cat (easiest)**
```bash
cat > who-are-you.sh << 'EOF'
#!/bin/bash

curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[] | select(.id == 70) | .name'
EOF
```

### **Method 2: Using nano**
```bash
nano who-are-you.sh
```
Then paste:
```bash
#!/bin/bash

curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[] | select(.id == 70) | .name'
```
Press `Ctrl+O`, `Enter`, `Ctrl+X`

---

## ⚙️ **Step 6: Make It Executable**

```bash
chmod +x who-are-you.sh
```

This gives the script permission to run.

---

## 🧪 **Step 7: Test It**

### **Run the script:**
```bash
./who-are-you.sh
```

Expected output (example):
```
"Batman"
```

### **Test with cat -e:**
```bash
./who-are-you.sh | cat -e
```

Expected output:
```
"Batman"$
```

The `$` shows the end of line (newline character).

---

## 📖 **Line-by-Line Explanation**

Let me break down the script:

```bash
#!/bin/bash
```
- **Shebang line** - Tells the system to use bash to run this script
- Must be the first line!

```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json
```
- `curl` - Download data from URL
- `-s` - Silent mode (no progress bar)
- Downloads the JSON file with all superheroes

```bash
| jq '.[] | select(.id == 70) | .name'
```
- `|` - Pipe operator (send curl output to jq)
- `.[]` - Loop through each superhero in the array
- `select(.id == 70)` - Keep only the hero with id 70
- `.name` - Extract just the name field

---

## 🔍 **Understanding jq Syntax**

### **The Array Loop: `.[]`**
```json
[{"id": 1}, {"id": 2}]
```
`.[]` gives you each object one by one:
```
{"id": 1}
{"id": 2}
```

### **The Select Filter: `select(.id == 70)`**
Keeps only objects where the condition is true:
```
{"id": 70, "name": "Batman"} ✅ Keep
{"id": 71, "name": "Robin"} ❌ Skip
```

### **The Field Access: `.name`**
```
{"id": 70, "name": "Batman"}
```
`.name` extracts:
```
"Batman"
```

---

## 🛠️ **Troubleshooting**

### **Problem: "curl: command not found"**
Install curl:
```bash
# Ubuntu/Debian
sudo apt install curl

# Mac
brew install curl
```

### **Problem: "jq: command not found"**
Install jq:
```bash
# Ubuntu/Debian
sudo apt install jq

# Mac
brew install jq
```

### **Problem: "Permission denied"**
Make it executable:
```bash
chmod +x who-are-you.sh
```

### **Problem: Wrong output format**
Make sure your script has:
1. `#!/bin/bash` at the top
2. Exactly the curl + jq command
3. No extra spaces or characters

---

## 🎓 **Learning Points**

### **What You Learned:**

1. **Bash scripting basics**
   - Creating executable scripts
   - Using shebang (`#!/bin/bash`)

2. **curl command**
   - Fetching data from URLs
   - Silent mode (`-s`)

3. **jq for JSON processing**
   - Array iteration (`.[]`)
   - Filtering (`select()`)
   - Field extraction (`.name`)

4. **Piping (`|`)**
   - Chaining commands together
   - Output of one becomes input of next

---

## 📊 **Visual Flow**

```
Internet
   │
   ├─ curl downloads JSON
   │   https://acad.learn2earn.ng/assets/superhero/all.json
   │
   ↓
[{id:1, name:"A-Bomb"}, {id:70, name:"Batman"}, {id:71, name:"Robin"}]
   │
   ├─ .[] loops through array
   │
   ↓
{id:1, name:"A-Bomb"}
{id:70, name:"Batman"}  ← select(.id == 70) keeps this
{id:71, name:"Robin"}
   │
   ├─ .name extracts name field
   │
   ↓
"Batman"
```

---

## 🧪 **Practice Exercises**

Try modifying the script:

### **1. Get superhero with id 1:**
```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[] | select(.id == 1) | .name'
```

### **2. Get both name and id for hero 70:**
```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[] | select(.id == 70) | {name, id}'
```

### **3. Get all superhero names:**
```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[].name'
```

---

## ✅ **Final Checklist**

- [ ] Script is named `who-are-you.sh`
- [ ] Script has `#!/bin/bash` at the top
- [ ] Script has the curl + jq command
- [ ] Script is executable (`chmod +x`)
- [ ] Output shows a name in quotes: `"SomeName"`
- [ ] `./who-are-you.sh | cat -e` shows `"SomeName"$`

---

## 🎉 **You Did It!**

You just learned:
- How to make API calls with curl
- How to parse JSON with jq
- How to create bash scripts
- How to chain commands with pipes

Now you know who you are! 🦸

---

**Ready to run it? Let me know what name you discover!** 😊
