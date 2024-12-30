# Hengo

Hengo 是一个模仿Gin的轻量级的 Go Web 框架，旨在简化 HTTP 服务的构建和 API 开发。它提供了简单的路由配置、JSON 绑定和响应处理，帮助开发者快速实现 Web 服务。

## 特性

- 简单易用的路由和请求处理
- 自动处理 JSON 请求和响应
- 日志记录与调试支持
- 支持自定义响应格式
- 提供 HTTP 状态码常量

## 安装

你可以通过 `go get` 安装 Hengo 框架：

```bash
go get github.com/kelin428/hengo
````

## 使用示例

### 基本示例

以下是一个简单的示例，展示如何使用 Hengo 创建一个 HTTP 服务并处理 GET 请求：

```go
type Name struct {
	Id string `json:"id"`
}

func main() {
	// 创建一个新的 Hengo 实例
	hen := hengo.New()

	// 设置一个 GET 路由
	hen.Get("/hello", func(c *hengo.Context) {
		var name Name
		c.BindJSON(&name)

		logConfig := log.NewConfig(false)
		logger := logConfig.LogHelp("log/log.go:30", "server/LocalHttpRequestFilter")
		logger.Info("test")

		c.JsonResponse(hengo.StatusOK, "Hello World", hengo.H{
			"name":    "Hen",
			"version": "0.1",
		})
	})

	// 启动 HTTP 服务器
	hen.Run(":8000")
}
```
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
