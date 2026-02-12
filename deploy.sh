#!/bin/bash
# 🦞 mypicoclaw 自动化部署脚本
# 适用系统: Debian/Ubuntu

# 基本信息 (建议在 GitHub fork 后修改此处的 URL)
REPO_URL="https://github.com/$(git remote get-url origin | cut -d: -f2 | cut -d. -f1)"
[ -z "$REPO_URL" ] && REPO_URL="https://github.com/your-username/mymypicoclaw"

echo "--- 准备部署 mypicoclaw 从: $REPO_URL ---"

# 1. 安装基础依赖
sudo apt update && sudo apt install -y curl ca-certificates tmux git golang-go

# 2. 克隆项目
git clone "$REPO_URL" ~/mypicoclaw
cd ~/mypicoclaw

# 3. 创建运行环境
mkdir -p ~/.mypicoclaw/workspace/sessions
mkdir -p ~/.mypicoclaw/workspace/memory
mkdir -p ~/.mypicoclaw/workspace/skills/search

# 4. 初始化配置 (如果不存在)
if [ ! -f ~/.mypicoclaw/config.json ]; then
    cp config.example.json ~/.mypicoclaw/config.json
    echo "[!] 配置文件已创建在 ~/.mypicoclaw/config.json，请手动编辑填入 API Key。"
fi

# 5. 编译
echo "--- 正在编译 mypicoclaw ---"
go build -o mypicoclaw ./cmd/mypicoclaw

# 6. 完成
echo "[DONE] 部署完成！"
echo "[HINT] 运行命令启动：nohup ./mypicoclaw gateway > mypicoclaw.log 2>&1 &"
