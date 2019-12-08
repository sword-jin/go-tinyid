package main

import (
	"github.com/rrylee/go-tinyid/internal"
	"github.com/rrylee/go-tinyid/server/http"
	"log"
	"os"
)

func main() {
	internal.Logger = log.New(os.Stdout, "[tiny-id]", 0)
	err := http.Run("/Users/rry/Code/go/src/github.com/rrylee/go-tinyid/server")
	if err != nil {
		panic(err)
	}
}
