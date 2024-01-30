package core

import (
	"encoding/json"
	"fmt"
	"pipebase/server/core/operations"
	"pipebase/server/enums"
	"pipebase/server/types"
)

func RouteOperationRequest(data []byte, genericRequest types.GenericRequest, session *types.Session) {
	switch genericRequest.Data.Type {
	case enums.CreateOperation:
		var SingleCreateRequest types.SingleCreateRecordRequestStruct
		err := json.Unmarshal(data, &SingleCreateRequest)
		if err != nil {
			fmt.Println("Error unmarshaling create request:", err)
			return
		}
		operations.SingleCreate(SingleCreateRequest)
	case enums.BulkCreateOperation:
		var BulkCreateRequest types.BulkCreateRecordRequestStruct
		err := json.Unmarshal(data, &BulkCreateRequest)
		if err != nil {
			fmt.Println("Error unmarshaling create request:", err)
			return
		}
		operations.BulkCreate(BulkCreateRequest)
	case enums.ReadAllOperation:
		var ReadAllRequest types.BulkReadRequestStruct
		err := json.Unmarshal(data, &ReadAllRequest)
		if err != nil {
			fmt.Println("Error unmarshaling create request:", err)
			return
		}
		operations.ReadAll(ReadAllRequest)
	case enums.UpdateOperation:
		var UpdateRequest types.UpdateRecordRequestStruct
		err := json.Unmarshal(data, &UpdateRequest)
		if err != nil {
			fmt.Println("Error unmarshaling create request:", err)
			return
		}
		operations.UpdateOne(UpdateRequest)
	}
}
