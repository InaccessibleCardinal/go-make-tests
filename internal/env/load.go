package env

import (
	"go-make-tests/internal/colors"
	"log"
	"os"
	"strings"
)

func LoadEnv() {
	bts, err := os.ReadFile(".env")
	LogFatalErr(err)

	lines := strings.Split(string(bts), "\n")

	for _, line := range lines {
		parts := strings.Split(line, "=")
		log.Println(colors.Blue("Getting environment ready..."))
		os.Setenv(parts[0], strings.TrimSpace(parts[1]))
	}
}

func LogFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
