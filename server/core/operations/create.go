package operations

import (
	"pipebase/server/helpers"
	"pipebase/server/types"
)

func BulkCreate(records types.BulkCreateRecordRequestStruct) error {
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

	data = append(data, records.Data.Record...)

	return helpers.WriteTableData(tableName, data)
}

func SingleCreate(record types.SingleCreateRecordRequestStruct) error {
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

	data = append(data, record.Data.Record)

	return helpers.WriteTableData(tableName, data)
}
