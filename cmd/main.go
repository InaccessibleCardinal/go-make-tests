package main

import (
	"go-make-tests/internal/env"
	"go-make-tests/internal/ui"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	env.LoadEnv()
}

func main() {
	ui.Run()
}
