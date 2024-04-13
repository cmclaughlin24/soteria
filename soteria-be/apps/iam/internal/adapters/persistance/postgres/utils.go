package postgres

import "strings"

func parsePostgresTextArray(deliveryMethods string) []string {
	if !(len(deliveryMethods) > 0) {
		return make([]string, 0)
	}

	return strings.Split(deliveryMethods[1:len(deliveryMethods)-1], ",")
}
