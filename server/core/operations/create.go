package operations

import (
	"encoding/json"
	"fmt"
	"pipebase/server/enums"
	"pipebase/server/helpers"
	"pipebase/server/types"
)

func HandleCreateRequest(jsonData []byte, incomingRequest interface{}) ([]byte, error) {
	requestMap := incomingRequest.(map[string]interface{})
	dataMap := requestMap["data"].(map[string]interface{})
	requestType := dataMap["type"]

	if requestType == string(enums.CreateOperation) {
		var singleRecord types.SingleCreateRecordRequestStruct

		err := json.Unmarshal(jsonData, &singleRecord)
		if err != nil {
			fmt.Println("Error unmarshaling create request:", err)
			return []byte("Error unmarshaling create request:"), err
		}

		err = singleCreate(singleRecord)

		if err != nil {
			fmt.Println("Unable to create record", err)
			return []byte("Unable to create record"), err
		}
	}

	if requestType == string(enums.BulkCreateOperation) {
		var bulkRecord types.BulkCreateRecordRequestStruct

		err := json.Unmarshal(jsonData, &bulkRecord)

		if err != nil {
			fmt.Println("Error unmarshaling create-bulk request:", err)
			return []byte("Error unmarshaling create-bulk request:"), err
		}

		err = bulkCreate(bulkRecord)

		if err != nil {
			fmt.Println("Unable to create bulk record", err)
			return []byte("Unable to create bulk record"), err
		}
	}

	return []byte("Create operation successfully processed"), nil
}

func bulkCreate(records types.BulkCreateRecordRequestStruct) error {
	tableName := records.Data.TableName

	if !helpers.CheckIfTableExists(tableName) {
		if err := helpers.CreateTableFile(tableName); err != nil {
			return err
		}
	}

	data, err := helpers.ReadTableData(tableName)

	if err != nil {
		return err
	}

	data = append(data, records.Data.Record...)

	return helpers.WriteTableData(tableName, data)
}

func singleCreate(record types.SingleCreateRecordRequestStruct) error {
	tableName := record.Data.TableName

	if !helpers.CheckIfTableExists(tableName) {
		if err := helpers.CreateTableFile(tableName); err != nil {
			return err
		}
	}

	data, err := helpers.ReadTableData(tableName)

	if err != nil {
		return err
	}

	data = append(data, record.Data.Record)

	return helpers.WriteTableData(tableName, data)
}
