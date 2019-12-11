# go-tinyid

[![Go Report Card](https://goreportcard.com/badge/github.com/rrylee/go-tinyid)](https://goreportcard.com/report/github.com/rrylee/go-tinyid)

ID Generator id生成器 分布式id生成系统，简单易用、高性能、高可用的id生成系统

fork from https://github.com/didi/tinyid (golang version)

### Server
./server/build.sh

```go
package main

import (
	"github/rrylee/go-tinyid/internal"
	"github/rrylee/go-tinyid/server/http"
	"log"
	"os"
)

func main() {
	internal.Logger = log.New(os.Stdout, "[tiny-id]", 0) // your logger component
	err := http.Run("/Users/rry/Code/github/go-tinyid/server")
	if err != nil {
		panic(err)
	}
}
```

### client 接入

```go
package main

import (
	"github/rrylee/go-tinyid/internal"
	"github/rrylee/go-tinyid/client"
	"github/rrylee/go-tinyid/client/config"
	"log"
	"os"
)

func main() {
	internal.Logger = log.New(os.Stdout, "[test]", 0)
	client.Init(&config.Config{
		TinyIdServer: []string{"http://127.0.0.1:8999"},
		TinyIdToken:  "test",
		Timeout:      1 * time.Second,
	})

	client := &client.TinyIdClient{}
	for i := 0; i < 100; i++ {
		fmt.Println(client.NextId("test"))
	}
}
```

### http 接入

```bash
$curl http://localhost:8999/tinyid/id/nextId?bizType=test&batchSize=10
${
    "Data": [
        1900091,
        1900092,
        1900093,
        1900094,
        1900095,
        1900096,
        1900097,
        1900098,
        1900099,
        1900100
    ],
    "Code": 200,
    "Message": ""
}

$ curl http://localhost:8999/tinyid/id/nextIdSimple?bizType=test&batchSize=10
$ 1900101,1900102,1900103,1900104,1900105,1900106,1900107,1900108,1900109,1900110
```
