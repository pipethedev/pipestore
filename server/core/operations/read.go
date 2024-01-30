package operations

import (
	"encoding/json"
	"fmt"
	"pipebase/server/enums"
	"pipebase/server/helpers"
	"pipebase/server/types"
)

func HandleReadRequest(jsonData []byte, incomingRequest interface{}) ([]byte, error) {
	requestMap := incomingRequest.(map[string]interface{})
	dataMap := requestMap["data"].(map[string]interface{})
	requestType := dataMap["type"]

	var response []byte

	if requestType == string(enums.ReadOneOperation) {
		var singleRead types.ReadOneRecordRequestStruct

		err := json.Unmarshal(jsonData, &singleRead)
		if err != nil {
			fmt.Println("Error unmarshaling read one request:", err)
			return nil, err
		}

		response, err = readOne(singleRead)

		if err != nil {
			fmt.Println("Unable to read one record", err)
			return nil, err
		}
	}

	if requestType == string(enums.ReadAllOperation) {
		var bulkData types.BulkReadRequestStruct

		err := json.Unmarshal(jsonData, &bulkData)
		if err != nil {
			fmt.Println("Error unmarshaling read all request:", err)
			return nil, err
		}

		response, err = readAll(bulkData)

		if err != nil {
			fmt.Println("Unable to read all records", err)
			return nil, err
		}
	}

	return response, nil
}

func readAll(request types.BulkReadRequestStruct) ([]byte, error) {
	data, err := helpers.ReadTableData(request.Data.TableName)

	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response := append(jsonData, '\n')

	return response, nil
}

func readOne(request types.ReadOneRecordRequestStruct) ([]byte, error) {
	tableName := request.Data.TableName

	if !helpers.CheckIfTableExists(tableName) {
		return nil, fmt.Errorf("table %s does not exist", tableName)
	}

	data, err := helpers.ReadTableData(tableName)
	if err != nil {
		return nil, err
	}

	index := -1
	for i, record := range data {
		if helpers.GetValueFromField(record, request.Data.Query.Field) == request.Data.Query.Value {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, fmt.Errorf("record not found in table %s with the provided query", tableName)
	}

	jsonData, err := json.Marshal(data[index])
	if err != nil {
		return nil, err
	}

	response := append(jsonData, '\n')

	return response, nil
}
