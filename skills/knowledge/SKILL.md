---
name: knowledge
description: Index and search through MyPicoClaw memory, sessions, and site monitor findings (Lightweight RAG).
metadata: {"nanobot":{"emoji":"🧠","requires":{"bins":["grep", "find"]}}}
---

# Knowledge Base (Lightweight RAG)

Search through your personal digital history.

## Search Paths
- **Memory**: `~/.mypicoclaw/workspace/memory/`
- **Sessions**: `~/.mypicoclaw/workspace/sessions/`
- **Monitor Drops**: `~/.mypicoclaw/workspace/skills/monitor/findings/` (if enabled)

## Capabilities

### 1. Keyword Search
Use `grep` to find relevant context in past sessions or memory files.
```bash
grep -rnEi "[Keyword]" ~/.mypicoclaw/workspace/
```

### 2. Chronological Retrieval
Find information by date.
```bash
find ~/.mypicoclaw/workspace/sessions/ -mtime -7 -name "*.md"
```

## Workflow
1. **Query Analysis**: Identify if the user is asking about something already discussed or "remembered".
2. **Context Retrieval**: Run search commands to find the relevant snippets.
3. **Synthesis**: Answer the user using the found context.

## Tips
- If the search returns too many results, ask for more specific keywords.
- Always check `MEMORY.md` first for high-level summaries.
