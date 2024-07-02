package helper

import (
	_ "embed"
	"strings"
)

//go:embed resources/application.properties
var properties string

func GetProperties(name string) string {
	if properties != "" {
		rows := strings.Split(properties, "\r\n")
		for _, row := range rows {
			if strings.Contains(row, "=") {
				data := strings.Split(row, "=")
				key := data[0]
				if key == name {
					return data[1]
				}
			}
		}
	}
	return ""
}
