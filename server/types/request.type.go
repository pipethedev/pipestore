package types

import (
	"server/enums"
	"time"
)

type KeyValue struct {
	Key   string
	Value interface{}
}

type AuthRequestStruct struct {
	Auth struct {
		Username string `json:"username"`
		APIKey   string `json:"apiKey"`
	} `json:"auth"`
}

type RecordWithTimestamps struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GenericRequest struct {
	Data struct {
		Type enums.Operation `json:"type" enum:"CREATE,BULK_CREATE,READ_ONE,READ_ALL,UPDATE,DELETE_ONE,DELETE_ALL" binding:"required"`
	} `json:"data"`
}

type SingleCreateRecordRequestStruct struct {
	Data struct {
		Type      enums.Operation        `json:"type" enum:"CREATE" binding:"required"`
		TableName string                 `json:"tableName"`
		Record    map[string]interface{} `json:"record"`
	} `json:"data"`
}

type BulkCreateRecordRequestStruct struct {
	Data struct {
		Type      enums.Operation          `json:"type" enum:"BULK_CREATE" binding:"required"`
		TableName string                   `json:"tableName"`
		Record    []map[string]interface{} `json:"record"`
	} `json:"data"`
}

type UpdateRecordRequestStruct struct {
	Data struct {
		Type      enums.Operation `json:"type" enum:"UPDATE" binding:"required"`
		TableName string          `json:"tableName"`
		Query     struct {
			Field string `json:"field"`
			Value string `json:"value"`
		} `json:"query"`

		Record interface{} `json:"record"`
	} `json:"data"`
}

type DeleteRecordRequestStruct struct {
	Data struct {
		Type      enums.Operation `json:"type" enum:"DELETE_ONE" binding:"required"`
		TableName string          `json:"tableName"`
		Query     struct {
			Field string `json:"field"`
			Value string `json:"value"`
		} `json:"query"`
	} `json:"data"`
}

type BulkDeleteRecordRequestStruct struct {
	Data struct {
		Type      enums.Operation `json:"type" enum:"DELETE_ALL" binding:"required"`
		TableName string          `json:"tableName"`
	} `json:"data"`
}

type ReadOneRecordRequestStruct struct {
	Data struct {
		Type      enums.Operation `json:"type" enum:"READ_ONE" binding:"required"`
		TableName string          `json:"tableName"`
		Query     struct {
			Field    string         `json:"field"`
			Operator enums.Operator `json:"operator"`
			Value    interface{}    `json:"value"`
		} `json:"query"`
	} `json:"data"`
}

type BulkReadRequestStruct struct {
	Data struct {
		Type      enums.Operation `json:"type" enum:"READ_ALL" binding:"required"`
		TableName string          `json:"tableName"`
	} `json:"data"`
}

type ConstraintRequestStruct struct {
	Data struct {
		Type       enums.Operation `json:"type" enum:"ADD_CONSTRAINT,REMOVE_CONSTRAINT"`
		TableName  string          `json:"tableName"`
		Constraint struct {
			Name  string `json:"name" enum:"UNIQUE"`
			Field string `json:"field"`
		} `json:"query"`
	} `json:"data"`
}
