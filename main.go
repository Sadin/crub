package main

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"os"
)

// system vars
var (
	logger log.Logger
)

func main() {
	fmt.Println("Crub")

	// set up logging
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowInfo())
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	level.Info(logger).Log("msg", "starting bot")
}
