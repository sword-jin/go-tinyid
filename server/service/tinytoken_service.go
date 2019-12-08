package service

import (
	"github.com/rrylee/go-tinyid/internal"
	"github.com/rrylee/go-tinyid/server/dao"
	"time"
)

var Token2bizTypes map[string]string

func AutoRefreshTokens() {
	go refresh()
}

func refresh() {
	internal.Logf("refresh token begin")

	tokens := dao.SelectAllToken()
	token2bizTypes := make(map[string]string, len(tokens))
	for _, token := range tokens {
		token2bizTypes[token.Token] = token.BizType
	}
	Token2bizTypes = token2bizTypes

	internal.Logf("refresh token success.")
	time.Sleep(1 * time.Minute)
}
