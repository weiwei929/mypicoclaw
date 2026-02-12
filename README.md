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

🦐 **MyPicoClaw** 是一款受 [nanobot](https://github.com/HKUDS/nanobot) 启发、完全由 Go 语言重写的超轻量级个人 AI 助手。它通过“自我进化”过程构建——由 AI 代理驱动了整个架构迁移和代码优化。

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
2026-02-09 🎉 MyPicoClaw 正式发布！仅用 1 天时间开发，为 $10 级硬件带来不到 10MB 内存占用的 AI 代理。🦐 皮皮虾，我们走！

## ✨ 特性

🪶 **超轻量级**：内存占用 <10MB —— 比常规核心功能缩小 99%。

💰 **极致低成本**：可在 $10 的硬件上高效运行 —— 比 Mac mini 便宜 98%。

⚡️ **闪电速度**：启动速度快 400 倍，即使在 0.6GHz 单核环境下也能在 1 秒内启动。

🌍 **真正的便携性**：支持 RISC-V、ARM 和 x86 的单一自包含二进制文件，一键运行！

🤖 **AI 自驱开发**：自主 Go 原生实现 —— 95% 的核心代码由 Agent 生成。

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
make deps

# 编译，无需安装
make build

# 为所有平台编译
make build-all

# 编译并安装
make install
```

## 🚀 部署预备清单 (Pre-Deployment Checklist)

为了实现“一气呵成”的部署体验，请在开始前确认以下事项：

1. **域名准备**：如果你打算使用 Caddy 访问，请确保域名已指向主 VPS。
2. **API Key**：
   - **Moonshot Global**: [获取地址](https://platform.moonshot.ai) (目前默认模型)
   - **Brave Search**: [获取地址](https://brave.com/search/api)
3. **大盘鸡 (STORAGE_VPS_HOST) 配对**：
   - 在主 VPS 上运行 `ssh-keygen`。
   - 运行 `ssh-copy-id root@STORAGE_VPS_HOST` 实现免密。
   - 确保大盘鸡已安装 `rsync`。

### 🚀 一键安装命令

```bash
curl -sSL https://raw.githubusercontent.com/weiwei929/mypicoclaw/main/deploy.sh | bash
```

### 🚀 快速开始

> [!TIP]
> 在 `~/.mypicoclaw/config.json` 中设置你的 API Key。
> 获取 Key：[Moonshot Global](https://platform.moonshot.ai) (Kimi) · [OpenRouter](https://openrouter.ai/keys) (LLM) · [智谱](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) (LLM)
> 联网搜索是 **可选** 的 - 获取免费的 [Brave Search API](https://brave.com/search/api) (每月 2000 次免费查询)

**1. 初始化**

```bash
MyPicoClaw onboard
```

**2. 配置** (`~/.mypicoclaw/config.json`)

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.MyPicoClaw/workspace",
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

**3. 获取 API Key**

- **LLM 供应商**: [Moonshot AI](https://platform.moonshot.ai) · [OpenRouter](https://openrouter.ai/keys) · [智谱](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) · [Anthropic](https://console.anthropic.com) · [OpenAI](https://platform.openai.com) · [Gemini](https://aistudio.google.com/api-keys)
- **联网搜索** (可选): [Brave Search](https://brave.com/search/api) - 提供免费档位 (2000 requests/month)

**4. 开始聊天**

```bash
MyPicoClaw agent -m "2+2 等于几？"
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
MyPicoClaw gateway
```
</details>

## ⚙️ 详细配置

配置文件路径：`~/.MyPicoClaw/config.json`

### 工作空间结构

MyPicoClaw 在你配置的工作空间（默认 `~/.MyPicoClaw/workspace`）中存储数据：

```
~/.MyPicoClaw/workspace/
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

| 命令 | 描述 |
|---------|-------------|
| `./mypicoclaw onboard` | 初始化配置与工作空间 |
| `./mypicoclaw agent -m "..."` | 与 Agent 进行单次对话 |
| `./mypicoclaw agent` | 进入交互式对话模式 |
| `./mypicoclaw gateway` | 启动网关（用于各聊天渠道） |
| `./mypicoclaw status` | 查看状态 |
| `./mypicoclaw cron list` | 列出所有定时任务 |
| `./mypicoclaw cron add ...` | 添加定时任务 |

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

---

## 📝 API 供应商对比

| 服务 | 免费档位 | 适用场景 |
|---------|-----------|-----------|
| **Moonshot** | 适配国际版 | 强力中文/英文支持 |
| **OpenRouter** | 200K tokens/月 | 尝试各种模型 (Claude, GPT-4 等) |
| **智谱 (Zhipu)** | 200K tokens/月 | 中国区访问流畅 |
| **Brave Search** | 2000 次/月 | 联网获取实时信息 |
| **Groq** | 有免费档位 | 极速推理 (Llama 3 等) |
