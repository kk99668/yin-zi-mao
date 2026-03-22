# API URL 更新设计

**日期**: 2026-03-22
**作者**: Claude
**状态**: 已批准

## 概述

将 Yin-Zi-Mao 程序的默认 API URL 从 `api.mzs2025.asia` 更改为 `api.yinzimao.com`。

## 目标

- 更新默认 API 基础 URL
- 更新默认回测 API 基础 URL
- 保持现有端口配置不变

## 修改范围

### 文件变更

**文件**: `internal/config/config.go`

**修改前** (第 15-16 行):
```go
DefaultAPIURL      = "https://api.mzs2025.asia:8003"
DefaultBacktestURL = "https://api.mzs2025.asia:8001/backtest"
```

**修改后**:
```go
DefaultAPIURL      = "https://api.yinzimao.com:8003"
DefaultBacktestURL = "https://api.yinzimao.com:8001/backtest"
```

### 影响分析

| 用户类型 | 影响 |
|---------|------|
| 新用户 | 登录时自动使用新的 API URL |
| 现有用户 | 配置文件中的 API URL 不变，下次登录或手动更新时生效 |

### 无需修改的组件

以下组件通过配置常量获取 URL，无需修改：
- `internal/api/client.go` - API 客户端
- `cmd/login.go` - 登录命令
- 所有使用 API 的命令

## 验证计划

1. 修改配置常量
2. 重新编译程序
3. 删除测试配置或使用新环境
4. 执行登录命令验证连接
5. 运行回测命令验证 API 功能

## 实施步骤

1. 编辑 `internal/config/config.go` 文件
2. 修改第 15 行 `DefaultAPIURL` 常量
3. 修改第 16 行 `DefaultBacktestURL` 常量
4. 运行 `make build` 重新编译
5. 测试登录和回测功能

## 回滚计划

如果需要回滚，将常量值恢复为原值并重新编译即可。
