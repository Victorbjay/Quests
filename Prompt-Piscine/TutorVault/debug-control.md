# Tutor Vault: Debug Control - Complete Guide

## 📚 Overview

This guide teaches you how to complete the **Debug Control** task from the 01-Edu Prompt Piscine. You'll learn to detect and fix AI hallucinations, anchor responses to data, and use delimiters effectively.

---

## 🎯 Learning Objectives

- Spot vague or fabricated information in AI outputs
- Rewrite prompts to add clear constraints
- Use delimiters and formatting to keep responses on topic
- Prevent AI hallucinations

---

## 📋 Exercise Breakdown

### Exercise 0: Environment Setup

**Task:** Set up Python.

**Steps:**
1. Check Python version: `python3 --version`
2. Create Hello World:

```python
print("Hello, World!")
```

**Expected Output:**
```
Hello, World!
```

---

### Exercise 1: Reducing False or Invented Facts

**Goal:** Prevent AI from making up information (hallucinations).

#### The Problem

Weak prompts that ask for fictional/unverified information cause AI to hallucinate.

**Weak Prompt:**
```
"Give the capitals of five imaginary countries."
```

**Issue:** AI invents fake countries and capitals that don't exist.

**Weak Output:**
- Zemaria → New Capital City
- Fantasiland → Dreamville
- Mirages → Oasis Town

All fabricated!

#### The Solution

Add "real" or "actual" to anchor to facts.

**Strong Prompt:**
```
"List five real countries and their capitals."
```

**Strong Output:**
- France → Paris
- Japan → Tokyo
- Brazil → Brasília
- Nigeria → Abuja
- Australia → Canberra

✅ Verifiable, factual information

---

### Exercise 2: Adding Source or Data Constraints

**Goal:** Force AI to use only provided material, not its training data.

#### Without Constraint

**Prompt:**
> "What does the article say about climate change?"

**Problem:** AI might answer from memory, not the article.

#### With Constraint

**Prompt:**
> "What does the article say about climate change? Use only the text provided below. Cite the source in your answer."

**Provided Text:**
> "Climate change is primarily caused by human activities, particularly the burning of fossil fuels. Carbon dioxide emissions have increased by 50% since the pre-industrial era."

**Constrained Output:**
> The article states that climate change is primarily caused by human activities, especially burning fossil fuels. According to the text, carbon dioxide emissions have increased by 50% since pre-industrial times. [Source: Provided text]

---

#### Key Phrases for Constraints

| Phrase | Purpose |
|--------|---------|
| "Use only the text provided" | Anchors to source |
| "Cite the source" | Requires attribution |
| "Do not use external knowledge" | Prevents training data leakage |
| "Based only on this information" | Keeps focus narrow |

---

### Exercise 3: Using Delimiters to Define Context

**Goal:** Use delimiters to clearly mark what AI should focus on.

#### What Are Delimiters?

Markers that define boundaries for AI context:

```
### text ###
```

```markdown
---
text
---
```

```
```
text
```
```

#### Without Delimiters

**Prompt:**
> "What is this text about?"

**Text:**
> "I love reading. Python is a programming language. Blue is my favorite color."

**Issue:** Mixed topics confuse the AI.

---

#### With Delimiters

**Prompt:**
> "What programming language is mentioned? Answer using only the text inside ###."

**Delimited Text:**
###
Python is a programming language.
###

**Output:**
> Python is the programming language mentioned.

---

#### Why Delimiters Work

1. **Clear boundaries** - AI knows exact context
2. **No context pollution** - Other text excluded
3. **Consistent results** - Same input = same focus
4. **Easy debugging** - See exactly what AI received

---

## 📊 Comparison Table

| Technique | Weak Version | Strong Version | Effect |
|-----------|-------------|---------------|--------|
| Fact Checking | "imaginary countries" | "real countries" | Prevents hallucination |
| Source | No constraint | "Use only provided text" | Anchors to source |
| Context | No delimiters | "### text ###" | Focuses AI |

---

## 🔑 Key Takeaways

1. **Hallucinations are real** - AI can make up facts when not constrained
2. **Be specific** - "real countries" vs "countries"
3. **Cite sources** - "Use only the text provided"
4. **Use delimiters** - Clear boundaries prevent confusion
5. **Test your prompts** - Always verify AI output accuracy

---

## 📁 File Structure

Your submission should look like:

```
debug-control/
├── Ex0_Env_Libraries.py           # Python setup
├── Ex1_Reducing_False_Facts.md     # Hallucination prevention
├── Ex2_Source_Constraints.md       # Source anchoring
└── Ex3_Using_Delimiters.md        # Context delimiting
```

---

## ⚠️ Important Notes

- Write in **your own words** - AI Detector scans all submissions
- For Exercises 2 & 3: Show **with AND without** constraints
- Document the **differences** in each comparison
- Include **analysis** explaining what you learned

---

## 📖 Resources

- [Learn Prompting - Avoiding Hallucinations](https://medium.com/@bijit211987/advanced-prompt-engineering-for-reducing-hallucination-bb2c8ce62fc6)
- [Prompt Engineering Guide - Robustness](https://www.promptlayer.com/glossary/prompt-robustness)
- [Preventing LLM Hallucination](https://cobusgreyling.medium.com/preventing-llm-hallucination-with-contextual-prompt-engineering-an-example-from-openai-7e7d58736162)

---

## 🎓 Quick Reference

```
# Prevent Hallucinations
"List [real/actual] countries..."
"Give me verified information about..."

# Anchor to Source
"Use only the text provided below"
"Cite the source in your answer"
"Based only on this information"

# Use Delimiters
### context ###
```context```
---context---
```

---

*Generated by iihlearn2earn for 01-Edu Learn2Earn*
