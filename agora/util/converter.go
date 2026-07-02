package util

// ExtractValuesFromArrayMap -
func ExtractValuesFromArrayMap(data []map[string]string, key string) []string {
	keys := []string{}
	for _, object := range data {
		keys = append(keys, object[key])
	}
	return keys
}
