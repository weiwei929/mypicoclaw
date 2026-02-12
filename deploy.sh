#!/bin/bash
# 🦞 PicoClaw 自动化部署脚本
# 适用系统: Debian/Ubuntu

# 基本信息 (建议在 GitHub fork 后修改此处的 URL)
REPO_URL="https://github.com/$(git remote get-url origin | cut -d: -f2 | cut -d. -f1)"
[ -z "$REPO_URL" ] && REPO_URL="https://github.com/your-username/mypicoclaw"

echo "--- 准备部署 PicoClaw 从: $REPO_URL ---"

# 1. 安装基础依赖
sudo apt update && sudo apt install -y curl ca-certificates tmux git golang-go

# 2. 克隆项目
git clone "$REPO_URL" ~/picoclaw
cd ~/picoclaw

# 3. 创建运行环境
mkdir -p ~/.picoclaw/workspace/sessions
mkdir -p ~/.picoclaw/workspace/memory
mkdir -p ~/.picoclaw/workspace/skills/search

# 4. 初始化配置 (如果不存在)
if [ ! -f ~/.picoclaw/config.json ]; then
    cp config.example.json ~/.picoclaw/config.json
    echo "[!] 配置文件已创建在 ~/.picoclaw/config.json，请手动编辑填入 API Key。"
fi

# 5. 编译
echo "--- 正在编译 PicoClaw ---"
go build -o picoclaw ./cmd/picoclaw

# 6. 完成
echo "[DONE] 部署完成！"
echo "[HINT] 运行命令启动：nohup ./picoclaw gateway > picoclaw.log 2>&1 &"
