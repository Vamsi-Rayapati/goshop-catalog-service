package utils

import "encoding/json"

func Map[T any, R any](input []T, format func(T) R) []R {
	output := make([]R, len(input))

	for index, value := range input {
		output[index] = format(value)
	}

	return output
}

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}
