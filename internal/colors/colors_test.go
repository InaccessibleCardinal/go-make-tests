package colors

import (
	"testing"
)

func TestBlue(t *testing.T) {
	want := "\u001b[36mtest\u001b[0m"
	got := Blue("test")
	if got != want {
		t.Errorf("Blue() = %v, want %v", got, want)
	}
}

func TestRed(t *testing.T) {
	want := "\u001b[31mtest\u001b[0m"
	got := Red("test")
	if got != want {
		t.Errorf("Red() = %v, want %v", got, want)
	}
}

func TestGreen(t *testing.T) {
	want := "\u001b[32mtest\u001b[0m"
	got := Green("test")
	if got != want {
		t.Errorf("Green() = %v, want %v", got, want)
	}
}
