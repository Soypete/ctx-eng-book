# Code Audit Skill

## Purpose
Audit and validate code examples in markdown files for a technical book. Ensures all code blocks are syntactically correct and executable.

## Trigger
Use this skill when:
- Writing or editing chapters with code examples
- Running a full book audit
- The user asks to validate code examples

## Workflow

### 1. Find Code Blocks
Search markdown files for fenced code blocks:
```
```language
code here
```
```

### 2. Extract Code
For each code block, extract:
- Language identifier (go, python, js, etc.)
- The code content
- File context (which .md file, chapter, module)

### 3. Validate Per Language

**Go:**
- Run `gofmt -l` to check syntax
- Try `go build` with the code (may need stub imports)

**Python:**
- Run `python -m py_compile` or `ruff check`

**JavaScript/TypeScript:**
- Run `node --check` or `tsc --noEmit`

**Shell:**
- Run `shellcheck` (or manual check for basic syntax)

### 4. Report Results

Output format:
```
## Code Audit Report

### Chapter: chapter-name

**File:** path/to/file.md

| Module | Language | Status | Issues |
|--------|----------|--------|--------|
| module-slug | go | ✅ VALID | - |
| module-slug | python | ❌ INVALID | syntax error at line 5 |

**Details:**
- `module-slug`: syntax error at line 5: unexpected token
```

### 5. Fix Suggestions
For each issue, suggest:
- The specific line with the error
- What the fix should be
- Example corrected code

## Tools Available
- grep: find code blocks in markdown
- read: read markdown files
- bash: run validators (gofmt, python, node, etc.)
- write: create temporary test files if needed

## Output
Return a structured audit report with:
1. Summary: total blocks, valid count, invalid count
2. Per-file breakdown
3. Detailed error messages with line numbers
4. Suggested fixes

## Book Context
This is "Context Engineering: Building Reliable AI Systems" - code examples should be:
- Syntactically correct
- idiomatic for the language
- Well-commented for educational value
- Match the book's examples directory if referenced