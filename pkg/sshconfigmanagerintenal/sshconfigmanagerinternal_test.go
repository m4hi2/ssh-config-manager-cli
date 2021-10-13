package sshconfigmanagerintenal

import (
	"os"
	"testing"
)

func TestFileBackUp(t *testing.T) {
	backUpFile("testdata/config")
	file, _ := os.ReadFile("testdata/config")
	backedUpFile, err := os.ReadFile("testdata/config.bak")
	if err != nil {
		t.Fatalf("Backup file not created")
	}

	if string(file) != string(backedUpFile) {
		t.Fatalf("Original file and backup file are not same")
	}
	os.Remove("testdata/config.bak")

}
