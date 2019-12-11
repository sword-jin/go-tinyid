package config

import "time"

type Config struct {
	TinyIdServer []string
	TinyIdToken  string
	Timeout      time.Duration
}
