package operations

import (
	"encoding/json"
	"fmt"
	"server/enums"
	"server/helpers"
	"server/types"
)

func HandleDeleteRequest(jsonData []byte, incomingRequest interface{}) ([]byte, error) {
	requestMap := incomingRequest.(map[string]interface{})
	dataMap := requestMap["data"].(map[string]interface{})
	requestType := dataMap["type"]

	if requestType == string(enums.DeleteAllOperation) {
		var bulkDeleteRequest types.BulkDeleteRecordRequestStruct

		err := json.Unmarshal(jsonData, &bulkDeleteRequest)
		if err != nil {
			fmt.Println("Error unmarshaling create request:", err)
			return []byte("Error unmarshaling create request:\n"), err
		}

		err = deleteAll(bulkDeleteRequest)

		if err != nil {
			fmt.Println("Unable to create record", err)
			return []byte("Unable to create record\n"), err
		}
	}

	if requestType == string(enums.DeleteOneOperation) {
		var deleteOneRequest types.DeleteRecordRequestStruct

		err := json.Unmarshal(jsonData, &deleteOneRequest)
		if err != nil {
			fmt.Println("Error unmarshaling delete-one request:", err)
			return []byte("Error unmarshaling delete-one request:"), err
		}

		err = deleteOne(deleteOneRequest)

		if err != nil {
			fmt.Println("Unable to create record", err)
			return []byte("Unable to create record\n"), err
		}
	}

	return []byte("Delete operation successfully processed\n"), nil
}

func deleteAll(deleteRequest types.BulkDeleteRecordRequestStruct) error {
	tableName := deleteRequest.Data.TableName

	if !helpers.CheckIfTableExists(tableName) {
		return fmt.Errorf("table %s does not exist", tableName)
	}

	emptyData := []interface{}{}

	return helpers.WriteTableData(tableName, emptyData)
}

func deleteOne(deleteRequest types.DeleteRecordRequestStruct) error {
	tableName := deleteRequest.Data.TableName

	if !helpers.CheckIfTableExists(tableName) {
		return fmt.Errorf("table %s does not exist", tableName)
	}

	data, err := helpers.ReadTableData(tableName)
	if err != nil {
		return err
	}

	index := -1
	for i, record := range data {
		if helpers.GetValueFromField(record, deleteRequest.Data.Query.Field) == deleteRequest.Data.Query.Value {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("record not found in table %s with the provided query", tableName)
	}

	data = append(data[:index], data[index+1:]...)

	return helpers.WriteTableData(tableName, data)
}
