---
name: monitor
description: Monitor specific websites (forums, video sites) for new content/files in the last 24 hours.
metadata: {"nanobot":{"emoji":"📡","requires":{"tools":["web_fetch", "web_search"]}}}
---

# Site Monitor

Use this skill to track updates on specific websites, especially forums and video platforms.

## Capabilities

- **Visit & Parse**: Fetch the homepage or "New Post" page of a site.
- **Timestamp Filtering**: Identify content posted within the last 24 hours.
- **List Aggregation**: Extract titles, links, and uploaders.

## Workflow

1. **Target Identification**: If the user provides a site name but not a specific URL, use `web_search` to find the "Latest" or "Archive" page.
2. **Content Fetching**: Use `web_fetch` to retrieve the page source.
3. **AI Filtering**: Analyze the text to find items from the last 24 hours.
   - Pay attention to relative timestamps like "2h ago", "12 hours ago", "Yesterday".
   - Compare absolute dates with the current time (provided in your system context).
4. **Report**: Return a clean list of findings including titles and URLs.

## Common Targets
- **Forums**: Reddit, Discuz, vBulletin, NodeBB.
- **Video Sites**: YouTube, Bilibili, PeerTube.
- **File Sharing**: Archive.org, specialized download forums.

## Tips
- For large forums, look for "New Posts" or "Latest Threads" links first.
- If a site is blocked or requires JavaScript that `web_fetch` can't handle, inform the user you might need a specialized scraper or RSS feed.
