package parser

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func TestCanReadConfigFile(t *testing.T) {
	_, err := os.ReadFile("testdata/config")

	if err != nil {
		t.Fatal("File can not be read")
	}
}

func TestExtractConfigPairs(t *testing.T) {
	configFile, err := os.Open("testdata/config")
	hostConfig := configPair{
		FieldName: "Host",
		Value:     "hello",
	}

	hostNameConfig := configPair{
		FieldName: "HostName",
		Value:     "10.0.0.1",
	}
	userConfig := configPair{
		FieldName: "User",
		Value:     "user",
	}
	portConfig := configPair{
		FieldName: "Port",
		Value:     "22",
	}
	if err != nil {
		t.Fatal("File Can not be read")
	}
	s := bufio.NewReader(configFile)
	value, _, _ := s.ReadLine()
	config := extractConfigPairs(string(value))

	if config != hostConfig {
		t.Fatal("Host Config Cannot be read")
	}

	value, _, _ = s.ReadLine()
	config = extractConfigPairs(string(value))
	if config != hostNameConfig {
		t.Fatal("HostName Config Cannot be read")
	}

	value, _, _ = s.ReadLine()
	config = extractConfigPairs(string(value))
	if config != portConfig {
		t.Fatal("Port Config Cannot be read")
	}

	value, _, _ = s.ReadLine()
	config = extractConfigPairs(string(value))
	if config != userConfig {
		t.Fatal("user Config Cannot be read")
	}
}

var sampleParsedConfigs = ParsedConfigMap{
	"hello": &Config{
		Host:     "hello",
		HostName: "10.0.0.1",
		Port:     "22",
		User:     "user",
	},
	"hello2": &Config{
		Host:     "hello2",
		HostName: "10.0.0.11",
		Port:     "22",
		User:     "user1",
	},
}

func TestParser(t *testing.T) {
	file, _ := os.ReadFile("testdata/config")
	parsedConfigs := Parse(string(file))

	if !reflect.DeepEqual(parsedConfigs, sampleParsedConfigs) {
		t.Fatalf("Configs are not equal. \n ParsedConfig is: %v \n and ExpectedParsedConfigs is: %v", parsedConfigs, sampleParsedConfigs)
	}
}

func TestComposer(t *testing.T) {
	sampleConfigFile, _ := os.ReadFile("testdata/config")
	sampleConfigFileString := string(sampleConfigFile)
	composedConfig := Compose(sampleParsedConfigs)

	if sampleConfigFileString != composedConfig {
		t.Fatalf("Composed Config is not up to the mark \nComposed Config: \n%s\n vs Sample Config: \n%s", composedConfig, sampleConfigFile)
	}
}
