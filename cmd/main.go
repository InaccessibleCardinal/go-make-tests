package main

import (
	"go-make-tests/internal/env"
	"go-make-tests/internal/ui"
)

func main() {
	env.LoadEnv()
	ui.Run()
}
