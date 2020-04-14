# 自定义 micro api 网关

> 由于涉及`Router`替换，而`micro/api`又有`internal`包，所以将这部分放在`micro`库内，尽量在`micro/api`基础上少做改动，方便持续升级维护

- API 网关路由支持服务筛选

## 使用

[example](/gateway/example/main.go)

```go
package main

import (
	"net/http"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/micro/v2/gateway/cmd"
	"github.com/micro/micro/v2/gateway/router"
)

func main() {
	// Router services filter
	opt := router.WithOption(
		router.WithFilter(func(req *http.Request) selector.Filter {
			return selector.FilterLabel("key", "val")
		}),
	)

	cmd.Init(
		opt,
		// micro option
	)
}
```

运行网关

```shell script
$ micro gateway
```
