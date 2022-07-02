package propertybasestests

import "strings"

func ConvertToRoman(arabic int) string {
	if arabic == 3 {
		return "III"
	}

	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
}
