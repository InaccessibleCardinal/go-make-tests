package env

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	// Create a temporary .env file for testing
	err := os.WriteFile(".env", []byte("KEY1=Value1\nKEY2=Value2"), 0644)
	if err != nil {
		t.Fatal("Error creating test .env file:", err)
	}
	defer os.Remove(".env")

	LoadEnv()

	if val := os.Getenv("KEY1"); val != "Value1" {
		t.Errorf("Expected KEY1 to be 'Value1', got '%s'", val)
	}

	if val := os.Getenv("KEY2"); val != "Value2" {
		t.Errorf("Expected KEY2 to be 'Value2', got '%s'", val)
	}
}
