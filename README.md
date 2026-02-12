<div align="center">
<img src="assets/logo.jpg" alt="MyPicoClaw" width="512">

<h1>MyPicoClaw: Go 语言编写的超高效 AI 助手</h1>

<h3>$10 硬件 · 10MB 内存 · 1秒启动 · 皮皮虾，我们走！</h3>

<p>
<img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/badge/Arch-x86__64%2C%20ARM64%2C%20RISC--V-blue" alt="Hardware">
<img src="https://img.shields.io/badge/license-MIT-green" alt="License">
</p>

</div>

---

🦐 **MyPicoClaw** 是一款受 [nanobot](https://github.com/HKUDS/nanobot) 启发、完全由 Go 语言重写的超轻量级个人 AI 助手。它通过"自我进化"过程构建——由 AI 代理驱动了整个架构迁移和代码优化。

⚡️ **在 $10 的硬件上以 <10MB 内存运行**：比 OpenClaw 节省 99% 的内存，比 Mac mini 便宜 98%！

<table align="center">
  <tr align="center">
    <td align="center" valign="top">
      <p align="center">
        <img src="assets/MyPicoClaw_mem.gif" width="360" height="240">
      </p>
    </td>
    <td align="center" valign="top">
      <p align="center">
        <img src="assets/licheervnano.png" width="400" height="240">
      </p>
    </td>
  </tr>
</table>

## 📢 新闻
2026-02-12 🛡️ v1.1 稳定版：新增 API 重试机制、优雅错误降级、systemd 生产部署支持。

2026-02-09 🎉 MyPicoClaw 正式发布！仅用 1 天时间开发，为 $10 级硬件带来不到 10MB 内存占用的 AI 代理。🦐 皮皮虾，我们走！

## ✨ 特性

🪶 **超轻量级**：内存占用 <10MB —— 比常规核心功能缩小 99%。

💰 **极致低成本**：可在 $10 的硬件上高效运行 —— 比 Mac mini 便宜 98%。

⚡️ **闪电速度**：启动速度快 400 倍，即使在 0.6GHz 单核环境下也能在 1 秒内启动。

🌍 **真正的便携性**：支持 RISC-V、ARM 和 x86 的单一自包含二进制文件，一键运行！

🤖 **AI 自驱开发**：自主 Go 原生实现 —— 95% 的核心代码由 Agent 生成。

🛡️ **生产级容错**：API 过载自动重试（指数退避），异常优雅降级，永不沉默。

|  | OpenClaw  | NanoBot | **MyPicoClaw** |
| --- | --- | --- |--- |
| **语言** | TypeScript | Python | **Go** |
| **内存占用** | >1GB |>100MB| **< 10MB** |
| **启动时间**</br>(0.8GHz 核心) | >500s | >30s |  **<1s** |
| **成本** | Mac Mini 599$ | 大多数 Linux SBC </br>~50$ |**任何 Linux 开发板**</br>**低至 10$** |
<img src="assets/compare.jpg" alt="MyPicoClaw" width="512">

## 🦾 演示
### 🛠️ 标准助手工作流
<table align="center">
  <tr align="center">
    <th><p align="center">🧩 全栈工程师</p></th>
    <th><p align="center">🗂️ 日志与计划管理</p></th>
    <th><p align="center">🔎 联网搜索与学习</p></th>
  </tr>
  <tr>
    <td align="center"><p align="center"><img src="assets/MyPicoClaw_code.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/MyPicoClaw_memory.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/MyPicoClaw_search.gif" width="240" height="180"></p></td>
  </tr>
  <tr>
    <td align="center">开发 • 部署 • 扩展</td>
    <td align="center">调度 • 自动化 • 记忆</td>
    <td align="center">发现 • 洞察 • 趋势</td>
  </tr>
</table>

### 🐜 创新的低功耗部署
MyPicoClaw 几乎可以部署在任何 Linux 设备上！

- $9.9 [LicheeRV-Nano](https://www.aliexpress.com/item/1005006519668532.html) E 或 W 版，极致迷你的家庭助手。
- $30~50 [NanoKVM](https://www.aliexpress.com/item/1005007369816019.html)，或 $100 [NanoKVM-Pro](https://www.aliexpress.com/item/1005010048471263.html)，用于自动化服务器维护。
- $50 [MaixCAM](https://www.aliexpress.com/item/1005008053333693.html) 或 $100 [MaixCAM2](https://www.kickstarter.com/projects/zepan/maixcam2) 智能监控。

🌟 更多部署案例等你探索！

## 📦 安装

### 使用预编译二进制文件安装

从 [Release](https://github.com/weiwei929/mypicoclaw/releases) 页面下载适合你平台的固件。

### 从源码安装（推荐用于开发，获取最新功能）

```bash
git clone https://github.com/weiwei929/mypicoclaw.git
cd mypicoclaw
go build -o mypicoclaw ./cmd/mypicoclaw
```

## 🚀 VPS 生产部署

> [!NOTE]
> MyPicoClaw 使用 Telegram Polling 模式，**不需要域名、Caddy 或开放任何端口**。

### 部署前准备

1. **API Key**：
   - **Moonshot Global**: [获取地址](https://platform.moonshot.ai) (默认模型)
   - **Brave Search** (可选): [获取地址](https://brave.com/search/api)
2. **远程节点配对** (如有多台 VPS)：
   ```bash
   ssh-keygen -t ed25519 -N ""
   ssh-copy-id root@<远程IP>
   ```

### 一键部署

```bash
git clone https://github.com/weiwei929/mypicoclaw.git
cd mypicoclaw
bash deploy/production.sh
```

部署脚本会自动完成：`git pull` → 编译 → 安装 systemd 服务 → 启动 → 验证状态。

### 后续更新

```bash
bash deploy/production.sh   # 自动 pull → 编译 → 重启
```

### 🚀 快速开始

> [!TIP]
> 在 `~/.mypicoclaw/config.json` 中设置你的 API Key。
> 获取 Key：[Moonshot Global](https://platform.moonshot.ai) (Kimi) · [OpenRouter](https://openrouter.ai/keys) (LLM) · [智谱](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) (LLM)
> 联网搜索是 **可选** 的 - 获取免费的 [Brave Search API](https://brave.com/search/api) (每月 2000 次免费查询)

**1. 初始化**

```bash
./mypicoclaw onboard
```

**2. 配置** (`~/.mypicoclaw/config.json`)

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.mypicoclaw/workspace",
      "model": "moonshot-v1-8k",
      "max_tokens": 8192,
      "temperature": 0.3,
      "max_tool_iterations": 20
    }
  },
  "providers": {
    "moonshot": {
      "api_key": "YOUR_MOONSHOT_API_KEY",
      "api_base": "https://api.moonshot.ai/v1"
    }
  },
  "tools": {
    "web": {
      "search": {
        "api_key": "YOUR_BRAVE_API_KEY",
        "max_results": 5
      }
    }
  }
}
```

**3. 开始聊天**

```bash
./mypicoclaw agent -m "2+2 等于几？"
```

就是这样！你只需 2 分钟就能拥有一个可以工作的 AI 助手。

---

## 💬 聊天应用支持

通过 Telegram、Discord 或飞书与你的 MyPicoClaw 对话。

| 渠道 | 设置难度 |
|---------|-------|
| **Telegram** | 简单 (只需要一个 Token) |
| **Discord** | 简单 (Bot Token + Intents) |
| **飞书 (Feishu)** | 简单 (WebSocket 模式) |
| **QQ** | 简单 (AppID + AppSecret) |
| **钉钉 (DingTalk)** | 中等 (应用凭证) |

<details>
<summary><b>Telegram</b> (推荐)</summary>

**1. 创建机器人**
- 在 Telegram 搜索 `@BotFather`
- 发送 `/newbot`，按提示操作
- 复制 Token

**2. 配置**
```json
{
  "channels": {
    "telegram": {
      "enabled": true,
      "token": "你的_BOT_TOKEN",
      "allowFrom": ["你的_USER_ID"]
    }
  }
}
```
> 在 Telegram 咨询 `@userinfobot` 获取你的用户 ID。

**3. 运行**
```bash
./mypicoclaw gateway
```
</details>

## ⚙️ 详细配置

配置文件路径：`~/.mypicoclaw/config.json`

### 工作空间结构

MyPicoClaw 在你配置的工作空间（默认 `~/.mypicoclaw/workspace`）中存储数据：

```
~/.mypicoclaw/workspace/
├── sessions/          # 对话会话与历史记录
├── memory/           # 长期记忆 (MEMORY.md)
├── cron/             # 定时任务数据库
├── skills/           # 自定义技能
├── AGENTS.md         # Agent 行为指南
├── IDENTITY.md       # Agent 身份定义
├── SOUL.md           # Agent 灵魂/个性定义
├── TOOLS.md          # 工具描述
└── USER.md           # 用户偏好信息
```

### 供应商支持 (Providers)

> [!NOTE]
> Groq 提供免费的 Whisper 语音转文字服务。如果配置了 Groq Key，Telegram 的语音消息将自动转换为文字。

| 供应商 | 用途 | 获取 Key |
|----------|---------|-------------|
| `moonshot` | LLM (Kimi 国际版直连) | [platform.moonshot.ai](https://platform.moonshot.ai) |
| `gemini` | LLM (Gemini 直连) | [aistudio.google.com](https://aistudio.google.com) |
| `zhipu` | LLM (智谱直连) | [bigmodel.cn](bigmodel.cn) |
| `openrouter` | LLM (推荐，支持所有模型) | [openrouter.ai](https://openrouter.ai) |
| `groq` | LLM + **语音转文字** (Whisper) | [console.groq.com](https://console.groq.com) |

## 📚 常用命令参考

### 应用命令

| 命令 | 描述 |
|---------|-------------|
| `./mypicoclaw onboard` | 初始化配置与工作空间 |
| `./mypicoclaw agent -m "..."` | 与 Agent 进行单次对话 |
| `./mypicoclaw agent` | 进入交互式对话模式 |
| `./mypicoclaw gateway` | 启动网关（用于各聊天渠道） |
| `./mypicoclaw status` | 查看状态 |
| `./mypicoclaw cron list` | 列出所有定时任务 |
| `./mypicoclaw cron add ...` | 添加定时任务 |

### 运维命令 (systemd 部署后)

| 命令 | 描述 |
|---------|-------------|
| `bash deploy/production.sh` | 一键更新部署（pull + build + restart） |
| `systemctl status mypicoclaw` | 查看服务状态 |
| `journalctl -u mypicoclaw -f` | 实时查看日志 |
| `systemctl restart mypicoclaw` | 重启服务 |
| `systemctl stop mypicoclaw` | 停止服务 |

---

## 🤝 贡献与路线图

欢迎 PR！代码库保持简洁易读。🤗

<img src="assets/wechat.png" alt="MyPicoClaw" width="512">

## 🐛 常见问题

### 联网搜索提示 "API 配置问题"
如果你还没有配置搜索 API Key，这是正常现象。MyPicoClaw 会提供参考链接供你手动搜索。
配置方法：
1. 在 [Brave Search API](https://brave.com/search/api) 获取免费 Key。
2. 填入 `config.json` 的 `tools.web.search.api_key` 中。

### API 报 "engine_overloaded" 错误
MyPicoClaw 内置了自动重试机制（指数退避 2s→4s→8s），大部分临时过载会自动恢复。如果持续失败，会返回友好的中文提示而不是沉默。

### 清除会话数据
如果遇到奇怪的 `tool_call_id not found` 错误，清除历史会话即可：
```bash
rm -rf ~/.mypicoclaw/sessions/*
systemctl restart mypicoclaw
```

---

## 📝 API 供应商对比

| 服务 | 免费档位 | 适用场景 |
|---------|-----------|-----------| 
| **Moonshot** | 适配国际版 | 强力中文/英文支持 |
| **OpenRouter** | 200K tokens/月 | 尝试各种模型 (Claude, GPT-4 等) |
| **智谱 (Zhipu)** | 200K tokens/月 | 中国区访问流畅 |
| **Brave Search** | 2000 次/月 | 联网获取实时信息 |
| **Groq** | 有免费档位 | 极速推理 (Llama 3 等) |
