package parser

import (
	"strings"
)

type Config struct {
	Host     string
	HostName string
	User     string
	Port     string
	IdFile   string
}

type configPair struct {
	FieldName string
	Value     string
}

type parsedConfigMap map[string]*Config

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

func Parse(fileContent string) parsedConfigMap {
	parsedConfigs := make(parsedConfigMap)
	configs := strings.Split(fileContent, "Host ")
	var configSlices []string
	for i, s := range configs {
		if i == 0 {
			continue
		}
		temp := "Host " + s
		configSlices = append(configSlices, temp)

	}

	for _, s := range configSlices {
		s = strings.TrimSpace(s)
		configLines := strings.Split(s, "\n")
		currentHost := ""
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
