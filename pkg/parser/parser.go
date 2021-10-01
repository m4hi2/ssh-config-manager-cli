package parser

import "strings"

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
