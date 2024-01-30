package helpers

func GetValueFromField(record interface{}, field string) string {
	if recordMap, ok := record.(map[string]interface{}); ok {
		if value, exists := recordMap[field]; exists {
			if strValue, ok := value.(string); ok {
				return strValue
			}
		}
	}
	return ""
}
