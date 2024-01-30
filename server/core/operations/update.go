package operations

import (
	"fmt"
	"pipebase/server/helpers"
	"pipebase/server/types"
)

func HandleUpdateRequest(request interface{}) {
	switch request := request.(type) {
	case types.UpdateRecordRequestStruct:
		updateOne(request)
	default:
		fmt.Println("Invalid request format for UpdateOperation")
	}
}

func updateOne(updateRequest types.UpdateRecordRequestStruct) error {
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
