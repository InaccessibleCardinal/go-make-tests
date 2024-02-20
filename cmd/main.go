package main

import (
	"go-make-tests/internal/env"
	"go-make-tests/internal/svc"
)

func main() {
	env.LoadEnv()
	language := "python"
	framework := "pytest"
	codeInput := `
	def add(n, m):
		return n + m
`
	outfile := "./out/add_test.py"
	svc.AskForTest(language, framework, codeInput, outfile)
}
