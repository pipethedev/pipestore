package operations

import (
	"encoding/json"
	"fmt"
	"server/enums"
	"server/helpers"
	"server/types"
	"time"
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
			return []byte("Error unmarshaling create request: \n"), err
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
			return []byte("Error unmarshaling create-bulk request: \n"), err
		}

		err = bulkCreate(bulkRecord)

		if err != nil {
			fmt.Println("Unable to create bulk record", err)
			return []byte("Unable to create bulk record\n"), err
		}
	}

	StartIndexing()

	return []byte("Create operation successfully processed\n"), nil
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

	for _, record := range records.Data.Record {
		record["id"] = helpers.GenerateUUID()
		record["createdAt"] = time.Now()
		record["updatedAt"] = time.Now()
	}

	rearrangedRecords := helpers.RearrangeRecords(records.Data.Record)
	interfaceRecords := make([]interface{}, len(rearrangedRecords))
	for i, r := range rearrangedRecords {
		interfaceRecords[i] = r
	}

	data = append(data, interfaceRecords...)

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
	record.Data.Record["id"] = helpers.GenerateUUID()
	record.Data.Record["createdAt"] = time.Now()
	record.Data.Record["updatedAt"] = time.Now()

	rearrangedRecord := helpers.RearrangeFields(record.Data.Record)
	data = append(data, rearrangedRecord)

	return helpers.WriteTableData(tableName, data)
}
