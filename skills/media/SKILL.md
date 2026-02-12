---
name: media
description: Download videos or files using yt-dlp and optionally transfer them to a configured storage VPS.
metadata: {"nanobot":{"emoji":"📽️","requires":{"bins":["yt-dlp", "rsync", "ssh"]}}}
---

# Media Downloader

Use this skill to download videos from URLs or direct files, with support for remote storage on your storage VPS.

## Target Storage
- **Host**: Configured via `MYPICOCLAW_STORAGE_VPS_HOST` or `config.json` → `storage_vps.host`
- **User**: Configured via `MYPICOCLAW_STORAGE_VPS_USER` (default: `root`)
- **Default Remote Path**: Configured via `MYPICOCLAW_STORAGE_VPS_PATH` (default: `/mnt/storage/pikpak/picoclaw_downloads`)

## Capabilities

1. **Local Download**: Use `yt-dlp` for videos or `curl -O` for direct files.
2. **Remote Transfer**: Use `rsync` to move finished downloads to the storage VPS.
3. **Status Tracking**: Report progress and final location.

## Instructions for Agent

### 1. Simple Download
If the user says "Download this video", run:
```bash
yt-dlp -f "bestvideo+bestaudio/best" --no-mtime [URL]
```

### 2. Download to Storage VPS
If the user specifies "Save to storage" or "Save to Big Chicken", follow these steps:
- **Step A: Download Locally**
  ```bash
  yt-dlp -o "%(title)s.%(ext)s" [URL]
  ```
- **Step B: Transfer via Rsync**
  Use the configured storage VPS host from your config:
  ```bash
  rsync -avz --remove-source-files [FILENAME] $STORAGE_USER@$STORAGE_HOST:$STORAGE_PATH/
  ```

## Setup Requirements (User Action)
- Install `yt-dlp` and `rsync` on the gateway VPS.
- **SSH Key**: Ensure the Gateway VPS can SSH into the storage VPS without a password prompt.
  - Run `ssh-keygen` (if not exists).
  - Run `ssh-copy-id $STORAGE_USER@$STORAGE_HOST`.

## Tips
- For YouTube, use `--proxy` if the VPS is in a restricted region.
- Use `--extract-audio --audio-format mp3` if the user only wants the sound.

## 🛡️ Advanced: Bypassing Anti-Leech (防盗链/地区限制)

If a download fails due to "Forbidden" or "Sign in to confirm your age", use these strategies:

### 1. Cookies (The most powerful way)
- **Problem**: Captcha or Login required.
- **Solution**: Export cookies from your browser (using extensions like "Get cookies.txt LOCALLY") and upload to `~/.mypicoclaw/cookies.txt`.
- **Command**:
  ```bash
  yt-dlp --cookies ~/.mypicoclaw/cookies.txt [URL]
  ```

### 2. User-Agent & Referer
- **Problem**: Basic anti-bot checks.
- **Example (Bilibili)**:
  ```bash
  yt-dlp --user-agent "Mozilla/5.0 ..." --referer "https://www.bilibili.com" [URL]
  ```

### 3. Region Bypassing
- Use a proxy if the video is restricted to a certain country:
  ```bash
  yt-dlp --proxy "http://user:pass@host:port" [URL]
  ```

### 4. Direct File Leech (Aria2)
For direct file links that check referers, use `aria2c`:
```bash
aria2c --referer="[URL]" "[FILE_LINK]"
```
