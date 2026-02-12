---
name: skill-creator
description: Create or update Skills for MyPicoClaw. Use when the user asks to add a new capability, define a new workflow, or package a new skill.
---

# Skill Creator

Create new skills to extend MyPicoClaw's capabilities.

## Quick Install Instructions

When the user asks to "create a new skill", follow this pattern:

### Step 1: Create the directory
```bash
mkdir -p skills/[skill-name]
```

### Step 2: Write SKILL.md
Use `write_file` to create `skills/[skill-name]/SKILL.md` with this template:

```markdown
---
name: [skill-name]
description: [What this skill does and when to use it]
metadata: {"nanobot":{"emoji":"🔧","requires":{"bins":["tool1"]}}}
---

# [Skill Name]

[Instructions for the agent on how to execute this skill]

## Usage
[Commands or workflows]
```

### Step 3: Confirm
Tell the user the skill is created and will be available after restart.

## Key Principles

1. **Concise is Key**: Only include information the agent doesn't already know.
2. **Frontmatter is Critical**: The `description` field is the primary trigger — make it comprehensive.
3. **Use exec tool**: Skills should instruct the agent to use shell commands via the `exec` tool.
4. **Optional resources**: Add `scripts/`, `references/`, or `assets/` subdirectories only if needed.

## Skill Naming
- Use lowercase letters, digits, and hyphens only (e.g., `my-skill`)
- Keep names short and descriptive
- Name the folder exactly after the skill name
