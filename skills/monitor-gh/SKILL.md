---
name: monitor-gh
description: Monitor GitHub repositories for new commits or releases in the last 24 hours.
metadata: {"nanobot":{"emoji":"🐙","requires":{"bins":["gh", "jq"]}}}
---

# GitHub Monitor

Check for activity on repositories you care about.

## Usage

### 1. Check Recent Commits
To check commits in the last 24 hours:
```bash
# Get current time in ISO8601 minus 24h (GNU/Linux compatible)
SINCE=$(date -u --date="24 hours ago" +"%Y-%m-%dT%H:%M:%SZ")
gh api repos/[owner]/[repo]/commits --jq '.[] | select(.commit.author.date > "'"$SINCE"'") | {sha: .sha[0:7], author: .commit.author.name, message: .commit.message, date: .commit.author.date}'
```

### 2. Check Recent Releases
```bash
gh api repos/[owner]/[repo]/releases --jq '.[] | select(.published_at > "'"$SINCE"'") | {tag: .tag_name, name: .name, date: .published_at}'
```

## Targets
- **Core**: `weiwei929/mypicoclaw`
- **Other Interests**: As specified by the user.

## Setup
Requires `gh` CLI to be authenticated.
