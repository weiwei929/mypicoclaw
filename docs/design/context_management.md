# 上下文管理优化方案 (Context Management Optimization Plan)

**状态**: ✅ 已实施
**负责人**: MyPicoClaw Team
**日期**: 2026-02-13

## 1. 问题定义 (Problem Definition)

### 现象 (Symptom)
在长对话或执行复杂任务时，Agent 会崩溃或返回 API 错误：
```text
Error: API error: {"error":{"message":"Invalid request: Your request exceeded model token limit: 8192 (requested: 9331)","type":"invalid_request_error"}}
```

### 根本原因 (Root Cause)
1.  **窗口限制**: 默认模型 `moonshot-v1-8k` 的上下文窗口较小（仅 8192 tokens）。
2.  **Token 膨胀**:
    - **系统提示词 (System Prompt)**: 身份定义 + 工具描述 + 记忆库，基线占用约 2000+ tokens。
    - **工具输出**: 搜索结果和网页内容非常冗长（产生巨大的 JSON 结构）。
    - **历史累积**: 每一轮对话都会无限制地堆积到上下文中，缺乏清理机制。
3.  **缺乏保险丝**: Agent 会盲目地发送全部历史记录，直到 API 拒绝请求。

---

## 2. 建议优化策略 (Proposed Optimization Strategy)

### A. 模型容量调整 (Model Capacity Tuning)
**目标**: 提升基线容量，以适应 Agentic 工作流。

- **动作**: 更新 `config.json` 默认配置。
- **选择**:
  - **标准版**: `moonshot-v1-32k` (**推荐默认**，容量翻 4 倍)。
  - **超大杯**: `moonshot-v1-128k` (适用于深度研究/代码编写任务)。
- **配置示例**:
  ```json
  "agents": {
    "defaults": {
      "model": "moonshot-v1-32k",
      "max_tokens": 32000
    }
  }
  ```

### B. 主动生命周期管理 (策略一 - 用户首选)
**目标**: 在溢出发生*之前*进行预防，避免崩溃。

- **机制**: **90% 阈值自动重置 (Auto-Reset)**。
- **逻辑流程**:
  1.  **估算**: 每次调用 LLM 前，计算 `used_tokens`（系统提示 + 历史记录 + 新输入）。
  2.  **检查**: 如果 `used_tokens > (max_tokens * 0.9)`：
      - **归档**: 将当前 `session.json` 重命名为 `sessions/archive/{session_id}_{timestamp}.json`。
      - **重置**: 清空消息历史列表（保留 System Prompt、记忆和技能）。
      - **通知**: 发送系统消息：*"⚠️ 会话记忆已满 (90%)，已自动归档旧话题并开启新对话。"*.
      - **继续**: 在全新的上下文中处理用户刚才发送的消息。

### C. 应急处理 (Reactive Fallback)
**目标**: 处理*单次*交互就超限的极端情况（例如：超长搜索结果）。

- **场景**: 即使历史记录为空，单次工具输出的内容也过大。
- **动作**:
  1.  **智能截断**: 对工具输出（Tool Outputs）实施严格的字符限制（例如：搜索摘要最多 2000 字符）。
  2.  **错误捕获**: 如果 API 返回 `context_length_exceeded`：
      - **丢弃工具输出**: 从历史中移除最后一条超长的工具结果。
      - **优雅降级**: 返回友好的错误提示，而不是直接崩溃。

---

## 3. 实施路线图 (Implementation Roadmap)

1.  **实现 Token 估算器**
    - 在 `pkg/utils` 中添加轻量级 tokenizer（或字符比例估算）。

2.  **更新核心逻辑 (`loop.go`)**
    - 在 `runAgentLoop` 中插入"飞行前检查"：
      ```go
      if estimatedTokens > maxTokens * 0.9 {
          al.sessions.ArchiveAndReset(sessionKey)
      }
      ```

3.  **配置迁移**
    - 更新 `default_config`，优先使用 32k 模型。

4.  **工具输出约束**
    - 审查 `web_search` 和 `read_url` 工具，强制执行输出长度限制。
