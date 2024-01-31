package helpers

import (
	"server/types"
	"sort"
	"strings"
)

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

func CompareIgnoreCase(a, b string) bool {
	return strings.EqualFold(strings.TrimSpace(a), strings.TrimSpace(b))
}

func RearrangeFields(record map[string]interface{}) map[string]interface{} {
	//Todo: Needs a re-work
	var pairs []types.KeyValue

	if id, ok := record["id"]; ok {
		pairs = append(pairs, types.KeyValue{Key: "id", Value: id})
	}

	for key, value := range record {
		if key != "id" && key != "createdAt" && key != "updatedAt" {
			pairs = append(pairs, types.KeyValue{Key: key, Value: value})
		}
	}

	pairs = append(pairs, types.KeyValue{Key: "createdAt", Value: record["createdAt"]})
	pairs = append(pairs, types.KeyValue{Key: "updatedAt", Value: record["updatedAt"]})

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})

	result := make(map[string]interface{})
	for _, pair := range pairs {
		result[pair.Key] = pair.Value
	}

	return result
}

func RearrangeRecords(records []map[string]interface{}) []map[string]interface{} {
	var rearrangedRecords []map[string]interface{}
	for _, record := range records {
		rearrangedRecord := RearrangeFields(record)
		rearrangedRecords = append(rearrangedRecords, rearrangedRecord)
	}
	return rearrangedRecords
}
