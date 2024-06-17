package main

import (
	"context"
	"github.com/liuzhaomax/ovo-user/internal/app"
	"github.com/liuzhaomax/ovo-user/internal/core"
)

func main() {
	app.Launch(
		context.Background(),
		app.SetConfigFile(core.LoadEnv()),
		app.SetWWWDir("www"),
	)
}
