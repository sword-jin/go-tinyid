package client

import (
	"fmt"
	"github.com/rrylee/go-tinyid/client/config"
	"github.com/rrylee/go-tinyid/internal"
	"log"
	"os"
	"testing"
	"time"
)

func TestTinyIdClient_NextId(t *testing.T) {
	internal.Logger = log.New(os.Stdout, "[test]", 0)
	Init(&config.Config{
		TinyIdServer: []string{"http://127.0.0.1:8999"},
		TinyIdToken:  "test",
		Timeout:      1 * time.Second,
	})

	client := &TinyIdClient{}
	for i := 0; i < 100; i++ {
		fmt.Println(client.NextId("test"))
	}
}
