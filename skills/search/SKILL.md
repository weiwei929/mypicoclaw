---
name: search
description: Search the web for current information, news, and deep-dive research.
homepage: https://brave.com/search/api
metadata: {"nanobot":{"emoji":"🔎","requires":{"tools":["web_search", "web_fetch"]}}}
---

# Web Search

Use this skill to find current information from the internet. You have access to both a high-level search engine and a deep-page fetcher.

## Workflow

1.  **Search**: Use `web_search` to find relevant URLs and summaries.
2.  **Fetch**: If a search result looks promising but the snippet is too short, use `web_fetch` to read the full content of the page.
3.  **Synthesize**: Combine information from multiple sources to provide a comprehensive answer.

## Tools

### web_search
Search the web via Brave Search API.
- `query`: The search term.
- `count`: Number of results (default 5, max 10).

### web_fetch
Fetch and extract raw text from a specific URL.
- `url`: The full URL to fetch.
- `maxChars`: Limit output size (default 50,000).

## Tips
- Be specific with your search queries.
- If the first search doesn't yield good results, try rephrasing or using specialized keywords.
- Always cite your sources with URLs.
