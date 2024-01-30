package operations

import (
	"encoding/json"
	"fmt"
	"pipebase/server/enums"
	"pipebase/server/helpers"
	"pipebase/server/types"
)

func HandleUpdateRequest(jsonData []byte, incomingRequest interface{}) ([]byte, error) {
	requestMap := incomingRequest.(map[string]interface{})
	dataMap := requestMap["data"].(map[string]interface{})
	requestType := dataMap["type"]

	if requestType == string(enums.UpdateOperation) {
		var updateRecord types.UpdateRecordRequestStruct

		err := json.Unmarshal(jsonData, &updateRecord)
		if err != nil {
			fmt.Println("Error unmarshaling update request:\n", err)
			return []byte("Error unmarshaling update request:\n"), err
		}

		err = updateOne(updateRecord)

		if err != nil {
			fmt.Println("Unable to update record", err)
			return []byte("Unable to update record"), err
		}
	}
	return []byte("Update operation successfully processed\n"), nil
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
