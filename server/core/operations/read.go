package operations

import (
	"encoding/json"
	"fmt"
	"server/enums"
	"server/helpers"
	"server/types"
	"sync"

	"github.com/blevesearch/bleve"
)

var index bleve.Index
var indexesMap sync.Map

func StartIndexing() {
	tables := []string{"users"}

	for _, tableName := range tables {
		indexMapping := bleve.NewIndexMapping()
		index, _ := bleve.NewMemOnly(indexMapping)

		tableData, err := helpers.ReadTableData(tableName)

		if err != nil {
			fmt.Printf("Error reading %s.json: %v\n", tableName, err)
			return
		}

		for _, record := range tableData {
			jsonData, _ := json.Marshal(record)
			err = index.Index(getRecordID(record), string(jsonData))
			if err != nil {
				fmt.Printf("Error indexing document for table %s: %v\n", tableName, err)
				return
			}
		}

		indexesMap.Store(tableName, index)
	}
	fmt.Println("Indexing complete")
}

func getRecordID(record interface{}) string {
	idKey := "id"

	if recordMap, ok := record.(map[string]interface{}); ok {
		if id, idExists := recordMap[idKey]; idExists {
			if idStr, isString := id.(string); isString {
				return idStr
			}
		}
	}
	return ""
}

func HandleReadRequest(jsonData []byte, incomingRequest interface{}) ([]byte, error) {
	requestMap := incomingRequest.(map[string]interface{})
	dataMap := requestMap["data"].(map[string]interface{})
	requestType := dataMap["type"]
	tableName := dataMap["tableName"]

	var response []byte

	indexInterface, found := indexesMap.Load(tableName)

	if !found {
		return nil, fmt.Errorf("index not found for table %s", tableName)
	}

	index, ok := indexInterface.(bleve.Index)
	if !ok {
		return nil, fmt.Errorf("invalid index type for table %s", tableName)
	}

	if requestType == string(enums.ReadOneOperation) {
		var singleRead types.ReadOneRecordRequestStruct

		err := json.Unmarshal(jsonData, &singleRead)
		if err != nil {
			fmt.Println("Error unmarshaling read one request:", err)
			return nil, err
		}

		response, err = readOne(singleRead, index)

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

		response, err = readAll(bulkData, index)

		if err != nil {
			fmt.Println("Unable to read all records", err)
			return nil, err
		}
	}

	return response, nil
}

func readAll(request types.BulkReadRequestStruct, index bleve.Index) ([]byte, error) {
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

func readOne(request types.ReadOneRecordRequestStruct, index bleve.Index) ([]byte, error) {
	tableName := request.Data.TableName

	if !helpers.CheckIfTableExists(tableName) {
		return nil, fmt.Errorf("table %s does not exist", tableName)
	}
	queryValue := request.Data.Query.Value

	query := bleve.NewQueryStringQuery(queryValue)

	search := bleve.NewSearchRequest(query)

	searchResults, err := index.Search(search)

	if err != nil {
		return nil, err
	}

	if len(searchResults.Hits) == 0 {
		return nil, fmt.Errorf("record not found in table %s with the provided query", tableName)
	}

	recordID := searchResults.Hits[0].ID

	doc, err := index.Document(recordID)
	if err != nil {
		return nil, err
	}

	var jsonField string
	for _, field := range doc.Fields {
		if field.Name() == "" {
			jsonField = string(field.Value())
			break
		}
	}

	var recordMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonField), &recordMap); err != nil {
		return nil, err
	}

	jsonData, err := json.MarshalIndent(recordMap, "", "  ")
	if err != nil {
		return nil, err
	}

	response := append(jsonData, '\n')

	return response, nil
}
