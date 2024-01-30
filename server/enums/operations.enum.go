package enums

type Operation string

const (
	CreateOperation Operation = "CREATE"

	BulkCreateOperation Operation = "BULK_CREATE"

	ReadOneOperation Operation = "READ_ONE"

	ReadAllOperation Operation = "READ_ALL"

	UpdateOperation Operation = "UPDATE"

	DeleteOneOperation Operation = "DELETE_ONE"

	DeleteAllOperation Operation = "DELETE_ALL"
)
