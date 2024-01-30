package operations

import (
	"encoding/json"
	"pipebase/server/helpers"
	"pipebase/server/types"
)

func ReadAll(request types.BulkReadRequestStruct) ([]byte, error) {
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

func ReadOne() {}
