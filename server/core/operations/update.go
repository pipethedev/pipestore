package operations

import (
	"fmt"
	"pipebase/server/helpers"
	"pipebase/server/types"
)

func UpdateOne(updateRequest types.UpdateRecordRequestStruct) error {
	tableName := updateRequest.Data.TableName

	if !helpers.CheckIfTableExists(tableName) {
		return fmt.Errorf("table %s does not exist", tableName)
	}

	data, err := helpers.ReadTableData(tableName)
	if err != nil {
		return err
	}

	index := -1
	for i, record := range data {
		if helpers.GetValueFromField(record, updateRequest.Data.Query.Field) == updateRequest.Data.Query.Value {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("record not found in table %s with the provided query", tableName)
	}

	data[index] = updateRequest.Data.Record

	return helpers.WriteTableData(tableName, data)
}
