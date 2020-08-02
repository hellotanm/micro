# 自定义 micro api 网关

> 由于涉及`Router`替换，而`micro/api`又有`internal`包，所以将这部分放在`micro`库内，尽量在`micro/client/api`基础上少做改动，方便持续升级维护

- API 网关路由支持服务筛选

## 使用

[example](/client/gateway/example/main.go)
```go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/micro/micro/v3/client/gateway"
	gwRouter "github.com/micro/micro/v3/client/gateway/router"
	"github.com/micro/micro/v3/cmd"
)

func main() {
	cmd.Register(
		gateway.Commands(
			gwRouter.WithFilter(func(req *http.Request) gwRouter.ServiceFilter {
				return gwRouter.FilterLabel("key", "val")
			}),
		),
	)

	if err := cmd.DefaultCmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

```shell script
$ go build -o micro main.go

$ micro --profile=local gateway
```
