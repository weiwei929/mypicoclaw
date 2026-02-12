#!/bin/bash
# MyPicoClaw Production Deployment Script
# Usage: bash deploy/production.sh

set -e

echo "🦞 MyPicoClaw 生产部署"
echo "======================"

# 1. Pull latest code
echo ""
echo "📥 Step 1: 拉取最新代码..."
cd /root/mypicoclaw
git pull

# 2. Build
echo ""
echo "🔨 Step 2: 编译..."
go build -p 1 -o mypicoclaw ./cmd/mypicoclaw
echo "   ✅ 编译成功"

# 3. Install systemd service
echo ""
echo "⚙️  Step 3: 安装 systemd 服务..."
cp deploy/mypicoclaw.service /etc/systemd/system/mypicoclaw.service
systemctl daemon-reload
echo "   ✅ 服务文件已安装"

# 4. Enable and start
echo ""
echo "🚀 Step 4: 启动服务..."
systemctl enable mypicoclaw
systemctl restart mypicoclaw
sleep 2

# 5. Verify
echo ""
echo "✅ Step 5: 验证状态..."
systemctl status mypicoclaw --no-pager -l

echo ""
echo "========================================="
echo "🦞 部署完成！"
echo ""
echo "常用命令："
echo "  查看状态:  systemctl status mypicoclaw"
echo "  实时日志:  journalctl -u mypicoclaw -f"
echo "  重启服务:  systemctl restart mypicoclaw"
echo "  停止服务:  systemctl stop mypicoclaw"
echo "  更新部署:  bash deploy/production.sh"
echo "========================================="
