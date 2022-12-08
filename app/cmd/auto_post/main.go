package main

import (
	"auto_post/app/cmd/auto_post/bundlefx"
	"go.uber.org/fx"
)

func main() {
	fx.New(bundlefx.Module).Run()
}
