package app

import (
	"os"
)

type Config struct {
	MuxAddr string
}

func (c *Config) Init() {
	c.MuxAddr = os.Getenv("MUX_ADDR")
}
