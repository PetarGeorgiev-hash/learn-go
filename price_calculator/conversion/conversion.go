package conversion

import "strconv"

func StringToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for index, value := range strings {
		floatPrice, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		floats[index] = floatPrice
	}
	return floats, nil
}
