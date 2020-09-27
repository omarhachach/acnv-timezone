package main

import (
	"github.com/omarhachach/bear"
)

// Config is extended from Bear config.
type Config struct {
	*bear.Config
	DB string `json:"db"`
}
