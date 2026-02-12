---
name: media
description: Download videos or files using yt-dlp and optionally transfer them to the storage VPS (STORAGE_VPS_HOST).
metadata: {"nanobot":{"emoji":"📽️","requires":{"bins":["yt-dlp", "rsync", "ssh"]}}}
---

# Media Downloader

Use this skill to download videos from URLs or direct files, with support for remote storage on your "Big Chicken" (大盘鸡) VPS.

## Target Storage
- **Host**: `STORAGE_VPS_HOST`
- **Default Remote Path**: `/mnt/storage/pikpak/picoclaw_downloads`
- **Storage Info**: ~10TB available via PikPak mount.

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
  ```bash
  rsync -avz --remove-source-files [FILENAME] root@STORAGE_VPS_HOST:/mnt/storage/pikpak/picoclaw_downloads/
  ```

## Setup Requirements (User Action)
- Install `yt-dlp` and `rsync` on the gateway VPS.
- **SSH Key**: Ensure the Gateway VPS can SSH into `root@STORAGE_VPS_HOST` without a password prompt.
  - Run `ssh-keygen` (if not exists).
  - Run `ssh-copy-id root@STORAGE_VPS_HOST`.

## Tips
- For YouTube, use `--proxy` if the VPS is in a restricted region.
- Use `--extract-audio --audio-format mp3` if the user only wants the sound.
