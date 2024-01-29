package core

import (
	"fmt"
	"log"
	"pipebase/server/core/operations"
	"pipebase/server/enums"
	"pipebase/server/types"
)

func RouteOperationRequest(request types.RecordRequestStruct, session *types.Session) {
	fmt.Println("Received data", request)

	fmt.Printf("Request type: %s", request.Data.Type)

	switch request.Data.Type {
	case enums.CreateOperation:
		operations.SingleCreate(request)
	case enums.BulkCreateOperation:
		//operations.BulkCreate()
	case enums.ReadOneOperation:
		operations.ReadOne()
	case enums.ReadAllOperation:
		operations.ReadAll(request, session)
	case enums.UpdateOperation:
		operations.UpdateOne()
	case enums.BulkUpdateOperation:
		operations.UpdateBulk()
	case enums.DeleteOneOperation:
		operations.DeleteOne()
	case enums.DeleteAllOperation:
		operations.DeleteAll()
	default:
		log.Println("Unknown operation:", request.Data.Type)
	}
}
