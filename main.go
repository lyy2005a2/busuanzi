package main

import (
	"github.com/lyy2005a2/busuanzi/config"
	"github.com/lyy2005a2/busuanzi/core"
	"github.com/lyy2005a2/busuanzi/process/redisutil"
	"github.com/lyy2005a2/busuanzi/process/webutil"
)

func main() {
	config.Init()
	redisutil.Init()

	core.InitExpire()

	webutil.Init()
}
