package parser

import (
	"fmt"
	"strings"
)

type Config struct {
	Host     string
	HostName string
	User     string
	Port     string
	// IdFile   string
}

type configPair struct {
	FieldName string
	Value     string
}

type ParsedConfigMap map[string]*Config

func extractConfigPairs(line string) configPair {
	configLine := strings.Split(line, " ")
	fieldName := strings.TrimSpace(configLine[0])
	value := strings.TrimSpace(configLine[1])
	config := configPair{
		FieldName: fieldName,
		Value:     value,
	}
	return config
}

func Parse(fileContent string) ParsedConfigMap {
	parsedConfigs := make(ParsedConfigMap)
	configs := strings.Split(fileContent, "Host ")
	var configSlices []string
	for _, s := range configs[1:] {
		temp := "Host " + s
		configSlices = append(configSlices, temp)

	}

	for _, s := range configSlices {
		s = strings.TrimSpace(s)
		configLines := strings.Split(s, "\n")
		var currentHost string
		for _, line := range configLines {
			if line == " " {
				continue
			}
			line = strings.TrimSpace(line)
			configPair := extractConfigPairs(line)

			switch configPair.FieldName {
			case "Host":
				currentHost = configPair.Value
				parsedConfigs[currentHost] = &Config{
					Host: currentHost,
				}
			case "HostName":
				parsedConfigs[currentHost].HostName = configPair.Value
			case "Port":
				parsedConfigs[currentHost].Port = configPair.Value
			case "User":
				parsedConfigs[currentHost].User = configPair.Value
			}
		}
	}
	return parsedConfigs
}

func Compose(parsedConfigs ParsedConfigMap) (sshConfig string) {
	for _, v := range parsedConfigs {
		sshConfig = sshConfig + fmt.Sprintf("Host %s\n\tHostName %s\n\tPort %s\n\tUser %s\n\n", v.Host, v.HostName, v.Port, v.User)
	}
	sshConfig = strings.TrimSpace(sshConfig)
	return
}
