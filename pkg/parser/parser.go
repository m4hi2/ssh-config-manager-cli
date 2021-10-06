package parser

import (
	"log"
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

func Parse(fileContent string) map[string]Config {
	parsedConfigs := make(map[string]Config)
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
		currentHost := "nil"
		for _, line := range configLines {
			if line == " " {
				continue
			}
			line = strings.TrimSpace(line)
			configPair := extractConfigPairs(line)

			if configPair.FieldName == "Host" {
				currentHost = configPair.Value
				parsedConfigs[currentHost] = Config{
					Host: currentHost,
				}
			}
			// TODO: write in if-else structure.
			if configPair.FieldName == "HostName" {
				parsedConfigs[currentHost] = Config{
					HostName: configPair.Value,
				}
			}
			if configPair.FieldName == "Port" {
				parsedConfigs[currentHost] = Config{
					Port: configPair.Value,
				}
			}
			if configPair.FieldName == "User" {
				parsedConfigs[currentHost] = Config{
					User: configPair.Value,
				}
			}
		}
	}
	return parsedConfigs
}

func rangePrinter(thingToPrint []string) {
	for i, s := range thingToPrint {
		log.Println(i, s)
	}
}
