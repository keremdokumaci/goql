package utils

import (
	"fmt"
	"os"
)

type TestMode string

const (
	INTEGRATION TestMode = "integration"
	UNIT        TestMode = "unit"
)

func GetTestMode() TestMode {
	mode := os.Getenv("TEST_MODE")
	switch mode {
	case string(INTEGRATION):
		return INTEGRATION
	case string(UNIT):
		return UNIT
	default:
		fmt.Println("using default test mode... (unit)")
		return UNIT
	}
}
