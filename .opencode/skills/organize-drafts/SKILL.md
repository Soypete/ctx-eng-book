---
name: organize-drafts
description: Analyzes draft content against the book outline and organizes it into the correct chapter/module files. Use this skill when the user has draft notes, writing session output, or raw content that needs to be placed in the right location. Also use when the user asks "where should I write next" or "what's the current state of the book?"
---

# Organize Drafts Skill

## Purpose

This skill manages the book's draft structure by:
1. Analyzing draft content against the outline
2. Placing content in the correct chapter/module files
3. Suggesting where to continue writing

## Trigger

Use when the user:
- Says "organize my drafts" or "where should I write next"
- Has notes/drafts that need placement
- Wants to see current book state
- Asks for a writing suggestion

## Workflow

### Step 1: Read the Outline

Read `/Users/soypete/code/misc/ctx-eng-book/chapters.md` to get the current 18-chapter structure.

### Step 2: Scan Existing Drafts

Find all draft files in:
```
cmd/authorpedro/books/ctx-eng-book/chapters/*/modules/*.md
```

List all existing modules to understand what's already drafted.

### Step 3: Analyze Content

For new content to organize:
- Read the content
- Match it against chapter topics in chapters.md
- Identify the best chapter/module placement

### Step 4: Suggest Next Action

Provide the user with:
1. **Current state**: What's drafted vs what's empty
2. **Recommendation**: Where to write next based on:
   - What's already drafted
   - What's blocked (needs research)
   - Dependencies between chapters
3. **Organization**: If placing new content, confirm the location

## Output Format

```
## Book State

| Chapter | Status | Notes |
|---------|--------|-------|
| Ch 1: Every Failure Is a Context Failure | draft | has 1 module |
| Ch 2: AI Is a Systems Problem | empty | - |
...

## Recommendation

Next: Chapter {N} — {Title}

Reason: {why this chapter is ready}

## Actions Needed

- [ ] {any prep work needed}
```

## File Placement

When organizing into chapters, create:
```
cmd/authorpedro/books/ctx-eng-book/chapters/{chapter-slug}/modules/{module-slug}.md
```

Chapter slugs:
- part-1-why-ai-systems-fail/chapter-1-every-failure-is-a-context-failure
- part-1-why-ai-systems-fail/chapter-2-ai-is-a-systems-problem
- part-2-how-models-use-context/chapter-3-attention
- part-2-how-models-use-context/chapter-4-in-context-learning
- part-2-how-models-use-context/chapter-5-tool-use
- part-3-context-is-data/chapter-6-memory-is-a-database
- part-3-context-is-data/chapter-7-context-is-a-query
- part-3-context-is-data/chapter-8-knowledge-graphs
- part-3-context-is-data/chapter-9-retrieval-beyond-vectors
- part-4-governance/chapter-10-personalization
- part-4-governance/chapter-11-stop-giving-agents-permissions
- part-4-governance/chapter-12-unix-philosophy
- part-5-orchestration/chapter-13-agents-are-workflows
- part-5-orchestration/chapter-14-cost-of-context
- part-5-orchestration/chapter-15-when-context-stops-working
- part-6-reliability/chapter-16-observability
- part-6-reliability/chapter-17-evaluating-ai-systems
- part-6-reliability/chapter-18-context-platform

## Handling Draft Content

If the user provides draft text during a /write session:
1. Summarize the key points
2. Ask which chapter it belongs to (or suggest based on analysis)
3. Create the module file if needed
4. Insert the content with minimal formatting

## No Existing Drafts?

If the chapter has no modules yet, create a placeholder:
```markdown
# {Module Title}

## Outline

- {topic 1}
- {topic 2}
- {topic 3}

## Status

draft | outline | empty

## Notes

{any notes about what this module needs}
```

---

After organizing, tell the user:
- What's been placed where
- The recommended next step
- Any gaps that need research before writing