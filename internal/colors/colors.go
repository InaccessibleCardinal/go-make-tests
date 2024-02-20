package colors

import "fmt"

func Blue(s string) string {
	return fmt.Sprintf("\u001b[36m%s\u001b[0m", s)
}

func Red(s string) string {
	return fmt.Sprintf("\u001b[31m%s\u001b[0m", s)
}

func Green(s string) string {
	return fmt.Sprintf("\u001b[32m%s\u001b[0m", s)
}