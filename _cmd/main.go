package main

import (
	"context"

	"github.com/Aninetix/core/ancore"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flg, config, logger := ancore.InitCore[anparam.Flags, anparam.Config]()
	AnCore := ancore.BootCore(flg, config, logger, ctx, cancel)

	AnCore.Run()

	<-ctx.Done()
}
