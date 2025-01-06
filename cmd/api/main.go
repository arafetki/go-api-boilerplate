package main

import (
	"os"
	"runtime/debug"

	"github.com/arafetki/go-echo-boilerplate/internal/app"
)

func main() {

	app := app.Init()

	if err := app.Run(); err != nil {
		trace := string(debug.Stack())
		app.Logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
