package handler

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(handlerTestRunner(m))
}
