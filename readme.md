# 弹幕机器人核心部分
- 环境：go 1.20+
### 开发调试方式
```bash
go mod tidy
go run test/test.go
```

### 抽奖功能配置
```go
package config

type Config struct {
	// ...
    LotteryEnable    bool              `json:",default=true"`  // 抽奖开关
    LotteryUrl       string            `json:",optional"`      // 抽奖地址
}
```
