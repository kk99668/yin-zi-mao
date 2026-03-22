# API URL 更新实现计划

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 将 Yin-Zi-Mao 程序的默认 API URL 从 `api.mzs2025.asia` 更改为 `api.yinzimao.com`

**Architecture:** 直接修改 `internal/config/config.go` 中的两个常量值，重新编译程序即可生效。

**Tech Stack:** Go 1.x, Make

---

## 文件结构

**修改文件:**
- `internal/config/config.go` - 配置常量定义

---

## Task 1: 修改 DefaultAPIURL 常量

**Files:**
- Modify: `internal/config/config.go:15`

- [ ] **Step 1: 编辑文件**

将第 15 行：
```go
DefaultAPIURL      = "https://api.mzs2025.asia:8003"
```

改为：
```go
DefaultAPIURL      = "https://api.yinzimao.com:8003"
```

- [ ] **Step 2: 验证修改**

Run: `grep "DefaultAPIURL" internal/config/config.go`
Expected: 显示 `https://api.yinzimao.com:8003`

- [ ] **Step 3: 提交更改**

```bash
git add internal/config/config.go
git commit -m "feat: 更新 DefaultAPIURL 为 api.yinzimao.com

- 将默认 API URL 从 api.mzs2025.asia:8003 更改为 api.yinzimao.com:8003
- 保持端口号 8003 不变

Co-Authored-By: Claude Opus 4.6 <noreply@anthropic.com>"
```

---

## Task 2: 修改 DefaultBacktestURL 常量

**Files:**
- Modify: `internal/config/config.go:16`

- [ ] **Step 1: 编辑文件**

将第 16 行：
```go
DefaultBacktestURL = "https://api.mzs2025.asia:8001/backtest"
```

改为：
```go
DefaultBacktestURL = "https://api.yinzimao.com:8001/backtest"
```

- [ ] **Step 2: 验证修改**

Run: `grep "DefaultBacktestURL" internal/config/config.go`
Expected: 显示 `https://api.yinzimao.com:8001/backtest`

- [ ] **Step 3: 提交更改**

```bash
git add internal/config/config.go
git commit -m "feat: 更新 DefaultBacktestURL 为 api.yinzimao.com

- 将默认回测 API URL 从 api.mzs2025.asia:8001/backtest 更改为 api.yinzimao.com:8001/backtest
- 保持端口号 8001 不变

Co-Authored-By: Claude Opus 4.6 <noreply@anthropic.com>"
```

---

## Task 3: 重新编译程序

**Files:**
- Build: `yin-zi-mao` (可执行文件)

- [ ] **Step 1: 清理旧构建**

```bash
make clean
```

Expected: 删除 build 目录中的旧文件

- [ ] **Step 2: 编译新版本**

```bash
make build
```

Expected: 编译成功，生成新的可执行文件

- [ ] **Step 3: 验证编译结果**

```bash
./yin-zi-mao version
```

Expected: 显示版本信息，程序可正常运行

---

## Task 4: 验证功能（可选，需要有效的 API 凭证）

**Files:**
- Test: 程序功能

- [ ] **Step 1: 测试登录功能**

```bash
./yin-zi-mao login --username <test_username> --password <test_password>
```

Expected: 登录成功，使用新的 API 地址

- [ ] **Step 2: 测试回测功能**

```bash
./yin-zi-mao backtest list
```

Expected: 能正确调用新 API 地址获取回测列表

---

## 验证检查清单

完成以上任务后，确认以下事项：

- [ ] `internal/config/config.go` 中两个常量已更新
- [ ] 程序已重新编译
- [ ] 新编译的程序可以正常运行
- [ ] (可选) 登录和 API 调用功能正常

---

## 回滚计划

如果需要回滚：

```bash
# 恢复代码
git revert HEAD~2..

# 重新编译
make build
```
