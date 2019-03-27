package app

import (
	"os"
	"strconv"
)

type Config struct {
	MuxAddr           string
	CVEWorkerPoolSize int
	CVEFeedsDir       string
}

func (c *Config) Init() error {

	c.MuxAddr = os.Getenv("MUX_ADDR")

	c.CVEFeedsDir = os.Getenv("CVE_FEEDS_DIR")

	cveWorkerPoolSize, err := strconv.Atoi(os.Getenv("CVE_WORKER_POOL_SIZE"))
	if err != nil {
		return err
	}
	c.CVEWorkerPoolSize = cveWorkerPoolSize

	return nil
}
