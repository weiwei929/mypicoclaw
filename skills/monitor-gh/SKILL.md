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
# Get current time in ISO8601 minus 24h
SINCE=$(date -u -v-24H +"%Y-%m-%dT%H:%M:%SZ" 2>/dev/null || date -u --date="24 hours ago" +"%Y-%m-%dT%H:%M:%SZ")
gh api repos/[owner]/[repo]/commits --ff since=$SINCE --jq '.[] | {sha: .sha, author: .commit.author.name, message: .commit.message, date: .commit.author.date}'
```

### 2. Check Recent Releases
```bash
gh api repos/[owner]/[repo]/releases --jq '.[] | select(.published_at > "'$SINCE'") | {tag: .tag_name, name: .name, date: .published_at}'
```

## Targets
- **Core**: `weiwei929/mypicoclaw`
- **Other Interests**: As specified by the user.

## Setup
Requires `gh` CLI to be authenticated.
