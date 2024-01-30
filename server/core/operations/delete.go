package operations

import (
	"fmt"
	"pipebase/server/helpers"
	"pipebase/server/types"

	"github.com/grahms/godantic"
)

func HandleDeleteRequest(jsonData []byte, request interface{}) {
	validator := godantic.Validate{}

	switch request := request.(type) {
	case types.BulkDeleteRecordRequestStruct:
		var bulkData types.BulkDeleteRecordRequestStruct

		err := validator.BindJSON(jsonData, &bulkData)

		if err != nil {
			fmt.Println("Error validating create request:", err)
			return
		}
		deleteAll(request)
	case types.DeleteRecordRequestStruct:
		var singleData types.DeleteRecordRequestStruct

		err := validator.BindJSON(jsonData, &singleData)

		if err != nil {
			fmt.Println("Error validating create request:", err)
			return
		}
		deleteOne(request)
	default:
		fmt.Println("Invalid request format for DeleteAllOperation")
	}
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
