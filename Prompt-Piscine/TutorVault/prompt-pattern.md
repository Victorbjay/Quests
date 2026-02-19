# Tutor Vault: Prompt Patterns - Complete Guide

## 📚 Overview

This guide teaches you how to complete the **Prompt Patterns** task from the 01-Edu Prompt Piscine. By the end, you'll understand how to transform weak prompts into powerful, structured AI interactions.

---

## 🎯 Learning Objectives

- Learn key prompt design patterns
- Transform single use-cases into reusable prompt templates
- Improve weak prompts by applying structured patterns
- Test prompts with multiple inputs for reliability

---

## 📋 Exercise Breakdown

### Exercise 0: Environment Setup

**Task:** Set up Python and verify it works.

**Steps:**
1. Check Python version: `python3 --version`
2. Create a simple Hello World program:

```python
print("Hello, World!")
```

**Expected Output:**
```
Hello, World!
```

---

### Exercise 1: Summarization Pattern

**Goal:** Create a reusable prompt for text summarization.

#### Step 1: Write a Weak Prompt

**Weak Prompt:**
```
"Summarize this text."
```

**Problem:** Too vague - AI gives inconsistent, unfocused results.

#### Step 2: Write a Refined Prompt

**Refined Prompt:**
```
"Summarize the following text in 3 bullet points, focusing on the main arguments and ignoring examples."
```

#### Step 3: Test with Different Inputs

**Input 1 (News Article):**
> "Apple today announced the new iPhone 16 Pro with advanced AI features, improved camera system, and longer battery life. The phone starts at $999 and will be available in stores next month. Analysts predict strong sales despite economic concerns."

**Weak Output:**
> "Apple announced a new iPhone with AI features, better camera, and longer battery. It costs $999 and comes out next month."

**Refined Output:**
- Apple released iPhone 16 Pro with AI features, better camera, and improved battery
- Price starts at $999, available next month
- Analysts forecast strong sales despite economic challenges

**Input 2 (Technical Article):**
> "React 19 introduces Server Components, which allow developers to render components on the server and stream them to the client. This reduces bundle size and improves initial load times. The new compiler eliminates the need for manual memoization, making code simpler and faster."

**Refined Output:**
- React 19 brings Server Components for server-side rendering
- New compiler removes manual memoization requirement
- These updates change React application development

#### Key Insight
The refined prompt specifies:
- ✅ Format (3 bullet points)
- ✅ Focus (main arguments, ignore examples)
- ✅ This leads to consistent, useful outputs

---

### Exercise 2: Data Extraction Template

**Goal:** Extract structured data from unstructured text.

#### Step 1: Create the Template Prompt

```
"Extract the following fields from the text: Name, Age, Location, Occupation. Return the output in JSON format."
```

#### Step 2: Test with Multiple Inputs

**Input 1:**
> "John Doe, age 29, lives in Paris and works as a software engineer."

**Output:**
```json
{
  "Name": "John Doe",
  "Age": 29,
  "Location": "Paris",
  "Occupation": "Software Engineer"
}
```

**Input 2:**
> "Sarah Johnson is a 34-year-old data scientist based in London."

**Output:**
```json
{
  "Name": "Sarah Johnson",
  "Age": 34,
  "Location": "London",
  "Occupation": "Data Scientist"
}
```

**Input 3:**
> "Michael Chen, 42, works as a product manager in Singapore."

**Output:**
```json
{
  "Name": "Michael Chen",
  "Age": 42,
  "Location": "Singapore",
  "Occupation": "Product Manager"
}
```

#### Consistency Check

| Field | Input 1 | Input 2 | Input 3 |
|-------|---------|---------|---------|
| Name | John Doe | Sarah Johnson | Michael Chen |
| Age | 29 | 34 | 42 |
| Location | Paris | London | Singapore |
| Occupation | Software Engineer | Data Scientist | Product Manager |

✅ All outputs are consistent and valid JSON!

---

### Exercise 3: Prompt Transformation

**Goal:** Strengthen weak prompts using patterns.

#### Starting Point (Weak Prompt)
```
"Write about climate change."
```

**Problem:** Too vague, no direction, output is generic.

---

#### Transformation 1: Instructional Pattern

**Prompt:**
```
"Write a 200-word explanation of climate change for high school learners."
```

**What changed:**
- ✅ Clear length (200 words)
- ✅ Specific audience (high school learners)
- ✅ Topic focus defined

**Expected Output:**
> Climate change refers to long-term shifts in global temperatures and weather patterns. Since the 1800s, human activities have been the main driver, primarily through burning fossil fuels...
>
> [Continues with educational content suitable for high schoolers]

---

#### Transformation 2: Comparative Pattern

**Prompt:**
```
"Compare the effects of climate change on coastal vs inland regions in 5 bullet points."
```

**What changed:**
- ✅ Clear comparison structure
- ✅ Two distinct categories to compare
- ✅ Specific format (5 bullet points)

**Expected Output:**
**Coastal Regions:**
- Sea level rise causes erosion and flooding
- Saltwater intrusion contaminates freshwater
- Increased hurricane intensity
- Economic impact on fishing/tourism
- Community displacement

**Inland Regions:**
- More frequent droughts affect agriculture
- Increased wildfire risk
- Heat waves in urban areas
- Changes in precipitation patterns
- Shift in agricultural zones

---

## 📊 Prompt Types Comparison

| Prompt Type | Meaning | Outcome |
|-------------|---------|---------|
| Vague Prompt | No clear direction | Broad, unfocused response |
| Instructional | Clear focus + structure | Organized, targeted reply |
| Comparative | Analytical comparison | Deeper critical analysis |

---

## 🔑 Key Takeaways

1. **Be Specific:** Weak prompts produce weak results
2. **Specify Format:** Tell AI exactly how you want output (bullet points, JSON, etc.)
3. **Define Audience:** "For beginners" vs "For experts" changes everything
4. **Test Multiple Times:** Always test with different inputs to ensure reliability
5. **Choose the Right Pattern:**
   - Need explanation? → Instructional
   - Need analysis? → Comparative
   - Need data? → Extraction
   - Need summary? → Summarization

---

## 📁 File Structure

Your final submission should look like:

```
prompt-patterns/
├── Ex0_Env_Libraries.py      # Python setup + Hello World
├── Ex1_Sum_Pattern.md        # Summarization pattern exercises
├── Ex2_Data_Ext_Template.md  # Data extraction with JSON
└── Ex3_Prompt_Transformation.md  # Prompt transformations
```

---

## ⚠️ Important Notes

- Write in **your own words** - AI Detector scans all submissions
- Include **analysis** for each exercise explaining what you learned
- Show **multiple inputs** for Exercises 1 and 2
- Document **differences** between weak and strong prompts

---

## 📖 Resources

- [Learn Prompting - Prompt Patterns](https://learnprompting.org/)
- [Prompt Engineering Guide](https://www.promptingguide.ai/)
- [Awesome ChatGPT Prompts](https://github.com/f/awesome-chatgpt-prompts)

---

*Generated for iihlearn2earn
