package conversion

import (
	"strconv"
)

func StringsToFloats(string []string) ([]float64, error) {
	floats := make([]float64, len(string))

	for stringIndex, stringVal := range string {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, err
		}

		floats[stringIndex] = floatVal
	}

	return floats, nil
}
