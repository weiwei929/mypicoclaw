#!/bin/bash
# 🦞 MyPicoClaw 终极自动化部署脚本
# 适用系统: Debian/Ubuntu

# 基本信息
REPO_URL="https://github.com/weiwei929/mypicoclaw"

echo "==========================================="
echo "   🦞 MyPicoClaw 部署套件 (Pre-Deployment) "
echo "==========================================="

# 1. 安装核心与技能依赖
echo "--- [1/5] 正在安装系统依赖 (Go, Git, yt-dlp, rsync, gh, jq) ---"
sudo apt update
sudo apt install -y curl ca-certificates tmux git golang-go rsync jq

# 安装 yt-dlp (推荐从 github 下载最新版以保障兼容性)
if ! command -v yt-dlp &> /dev/null; then
    sudo curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /usr/local/bin/yt-dlp
    sudo chmod a+rx /usr/local/bin/yt-dlp
fi

# 2. 克隆/更新项目
if [ -d "~/mypicoclaw" ]; then
    echo "--- [2/5] 目录已存在，正在拉取最新代码 ---"
    cd ~/mypicoclaw && git pull
else
    echo "--- [2/5] 正在克隆项目仓库 ---"
    git clone "$REPO_URL" ~/mypicoclaw
    cd ~/mypicoclaw
fi

# 3. 初始化运行环境与技能
echo "--- [3/5] 初始化工作空间与技能目录 ---"
MY_HOME="$HOME/.mypicoclaw"
mkdir -p "$MY_HOME/workspace/sessions"
mkdir -p "$MY_HOME/workspace/memory"
mkdir -p "$MY_HOME/workspace/skills"

# 自动同步内置技能到工作空间
cp -r skills/* "$MY_HOME/workspace/skills/"

# 4. 配置文件生成向导
echo "--- [4/5] 检查配置文件 ---"
CONF_FILE="$MY_HOME/config.json"
if [ ! -f "$CONF_FILE" ]; then
    cp config.example.json "$CONF_FILE"
    echo "[!] 配置文件已创建: $CONF_FILE"
    echo "[?] 请记得填入你的 Moonshot (Kimi) 和 Brave Search API Key。"
else
    echo "[OK] 配置文件已存在，跳过初始化。"
fi

# 5. 编译
echo "--- [5/5] 正在编译二进制文件 ---"
go build -o mypicoclaw ./cmd/mypicoclaw

echo "==========================================="
echo " 🎉 MyPicoClaw 准备就绪！"
echo "==========================================="
echo "💡 下一步建议 (可选)："
echo "   1. 配置大盘鸡免密: ssh-copy-id root@STORAGE_VPS_HOST"
echo "   2. 修改配置: nano $CONF_FILE"
echo "   3. 启动服务: nohup ./mypicoclaw gateway > mypicoclaw.log 2>&1 &"
echo "==========================================="
