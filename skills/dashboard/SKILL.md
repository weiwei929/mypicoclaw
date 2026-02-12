---
name: dashboard
description: Show system health and resource usage for the Gateway VPS and a configured Storage VPS.
metadata: {"nanobot":{"emoji":"📊","requires":{"bins":["ssh", "df", "free", "uptime"]}}}
---

# System Dashboard

Check the pulses of your "Little Chickens" (VPS nodes).

## Monitored Nodes
1. **Gateway Node**: Current local host.
2. **Storage Node (Big Chicken)**: Configured via `MYPICOCLAW_STORAGE_VPS_HOST` or `config.json` → `storage_vps.host`

## Usage

### 1. Summary Report
When asked "How are my servers doing?" or "Dashboard", perform the following:

- **Local (Gateway) Check**:
  ```bash
  uptime -p
  free -h | grep Mem
  df -h / | tail -1
  ```
- **Remote (Storage) Check** (use configured host):
  ```bash
  ssh $STORAGE_USER@$STORAGE_HOST "uptime -p; free -h | grep Mem; df -h /mnt/storage/pikpak | tail -1"
  ```

### 2. Detailed Storage Check
When asked about disk space on the Big Chicken:
```bash
ssh $STORAGE_USER@$STORAGE_HOST "df -h / /mnt/storage/pikpak"
```

## AI Output Format
Present the data in a clean Markdown table:

| Node | Uptime | Memory Usage | Disk (Used/Total) |
|------|--------|--------------|-------------------|
| Gateway | ... | ... | ... |
| Storage | ... | ... | ... |

## Setup Note
Requires SSH passwordless login to the configured storage VPS. Set the host via `MYPICOCLAW_STORAGE_VPS_HOST` env var.
