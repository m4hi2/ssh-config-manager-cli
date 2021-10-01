package parser

import (
	"os"
	"testing"
)

func TestCanReadConfigFile(t *testing.T) {
	_, err := os.ReadFile("testdata/config")

	if err != nil {
		t.Fatal("File can not be read")
	}
}
