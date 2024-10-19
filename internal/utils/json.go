package utils

import "encoding/json"

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// форматирование json-строки
func FormatJson(data string, indent string) string {
	var obj map[string]interface{}
	json.Unmarshal([]byte(data), &obj)

	result, _ := json.MarshalIndent(obj, "", indent)
	return string(result)
}
