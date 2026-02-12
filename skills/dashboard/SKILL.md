---
name: dashboard
description: Show system health and resource usage for the Gateway VPS and the Storage VPS (STORAGE_VPS_HOST).
metadata: {"nanobot":{"emoji":"📊","requires":{"bins":["ssh", "df", "free", "uptime"]}}}
---

# System Dashboard

Check the pulses of your "Little Chickens" (VPS nodes).

## Monitored Nodes
1. **Gateway Node**: Current local host.
2. **Storage Node (Big Chicken)**: `STORAGE_VPS_HOST`

## Usage

### 1. Summary Report
When asked "How are my servers doing?" or "Dashboard", perform the following:

- **Local (Gateway) Check**:
  ```bash
  uptime -p
  free -h | grep Mem
  df -h / | tail -1
  ```
- **Remote (Storage) Check**:
  ```bash
  ssh root@STORAGE_VPS_HOST "uptime -p; free -h | grep Mem; df -h /mnt/storage/pikpak | tail -1"
  ```

### 2. Detailed Storage Check
When asked about disk space on the Big Chicken:
```bash
ssh root@STORAGE_VPS_HOST "df -h / /mnt/storage/pikpak"
```

## AI Output Format
Present the data in a clean Markdown table:

| Node | Uptime | Memory Usage | Disk (Used/Total) |
|------|--------|--------------|-------------------|
| Gateway | ... | ... | ... |
| Storage | ... | ... | ... |

## Setup Note
Requires SSH passwordless login to `root@STORAGE_VPS_HOST`.
