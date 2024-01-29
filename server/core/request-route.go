package core

import (
	"log"
	"pipebase/server/core/operations"
	"pipebase/server/enums"
	"pipebase/server/types"
)

func RouteOperationRequest(request types.RecordRequestStruct) {
	switch request.Type {
	case enums.CreateOperation:
		operations.SingleCreate()
	case enums.BulkCreateOperation:
		operations.BulkCreate()
	case enums.ReadOneOperation:
		operations.ReadOne()
	case enums.ReadAllOperation:
		operations.ReadAll()
	case enums.UpdateOperation:
		operations.UpdateOne()
	case enums.BulkUpdateOperation:
		operations.UpdateBulk()
	case enums.DeleteOneOperation:
		operations.DeleteOne()
	case enums.DeleteAllOperation:
		operations.DeleteAll()
	default:
		log.Println("Unknown operation:", request.Type)
	}
}
