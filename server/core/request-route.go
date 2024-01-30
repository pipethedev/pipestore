package core

import (
	"encoding/json"
	"fmt"
	"pipebase/server/core/operations"
	"pipebase/server/enums"
	"pipebase/server/types"
)

func RouteOperationRequest(data []byte, genericRequest types.GenericRequest, session *types.Session) {
	var request interface{}
	var response []byte

	err := json.Unmarshal(data, &request)
	if err != nil {
		fmt.Println("Error unmarshaling request:", err)
		return
	}

	switch genericRequest.Data.Type {
	case enums.CreateOperation, enums.BulkCreateOperation:
		operations.HandleCreateRequest(data, request)
	case enums.ReadOneOperation, enums.ReadAllOperation:
		response, _ = operations.HandleReadRequest(data, request)
	case enums.UpdateOperation:
		operations.HandleUpdateRequest(request)
	case enums.DeleteOneOperation, enums.DeleteAllOperation:
		operations.HandleDeleteRequest(data, request)
	default:
		fmt.Println("Unknown operation:", genericRequest.Data.Type)
	}

	_, _ = session.Conn.Write(response)
}
