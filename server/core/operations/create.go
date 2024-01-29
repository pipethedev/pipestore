package operations

import (
	"fmt"
	"pipebase/server/helpers"
	"pipebase/server/types"
)

func BulkCreate(records []types.RecordRequestStruct) error {
	if len(records) == 0 {
		return fmt.Errorf("no records provided for bulk create")
	}

	tableName := records[0].Data.TableName

	if err := helpers.CreateTableFile(tableName); err != nil {
		return err
	}

	data, err := helpers.ReadTableData(tableName)
	if err != nil {
		return err
	}

	for _, record := range records {
		data = append(data, record.Data.Record)
	}

	return helpers.WriteTableData(tableName, data)
}

func SingleCreate(record types.RecordRequestStruct) error {
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
