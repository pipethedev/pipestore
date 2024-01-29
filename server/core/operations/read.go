package operations

import (
	"encoding/json"
	"log"
	"pipebase/server/helpers"
	"pipebase/server/types"
)

func ReadAll(request types.RecordRequestStruct, session *types.Session) ([]byte, error) {
	data, err := helpers.ReadTableData(request.Data.TableName)

	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response := append(jsonData, '\n')

	_, err = session.Conn.Write(response)
	if err != nil {
		log.Fatalln("Error writing response:", err)
		return nil, err
	}

	return response, nil
}

func ReadOne() {}
