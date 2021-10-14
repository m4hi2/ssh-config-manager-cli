package sshconfigmanagerintenal

import (
	"log"
	"os"
	"testing"

	"github.com/m4hi2/ssh-config-manager-cli/pkg/parser"
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

func TestLowLevelAdd(t *testing.T) {
	sampleConfig := parser.Config{
		Host:     "gg",
		HostName: "hallo",
		Port:     "22",
		User:     "m4hi2",
	}
	Add(sampleConfig, "testdata/config")

	fileSample, _ := os.ReadFile("testdata/config_add_sample")
	fileAdded, _ := os.ReadFile("testdata/config")

	if string(fileAdded) != string(fileSample) {
		addCleanUp("testdata/config")
		log.Fatalln("Config file not created properly.")

	}

	addCleanUp("testdata/config")
}

func addCleanUp(filePath string) {
	bakFilePath := filePath + ".bak"
	backUpFile, _ := os.ReadFile(bakFilePath)
	os.WriteFile(filePath, []byte(backUpFile), 0644)
	os.Remove(bakFilePath)

}
